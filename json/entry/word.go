// © 2018 Imhotep Software LLC. All rights reserved.

// Package entry provides for JSON encoding/decoding of a dictionary entry.
package entry

import (
	"encoding/json"
)

// Word from a dictionary load
type Word struct {
	Dictionary string `json:"dictionary"`
	Location   string `json:"location"`
	Word       string `json:"random_word"`
}

// ToJSON converts entry into raw json
func (w *Word) ToJSON() ([]byte, error) {
	return json.Marshal(w)
}

// FromJSON hydrates an entry from raw json
func (w *Word) FromJSON(bb []byte) error {
	return json.Unmarshal(bb, &w)
}
