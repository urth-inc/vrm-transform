package interfaces

type ConvertToKtx2ImageDependenciesInterface interface {
	UUIDGenerator() string
	ContentTypeDetector(data []byte) string
	ImageSizer(data []byte) (int, int, error)
	CommandExecutor(name string, args ...string) error
	ParamsGenerator(mode string, width, height int, inputPath, outputPath string, isSRGB bool, etc1sQuality, uastcQuality, zstdLevel int) []string
	FileReader(filePath string) ([]byte, error)
	FileCreator(filePath string) (File, error)
	FileRemover(filePath string) error
}

type ConvertToKtx2TextureDependenciesInterface interface {
	ConvertToKtx2Image(deps ConvertToKtx2ImageDependenciesInterface, ktx2Mode string, buf []byte, isSRGB bool, etc1sQuality int, uastcQuality int, zstdLevel int) ([]byte, error)
}
