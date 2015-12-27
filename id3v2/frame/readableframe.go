package frame

// NewReadableFrame creates a new ReadableFrame with
// the specified parameters
func NewReadableFrame(frame string, read func() interface{}) ReadableFrame {
	return ReadableFrame{
		Frame: frame,
		Read:  read,
	}
}

// ReadableFrame represents a frame that can be read
type ReadableFrame struct {
	Frame string
	Read  func() interface{}
}

// ID returns the frame identifier
func (f ReadableFrame) ID() string {
	return f.Frame
}
