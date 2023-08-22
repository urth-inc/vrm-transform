package glb

import (
	"fmt"
	"golang.org/x/exp/slices"
	"net/http"
	"os"
	"os/exec"

	"github.com/google/uuid"
	"github.com/h2non/bimg"
	"github.com/qmuntal/gltf"
	"github.com/urth-inc/vrm-transform/internal/fileUtil"
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

func toKtx2Image(buf []byte) (image []byte, err error) {
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

	cmd := exec.Command("toktx", "--bcmp", "--threads", "2", "--2d", "--genmipmap", "--t2", "--encode", "etc1s", "--clevel", "1", "--qlevel", "255", outputPath, inputPath)

	err = cmd.Run()
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

func (g *GLB) ToKtx2Texture() (err error) {
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
			img, err := toKtx2Image(data)
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
