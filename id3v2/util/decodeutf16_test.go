package util_test

import (
	"testing"

	"golang.org/x/text/encoding/unicode"

	"github.com/CentaurWarchief/mp3/id3v2/util"
	"github.com/stretchr/testify/assert"
)

func TestDecodeInvalidUTF16(t *testing.T) {
	decoded := util.DecodeUTF16(
		[]byte{},
		unicode.UTF16(unicode.BigEndian, unicode.ExpectBOM),
	)

	assert.Equal(t, "", decoded)
}

func TestDecodeUTF16(t *testing.T) {
	decoded := util.DecodeUTF16(
		[]byte{
			255, 254, 73, 0, 110, 0, 32, 0, 65, 0,
			110, 0, 121, 0, 32, 0, 84, 0, 111, 0, 110,
			0, 103, 0, 117, 0, 101, 0,
		},
		unicode.UTF16(unicode.BigEndian, unicode.ExpectBOM),
	)

	assert.Equal(t, "In Any Tongue", decoded)
}
