package entry_test

import (
	"encoding/json"
	"testing"

	"github.com/gopherland/labs/json/entry"
	"github.com/stretchr/testify/assert"
)

const raw = `{"dictionary":"artists","location":"/tmp/assets","random_word":"bumblebeetuna"}`

func TestDicMarshal(t *testing.T) {
	e := entry.Word{
		Dictionary: "artists",
		Location:   "dictionary1",
		Word:       "bumblebeetuna",
	}
	bb, err := json.Marshal(e)

	assert.Nil(t, err)
	assert.Equal(t, string(raw), string(bb))
}

// func TestDicUnmarshal(t *testing.T) {
// 	var e entry.Word
// 	err := json.Unmarshal([]byte(raw), &e)

// 	assert.Nil(t, err)
// 	assert.Equal(t, "artists", e.Dictionary)
// 	assert.Equal(t, "/tmp/assets", e.Location)
// 	assert.Equal(t, "bumblebeetuna", e.Word)
// }

// func TestDicUnmarshallFail(t *testing.T) {
// 	var e entry.Word
// 	err := json.Unmarshal([]byte("fred"), &e)
// 	assert.NotNil(t, err)
// }
