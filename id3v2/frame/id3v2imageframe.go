package frame

type ID3v2ImageFrame struct {
	Binary      []byte
	MIMEType    string
	Description string
	Type        []byte
	Encoding    []byte
}
