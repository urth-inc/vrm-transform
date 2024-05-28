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

func TestToKtx2Texture(t *testing.T) {
	file, err := os.Open("../../test/Duck.glb")

	if err != nil {
		fmt.Println("File open error:", err)
		return
	}

	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("File read error:", err)
		return
	}

	glb, err := ReadBinary(buf)
	if err != nil {
		fmt.Println("Glb read error:", err)
		return
	}

	err = glb.ToKtx2Texture("uastc", -1, 2, 3)
	// err = myglb.ToKtx2Texture("etc1s", 128, -1, -1)

	if err != nil {
		fmt.Println("ToKtx2Texture error:", err)
	}

	res, err := WriteBinary(glb)
	if err != nil {
		fmt.Println("Glb write error:", err)
		return
	}

	if res == nil {
		fmt.Println("Glb write error: res is nil")
		return
	}
}
