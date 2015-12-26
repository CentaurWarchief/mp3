package id3v2

import (
	"io"

	"github.com/CentaurWarchief/mp3/id3v2/frame"
)

// Given a reader captures a list of frames until EOF
type ID3v2FrameCapturer func(r io.Reader) []frame.ID3v2CapturedFrame
