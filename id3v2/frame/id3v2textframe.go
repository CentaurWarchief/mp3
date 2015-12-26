package frame

func NewTextFrame(frame, text string) ID3v2TextFrame {
	return ID3v2TextFrame{
		Frame: frame,
		Text:  text,
	}
}

type ID3v2TextFrame struct {
	Frame string
	Text  string
}

func (f ID3v2TextFrame) ID() string {
	return f.Frame
}
