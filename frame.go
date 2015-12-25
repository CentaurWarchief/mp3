package mp3

type Frame struct {
	Version       string
	Layer         string
	Frequency     int
	Samples       int
	Bitrate       int
	SampleRate    int
	Mode          int
	ModeExtension int
	Private       int
	Copyright     int
	Emphasis      int
	Padding       int
}

func (f Frame) Size() int {
	if f.Layer == "1" {
		return ((12 * f.Bitrate * 1000 / f.SampleRate) + f.Padding) * 4
	}

	return ((144 * f.Bitrate * 1000 / f.SampleRate) + f.Padding)
}

func (f Frame) ModeText() string {
	switch f.Mode {
	case 0:
		return "Stereo"
	case 1:
		return "Joint Stereo"
	case 2:
		return "Dual Channel"
	case 3:
		return "Mono"
	}

	return ""
}
