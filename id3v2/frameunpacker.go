package id3v2

import "github.com/CentaurWarchief/mp3/id3v2/frame"

// FrameUnpacker defines an interface for unpacking ID3v2 frames
type FrameUnpacker interface {
	Unpack(frame frame.Frame, body []byte) interface{}
	CanUnpack(frame frame.Frame) bool
}
