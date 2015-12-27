package unpack

import (
	"bytes"

	"github.com/CentaurWarchief/mp3/id3v2/frame"
)

// AttachedPictureFrameUnpacker allows to unpack an attached
// picture frame. It handles both PIC (ID3v2.2) and APIC
// frames
type AttachedPictureFrameUnpacker struct {
}

// CanUnpack checks whether the given frame is a PIC v2.2 or APIC v2.3 and v2.4
func (u AttachedPictureFrameUnpacker) CanUnpack(frame frame.Frame) bool {
	return frame.ID() == "PIC" || frame.ID() == "APIC"
}

// Unpack unpacks the given attached picture frame
func (u AttachedPictureFrameUnpacker) Unpack(f frame.Frame, b []byte) interface{} {
	// http://id3.org/id3v2.3.0#Attached_picture
	n0 := 0
	n1 := 1 + bytes.IndexByte(b[(1+n0):], 0x00)
	n2 := 2 + bytes.IndexByte(b[(3+n1):], 0x00)

	// TODO: handle description encoding properly
	return frame.ImageFrame{
		Encoding:    b[n0],
		MIMEType:    string(b[1:n1]),
		Type:        b[(n1 + 2)],
		Description: string(b[(2 + n1):(n1 + n2 + 1)]),
		Binary:      b[(n1 + n2 + 2):],
	}
}
