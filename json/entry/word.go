// © 2018 Imhotep Software LLC. All rights reserved.

// Package entry provides for JSON encoding/decoding of a dictionary entry.
package entry

// Entry from a dictionary load.
type Word struct {
	// YOUR CODE...
}

// ToJSON converts entry into raw json.
func (w *Word) ToJSON() ([]byte, error) {
	// YOUR CODE...
	return []byte{}, nil
}

// FromJSON hydrates an entry from raw json.
func (w *Word) FromJSON(bb []byte) error {
	// YOUR CODE...
	return nil
}
