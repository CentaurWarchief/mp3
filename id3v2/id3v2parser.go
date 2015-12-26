package id3v2

import "io"

func NewEmptyID3v2Parser() *ID3v2Parser {
	return &ID3v2Parser{
		make(map[int]ID3v2FrameCapturer),
	}
}

func NewID3v2Parser() *ID3v2Parser {
	p := &ID3v2Parser{
		make(map[int]ID3v2FrameCapturer),
	}

	p.AddVersionedFrameCapturer(2, ID3v22FrameCapturer)
	p.AddVersionedFrameCapturer(3, ID3v23FrameCapturer)
	p.AddVersionedFrameCapturer(4, ID3v24FrameCapturer)

	return p
}

type ID3v2Parser struct {
	byVersion map[int]ID3v2FrameCapturer
}

func (p *ID3v2Parser) AddVersionedFrameCapturer(
	major int,
	capturer ID3v2FrameCapturer,
) {
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
