package capture_test

import (
	"bytes"
	"testing"

	"github.com/CentaurWarchief/mp3/id3v2/capture"
	"github.com/stretchr/testify/assert"
)

func TestCaptureID3v22FramesEOF(t *testing.T) {
	reader := bytes.NewReader([]byte{})
	frames := capture.ID3v22FrameCapturer(reader)

	assert.Len(t, frames, 0)
}

func TestCaptureID3v22Frames(t *testing.T) {
	reader := bytes.NewReader([]byte{
		84, 84, 50, 0, 0, 0,
		84, 80, 49, 0, 0, 0,
		84, 89, 69, 0, 0, 0,
		84, 73, 57, 0, 0, 0,
		84, 65, 76, 0, 0, 0,
	})

	frames := capture.ID3v22FrameCapturer(reader)

	assert.Len(t, frames, 4)

	assert.Equal(t, frames[0].Frame, "TT2")
	assert.Equal(t, frames[1].Frame, "TP1")
	assert.Equal(t, frames[2].Frame, "TYE")
	assert.Equal(t, frames[3].Frame, "TAL")

	assert.Equal(t, uint64(0), frames[0].Size)
	assert.Equal(t, uint64(0), frames[1].Size)
	assert.Equal(t, uint64(0), frames[2].Size)
	assert.Equal(t, uint64(0), frames[3].Size)

	assert.Equal(t, 16, frames[0].Position)
	assert.Equal(t, 22, frames[1].Position)
	assert.Equal(t, 28, frames[2].Position)
	assert.Equal(t, 40, frames[3].Position)
}
