package mp3

type MP3 struct {
	File string
	Size int64
}

func New(file string, size int64) MP3 {
	return MP3{
		File: file,
		Size: size,
	}
}
