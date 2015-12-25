package id3v2

import "io"

type ID3v2FrameCapturer func(r io.Reader) ID3v2FrameList
