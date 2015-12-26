package id3v2

import (
	"io"

	"github.com/CentaurWarchief/mp3/id3v2/capture"
)

// Creates an empty ID3v2Parser without adding any frame
// capturer. You must add your own capture strategy(ies)
func NewEmptyID3v2Parser() *ID3v2Parser {
	return &ID3v2Parser{
		make(map[int]ID3v2FrameCapturer),
	}
}

// Creates a new ID3v2Parser adding by default capture
// strategies for both v2.2/v2.3 and v2.4 frames
func NewID3v2Parser() *ID3v2Parser {
	p := &ID3v2Parser{
		make(map[int]ID3v2FrameCapturer),
	}

	p.AddVersionedFrameCapturer(2, capture.ID3v22FrameCapturer)
	p.AddVersionedFrameCapturer(3, capture.ID3v23FrameCapturer)
	p.AddVersionedFrameCapturer(4, capture.ID3v24FrameCapturer)

	return p
}

type ID3v2Parser struct {
	byVersion map[int]ID3v2FrameCapturer
}

// Adds a new frame capturer for the given major version of ID3v2
func (p *ID3v2Parser) AddVersionedFrameCapturer(major int, capturer ID3v2FrameCapturer) {
	p.byVersion[major] = capturer
}

func (p ID3v2Parser) Parse(r io.Reader) error {
	header, err := ParseID3v2Header(r)

	if err != nil {
		return err
	}

	if _, ok := p.byVersion[header.MajorVersion]; !ok {
		return ErrUnsupportedVersion
	}

	// capturer := p.byVersion[header.MajorVersion]
	// captured := capturer(io.LimitReader(r, int64(header.Size)))

	return nil
}
