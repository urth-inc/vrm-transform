package imageUtil

import (
	"github.com/h2non/bimg"
)

func GetImageSize(buf []byte) (width, height int, err error) {
	newImage := bimg.NewImage(buf)

	size, err := newImage.Size()
	if err != nil {
		return 0, 0, err
	}

	return size.Width, size.Height, nil
}
