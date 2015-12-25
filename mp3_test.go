package mp3_test

import (
	"testing"

	"github.com/CentaurWarchief/mp3"
	"github.com/stretchr/testify/assert"
)

func TestNewMP3(t *testing.T) {
	mp3 := mp3.New("res/Windmill.mp3", 147301)

	assert.Equal(t, "res/Windmill.mp3", mp3.File)
	assert.Equal(t, int64(147301), mp3.Size)
}
