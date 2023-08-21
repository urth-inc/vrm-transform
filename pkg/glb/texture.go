package glb

import (
	"golang.org/x/exp/slices"
	"os/exec"

	"github.com/google/uuid"
	"github.com/h2non/bimg"
	"github.com/qmuntal/gltf"
)

func resizeImage(buf []byte, width, height int) (image []byte, err error) {
	// TODO: check if image is PNG or JPEG
	// TODO: check if image is larger than width and height
	image, err = bimg.NewImage(buf).Enlarge(width, height)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func toKtx2Image(buf []byte) (image []byte, err error) {
    // バージョン4のランダムなUUIDを生成します
    var path: string = uuid.New()

	// path := fmt.Sprintf("image_%d", idx)
	// dumpBin(img, path)
	// dumpBin(b, path)

	fmt.Println("path:", path)

	cmd := exec.Command("toktx", "--2d", "--genmipmap", "--target_type", "RGBA", "--t2", "--encode", "etc1s", "--clevel", "5", "--qlevel", "255", path, path)
	// cmd := exec.Command("toktx", "--2d", "--genmipmap", "--bcmp", "--target_type", "RGBA", "--t2", "--encode", "uastc", path, path)

	// env := os.Environ()
	// env = append(env, "PATH=/home/kira/Downloads/apps/KTX-Software-4.2.1-Linux-x86_64/bin:"+os.Getenv("PATH"))
	// env = append(env, "LD_LIBRARY_PATH=/home/kira/Downloads/apps/KTX-Software-4.2.1-Linux-x86_64/lib:"+os.Getenv("LD_LIBRARY_PATH"))
	// cmd.Env = env

	// err := cmd.Run()
	// if err != nil {
	// >fmt.Println(err)
	// }

	// ktx := readFile(path + ".ktx2")
	// newBin = append(newBin, ktx...)
	// doc.BufferViews[idx].ByteLength = uint32(len(ktx))
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

func (g *GLB) ToKTX2Texture() (err error) {
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

	g.BIN = newBin

	return nil
}
