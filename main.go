package main

import (
	"fmt"
	"github.com/urth-inc/vrm-transform/pkg/glb"
	"github.com/urth-inc/vrm-transform/pkg/vrm"
	"io/ioutil"
	"os"
)

func main() {
	filePath := "./assets/avatar02.vrm"
	// filePath := "./assets/BoxFox.glb"
	// filePath := "./assets/kemomimi.vrm"
	// filePath := "./assets/avatar01_0806.vrm"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("File open error:", err)
		return
	}
	defer file.Close()

	// ファイルの内容を[]byteに読み込む
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("File read error:", err)
		return
	}

	fmt.Println(vrm.IsVRM(fileData))

	myglb, err := glb.ReadBinary(fileData)
	if err != nil {
		fmt.Println("File read error:", err)
		return
	}

	// err = myglb.ResizeTexture(128, 128)
	// if err != nil {
	// 	fmt.Println("File read error:", err)
	// }

	err = myglb.ToKtx2Texture()
	if err != nil {
		fmt.Println("File read error:", err)
	}

	buf, err := glb.WriteBinary(myglb)
	if err != nil {
		fmt.Println("File write error:", err)
		return
	}

	// filePath = "output.glb"
	filePath = "avatar.vrm"
	file, err = os.Create(filePath)
	if err != nil {
		fmt.Println("File create error:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(buf)
	if err != nil {
		fmt.Println("File write error:", err)
		return
	}
}
