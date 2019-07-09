package main

import (
	"encoding/json"
	"net/http"

	"github.com/williammartin/storywriter/story"
)

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/write", write)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func write(w http.ResponseWriter, r *http.Request) {
	var draft draft
	if err := json.NewDecoder(r.Body).Decode(&draft); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("please provide a well-formed body"))
		return
	}

	interpolator := story.Interpolator{}
	story := interpolator.Interpolate(draft.Template, draft.Words)

	w.WriteHeader(200)
	w.Write([]byte(story))
}

type draft struct {
	Template string   `json:"template"`
	Words    []string `json:"words"`
}
