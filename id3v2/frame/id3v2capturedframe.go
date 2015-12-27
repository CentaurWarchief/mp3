package frame

import "io"

// ID3v2CapturedFrame represents a captured frame at the
// specified position with its size
type ID3v2CapturedFrame struct {
	Frame    string
	Size     uint64
	Position int
}

// ID returns the frame identifier
func (f ID3v2CapturedFrame) ID() string {
	return f.Frame
}

// AsReadable turns the captured frame in a
// readable frame (ID3v2ReadableFrame)
func (f ID3v2CapturedFrame) AsReadable(
	r io.ReaderAt,
	u func(body []byte) interface{},
) ID3v2ReadableFrame {
	return NewID3v2ReadableFrame(f.Frame, func() interface{} {
		body := make([]byte, f.Size)

		if _, err := r.ReadAt(body, int64(f.Position)); err != nil {
			return nil
		}

		return u(body)
	})
}
