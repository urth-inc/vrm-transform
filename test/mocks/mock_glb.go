// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/glb/texture.go
//
// Generated by this command:
//
//	mockgen -source=./pkg/glb/texture.go
//

// Package mock_glb is a generated GoMock package.
package mock_glb

import (
	os "os"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockKtx2ConversionDependenciesInterface is a mock of Ktx2ConversionDependenciesInterface interface.
type MockKtx2ConversionDependenciesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockKtx2ConversionDependenciesInterfaceMockRecorder
}

// MockKtx2ConversionDependenciesInterfaceMockRecorder is the mock recorder for MockKtx2ConversionDependenciesInterface.
type MockKtx2ConversionDependenciesInterfaceMockRecorder struct {
	mock *MockKtx2ConversionDependenciesInterface
}

// NewMockKtx2ConversionDependenciesInterface creates a new mock instance.
func NewMockKtx2ConversionDependenciesInterface(ctrl *gomock.Controller) *MockKtx2ConversionDependenciesInterface {
	mock := &MockKtx2ConversionDependenciesInterface{ctrl: ctrl}
	mock.recorder = &MockKtx2ConversionDependenciesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKtx2ConversionDependenciesInterface) EXPECT() *MockKtx2ConversionDependenciesInterfaceMockRecorder {
	return m.recorder
}

// CommandExecutor mocks base method.
func (m *MockKtx2ConversionDependenciesInterface) CommandExecutor(name string, args ...string) error {
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
func (mr *MockKtx2ConversionDependenciesInterfaceMockRecorder) CommandExecutor(name any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandExecutor", reflect.TypeOf((*MockKtx2ConversionDependenciesInterface)(nil).CommandExecutor), varargs...)
}

// ContentTypeDetector mocks base method.
func (m *MockKtx2ConversionDependenciesInterface) ContentTypeDetector(data []byte) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContentTypeDetector", data)
	ret0, _ := ret[0].(string)
	return ret0
}

// ContentTypeDetector indicates an expected call of ContentTypeDetector.
func (mr *MockKtx2ConversionDependenciesInterfaceMockRecorder) ContentTypeDetector(data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContentTypeDetector", reflect.TypeOf((*MockKtx2ConversionDependenciesInterface)(nil).ContentTypeDetector), data)
}

// FileCreator mocks base method.
func (m *MockKtx2ConversionDependenciesInterface) FileCreator(filePath string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileCreator", filePath)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FileCreator indicates an expected call of FileCreator.
func (mr *MockKtx2ConversionDependenciesInterfaceMockRecorder) FileCreator(filePath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileCreator", reflect.TypeOf((*MockKtx2ConversionDependenciesInterface)(nil).FileCreator), filePath)
}

// FileReader mocks base method.
func (m *MockKtx2ConversionDependenciesInterface) FileReader(filePath string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileReader", filePath)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FileReader indicates an expected call of FileReader.
func (mr *MockKtx2ConversionDependenciesInterfaceMockRecorder) FileReader(filePath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileReader", reflect.TypeOf((*MockKtx2ConversionDependenciesInterface)(nil).FileReader), filePath)
}

// FileRemover mocks base method.
func (m *MockKtx2ConversionDependenciesInterface) FileRemover(filePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileRemover", filePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// FileRemover indicates an expected call of FileRemover.
func (mr *MockKtx2ConversionDependenciesInterfaceMockRecorder) FileRemover(filePath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileRemover", reflect.TypeOf((*MockKtx2ConversionDependenciesInterface)(nil).FileRemover), filePath)
}

// ImageSizer mocks base method.
func (m *MockKtx2ConversionDependenciesInterface) ImageSizer(data []byte) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageSizer", data)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImageSizer indicates an expected call of ImageSizer.
func (mr *MockKtx2ConversionDependenciesInterfaceMockRecorder) ImageSizer(data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageSizer", reflect.TypeOf((*MockKtx2ConversionDependenciesInterface)(nil).ImageSizer), data)
}

// ParamsGenerator mocks base method.
func (m *MockKtx2ConversionDependenciesInterface) ParamsGenerator(mode string, width, height int, inputPath, outputPath string, isSRGB bool, etc1sQuality, uastcQuality, zstdLevel int) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParamsGenerator", mode, width, height, inputPath, outputPath, isSRGB, etc1sQuality, uastcQuality, zstdLevel)
	ret0, _ := ret[0].([]string)
	return ret0
}

// ParamsGenerator indicates an expected call of ParamsGenerator.
func (mr *MockKtx2ConversionDependenciesInterfaceMockRecorder) ParamsGenerator(mode, width, height, inputPath, outputPath, isSRGB, etc1sQuality, uastcQuality, zstdLevel any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParamsGenerator", reflect.TypeOf((*MockKtx2ConversionDependenciesInterface)(nil).ParamsGenerator), mode, width, height, inputPath, outputPath, isSRGB, etc1sQuality, uastcQuality, zstdLevel)
}

// UUIDGenerator mocks base method.
func (m *MockKtx2ConversionDependenciesInterface) UUIDGenerator() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UUIDGenerator")
	ret0, _ := ret[0].(string)
	return ret0
}

// UUIDGenerator indicates an expected call of UUIDGenerator.
func (mr *MockKtx2ConversionDependenciesInterfaceMockRecorder) UUIDGenerator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UUIDGenerator", reflect.TypeOf((*MockKtx2ConversionDependenciesInterface)(nil).UUIDGenerator))
}
