package capture

import (
	"io"
	"io/ioutil"

	. "github.com/CentaurWarchief/mp3/id3v2/frame"
)

const (
	ID3v23FrameHeaderSize = 10
)

func ID3v23FrameCapturer(r io.Reader) (frames []ID3v2CapturedFrame) {
	position := ID3v23FrameHeaderSize

	for {
		frame := make([]byte, ID3v23FrameHeaderSize)

		if _, err := r.Read(frame); err != nil {
			break
		}

		size := (uint64(frame[4]) << 0x18)
		size |= (uint64(frame[5]) << 0x10)
		size |= (uint64(frame[6]) << 0x08)
		size |= (uint64(frame[7]))

		io.CopyN(ioutil.Discard, r, int64(size))

		if !IsValidFrameName(frame[:4]) {
			position += ID3v23FrameHeaderSize + int(size)
			continue
		}

		frames = append(frames, ID3v2CapturedFrame{
			Frame:    string(frame[:4]),
			Size:     size,
			Position: ID3v23FrameHeaderSize + position,
		})

		position += ID3v23FrameHeaderSize + int(size)
	}

	return frames
}
