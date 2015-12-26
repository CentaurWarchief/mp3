package frame_test

import (
	"testing"

	"github.com/CentaurWarchief/mp3/id3v2/frame"
	"github.com/stretchr/testify/assert"
)

func TestIsValidFrameName(t *testing.T) {
	assert.True(t, frame.IsValidFrameName([]byte("TIT1")))
	assert.True(t, frame.IsValidFrameName([]byte("TIT2")))
	assert.True(t, frame.IsValidFrameName([]byte("TIT3")))
	assert.True(t, frame.IsValidFrameName([]byte("TPE1")))
	assert.True(t, frame.IsValidFrameName([]byte("TPE2")))
	assert.True(t, frame.IsValidFrameName([]byte("TPE3")))
	assert.True(t, frame.IsValidFrameName([]byte("TPE4")))
	assert.True(t, frame.IsValidFrameName([]byte("TRCK")))
	assert.False(t, frame.IsValidFrameName([]byte("#")))
	assert.False(t, frame.IsValidFrameName([]byte("")))
	assert.False(t, frame.IsValidFrameName([]byte("TPE5")))
	assert.False(t, frame.IsValidFrameName([]byte("TPE6")))
	assert.False(t, frame.IsValidFrameName([]byte("TPE7")))
	assert.False(t, frame.IsValidFrameName([]byte("TPE8")))
	assert.False(t, frame.IsValidFrameName([]byte("TPE9")))
	assert.False(t, frame.IsValidFrameName([]byte("TPE?")))
}
