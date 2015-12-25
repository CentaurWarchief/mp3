package id3v2_test

import (
	"testing"

	"github.com/CentaurWarchief/mp3/id3v2"
	"github.com/stretchr/testify/assert"
)

func TestFrameSizePositionID(t *testing.T) {
	f := id3v2.ID3v2FrameSizePosition{
		Tag: "TRCK",
	}

	assert.Equal(t, "TRCK", f.ID())
}
