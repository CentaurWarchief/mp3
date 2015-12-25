package mp3

import (
	"fmt"
	"io"
	"math"
	"os"
)

var (
	versions = map[byte]string{
		0x0: "2.5",
		0x1: "x",
		0x2: "2",
		0x3: "1",
	}

	layers = map[byte]string{
		0x0: "x",
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

var ModeList = []string{
	"Stereo",
	"Joint Stereo",
	"Dual Channel",
	"Mono",
}

type Frame struct {
	Version       string
	Layer         string
	Frequency     int
	Samples       int
	Bitrate       int
	SampleRate    int
	Size          int
	Mode          int
	ModeExtension int
	ModeText      string
}

func round(f float64) int {
	return int(math.Floor(f + .5))
}

func frame(b []byte) *Frame {
	frame := &Frame{}

	version := versions[((b[1] & 0x18) >> 3)]
	layer := layers[((b[1] & 0x06) >> 1)]

	frame.Version = version
	frame.Layer = layer
	frame.Samples = samples[version][layer]

	key := fmt.Sprintf("%s%s", version, layer)

	if version == "2.5" {
		key = fmt.Sprintf("2%s", layer)
		frame.Samples = samples["2"][layer]
	}

	index := ((b[2] & 0xF0) >> 4)
	bitrate := 0

	if bitrates[key][index] != 0 {
		bitrate = bitrates[key][index]
	}

	rate := rates[version][((b[2] & 0x0C) >> 2)]

	mode := int((b[3] & 0x0C) >> 6)

	frame.Frequency = int((b[2] >> 2) & 0x3)
	frame.Mode = mode
	frame.ModeExtension = int((b[3] & 0x30) >> 4)
	frame.ModeText = ModeList[mode]
	frame.SampleRate = rate
	frame.Bitrate = bitrate
	frame.Size = computeFrameSize(layer, bitrate, rate, int(((b[2] & 0x02) >> 1)))

	return frame
}

func computeFrameSize(layer string, bitrate, sample, padding int) int {
	if layer == "1" {
		return ((12 * bitrate * 1000 / sample) + padding) * 4
	}

	return ((144 * bitrate * 1000 / sample) + padding)
}

func Duration(r io.ReadSeeker) int {
	var duration float64

	r.Seek(0, os.SEEK_SET)

	block := make([]byte, 100)

	r.Read(block)

	if string(block[0:3]) == "ID3" {
		flags := block[5]
		footer := (flags & 0x10) != 0

		z0 := block[6]
		z1 := block[7]
		z2 := block[8]
		z3 := block[9]

		size := 0

		if (z0&0x80) == 0 && (z1&0x80) == 0 && (z2&0x80) == 0 && (z3&0x80) == 0 {
			tag := (int((z0 & 0x7F)) * 2097152)
			tag = tag + (int((z1 & 0x7F)) * 16384)
			tag = tag + (int((z2 & 0x7F)) * 128)
			tag = tag + (int(z3 & 0x7F))

			size = 10 + tag

			if footer {
				size = 10 + size
			}
		}

		if _, err := r.Seek(int64(size), os.SEEK_SET); err != nil {
			return 0
		}
	}

	for {
		block := make([]byte, 10)

		if n, err := r.Read(block); err != nil || n == 0 {
			break
		}

		if block[0] == '\xFF' && ((block[1] & 0xE0) != 0) {
			frame := frame(block[0:4])

			if frame.Size == 0 {
				return round(duration)
			}

			duration += float64(frame.Samples) / float64(frame.SampleRate)

			r.Seek(int64((frame.Size - 10)), os.SEEK_CUR)

			continue
		}

		if string(block[0:3]) == "TAG" {
			if _, err := r.Seek(int64((128 - 10)), os.SEEK_CUR); err != nil {
				return 0
			}

			continue
		}

		if _, err := r.Seek(-int64(9), os.SEEK_CUR); err != nil {
			break
		}
	}

	return round(duration)
}
