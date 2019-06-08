// © 2019 Imhotep Software LLC. All rights reserved.

// Package entry provides for JSON encoding/decoding of a dictionary entry.
package entry

import (
	"encoding/json"
	"fmt"
)

const prefix = "/dictionary/assets"

// Word from a dictionary load
type Word struct {
	Dictionary string
	Location   string
	Word       string
}

// JSONWord a JSON representation of a word
type JSONWord struct {
	Dictionary string `json:"dictionary"`
	Location   string `json:"location"`
	Word       string `json:"random_word"`
}

func (w *JSONWord) toWord() *Word {
	return &Word{
		Dictionary: w.Dictionary,
		Location:   prefix + w.Location,
		Word:       w.Word,
	}
}

// MarshalJSON converts entry into raw json
func (w *Word) MarshalJSON() ([]byte, error) {
	fmt.Println("M")
	return json.Marshal(w)
}

// UnmarshalJSON hydrates an entry from raw json
func (w *Word) UnmarshalJSON(bb []byte) error {
	var wo JSONWord
	fmt.Println("UN")
	if err := json.Unmarshal(bb, &wo); err != nil {
		return err
	}
	w = wo.toWord()
	return nil
}
