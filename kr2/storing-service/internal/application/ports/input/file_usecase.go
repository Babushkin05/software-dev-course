package input

type FileUseCase interface {
	Upload(filename string, content []byte) (string, error)
	Download(fileID string) (filename string, content []byte, err error)
}
