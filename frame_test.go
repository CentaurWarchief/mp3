package mp3_test

import (
	"testing"

	"github.com/CentaurWarchief/mp3"
	"github.com/stretchr/testify/assert"
)

func TestFrameSize(t *testing.T) {
	for layer, bspe := range map[string][]int{
		"1": []int{128, 48000, 0, 128},
		"2": []int{320, 48000, 0, 960},
		"3": []int{320, 48000, 0, 960},
	} {
		frame := mp3.Frame{
			Layer:      layer,
			Bitrate:    bspe[0],
			SampleRate: bspe[1],
			Padding:    bspe[2],
		}

		assert.Equal(t, bspe[3], frame.Size())
	}
}

func TestModeText(t *testing.T) {
	for mode, text := range map[int]string{
		0: "Stereo",
		1: "Joint Stereo",
		2: "Dual Channel",
		3: "Mono",
		4: "",
	} {
		frame := mp3.Frame{
			Mode: mode,
		}

		assert.Equal(t, text, frame.ModeText())
	}
}
