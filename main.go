package main

import (
	"fmt"
	"mime"
)

func main() {
	mimetype := Accepts("text", "json")

	fmt.Println(mimetype)
}

func Accepts(files ...string) string {
	for _, file := range files {
		mimetype := mime.TypeByExtension("." + file)
		if mimetype != "" {
			return mimetype
		}
	}
	return "application/octet-stream"
}
