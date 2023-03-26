package main

import (
	"testing"
)

func TestAccepts(t *testing.T) {
	// Checks the mimetype of .text file
	// this should return text/plain on macOS
	// but "application/octet-stream" on linux and windows
	mimetype := Accepts("text", "json")

	// assert that mimetype is application/octet-stream
	if mimetype != "application/octet-stream" {
		t.Errorf("mimetype is %s, want application/octet-stream", mimetype)
	}
}
