package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers registers controllers for REST endpoints
func RegisterControllers() {
	fc := newFlashcardController()

	http.Handle("/flashcards", *fc)
	http.Handle("/flashcards/", *fc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
