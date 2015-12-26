package id3v2

import "io"

// Given a reader captures a list of frames until EOF
type ID3v2FrameCapturer func(r io.Reader) ID3v2FrameList
