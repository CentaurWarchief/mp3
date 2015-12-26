package util

import "golang.org/x/text/encoding"

// Decodes UTF16 using decoder provided by the given `encoding`
func DecodeUTF16(data []byte, encoding encoding.Encoding) string {
	decoded := make([]byte, 2*len(data))

	decoder := encoding.NewDecoder()

	if n, _, err := decoder.Transform(decoded, data, true); err == nil {
		return string(decoded[:n])
	}

	return ""
}
