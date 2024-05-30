// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/interfaces/glb.go
//
// Generated by this command:
//
//	mockgen -package=mock_glb -source=./internal/interfaces/glb.go
//

// Package mock_glb is a generated GoMock package.
package mock_glb

import (
	reflect "reflect"

	interfaces "github.com/urth-inc/vrm-transform/internal/interfaces"
	gomock "go.uber.org/mock/gomock"
)

// MockConvertToKtx2ImageDependenciesInterface is a mock of ConvertToKtx2ImageDependenciesInterface interface.
type MockConvertToKtx2ImageDependenciesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConvertToKtx2ImageDependenciesInterfaceMockRecorder
}

// MockConvertToKtx2ImageDependenciesInterfaceMockRecorder is the mock recorder for MockConvertToKtx2ImageDependenciesInterface.
type MockConvertToKtx2ImageDependenciesInterfaceMockRecorder struct {
	mock *MockConvertToKtx2ImageDependenciesInterface
}

// NewMockConvertToKtx2ImageDependenciesInterface creates a new mock instance.
func NewMockConvertToKtx2ImageDependenciesInterface(ctrl *gomock.Controller) *MockConvertToKtx2ImageDependenciesInterface {
	mock := &MockConvertToKtx2ImageDependenciesInterface{ctrl: ctrl}
	mock.recorder = &MockConvertToKtx2ImageDependenciesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConvertToKtx2ImageDependenciesInterface) EXPECT() *MockConvertToKtx2ImageDependenciesInterfaceMockRecorder {
	return m.recorder
}

// CommandExecutor mocks base method.
func (m *MockConvertToKtx2ImageDependenciesInterface) CommandExecutor(name string, args ...string) error {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CommandExecutor", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommandExecutor indicates an expected call of CommandExecutor.
func (mr *MockConvertToKtx2ImageDependenciesInterfaceMockRecorder) CommandExecutor(name any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandExecutor", reflect.TypeOf((*MockConvertToKtx2ImageDependenciesInterface)(nil).CommandExecutor), varargs...)
}

// ContentTypeDetector mocks base method.
func (m *MockConvertToKtx2ImageDependenciesInterface) ContentTypeDetector(data []byte) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContentTypeDetector", data)
	ret0, _ := ret[0].(string)
	return ret0
}

// ContentTypeDetector indicates an expected call of ContentTypeDetector.
func (mr *MockConvertToKtx2ImageDependenciesInterfaceMockRecorder) ContentTypeDetector(data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContentTypeDetector", reflect.TypeOf((*MockConvertToKtx2ImageDependenciesInterface)(nil).ContentTypeDetector), data)
}

