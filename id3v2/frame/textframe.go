package frame

// NewTextFrame creates a new TextFrame
func NewTextFrame(frame, text string) TextFrame {
	return TextFrame{
		Frame: frame,
		Text:  text,
	}
}

// TextFrame represents a simple ID3v2 text frame
type TextFrame struct {
	Frame string
	Text  string
}

// ID returns the frame identifier
func (f TextFrame) ID() string {
	return f.Frame
}
