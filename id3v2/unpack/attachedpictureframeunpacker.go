package unpack

import "github.com/CentaurWarchief/mp3/id3v2/frame"

// AttachedPictureFrameUnpacker allows to unpack an attached
// picture frame
type AttachedPictureFrameUnpacker struct {
}

// CanUnpack checks whether the given frame is an attached
// picture frame PIC v2.2 or APIC v2.3/v2.4
func (u AttachedPictureFrameUnpacker) CanUnpack(frame frame.Frame) bool {
	return frame.ID() == "PIC" || frame.ID() == "APIC"
}

// Unpack unpacks the given frame
func (u AttachedPictureFrameUnpacker) Unpack(
	frame frame.Frame,
	body []byte,
) interface{} {
	return nil
}
