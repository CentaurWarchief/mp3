package frame_test

import (
	"testing"

	"github.com/CentaurWarchief/mp3/id3v2/frame"
	"github.com/stretchr/testify/assert"
)

func TestID3v2CapturedFrameID(t *testing.T) {
	frame := frame.ID3v2CapturedFrame{
		Frame: "TIT2",
	}

	assert.Equal(t, "TIT2", frame.ID())
}
