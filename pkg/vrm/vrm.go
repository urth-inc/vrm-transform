package vrm

import (
	"encoding/binary"
	"encoding/json"
	"github.com/qmuntal/gltf"
)

const (
	ChunkTypeJSON     = 0x4e4f534a
	GLB_HEADER_LENGTH = 12
	GLB_MAGIC         = 0x46546C67
)

func isGLB(view []byte) bool {
	var magic uint32 = binary.LittleEndian.Uint32(view[:4])
	var version uint32 = binary.LittleEndian.Uint32(view[4:8])

	return magic == GLB_MAGIC && version == 2
}

func IsVRM(glb []byte) bool {
	if !isGLB(glb) {
		return false
	}

	// Decode JSON chunk.
	var offset uint32 = GLB_HEADER_LENGTH
	var jsonByteLength uint32 = binary.LittleEndian.Uint32(glb[offset : offset+4])
	offset += 4

	var jsonChunkType uint32 = binary.LittleEndian.Uint32(glb[offset : offset+4])
	offset += 4

	if jsonChunkType != ChunkTypeJSON {
		return false
	}

	var jsonDocument gltf.Document
	err := json.Unmarshal(glb[offset:offset+jsonByteLength], &jsonDocument)
	if err != nil {
		return false
	}

	var vrmExt bool = jsonDocument.Extensions["VRM"] != nil || jsonDocument.Extensions["VRMC_vrm"] != nil

	return vrmExt
}
