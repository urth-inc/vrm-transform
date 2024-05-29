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
)

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

func toKtx2Image(ktx2Mode string, buf []byte, isSRGB bool, etc1sQuality int, uastcQuality int, zstdLevel int) (image []byte, err error) {
	var mimeType string = http.DetectContentType(buf)

	var inputPath string = "/tmp/" + uuid.New().String()
	var outputPath string = "/tmp/" + uuid.New().String()
	if mimeType == "image/png" {
		inputPath += ".png"
	} else if mimeType == "image/jpeg" {
		inputPath += ".jpg"
	} else {
		return nil, fmt.Errorf("invalid image type: %s", mimeType)
	}

	file, err := os.Create(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.Write(buf)
	if err != nil {
		return nil, err
	}

	var width, height int
	width, height, err = imageUtil.GetImageSize(buf)
	if err != nil {
		return nil, err
	}

	var params []string = getKtx2Params(ktx2Mode, width, height, inputPath, outputPath, isSRGB, etc1sQuality, uastcQuality, zstdLevel)

	err = exec.Command("toktx", params...).Run()
	if err != nil {
		return nil, err
	}

	outputPath += ".ktx2"

	ktx2file, err := fileUtil.ReadFile(outputPath)
	if err != nil {
		return nil, err
	}

	os.Remove(inputPath)
	os.Remove(outputPath)

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

// estimate whether the texture should be sRGB or not from the texture slot
// should sRGB: baseColorTexture, emissiveTexture
// should linear: normalTexture, occlusionTexture, metallicRoughnessTexture
func getIsSrgbMap(gltfDocument gltf.Document) map[uint32]bool {
	isSRGBs := make(map[uint32]bool)

	// Helper function to add texture indices to the map
	addToMap := func(texture *gltf.TextureInfo) {
		if texture != nil {
			isSRGBs[uint32(texture.Index)] = true
		}
	}

	// Golang return zero-value for non-exist key
	// We just need to set true for sRGB texture
	for _, material := range gltfDocument.Materials {
		if material.PBRMetallicRoughness != nil {
			addToMap(material.PBRMetallicRoughness.BaseColorTexture)
		}
		addToMap(material.EmissiveTexture)
	}

	return isSRGBs
}

func getBufferViewIndex2TextureIndex(gltfDocument gltf.Document) map[uint32][]uint32 {
	imageToBufferView := make(map[uint32]uint32)

	for imageIndex, image := range gltfDocument.Images {
		if image.BufferView != nil {
			imageToBufferView[uint32(imageIndex)] = *image.BufferView
		}
	}

	bufferViewToTextures := make(map[uint32][]uint32)
	for textureIndex, texture := range gltfDocument.Textures {
		imageIndex := texture.Source
		if bufferViewIndex, exists := imageToBufferView[*imageIndex]; exists {
			bufferViewToTextures[uint32(bufferViewIndex)] = append(bufferViewToTextures[uint32(bufferViewIndex)], uint32(textureIndex))
		}
	}

	return bufferViewToTextures
}

func (g *GLB) ToKtx2Texture(ktx2Mode string, etc1sQuality int, uastcQuality int, zstdLevel int) (err error) {
	var jsonDocument gltf.Document = g.GltfDocument
	var bin []byte = g.BIN

	imagesBufferViews := make([]uint32, 0)

	// isSRGBs := getIsSrgbMap(jsonDocument)
	// bufferViewToTexture := getBufferViewIndex2TextureIndex(jsonDocument)

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
			isSRGB := false

			// some models lie about the texture slot, so we need to comment out this part and always treat as non sRGB
			// if textures, exists := bufferViewToTexture[uint32(idx)]; exists {
			// 	for _, textureIndex := range textures {
			// 		if isSRGBs[textureIndex] {
			// 			isSRGB = true
			// 		}
			// 	}
			// }

			img, err := toKtx2Image(ktx2Mode, data, isSRGB, etc1sQuality, uastcQuality, zstdLevel)
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
