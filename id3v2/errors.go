package id3v2

import "errors"

var (
	ErrInvalidSize              = errors.New("invalid size")
	ErrID3HeaderNotFound        = errors.New("ID3 header not found")
	ErrUnrecognizedID3v2Version = errors.New("unrecognized ID3v2 version")
)
