package id3v2

type ID3v2FrameSizePosition struct {
	Tag      string
	Size     uint64
	Position int
}

func (f ID3v2FrameSizePosition) ID() string {
	return f.Tag
}
