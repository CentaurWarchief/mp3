package id3v2_test

import (
	"testing"

	"github.com/CentaurWarchief/mp3/id3v2"
	"github.com/stretchr/testify/assert"
)

func TestHasFrame(t *testing.T) {
	l := id3v2.ID3v2FrameList{
		id3v2.ID3v2Frame{ID: "TYER"},
		id3v2.ID3v2Frame{ID: "TYER"},
		id3v2.ID3v2Frame{ID: "TYER"},
	}

	assert.True(t, l.HasFrame("TYER"))
	assert.False(t, l.HasFrame("TRCK"))
}

func TestGetFrames(t *testing.T) {
	l := id3v2.ID3v2FrameList{
		id3v2.ID3v2Frame{ID: "TRCK"},
		id3v2.ID3v2Frame{ID: "TRCK"},
		id3v2.ID3v2Frame{ID: "TRCK"},
	}

	assert.Len(t, l.Frames("APIC"), 0)
	assert.Len(t, l.Frames("TRCK"), 3)
}
