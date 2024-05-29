package glb

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestToKtx2Image(t *testing.T) {
	file, err := os.Open("../../test/Duck_baseColorTexture.png")

	if err != nil {
		t.Errorf("Open file error: %v", err)
	}

	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf("Read file error: %v", err)
	}

	etc1sFile, err := toKtx2Image("etc1s", buf, false, 128, -1, -1)

	if err != nil {
		t.Errorf("toKtx2Image error: %v", err)
	}

	if etc1sFile == nil {
		t.Errorf("toKtx2Image error: ktx2file is nil")
	}

	uastcFile, err := toKtx2Image("uastc", buf, false, -1, 2, 3)
	if err != nil {
		t.Errorf("toKtx2Image error: %v", err)
	}

	if uastcFile == nil {
		t.Errorf("toKtx2Image error: ktx2file is nil")
	}
}

func prepareGLB(filePath string) (*GLB, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("File open error: %v", err)
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("File read error: %v", err)
	}

	glb, err := ReadBinary(buf)
	if err != nil {
		return nil, fmt.Errorf("GLB read error: %v", err)
	}

	return &glb, nil
}

func TestToKtx2TextureUASTC(t *testing.T) {
	glb, err := prepareGLB("../../test/Duck.glb")

	if err != nil {
		fmt.Println("Glb read error:", err)
		return
	}

	err = glb.ToKtx2Texture("uastc", -1, 2, 3)

	if err != nil {
		fmt.Println("ToKtx2Texture error:", err)
	}

	res, err := WriteBinary(*glb)
	if err != nil {
		fmt.Println("Glb write error:", err)
		return
	}

	if res == nil {
		fmt.Println("Glb write error: res is nil")
		return
	}
}

func TestToKtx2TextureETC1S(t *testing.T) {
	glb, err := prepareGLB("../../test/Duck.glb")

	if err != nil {
		fmt.Println("Glb read error:", err)
		return
	}

	err = glb.ToKtx2Texture("etc1s", 128, -1, -1)

	if err != nil {
		fmt.Println("ToKtx2Texture error:", err)
	}

	res, err := WriteBinary(*glb)
	if err != nil {
		fmt.Println("Glb write error:", err)
		return
	}

	if res == nil {
		fmt.Println("Glb write error: res is nil")
		return
	}
}

func TestGetKtx2Params(t *testing.T) {
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
