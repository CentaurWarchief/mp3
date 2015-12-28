package unpack

import (
	"bytes"

	"github.com/CentaurWarchief/mp3/id3v2/frame"
)

// AttachedPictureFrameUnpacker allows to unpack an attached picture
// frame. It handles both PIC (v2.2) and APIC (v2.3 and v2.4) frames
type AttachedPictureFrameUnpacker struct {
}

// CanUnpack checks whether the given frame is a PIC v2.2 or APIC v2.3 and v2.4
func (u AttachedPictureFrameUnpacker) CanUnpack(frame frame.Frame) bool {
	return frame.ID() == "PIC" || frame.ID() == "APIC"
}

// Unpack unpacks the given picture/attached picture frame into a ImageFrame
func (u AttachedPictureFrameUnpacker) Unpack(f frame.Frame, b []byte) interface{} {
	// http://id3.org/id3v2.3.0#Attached_picture
	encoding := b[0]

	var mime int
	var description int

	if mime = bytes.IndexByte(b[1:], 0x0); mime == -1 {
		return nil
	}

	if description = bytes.IndexByte(b[(mime+2):], 0x0); description == -1 {
		return nil
	}

	// TODO: handle description encoding properly
	return frame.ImageFrame{
		Encoding:    encoding,
		MIMEType:    string(b[1:(mime + 1)]),
		Type:        int(b[(mime + 2)]),
		Description: string(b[(mime + 3):(mime + 2 + description)]),
		Binary:      b[(mime + 2 + description + 1):],
	}
}
