package id3v2_test

import (
	"testing"

	. "github.com/CentaurWarchief/mp3/id3v2"
	"github.com/stretchr/testify/assert"
)

func TestHasTag(t *testing.T) {
	v23 := ID3v2TagList{
		"TRCK": "TRCK",
		"APIC": "APIC",
	}

	assert.Len(t, v23, 2)
	assert.True(t, v23.HasTag("TRCK"))
	assert.True(t, v23.HasTag("APIC"))
	assert.False(t, v23.HasTag("TAG"))
	assert.False(t, v23.HasTag("NOPE"))
}
