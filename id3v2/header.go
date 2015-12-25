package id3v2

type Header struct {
	MajorVersion      int
	MinorVersion      int
	Unsynchronization bool
	Extended          bool
	Experimental      bool
	Footer            bool
	Size              uint64
}
