package id3v2

import (
	"io"
	"io/ioutil"
)

const (
	ID3v22FrameHeaderSize = 10
)

func ID3v22FrameCapturer(r io.Reader) (frames ID3v2FrameList) {
	position := ID3v22FrameHeaderSize

	for {
		frame := make([]byte, 6)

		if _, err := r.Read(frame); err != nil {
			break
		}

		size := (int(frame[3]) << 0x10)
		size |= (int(frame[4]) << 0x08)
		size |= (int(frame[5]))

		tag := string(frame[0:3])

		io.CopyN(ioutil.Discard, r, int64(4))
		io.CopyN(ioutil.Discard, r, int64(size))

		if !ID3v22TagList.HasTag(tag) {
			continue
		}

		frames = append(frames, ID3v2FrameSizePosition{
			Tag:          tag,
			Size:         uint64(size),
			Position:     position,
			BodyPosition: (ID3v22FrameHeaderSize - 4) + position,
		})

		position += (ID3v22FrameHeaderSize + size)
	}

	return frames
}
