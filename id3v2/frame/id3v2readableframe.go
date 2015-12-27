package frame

func NewID3v2ReadableFrame(frame string, read func() interface{}) ID3v2ReadableFrame {
	return ID3v2ReadableFrame{
		Frame: frame,
		Read:  read,
	}
}

type ID3v2ReadableFrame struct {
	Frame string
	Read  func() interface{}
}

func (f ID3v2ReadableFrame) ID() string {
	return f.Frame
}
