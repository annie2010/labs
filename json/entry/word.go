// © 2019 Imhotep Software LLC. All rights reserved.

// Package entry provides for JSON encoding/decoding of a dictionary entry.
package entry

import (
	"encoding/json"
)

const prefix = "/dictionary/assets"

// Word from a dictionary load
type (
	Word struct {
		Dictionary string `json:"dictionary"`
		Location   string `json:"location"`
		Word       string `json:"random_word"`
	}

	aWord Word
)

// Add compiler check to ensure our word is marshallable.
var _ json.Marshaler = &Word{}

// MarshalJSON converts entry into raw json
func (w *Word) MarshalJSON() ([]byte, error) {
	return json.Marshal(aWord(*w))
}

// UnmarshalJSON hydrates an entry from raw json
func (w *Word) UnmarshalJSON(bb []byte) error {
	var aw aWord
	if err := json.Unmarshal(bb, &aw); err != nil {
		return err
	}
	*w = Word(aw)

	return nil
}
