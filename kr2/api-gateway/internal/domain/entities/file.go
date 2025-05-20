package entities

type File struct {
	ID       string
	Name     string
	Content  []byte
	MimeType string
}
