package frame_test

import (
	"testing"

	"github.com/CentaurWarchief/mp3/id3v2/frame"
	"github.com/stretchr/testify/assert"
)

func TestCapturedFrameID(t *testing.T) {
	frame := frame.CapturedFrame{
		Frame: "TIT2",
	}

	assert.Equal(t, "TIT2", frame.ID())
}
