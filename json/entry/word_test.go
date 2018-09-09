package entry_test

import (
	"testing"

	"github.com/gopherland/labs/json/entry"
	"github.com/stretchr/testify/assert"
)

const raw = `{"dictionary":"artists","location":"/tmp/assets","random_word":"bumblebeetuna"}`

func TestToJSON(t *testing.T) {
	e := entry.Word{
		Dictionary: "artists",
		Location:   "/tmp/assets",
		Word:       "bumblebeetuna",
	}
	bb, err := e.ToJSON()

	assert.Nil(t, err)
	assert.Equal(t, string(raw), string(bb))
}

func TestFromJSON(t *testing.T) {
	var e entry.Word
	err := e.FromJSON([]byte(raw))

	assert.Nil(t, err)
	assert.Equal(t, "artists", e.Dictionary)
	assert.Equal(t, "/tmp/assets", e.Location)
	assert.Equal(t, "bumblebeetuna", e.Word)
}

func TestFromJSONFail(t *testing.T) {
	var e entry.Word
	err := e.FromJSON([]byte("fred"))
	assert.NotNil(t, err)
}
