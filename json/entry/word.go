// © 2019 Imhotep Software LLC. All rights reserved.

// Package entry provides for JSON encoding/decoding of a dictionary entry.
package entry

import (
	"encoding/json"
)

// Word reprsents a dictionary word.
type Word struct {
	// YOUR CODE...
}

// Make compiler check our marshaller interface is implemented.
var _ json.Marshaler = &Word{}
