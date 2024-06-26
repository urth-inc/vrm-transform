package glb

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"golang.org/x/exp/slices"

	"github.com/google/uuid"
	"github.com/h2non/bimg"
	"github.com/qmuntal/gltf"
	"github.com/urth-inc/vrm-transform/internal/fileUtil"
	"github.com/urth-inc/vrm-transform/internal/imageUtil"
	"github.com/urth-inc/vrm-transform/internal/interfaces"
)

type DefaultConvertToKtx2ImageDependencies struct{}

func (d *DefaultConvertToKtx2ImageDependencies) UUIDGenerator() string {
	return uuid.New().String()
}

func (d *DefaultConvertToKtx2ImageDependencies) ContentTypeDetector(data []byte) string {
	return http.DetectContentType(data)
}

func (d *DefaultConvertToKtx2ImageDependencies) ImageSizer(data []byte) (int, int, error) {
	return imageUtil.GetImageSize(data)
}

func (d *DefaultConvertToKtx2ImageDependencies) CommandExecutor(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	return cmd.Run()
}

func (d *DefaultConvertToKtx2ImageDependencies) ParamsGenerator(mode string, width, height int, inputPath, outputPath string, isSRGB bool, etc1sQuality, uastcQuality, zstdLevel int) []string {
	return getKtx2Params(mode, width, height, inputPath, outputPath, isSRGB, etc1sQuality, uastcQuality, zstdLevel)
}

func (d *DefaultConvertToKtx2ImageDependencies) FileReader(filePath string) ([]byte, error) {
	return fileUtil.ReadFile(filePath)
}

func (d *DefaultConvertToKtx2ImageDependencies) FileCreator(filePath string) (interfaces.File, error) {
	return os.Create(filePath)
}

func (d *DefaultConvertToKtx2ImageDependencies) FileRemover(filePath string) error {
	return os.Remove(filePath)
}

type DefaultConvertToKtx2TextureDependencies struct{}

func (d *DefaultConvertToKtx2TextureDependencies) ConvertToKtx2Image(deps interfaces.ConvertToKtx2ImageDependenciesInterface, ktx2Mode string, buf []byte, isSRGB bool, etc1sQuality int, uastcQuality int, zstdLevel int) ([]byte, error) {
	return convertToKtx2Image(deps, ktx2Mode, buf, isSRGB, etc1sQuality, uastcQuality, zstdLevel)
}

