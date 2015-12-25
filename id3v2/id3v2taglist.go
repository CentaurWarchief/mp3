package id3v2

type ID3v2TagList map[string]string

func (l ID3v2TagList) HasTag(tag string) bool {
	if len(tag) != 4 {
		return false
	}

	if _, ok := l[tag]; ok {
		return true
	}

	return false
}
