package capture_test

import (
	"bytes"
	"testing"

	"github.com/CentaurWarchief/mp3/id3v2/capture"
	"github.com/stretchr/testify/assert"
)

func TestCaptureID3v23FramesEOF(t *testing.T) {
	frames := capture.ID3v23FrameCapturer(bytes.NewReader([]byte{}))

	assert.Len(t, frames, 0)
}

func TestCaptureID3v23Frames(t *testing.T) {
	reader := bytes.NewReader([]byte{
		84, 73, 84, 50, 0, 0, 0, 0, 0, 0,
		84, 80, 69, 49, 0, 0, 0, 0, 0, 0,
		84, 65, 76, 66, 0, 0, 0, 0, 0, 0,
		84, 67, 79, 78, 0, 0, 0, 0, 0, 0,
		84, 82, 67, 75, 0, 0, 0, 0, 0, 0,
		84, 89, 69, 82, 0, 0, 0, 0, 0, 0,
		84, 80, 69, 50, 0, 0, 0, 0, 0, 0,
		84, 73, 84, 53, 0, 0, 0, 0, 0, 0,
		84, 73, 84, 54, 0, 0, 0, 0, 0, 0,
		84, 73, 84, 55, 0, 0, 0, 0, 0, 0,
		84, 73, 84, 56, 0, 0, 0, 0, 0, 0,
		84, 73, 84, 57, 0, 0, 0, 0, 0, 0,
		65, 80, 73, 67, 0, 0, 0, 0, 0, 0,
		65, 80, 73, 67, 0, 0, 0, 0, 0, 0,
	})

	frames := capture.ID3v23FrameCapturer(reader)

	assert.Len(t, frames, 9)

	for i, frame := range []string{
		"TIT2", "TPE1", "TALB",
		"TCON", "TRCK", "TYER",
		"TPE2", "APIC", "APIC",
	} {
		assert.Equal(t, frame, frames[i].Frame)
	}

	for i, position := range []int{20, 30, 40, 50, 60, 70, 80, 140, 150} {
		assert.Equal(t, position, frames[i].Position)
	}
}
