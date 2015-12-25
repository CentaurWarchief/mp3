package id3v2_test

import (
	"testing"

	"github.com/CentaurWarchief/mp3/id3v2"
	"github.com/stretchr/testify/assert"
)

type SimpleFrame struct {
	Tag string
}

func (f SimpleFrame) ID() string {
	return f.Tag
}

func TestSingleFrame(t *testing.T) {
	l := id3v2.ID3v2FrameList{
		SimpleFrame{Tag: "TRCK"},
		SimpleFrame{Tag: "TYER"},
		SimpleFrame{Tag: "TRCK"},
	}

	f, err := l.SingleFrame("TRCK")

	assert.NotNil(t, f)
	assert.Nil(t, err)

	f, err = l.SingleFrame("TIT2")

	assert.Nil(t, f)
	assert.Equal(t, id3v2.ErrFrameNotFound, err)
}

func TestHasFrame(t *testing.T) {
	l := id3v2.ID3v2FrameList{
		SimpleFrame{Tag: "TYER"},
		SimpleFrame{Tag: "TYER"},
		SimpleFrame{Tag: "TYER"},
	}

	assert.True(t, l.HasFrame("TYER"))
	assert.False(t, l.HasFrame("TRCK"))
}

func TestGetFrames(t *testing.T) {
	l := id3v2.ID3v2FrameList{
		SimpleFrame{Tag: "TRCK"},
		SimpleFrame{Tag: "TRCK"},
		SimpleFrame{Tag: "TRCK"},
	}

	assert.Len(t, l.Frames("APIC"), 0)
	assert.Len(t, l.Frames("TRCK"), 3)
}
