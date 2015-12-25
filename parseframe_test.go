package mp3_test

import (
	"testing"

	"github.com/CentaurWarchief/mp3"
	"github.com/stretchr/testify/assert"
)

func TestParseFrame(t *testing.T) {
	frames := []mp3.Frame{
		mp3.Frame{
			Version:       "1",
			Layer:         "3",
			Frequency:     1,
			Samples:       1152,
			Bitrate:       320,
			SampleRate:    48000,
			Mode:          0,
			ModeExtension: 2,
			Private:       0,
			Copyright:     0,
			Emphasis:      0,
			Padding:       0,
		},
		mp3.Frame{
			Version:       "1",
			Layer:         "3",
			Frequency:     0,
			Samples:       1152,
			Bitrate:       320,
			SampleRate:    44100,
			Mode:          0,
			ModeExtension: 0,
			Private:       0,
			Copyright:     0,
			Emphasis:      0,
			Padding:       1,
		},
		mp3.Frame{
			Version:       "1",
			Layer:         "3",
			Frequency:     0,
			Samples:       1152,
			Bitrate:       320,
			SampleRate:    44100,
			Mode:          0,
			ModeExtension: 0,
			Private:       0,
			Copyright:     0,
			Emphasis:      0,
			Padding:       0,
		},
	}

	for i, block := range [][]byte{
		[]byte{255, 251, 228, 100},
		[]byte{255, 251, 226, 0},
		[]byte{255, 251, 224, 0},
	} {
		frame := mp3.ParseFrame(block)

		assert.Equal(t, frames[i].Version, frame.Version)
		assert.Equal(t, frames[i].Layer, frame.Layer)
		assert.Equal(t, frames[i].Frequency, frame.Frequency)
		assert.Equal(t, frames[i].Samples, frame.Samples)
		assert.Equal(t, frames[i].Bitrate, frame.Bitrate)
		assert.Equal(t, frames[i].SampleRate, frame.SampleRate)
		assert.Equal(t, frames[i].Mode, frame.Mode)
		assert.Equal(t, frames[i].ModeExtension, frame.ModeExtension)
		assert.Equal(t, frames[i].Private, frame.Private)
		assert.Equal(t, frames[i].Copyright, frame.Copyright)
		assert.Equal(t, frames[i].Emphasis, frame.Emphasis)
		assert.Equal(t, frames[i].Padding, frame.Padding)
	}
}
