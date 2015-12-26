package frame_test

import (
	"testing"

	"github.com/CentaurWarchief/mp3/id3v2/frame"
	"github.com/stretchr/testify/assert"
)

func TestNewTextFrame(t *testing.T) {
	frame := frame.NewTextFrame("TIT2", "In Any Tongue")

	assert.NotNil(t, frame)
	assert.Equal(t, "TIT2", frame.ID())
	assert.Equal(t, "TIT2", frame.Frame)
	assert.Equal(t, "In Any Tongue", frame.Text)
}
