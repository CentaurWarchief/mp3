package mp3

import (
	"io"
	"math"
	"os"
)

func round(f float64) int {
	return int(math.Floor(f + .5))
}

func Duration(r io.ReadSeeker) int {
	var duration float64

	r.Seek(0, os.SEEK_SET)
	defer r.Seek(0, os.SEEK_SET)

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

		if block[0] == 0 || block[0] == 32 {
			break
		}

		if block[0] == '\xFF' && ((block[1] & 0xE0) != 0) {
			frame := ParseFrame(block[0:4])

			size := frame.Size()

			if size == 0 {
				return round(duration)
			}

			duration += float64(frame.Samples) / float64(frame.SampleRate)

			r.Seek(int64((size - 10)), os.SEEK_CUR)

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
