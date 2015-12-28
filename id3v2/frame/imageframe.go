package frame

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
)

// ImageFrame represents an attached picture (APIC)
// or a picture (PIC) frame
type ImageFrame struct {
	Encoding    byte
	MIMEType    string
	Type        int
	Description string
	Binary      []byte
}

// WriteTo writes the entire image to the given writer (io.Writer)
func (f ImageFrame) WriteTo(w io.Writer) (int64, error) {
	return io.Copy(
		w,
		bytes.NewReader(f.Binary),
	)
}

// AsDataURI returns the image as data URI (http://tools.ietf.org/html/rfc2397)
// string
func (f ImageFrame) AsDataURI() string {
	return fmt.Sprintf(
		"data:%s;base64,%s",
		f.MIMEType,
		base64.StdEncoding.EncodeToString(f.Binary),
	)
}
