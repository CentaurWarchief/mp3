package id3v2

import (
	"io"

	"github.com/CentaurWarchief/mp3/id3v2/frame"
)

// FrameCapturer captures all ID3v2 frames from the given reader
// until it reach its EOF
type FrameCapturer func(r io.Reader) []frame.CapturedFrame
