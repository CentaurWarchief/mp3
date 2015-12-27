package unpack

import (
	"github.com/CentaurWarchief/mp3/id3v2/frame"
	"github.com/CentaurWarchief/mp3/id3v2/util"

	"golang.org/x/text/encoding/unicode"
)

// $00 - ISO-8859-1
// $01 - UCS-2 (UTF-16 encoded with BOM)
// $02 - UTF16-BE (encoded without BOM)
// $03 - UTF-8
const (
	ISO88591 = 0x00
	UTF16    = 0x01
	UTF16BE  = 0x02
	UTF8     = 0x03
)

// TextFrameUnpacker allows to unpack a simple text frame. It also
// supports all specified encodings by ID3v2 standard
type TextFrameUnpacker struct {
}

// CanUnpack checks whether the given frame is a text frame. Usually
// all text frames starts with `T` with an exception of `TXXX`
func (u TextFrameUnpacker) CanUnpack(frame frame.Frame) bool {
	return frame.ID()[0] == 'T' && frame.ID() != "TXXX"
}

// Unpack unpacks the given frame
func (u TextFrameUnpacker) Unpack(f frame.Frame, body []byte) interface{} {
	if body[0] == ISO88591 {
		utf8 := make([]rune, len(body[1:]))

		for i, b := range body[1:] {
			utf8[i] = rune(b)
		}

		return frame.NewTextFrame(f.ID(), string(utf8))
	}

	if body[0] == UTF16 {
		return frame.NewTextFrame(
			f.ID(),
			util.DecodeUTF16(
				body[1:],
				unicode.UTF16(unicode.BigEndian, unicode.ExpectBOM),
			),
		)
	}

	if body[0] == UTF16BE {
		return frame.NewTextFrame(
			f.ID(),
			util.DecodeUTF16(
				body[1:],
				unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM),
			),
		)
	}

	if body[0] == UTF8 {
		return frame.NewTextFrame(f.ID(), string(body[1:]))
	}

	return nil
}
