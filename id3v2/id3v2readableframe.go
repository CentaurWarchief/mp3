package id3v2

type ID3v2ReadableFrame struct {
	Tag  string
	Read func() interface{}
}

func (f ID3v2ReadableFrame) ID() string {
	return f.Tag
}
