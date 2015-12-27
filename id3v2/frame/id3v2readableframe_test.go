package frame_test

import (
	"testing"

	"github.com/CentaurWarchief/mp3/id3v2/frame"
	"github.com/stretchr/testify/assert"
)

func TestNewID3v2ReadableFrame(t *testing.T) {
	frame := frame.NewID3v2ReadableFrame("TRCK", func() interface{} {
		return nil
	})

	assert.NotNil(t, frame)
	assert.NotNil(t, frame.Read)
	assert.Equal(t, "TRCK", frame.ID())
	assert.Equal(t, "TRCK", frame.Frame)
}
