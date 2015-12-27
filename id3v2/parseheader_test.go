package id3v2_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/CentaurWarchief/mp3/id3v2"
	"github.com/stretchr/testify/assert"
)

func TestParseHeaderEOF(t *testing.T) {
	reader := bytes.NewReader([]byte{0, 0, 0})
	reader.Seek(10, os.SEEK_SET)

	header, err := id3v2.ParseHeader(reader)

	assert.Nil(t, header)
	assert.Equal(t, io.EOF, err)
}

func TestParseInvalidHeader(t *testing.T) {
	reader := bytes.NewReader([]byte{
		32, 32, 32, 32, 32,
		32, 32, 32, 32, 32,
	})

	header, err := id3v2.ParseHeader(reader)

	assert.Nil(t, header)
	assert.Equal(t, id3v2.ErrHeaderNotFound, err)
}

func TestParseUnrecognizedMajorVersion(t *testing.T) {
	reader := bytes.NewReader([]byte{73, 68, 51, 5, 0, 0, 0, 6, 96, 117})

	header, err := id3v2.ParseHeader(reader)

	assert.Nil(t, header)
	assert.Equal(t, id3v2.ErrUnrecognizedVersion, err)
}

func TestParseInvalidSynchSafeSize(t *testing.T) {
	reader := bytes.NewReader([]byte{73, 68, 51, 3, 0, 0, 128, 128, 128, 128})

	header, err := id3v2.ParseHeader(reader)

	assert.Nil(t, header)
	assert.Equal(t, id3v2.ErrInvalidSize, err)
}

func TestParseID3v2Header(t *testing.T) {
	for _, headerBytes := range [][]byte{
		[]byte{73, 68, 51, 2, 0, 0, 0, 6, 96, 117},
		[]byte{73, 68, 51, 3, 0, 0, 0, 6, 96, 117},
		[]byte{73, 68, 51, 4, 0, 0, 0, 6, 96, 117},
	} {
		reader := bytes.NewReader(headerBytes)
		header, err := id3v2.ParseHeader(reader)

		assert.Nil(t, err)
		assert.Equal(t, int(headerBytes[3]), header.MajorVersion)
		assert.Equal(t, 0, header.MinorVersion)
		assert.False(t, header.Unsynchronization)
		assert.False(t, header.Extended)
		assert.False(t, header.Experimental)
		assert.False(t, header.Footer)
		assert.Equal(t, uint64(111479), header.Size)
	}
}
