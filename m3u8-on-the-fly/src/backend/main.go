package main

import (
	"net/http"
	"strings"

	"../signM3u8"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	videoName := strings.TrimPrefix(path, "/")
	m3u8, err := signM3u8.SignM3u8(videoName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("not exist"))
	} else {
		w.Write([]byte(m3u8))
	}
}
func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
