package id3v2

type ID3v2Header struct {
	MajorVersion      int
	MinorVersion      int
	Unsynchronization bool
	Extended          bool
	Experimental      bool
	Footer            bool
	Size              uint64
}
