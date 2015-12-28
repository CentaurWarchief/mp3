package unpack

import (
	"github.com/CentaurWarchief/mp3/id3v2/frame"
	"github.com/CentaurWarchief/mp3/id3v2/util"

	"golang.org/x/text/encoding/unicode"
)

// https://en.wikipedia.org/wiki/ID3#ID3v2
// $00 - ISO-8859-1
// $01 - UCS-2 (UTF-16 encoded with BOM)
// $02 - UTF16-BE (encoded without BOM)
// $03 - UTF-8
const (
	ISO88591 byte = iota
	UTF16
	UTF16BE
	UTF8
)

// TextFrameUnpacker allows to unpack a simple text frame. It also
// supports all specified encodings by ID3v2 standard
type TextFrameUnpacker struct {
}

// CanUnpack checks if the given frame is a text frame. Usually
// all text frames starts with T with exception of TXXX as per
// specification
func (u TextFrameUnpacker) CanUnpack(frame frame.Frame) bool {
	return frame.ID()[0] == 'T' && frame.ID() != "TXXX"
}

// Unpack unpacks the given text frame respecting its encoding
// (ISO-8859-1, UTF-8, UTF-16 or UTF16-BE)
func (u TextFrameUnpacker) Unpack(f frame.Frame, b []byte) interface{} {
	if b[0] == ISO88591 {
		utf8 := make([]rune, len(b[1:]))

		for i, b := range b[1:] {
			utf8[i] = rune(b)
		}

		return frame.NewTextFrame(f.ID(), string(utf8))
	}

	if b[0] == UTF16 {
		return frame.NewTextFrame(f.ID(), util.DecodeUTF16(
			b[1:],
			unicode.UTF16(unicode.BigEndian, unicode.ExpectBOM),
		))
	}

	if b[0] == UTF16BE {
		return frame.NewTextFrame(f.ID(), util.DecodeUTF16(
			b[1:],
			unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM),
		))
	}

	if b[0] == UTF8 {
		return frame.NewTextFrame(f.ID(), string(b[1:]))
	}

	return nil
}
