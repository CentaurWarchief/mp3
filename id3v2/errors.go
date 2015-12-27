package id3v2

import "errors"

var (
	// ErrInvalidSize is the error returned by ParseHeader when
	// it encounters a invalid zero bit
	ErrInvalidSize = errors.New("invalid size")
	// ErrHeaderNotFound is the error returned when ParseHeader
	// cannot find ID3 header mark
	ErrHeaderNotFound = errors.New("header not found")
	// ErrUnrecognizedVersion is the error returned by ParseHeader
	// when the major version of ID3v2 is not recognized by this
	// library or even is a documented version by ID3v2 standard
	ErrUnrecognizedVersion = errors.New("unrecognized version")
	// ErrUnsupportedVersion is the error returned by Parser#Parse
	// when there's no capturer for the given ID3v2 major version
	ErrUnsupportedVersion = errors.New("unsupported version")
)
