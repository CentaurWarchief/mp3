package frame

type ID3v2CapturedFrame struct {
	Frame    string
	Size     uint64
	Position int
}

func (f ID3v2CapturedFrame) ID() string {
	return f.Frame
}
