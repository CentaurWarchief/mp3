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

	size := uint64(0)

	for _, b := range block[6:] {
		if (b & (1 << 7)) != 0 {
			return nil, ErrInvalidSize
		}

		size |= ((size << 7) | uint64(b))
	}

	return &Header{
		MajorVersion:      int(block[3]),
		MinorVersion:      int(block[4]),
		Unsynchronization: (block[5] & 1 << 7) != 0,
		Extended:          (block[5] & 1 << 6) != 0,
		Experimental:      (block[5] & 1 << 5) != 0,
		Footer:            (block[5] & 1 << 4) != 0,
		Size:              size,
	}, nil
}
