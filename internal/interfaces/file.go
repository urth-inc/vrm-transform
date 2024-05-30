package interfaces

type File interface {
	Write(p []byte) (n int, err error)
	Close() error
}
