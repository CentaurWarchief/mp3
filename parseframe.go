package mp3

import "fmt"

var (
	versions = map[byte]string{
		0x0: "2.5",
		0x1: "",
		0x2: "2",
		0x3: "1",
	}

	layers = map[byte]string{
		0x0: "",
		0x1: "3",
		0x2: "2",
		0x3: "1",
	}

	bitrates = map[string][]int{
		// MPEG 1.0
		"11": []int{0, 32, 64, 96, 128, 160, 192, 224, 256, 288, 320, 352, 384, 416, 448},
		"12": []int{0, 32, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320, 384},
		"13": []int{0, 32, 40, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320},

		// MPEG 2.0
		"21": []int{0, 32, 48, 56, 64, 80, 96, 112, 128, 144, 160, 176, 192, 224, 256},
		"22": []int{0, 8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 144, 160},
		"23": []int{0, 8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 144, 160},
	}

	rates = map[string][]int{
		"1":   []int{44100, 48000, 32000, 50000},
		"2":   []int{22050, 24000, 16000, 50000},
		"2.5": []int{11025, 12000, 8000, 50000},
	}

	samples = map[string]map[string]int{
		"1": map[string]int{
			"1": 384,
			"2": 1152,
			"3": 1152,
		},
		"2": map[string]int{
			"1": 384,
			"2": 1152,
			"3": 576,
		},
	}
)

// ParseFrame parses a MPEG audio frame header
// http://www.codeproject.com/Articles/8295/MPEG-Audio-Frame-Header
func ParseFrame(b []byte) *Frame {
	frame := &Frame{}

	version := versions[((b[1] & 0x18) >> 3)]
	layer := layers[((b[1] & 0x06) >> 1)]

	key := fmt.Sprintf("%s%s", version, layer)

	frame.Samples = samples[version][layer]

	if version == "2.5" {
		key = fmt.Sprintf("2%s", layer)
		frame.Samples = samples["2"][layer]
	}

	index := ((b[2] & 0xF0) >> 4)
	bitrate := 0

	if bitrates[key][index] != 0 {
		bitrate = bitrates[key][index]
	}

	frame.Version = version
	frame.Layer = layer
	frame.Frequency = int((b[2] >> 2) & 0x3)
	frame.Bitrate = bitrate
	frame.SampleRate = rates[version][((b[2] & 0x0C) >> 2)]
	frame.Mode = int((b[3] & 0x0C) >> 6)
	frame.ModeExtension = int((b[3] & 0x30) >> 4)
	frame.Private = int((b[2] & 0x01))
	frame.Copyright = int((b[2] & 0x08) >> 3)
	frame.Emphasis = int((b[3] & 0x03))
	frame.Padding = int(((b[2] & 0x02) >> 1))

	return frame
}
