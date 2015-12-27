package id3v2

import "errors"

var (
	// ErrInvalidSize is the error returned when the entire size
	// of ID3 header cannot be measured
	ErrInvalidSize = errors.New("invalid size")
	// ErrHeaderNotFound is the error returned when ID3 header
	// mark was not found
	ErrHeaderNotFound = errors.New("header not found")
	// ErrUnrecognizedVersion is the error returned when the major
	// version of ID3v2 is not recognized by this library or even
	// it's a documented version by the ID3v2 standard
	ErrUnrecognizedVersion = errors.New("unrecognized version")
	// ErrUnsupportedVersion is the error returned when there's no
	// frame capturer for the given ID3v2 major version
	ErrUnsupportedVersion = errors.New("unsupported version")
)
