package id3v2_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/CentaurWarchief/mp3/id3v2"
	"github.com/stretchr/testify/assert"
)

func TestParseID3HeaderEOF(t *testing.T) {
	reader := bytes.NewReader([]byte{0, 0, 0})
	reader.Seek(10, os.SEEK_SET)

	header, err := id3v2.ParseID3v2Header(reader)

	assert.Nil(t, header)
	assert.Equal(t, io.EOF, err)
}

func TestParseNonID3Header(t *testing.T) {
	reader := bytes.NewReader([]byte{
		32, 32, 32, 32, 32,
		32, 32, 32, 32, 32,
	})

	header, err := id3v2.ParseID3v2Header(reader)

	assert.Nil(t, header)
	assert.Equal(t, id3v2.ErrID3HeaderNotFound, err)
}

func TestParseUnrecognizedID3MajorVersion(t *testing.T) {
	reader := bytes.NewReader([]byte{73, 68, 51, 5, 0, 0, 0, 6, 96, 117})

	header, err := id3v2.ParseID3v2Header(reader)

	assert.Nil(t, header)
	assert.Equal(t, id3v2.ErrUnrecognizedVersion, err)
}

func TestParseID3v2Header(t *testing.T) {
	reader := bytes.NewReader([]byte{
		73, 68, 51, 3, 0,
		0, 0, 6, 96, 117,
	})

	header, err := id3v2.ParseID3v2Header(reader)

	assert.Nil(t, err)
	assert.Equal(t, 3, header.MajorVersion)
	assert.Equal(t, 0, header.MinorVersion)
	assert.False(t, header.Unsynchronization)
	assert.False(t, header.Extended)
	assert.False(t, header.Experimental)
	assert.False(t, header.Footer)
	assert.Equal(t, uint64(111479), header.Size)
}
