package glb

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/qmuntal/gltf"
)

const (
	ChunkTypeJSON     = 0x4e4f534a
	ChunkTypeBIN      = 0x004e4942
	GLB_HEADER_LENGTH = 12
	GLB_MAGIC         = 0x46546C67
)

type GLB struct {
	GltfDocument gltf.Document
	BIN          []byte
}

func isGLB(view []byte) bool {
	var magic uint32 = binary.LittleEndian.Uint32(view[:4])
	var version uint32 = binary.LittleEndian.Uint32(view[4:8])

	return magic == GLB_MAGIC && version == 2
}

func ReadBinary(input_glb []byte) (glb GLB, err error) {
	if !isGLB(input_glb) {
		err = errors.New("not a glTF2.0 file")
		return glb, err
	}

	var jsonDocument gltf.Document
	var bin []byte

	// Decode JSON chunk.
	var offset uint32 = GLB_HEADER_LENGTH
	var jsonByteLength uint32 = binary.LittleEndian.Uint32(input_glb[offset : offset+4])
	offset += 4

	var jsonChunkType uint32 = binary.LittleEndian.Uint32(input_glb[offset : offset+4])
	offset += 4

	if jsonChunkType != ChunkTypeJSON {
		err = errors.New("Missing required GLB JSON chunk.")
		return glb, err
	}

	err = json.Unmarshal(input_glb[offset:offset+jsonByteLength], &jsonDocument)
	if err != nil {
		err = errors.New("JSON Unmarshal error")
		return glb, err
	}

	offset += jsonByteLength

	// Decode BIN chunk.
	var glbByteLength uint32 = binary.LittleEndian.Uint32(input_glb[8:12])
	if glbByteLength <= offset {
		return glb, err
	}

	var binChunkLength uint32 = binary.LittleEndian.Uint32(input_glb[offset : offset+4])
	offset += 4

	var binChunkType uint32 = binary.LittleEndian.Uint32(input_glb[offset : offset+4])
	offset += 4

	if binChunkType != ChunkTypeBIN {
		err = errors.New("Expected GLB BIN in second chunk.")
		return glb, err
	}

	bin = input_glb[offset : offset+binChunkLength]
	return GLB{
		GltfDocument: jsonDocument,
		BIN:          bin,
	}, err
}

func WriteBinary(glb GLB) (output_glb []byte, err error) {
	jsonData, err := json.Marshal(glb.GltfDocument)
	if err != nil {
		return nil, err
	}
	binData := glb.BIN

	var jsonPadding uint32 = (4 - uint32(len(jsonData))%4) % 4
	var binPadding uint32 = (4 - uint32(len(binData))%4) % 4

	var magic uint32 = GLB_MAGIC
	var version uint32 = 2
	var length uint32 = GLB_HEADER_LENGTH + 8 + uint32(len(jsonData)) + jsonPadding + 8 + uint32(len(binData)) + binPadding

	output_glb = make([]byte, length, length)

	var offset uint32 = 0

	binary.LittleEndian.PutUint32(output_glb[offset:offset+4], magic)
	offset += 4

	binary.LittleEndian.PutUint32(output_glb[offset:offset+4], version)
	offset += 4

	binary.LittleEndian.PutUint32(output_glb[offset:offset+4], length)
	offset += 4

	var jsonChunkLength uint32 = uint32(len(jsonData)) + jsonPadding
	binary.LittleEndian.PutUint32(output_glb[offset:offset+4], jsonChunkLength)
	offset += 4

	var jsonChunkType uint32 = ChunkTypeJSON
	binary.LittleEndian.PutUint32(output_glb[offset:offset+4], jsonChunkType)
	offset += 4

	copy(output_glb[offset:offset+uint32(len(jsonData))], jsonData)
	offset += uint32(len(jsonData))

	if jsonPadding != 0 {
		pad := [3]byte{' ', ' ', ' '}
		copy(output_glb[offset:offset+jsonPadding], pad[:jsonPadding])
		offset += jsonPadding
	}

	var binaryChunkLength uint32 = uint32(len(binData)) + binPadding
	binary.LittleEndian.PutUint32(output_glb[offset:offset+4], binaryChunkLength)
	offset += 4

	var binaryChunkType uint32 = ChunkTypeBIN
	binary.LittleEndian.PutUint32(output_glb[offset:offset+4], binaryChunkType)
	offset += 4

	copy(output_glb[offset:offset+uint32(len(binData))], binData)
	offset += uint32(len(binData))

	if binPadding != 0 {
		pad := [3]byte{' ', ' ', ' '}
		copy(output_glb[offset:offset+binPadding], pad[:binPadding])
		offset += binPadding
	}

	return output_glb, nil
}
