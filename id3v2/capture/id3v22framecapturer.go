package capture

import (
	"io"
	"io/ioutil"

	. "github.com/CentaurWarchief/mp3/id3v2/frame"
)

const (
	ID3v22FrameHeaderSize = 10
)

func ID3v22FrameCapturer(r io.Reader) (frames []ID3v2CapturedFrame) {
	position := ID3v22FrameHeaderSize

	for {
		frame := make([]byte, 6)

		if _, err := r.Read(frame); err != nil {
			break
		}

		size := (int(frame[3]) << 0x10)
		size |= (int(frame[4]) << 0x08)
		size |= (int(frame[5]))

		io.CopyN(ioutil.Discard, r, int64(4))
		io.CopyN(ioutil.Discard, r, int64(size))

		if !IsValidFrameName(frame[:3]) {
			continue
		}

		frames = append(frames, ID3v2CapturedFrame{
			Frame:    string(frame[:3]),
			Size:     uint64(size),
			Position: (ID3v22FrameHeaderSize - 4) + position,
		})

		position += (ID3v22FrameHeaderSize + size)
	}

	return frames
}
