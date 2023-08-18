package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"vrm-transform/pkg/glb"
)

func main() {
	filePath := "./assets/avatar01_0806.vrm"
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

	myglb, err := glb.ReadBinary(fileData)
	if err != nil {
		fmt.Println("File read error:", err)
		return
	}

	err = myglb.ResizeTexture(128, 128)
	if err != nil {
		fmt.Println("File read error:", err)
	}

	buf, err := glb.WriteBinary(myglb)
	if err != nil {
		fmt.Println("File write error:", err)
		return
	}

	filePath = "output.glb"
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
