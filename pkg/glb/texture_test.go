package glb

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	mock_glb "github.com/urth-inc/vrm-transform/test/mocks"

	"go.uber.org/mock/gomock"
)

func TestGetKtx2Params(t *testing.T) {
	// TODO: add other test cases
	cases := []struct {
		name           string
		ktx2Mode       string
		width          int
		height         int
		inputPath      string
		outputPath     string
		isSRGB         bool
		etc1sQuality   int
		uastcQuality   int
		zstdLevel      int
		expectedParams []string
	}{
		{
			name:           "ETC1Sデフォルト品質",
			ktx2Mode:       "etc1s",
			width:          1024,
			height:         512,
			inputPath:      "input.ktx",
			outputPath:     "output.ktx",
			isSRGB:         false,
			etc1sQuality:   300,
			uastcQuality:   2,
			zstdLevel:      3,
			expectedParams: []string{"--genmipmap", "--t2", "--assign_oetf", "linear", "--assign_primaries", "none", "--encode", "etc1s", "--clevel", "1", "--qlevel", "128", "output.ktx", "input.ktx"},
		},
		{
			name:           "UASTCリサイズあり",
			ktx2Mode:       "uastc",
			width:          1027,
			height:         513,
			inputPath:      "input.ktx",
			outputPath:     "output.ktx",
			isSRGB:         true,
			etc1sQuality:   50,
			uastcQuality:   5,
			zstdLevel:      23,
			expectedParams: []string{"--genmipmap", "--t2", "--encode", "uastc", "--uastc_quality", "2", "--zcmp", "3", "--resize", "1028x516", "output.ktx", "input.ktx"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := getKtx2Params(c.ktx2Mode, c.width, c.height, c.inputPath, c.outputPath, c.isSRGB, c.etc1sQuality, c.uastcQuality, c.zstdLevel)
			if len(got) != len(c.expectedParams) {
				t.Errorf("Expected params length %d, got %d", len(c.expectedParams), len(got))
			}
			for i, expected := range c.expectedParams {
				if got[i] != expected {
					t.Errorf("Expected param at index %d to be %s, got %s", i, expected, got[i])
				}
			}
		})
	}
}

// TestConvertToKtx2Image tests the convertToKtx2Image utility function
func TestToKtx2Image(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFile := mock_glb.NewMockFile(ctrl)
	mockFile.EXPECT().Write(gomock.Any()).Return(0, nil)
	mockFile.EXPECT().Close().Return(nil)

	mockDeps := mock_glb.NewMockConvertToKtx2ImageDependenciesInterface(ctrl)

	// TODO: add other test cases
	// Test data
	testData := []byte("test image data")
	uuid := "unique-id"
	inputPath := "/tmp/" + uuid + ".png"
	outputPath := "/tmp/" + uuid
	mode := "UASTC"
	isSRGB := false
	etc1sQuality, uastcQuality, zstdLevel := 128, 3, 4
	width, height := 1024, 1024

	mockDeps.EXPECT().ContentTypeDetector(testData).Return("image/png")
	mockDeps.EXPECT().UUIDGenerator().Return(uuid).Times(2)
	mockDeps.EXPECT().FileCreator(inputPath).Return(mockFile, nil)
	mockDeps.EXPECT().ImageSizer(testData).Return(width, height, nil)
	mockDeps.EXPECT().ParamsGenerator(mode, width, height, inputPath, outputPath, isSRGB, etc1sQuality, uastcQuality, zstdLevel).Return([]string{"toktx", "--t2", outputPath, inputPath})
	mockDeps.EXPECT().CommandExecutor("toktx", gomock.Any()).Return(nil)
	mockDeps.EXPECT().FileReader(outputPath+".ktx2").Return([]byte("ktx2 image data"), nil)
	mockDeps.EXPECT().FileRemover(inputPath).Return(nil)
	mockDeps.EXPECT().FileRemover(outputPath + ".ktx2").Return(nil)

	result, err := convertToKtx2Image(mockDeps, mode, testData, isSRGB, etc1sQuality, uastcQuality, zstdLevel)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := []byte("ktx2 image data")
	if !bytes.Equal(result, expectedResult) {
		t.Errorf("Expected result %v, got %v", expectedResult, result)
	}
}

// TestToKtx2Texture tests the ToKtx2Texture method of GLB
func TestToKtx2Texture(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeps := mock_glb.NewMockConvertToKtx2TextureDependenciesInterface(ctrl)

	// TODO: add other test cases
	// TODO: remove external dependency on test
	file, err := os.Open("../../test/Duck.glb")
	if err != nil {
		fmt.Println("File open error:", err)
		return
	}
	defer file.Close()

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("File read error:", err)
		return
	}

	test_glb, err := ReadBinary(fileData)
	if err != nil {
		fmt.Println("File read error:", err)
		return
	}

	// Set up expectations for the mock object
	mockDeps.EXPECT().
		ConvertToKtx2Image(gomock.Any(), "uastc", gomock.Any(), false, -1, 2, 3).
		Return([]byte{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, nil).
		Times(1)

	// Execute the method under test
	err = test_glb.ToKtx2Texture(mockDeps, "uastc", -1, 2, 3)

	// Assert that there was no error and the expected changes were made
	assert.NoError(t, err)
	assert.Equal(t, uint32(0x18ea2), test_glb.GltfDocument.Buffers[0].ByteLength)
}
