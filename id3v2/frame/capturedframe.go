package frame

import "io"

// CapturedFrame represents a captured frame at the
// specified position with its size
type CapturedFrame struct {
	Frame    string
	Size     uint64
	Position int
}

// ID returns the frame identifier
func (f CapturedFrame) ID() string {
	return f.Frame
}

// AsReadable turns this captured frame into a
// readable frame (ReadableFrame)
func (f CapturedFrame) AsReadable(r io.ReaderAt, u func(body []byte) interface{}) ReadableFrame {
	return NewReadableFrame(f.Frame, func() interface{} {
		body := make([]byte, f.Size)

		if _, err := r.ReadAt(body, int64(f.Position)); err != nil {
			return nil
		}

		return u(body)
	})
}
