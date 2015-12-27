package frame

// NewID3v2ReadableFrame creates a new ID3v2ReadableFrame
// with the specified parameters
func NewID3v2ReadableFrame(frame string, read func() interface{}) ID3v2ReadableFrame {
	return ID3v2ReadableFrame{
		Frame: frame,
		Read:  read,
	}
}

// ID3v2ReadableFrame represents a frame that can be read
type ID3v2ReadableFrame struct {
	Frame string
	Read  func() interface{}
}

// ID returns the frame identifier
func (f ID3v2ReadableFrame) ID() string {
	return f.Frame
}
