package id3v2

// Header represents a valid ID3v2 header
type Header struct {
	MajorVersion      int
	MinorVersion      int
	Unsynchronization bool
	Extended          bool
	Experimental      bool
	Footer            bool
	Size              uint64
}
