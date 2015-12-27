package capture

import (
	"io"
	"io/ioutil"

	"github.com/CentaurWarchief/mp3/id3v2/frame"
)

// ID3v22FrameCapturer captures all v2.2 frames from the given
// reader until reach its EOF
func ID3v22FrameCapturer(r io.Reader) (frames []frame.CapturedFrame) {
	position := 10

	for {
		block := make([]byte, 6)

		if _, err := r.Read(block); err != nil {
			break
		}

		size := (int(block[3]) << 0x10)
		size |= (int(block[4]) << 0x08)
		size |= (int(block[5]))

		io.CopyN(ioutil.Discard, r, int64(size))

		if !frame.IsValidFrameName(block[:3]) {
			position += 6 + size
			continue
		}

		frames = append(frames, frame.CapturedFrame{
			Frame:    string(block[:3]),
			Size:     uint64(size),
			Position: 6 + position,
		})

		position += 6 + size
	}

	return frames

}
