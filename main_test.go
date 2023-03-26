package main

import (
	"testing"
)

func TestAccepts(t *testing.T) {
	// Checks the mimetype of .text files
	mimetype := Accepts("text", "json")

	// assert that mimetype is application/json
	if mimetype != "application/json" {
		t.Errorf("mimetype is %s, want application/json", mimetype)
	}
}
