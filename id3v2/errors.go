package id3v2

import "errors"

var (
	ErrInvalidSize         = errors.New("invalid size")
	ErrID3HeaderNotFound   = errors.New("ID3 header not found")
	ErrUnrecognizedVersion = errors.New("unrecognized version")
	ErrUnsupportedVersion  = errors.New("unsupported version")
	ErrFrameNotFound       = errors.New("frame not found")
)
