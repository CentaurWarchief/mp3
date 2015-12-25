package id3v2

import "io"

func ParseID3v2Header(r io.Reader) (*Header, error) {
	block := make([]byte, 10)

	if _, err := r.Read(block); err != nil {
		return nil, err
	}

	if string(block[:3]) != "ID3" {
		return nil, ErrID3HeaderNotFound
	}

	major := int(block[3])

	switch major {
	case 2:
	case 3:
	case 4:
		break
	default:
		return nil, ErrUnrecognizedVersion
	}

	size := uint64(0)

	// https://en.wikipedia.org/wiki/Synchsafe
	// http://stackoverflow.com/a/5652842
	for _, b := range block[6:] {
		if (b & (1 << 7)) != 0 {
			return nil, ErrInvalidSize
		}

		size |= ((size << 7) | uint64(b))
	}

	return &Header{
		MajorVersion:      major,
		MinorVersion:      int(block[4]),
		Unsynchronization: (block[5] & 0x80) != 0,
		Extended:          (block[5] & 0x40) != 0,
		Experimental:      (block[5] & 0x20) != 0,
		Footer:            (block[5] & 0x10) != 0,
		Size:              size,
	}, nil
}
