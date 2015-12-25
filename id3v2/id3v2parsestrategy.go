package id3v2

import "io"

type ID3v2ParseStrategy func(r io.Reader) ID3v2FrameList
