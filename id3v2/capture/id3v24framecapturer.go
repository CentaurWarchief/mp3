package capture

import (
	"io"
	"io/ioutil"

	"github.com/CentaurWarchief/mp3/id3v2/frame"
)

// ID3v24FrameCapturer captures all v2.4 frames from the given
// reader until reach its EOF
func ID3v24FrameCapturer(r io.Reader) (frames []frame.ID3v2CapturedFrame) {
	position := 10

	for {
		block := make([]byte, 10)

		if _, err := r.Read(block); err != nil {
			break
		}

		size := int32(0)

		size |= (int32((block[4] & 0x7F)) << 0x15) // 21
		size |= (int32((block[5] & 0x7F)) << 0x0E) // 14
		size |= (int32((block[6] & 0x7F)) << 0x07) // 07
		size |= (int32((block[7] & 0x7F)))

		io.CopyN(ioutil.Discard, r, int64(size))

		if !frame.IsValidFrameName(block[:4]) {
			continue
		}

		frames = append(frames, frame.ID3v2CapturedFrame{
			Frame:    string(block[:4]),
			Size:     uint64(size),
			Position: 10 + position,
		})

		position += 10 + int(size)
	}

	return frames
}
