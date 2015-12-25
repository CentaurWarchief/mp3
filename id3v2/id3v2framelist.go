package id3v2

type ID3v2FrameList []ID3v2Frame

func (l ID3v2FrameList) HasFrame(ID string) bool {
	for _, frame := range l {
		if frame.ID == ID {
			return true
		}
	}

	return false
}

func (l ID3v2FrameList) Frames(ID string) (frames []ID3v2Frame) {
	for _, frame := range l {
		if frame.ID != ID {
			continue
		}

		frames = append(frames, frame)
	}

	return frames
}
