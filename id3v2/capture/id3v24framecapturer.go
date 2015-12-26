package capture

import (
	"io"
	"io/ioutil"

	. "github.com/CentaurWarchief/mp3/id3v2/frame"
)

const (
	ID3v24FrameHeaderSize = 10
)

func ID3v24FrameCapturer(r io.Reader) (frames []ID3v2CapturedFrame) {
	position := ID3v24FrameHeaderSize

	for {
		frame := make([]byte, ID3v24FrameHeaderSize)

		if _, err := r.Read(frame); err != nil {
			break
		}

		size := int32(0)

		size |= (int32((frame[4] & 0x7F)) << 0x15) // 21
		size |= (int32((frame[5] & 0x7F)) << 0x0E) // 14
		size |= (int32((frame[6] & 0x7F)) << 0x07) // 07
		size |= (int32((frame[7] & 0x7F)))

		io.CopyN(ioutil.Discard, r, int64(size))

		if !IsValidFrameName(frame[:4]) {
			continue
		}

		frames = append(frames, ID3v2CapturedFrame{
			Frame:    string(frame[:4]),
			Size:     uint64(size),
			Position: ID3v24FrameHeaderSize + position,
		})

		position += (ID3v24FrameHeaderSize + int(size))
	}

	return frames
}
