package capture

import (
	"io"
	"io/ioutil"

	"github.com/CentaurWarchief/mp3/id3v2/frame"
)

// ID3v23FrameCapturer captures all v2.3 frames from the given
// reader until reach its EOF
func ID3v23FrameCapturer(r io.Reader) (frames []frame.ID3v2CapturedFrame) {
	position := 10

	for {
		block := make([]byte, 10)

		if _, err := r.Read(block); err != nil {
			break
		}

		size := (uint64(block[4]) << 0x18)
		size |= (uint64(block[5]) << 0x10)
		size |= (uint64(block[6]) << 0x08)
		size |= (uint64(block[7]))

		io.CopyN(ioutil.Discard, r, int64(size))

		if !frame.IsValidFrameName(block[:4]) {
			position += 10 + int(size)
			continue
		}

		frames = append(frames, frame.ID3v2CapturedFrame{
			Frame:    string(block[:4]),
			Size:     size,
			Position: 10 + position,
		})

		position += 10 + int(size)
	}

	return frames
}
