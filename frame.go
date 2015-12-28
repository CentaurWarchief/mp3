package mp3

// Frame represents a MPEG audio frame
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

// Size returns the frame size based on its layer, bitrate, sample rate
// and padding
func (f Frame) Size() int {
	if f.Layer == "1" {
		return ((12 * f.Bitrate * 1000 / f.SampleRate) + f.Padding) * 4
	}

	return ((144 * f.Bitrate * 1000 / f.SampleRate) + f.Padding)
}

// ModeText returns the frame mode as string
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
