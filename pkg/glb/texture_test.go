package glb

import (
	"testing"
)

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