func resizeImage(buf []byte, width, height int) (image []byte, err error) {
	newImage := bimg.NewImage(buf)

	size, err := newImage.Size()
	if err != nil {
		return nil, err
	}

	if size.Width <= width && size.Height <= height {
		return buf, nil
	}

	image, err = newImage.Enlarge(width, height)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func getKtx2Params(ktx2Mode string, width int, height int, inputPath string, outputPath string, isSRGB bool, etc1sQuality int, uastcQuality int, zstdLevel int) []string {
	if etc1sQuality < 1 || etc1sQuality > 255 {
		etc1sQuality = 128
	}
	if uastcQuality < 0 || uastcQuality > 4 {
		uastcQuality = 2
	}
	if zstdLevel < 1 || zstdLevel > 22 {
		zstdLevel = 3
	}

	var params []string = make([]string, 0)

	// ref: https://github.khronos.org/KTX-Software/ktxtools/toktx.html
	params = append(params, "--genmipmap")
	params = append(params, "--t2")

	if !isSRGB {
		params = append(params, "--assign_oetf", "linear", "--assign_primaries", "none")
	}

	switch ktx2Mode {
	case "etc1s":
		params = append(params, "--encode", "etc1s")
		params = append(params, "--clevel", "1")
		params = append(params, "--qlevel", strconv.Itoa(etc1sQuality))
	default:
		params = append(params, "--encode", "uastc")
		params = append(params, "--uastc_quality", strconv.Itoa(uastcQuality))
		params = append(params, "--zcmp", strconv.Itoa(zstdLevel))
	}

	if width%4 != 0 || height%4 != 0 {
		width += (4 - width%4) % 4
		height += (4 - height%4) % 4
		params = append(params, "--resize", strconv.Itoa(width)+"x"+strconv.Itoa(height))
	}

	params = append(params, outputPath, inputPath)

	return params
}

func convertToKtx2Image(deps interfaces.ConvertToKtx2ImageDependenciesInterface, ktx2Mode string, buf []byte, isSRGB bool, etc1sQuality int, uastcQuality int, zstdLevel int) (image []byte, err error) {
	var mimeType string = deps.ContentTypeDetector(buf)
	var inputPath string = "/tmp/" + deps.UUIDGenerator()
	var outputPath string = "/tmp/" + deps.UUIDGenerator()

	if mimeType == "image/png" {
		inputPath += ".png"
	} else if mimeType == "image/jpeg" {
		inputPath += ".jpg"
	} else {
		return nil, fmt.Errorf("invalid image type: %s", mimeType)
	}

	file, err := deps.FileCreator(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.Write(buf)
	if err != nil {
		return nil, err
	}

	width, height, err := deps.ImageSizer(buf)
	if err != nil {
		return nil, err
	}

	params := deps.ParamsGenerator(ktx2Mode, width, height, inputPath, outputPath, isSRGB, etc1sQuality, uastcQuality, zstdLevel)
	err = deps.CommandExecutor("toktx", params...)
	if err != nil {
		return nil, err
	}

	// toktx add .ktx2 extension automatically
	outputPath += ".ktx2"

	ktx2file, err := deps.FileReader(outputPath)
	if err != nil {
		return nil, err
	}

	if err = deps.FileRemover(inputPath); err != nil {
		return nil, err
	}
	if err = deps.FileRemover(outputPath); err != nil {
		return nil, err
	}

	return ktx2file, nil
}

func (g *GLB) ResizeTexture(width, height int) (err error) {
	var jsonDocument gltf.Document = g.GltfDocument
	var bin []byte = g.BIN

	imagesBufferViews := make([]uint32, 0)

	for _, image := range jsonDocument.Images {
		imagesBufferViews = append(imagesBufferViews, *image.BufferView)
	}

	var newBin []byte = make([]byte, 0)

	var offset uint32 = 0
	for idx, bufferView := range jsonDocument.BufferViews {
		var byteOffset uint32 = 0
		if idx != 0 {
			byteOffset = bufferView.ByteOffset
		}

		data := bin[byteOffset : byteOffset+bufferView.ByteLength]

		if slices.Contains(imagesBufferViews, uint32(idx)) {
			img, err := resizeImage(data, width, height)
			if err != nil {
				return err
			}

			newBin = append(newBin, img...)
			jsonDocument.BufferViews[idx].ByteLength = uint32(len(img))
		} else {
			newBin = append(newBin, data...)
			jsonDocument.BufferViews[idx].ByteLength = uint32(len(data))
		}

		jsonDocument.BufferViews[idx].ByteOffset = offset
		offset += jsonDocument.BufferViews[idx].ByteLength
	}

	g.BIN = newBin
	g.GltfDocument.Buffers[0].ByteLength = uint32(len(newBin))

	return nil
}

func (g *GLB) ToKtx2Texture(deps interfaces.ConvertToKtx2TextureDependenciesInterface, ktx2Mode string, etc1sQuality int, uastcQuality int, zstdLevel int) (err error) {
	var jsonDocument gltf.Document = g.GltfDocument
	var bin []byte = g.BIN

	imagesBufferViews := make([]uint32, 0)

	for _, image := range jsonDocument.Images {
		imagesBufferViews = append(imagesBufferViews, *image.BufferView)
		image.MimeType = "image/ktx2"
	}

	var newBin []byte = make([]byte, 0)

	var offset uint32 = 0
	for idx, bufferView := range jsonDocument.BufferViews {
		var byteOffset uint32 = 0
		if idx != 0 {
			byteOffset = bufferView.ByteOffset
		}

		data := bin[byteOffset : byteOffset+bufferView.ByteLength]

		if slices.Contains(imagesBufferViews, uint32(idx)) {
			// some models lie about the texture slot, so we always treat texture as non-color
			isSRGB := false

			img, err := deps.ConvertToKtx2Image(&DefaultConvertToKtx2ImageDependencies{}, ktx2Mode, data, isSRGB, etc1sQuality, uastcQuality, zstdLevel)

			if err != nil {
				return err
			}

			newBin = append(newBin, img...)
			jsonDocument.BufferViews[idx].ByteLength = uint32(len(img))
		} else {
			newBin = append(newBin, data...)
			jsonDocument.BufferViews[idx].ByteLength = uint32(len(data))
		}

		jsonDocument.BufferViews[idx].ByteOffset = offset
		offset += jsonDocument.BufferViews[idx].ByteLength
	}

	for idx, texture := range g.GltfDocument.Textures {
		g.GltfDocument.Textures[idx].Extensions = map[string]interface{}{
			"KHR_texture_basisu": map[string]interface{}{
				"source": texture.Source,
			},
		}
		g.GltfDocument.Textures[idx].Source = nil
	}

	g.GltfDocument.ExtensionsUsed = append(jsonDocument.ExtensionsUsed, "KHR_texture_basisu")
	g.GltfDocument.ExtensionsRequired = append(jsonDocument.ExtensionsRequired, "KHR_texture_basisu")

	g.BIN = newBin
	g.GltfDocument.Buffers[0].ByteLength = uint32(len(newBin))

	return nil
}
