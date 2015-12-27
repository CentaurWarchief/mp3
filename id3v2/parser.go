package id3v2

import (
	"io"

	"github.com/CentaurWarchief/mp3/id3v2/capture"
	"github.com/CentaurWarchief/mp3/id3v2/frame"
	"github.com/CentaurWarchief/mp3/id3v2/unpack"
)

// NewEmptyParser creates an empty ID3v2Parser without adding
// any frame capturer nor unpacker. You must add your own capture
// strategy(ies) and unpack(ers)
func NewEmptyParser() *Parser {
	return &Parser{
		make(map[int]FrameCapturer),
		make([]FrameUnpacker, 0),
	}
}

// NewParser creates a new Parser adding by default
// capture strategies for both v2.2/v2.3 and v2.4 frames
func NewParser() *Parser {
	p := &Parser{
		make(map[int]FrameCapturer),
		make([]FrameUnpacker, 0),
	}

	p.AddVersionedFrameCapturer(2, capture.ID3v22FrameCapturer)
	p.AddVersionedFrameCapturer(3, capture.ID3v23FrameCapturer)
	p.AddVersionedFrameCapturer(4, capture.ID3v24FrameCapturer)

	p.AddUnpacker(unpack.TextFrameUnpacker{})
	p.AddUnpacker(unpack.AttachedPictureFrameUnpacker{})

	return p
}

// Parser is a versioned ID3v2 frame parser. It exposes a simple
// API for parsing all valid frames from a native `io.Reader`
type Parser struct {
	byVersion map[int]FrameCapturer
	unpacker  []FrameUnpacker
}

// AddVersionedFrameCapturer adds a new frame capturer
// for the given major version of ID3v2
func (p *Parser) AddVersionedFrameCapturer(major int, capturer FrameCapturer) {
	p.byVersion[major] = capturer
}

// AddUnpacker adds a new frame unpacker
func (p *Parser) AddUnpacker(unpacker FrameUnpacker) {
	p.unpacker = append(p.unpacker, unpacker)
}

// Parse parses all valid ID3v2 frames and returns
// all readable frames or an error
func (p Parser) Parse(r io.Reader) (frames []frame.ReadableFrame, err error) {
	header, err := ParseHeader(r)

	if err != nil {
		return nil, err
	}

	if _, ok := p.byVersion[header.MajorVersion]; !ok {
		return nil, ErrUnsupportedVersion
	}

	capturer := p.byVersion[header.MajorVersion]
	captured := capturer(io.LimitReader(r, int64(header.Size)))

	for _, f := range captured {
		for _, u := range p.unpacker {
			if !u.CanUnpack(f) {
				continue
			}

			frames = append(
				frames,
				f.AsReadable(r.(io.ReaderAt), wrap(f, u)),
			)
		}
	}

	return frames, nil
}

func wrap(f frame.CapturedFrame, u FrameUnpacker) func(body []byte) interface{} {
	return func(body []byte) interface{} {
		return u.Unpack(f, body)
	}
}