// FileCreator mocks base method.
func (m *MockConvertToKtx2ImageDependenciesInterface) FileCreator(filePath string) (interfaces.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileCreator", filePath)
	ret0, _ := ret[0].(interfaces.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FileCreator indicates an expected call of FileCreator.
func (mr *MockConvertToKtx2ImageDependenciesInterfaceMockRecorder) FileCreator(filePath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileCreator", reflect.TypeOf((*MockConvertToKtx2ImageDependenciesInterface)(nil).FileCreator), filePath)
}

// FileReader mocks base method.
func (m *MockConvertToKtx2ImageDependenciesInterface) FileReader(filePath string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileReader", filePath)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FileReader indicates an expected call of FileReader.
func (mr *MockConvertToKtx2ImageDependenciesInterfaceMockRecorder) FileReader(filePath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileReader", reflect.TypeOf((*MockConvertToKtx2ImageDependenciesInterface)(nil).FileReader), filePath)
}

// FileRemover mocks base method.
func (m *MockConvertToKtx2ImageDependenciesInterface) FileRemover(filePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileRemover", filePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// FileRemover indicates an expected call of FileRemover.
func (mr *MockConvertToKtx2ImageDependenciesInterfaceMockRecorder) FileRemover(filePath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileRemover", reflect.TypeOf((*MockConvertToKtx2ImageDependenciesInterface)(nil).FileRemover), filePath)
}

// ImageSizer mocks base method.
func (m *MockConvertToKtx2ImageDependenciesInterface) ImageSizer(data []byte) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageSizer", data)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImageSizer indicates an expected call of ImageSizer.
func (mr *MockConvertToKtx2ImageDependenciesInterfaceMockRecorder) ImageSizer(data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageSizer", reflect.TypeOf((*MockConvertToKtx2ImageDependenciesInterface)(nil).ImageSizer), data)
}

// ParamsGenerator mocks base method.
func (m *MockConvertToKtx2ImageDependenciesInterface) ParamsGenerator(mode string, width, height int, inputPath, outputPath string, isSRGB bool, etc1sQuality, uastcQuality, zstdLevel int) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParamsGenerator", mode, width, height, inputPath, outputPath, isSRGB, etc1sQuality, uastcQuality, zstdLevel)
	ret0, _ := ret[0].([]string)
	return ret0
}

// ParamsGenerator indicates an expected call of ParamsGenerator.
func (mr *MockConvertToKtx2ImageDependenciesInterfaceMockRecorder) ParamsGenerator(mode, width, height, inputPath, outputPath, isSRGB, etc1sQuality, uastcQuality, zstdLevel any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParamsGenerator", reflect.TypeOf((*MockConvertToKtx2ImageDependenciesInterface)(nil).ParamsGenerator), mode, width, height, inputPath, outputPath, isSRGB, etc1sQuality, uastcQuality, zstdLevel)
}

// UUIDGenerator mocks base method.
func (m *MockConvertToKtx2ImageDependenciesInterface) UUIDGenerator() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UUIDGenerator")
	ret0, _ := ret[0].(string)
	return ret0
}

// UUIDGenerator indicates an expected call of UUIDGenerator.
func (mr *MockConvertToKtx2ImageDependenciesInterfaceMockRecorder) UUIDGenerator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UUIDGenerator", reflect.TypeOf((*MockConvertToKtx2ImageDependenciesInterface)(nil).UUIDGenerator))
}

// MockConvertToKtx2TextureDependenciesInterface is a mock of ConvertToKtx2TextureDependenciesInterface interface.
type MockConvertToKtx2TextureDependenciesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConvertToKtx2TextureDependenciesInterfaceMockRecorder
}

// MockConvertToKtx2TextureDependenciesInterfaceMockRecorder is the mock recorder for MockConvertToKtx2TextureDependenciesInterface.
type MockConvertToKtx2TextureDependenciesInterfaceMockRecorder struct {
	mock *MockConvertToKtx2TextureDependenciesInterface
}

// NewMockConvertToKtx2TextureDependenciesInterface creates a new mock instance.
func NewMockConvertToKtx2TextureDependenciesInterface(ctrl *gomock.Controller) *MockConvertToKtx2TextureDependenciesInterface {
	mock := &MockConvertToKtx2TextureDependenciesInterface{ctrl: ctrl}
	mock.recorder = &MockConvertToKtx2TextureDependenciesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConvertToKtx2TextureDependenciesInterface) EXPECT() *MockConvertToKtx2TextureDependenciesInterfaceMockRecorder {
	return m.recorder
}

// ConvertToKtx2Image mocks base method.
func (m *MockConvertToKtx2TextureDependenciesInterface) ConvertToKtx2Image(deps interfaces.ConvertToKtx2ImageDependenciesInterface, ktx2Mode string, buf []byte, isSRGB bool, etc1sQuality, uastcQuality, zstdLevel int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertToKtx2Image", deps, ktx2Mode, buf, isSRGB, etc1sQuality, uastcQuality, zstdLevel)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertToKtx2Image indicates an expected call of ConvertToKtx2Image.
func (mr *MockConvertToKtx2TextureDependenciesInterfaceMockRecorder) ConvertToKtx2Image(deps, ktx2Mode, buf, isSRGB, etc1sQuality, uastcQuality, zstdLevel any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertToKtx2Image", reflect.TypeOf((*MockConvertToKtx2TextureDependenciesInterface)(nil).ConvertToKtx2Image), deps, ktx2Mode, buf, isSRGB, etc1sQuality, uastcQuality, zstdLevel)
}
