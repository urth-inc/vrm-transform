package main

import (
	"fmt"
	"github.com/qmuntal/gltf"
	"github.com/urth-inc/vrm-transform/pkg/glb"
	"github.com/urth-inc/vrm-transform/pkg/vrm"

	"io/ioutil"
	"os"
)

func jsonDump(g glb.GLB, path string) {
	doc := g.GltfDocument
	gltf.Save(&doc, "./"+path)
}

// this script is for debug
func main() {
	filePath := "./assets/test.glb"

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

	fmt.Println("isVRM:", vrm.IsVRM(fileData))

	myglb, err := glb.ReadBinary(fileData)
	if err != nil {
		fmt.Println("File read error:", err)
		return
	}

	jsonDump(myglb, "before.json")
	err = myglb.ResizeTexture(1024, 1024)
	if err != nil {
		fmt.Println("File read error:", err)
	}

	err = myglb.ToKtx2Texture()
	if err != nil {
		fmt.Println("File read error:", err)
	}
	jsonDump(myglb, "after.json")

	buf, err := glb.WriteBinary(myglb)
	if err != nil {
		fmt.Println("File write error:", err)
		return
	}

	filePath = "output.glb"
	// filePath = "avatar.vrm"
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

	gltf.Open("output.glb")
}
