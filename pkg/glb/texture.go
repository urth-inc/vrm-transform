package glb

import (
	"golang.org/x/exp/slices"

	"github.com/h2non/bimg"
	"github.com/qmuntal/gltf"
)

func resizeImage(buf []byte, width, height int) (image []byte, err error) {
	image, err = bimg.NewImage(buf).Enlarge(width, height)
	if err != nil {
		return nil, err
	}

	return image, nil
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

	return nil
}
