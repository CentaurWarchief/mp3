package id3v2_test

import (
	"io"
	"reflect"
	"testing"

	"github.com/CentaurWarchief/mp3/id3v2"
	"github.com/CentaurWarchief/mp3/id3v2/frame"
	"github.com/stretchr/testify/assert"
)

func countOfRegisteredCaptures(
	parser *id3v2.ID3v2Parser,
) int {
	return reflect.Indirect(reflect.ValueOf(*parser)).FieldByName("byVersion").Len()
}

func TestNewID3v2Parser(t *testing.T) {
	assert.Equal(t, 3, countOfRegisteredCaptures(id3v2.NewID3v2Parser()))
}

func TestAddVersionedFrameCapturer(t *testing.T) {
	parser := id3v2.NewEmptyID3v2Parser()

	assert.Equal(t, 0, countOfRegisteredCaptures(parser))

	parser.AddVersionedFrameCapturer(3, func(r io.Reader) []frame.ID3v2CapturedFrame {
		return nil
	})

	assert.Equal(t, 1, countOfRegisteredCaptures(parser))
}
