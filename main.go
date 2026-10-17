package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed public
var public embed.FS

func main() {
	publicFS, err := fs.Sub(public, "public")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(publicFS)))

	port := ":8080"

	log.Printf("Server has started. Listening on http://127.0.0.1%s...\n", port)
	log.Println("Press Ctrl+C to exit.")

	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
