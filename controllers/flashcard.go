package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/flashcards/webservice/models"
)

type flashcardController struct {
	flashcardIDPattern *regexp.Regexp
}

func (fc flashcardController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/flashcards" {
		switch r.Method {
		case http.MethodGet:
			fc.getAll(w, r)
		case http.MethodPost:
			fc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := fc.flashcardIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			fc.get(id, w)
		case http.MethodPut:
			fc.put(id, w, r)
		case http.MethodDelete:
			fc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func newFlashcardController() *flashcardController {
	return &flashcardController{
		flashcardIDPattern: regexp.MustCompile(`^/flashcards/(\d+)/?`),
	}
}

func (fc *flashcardController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetFlashCards(), w)
}

func (fc *flashcardController) get(id int, w http.ResponseWriter) {
	f, err := models.GetFlashCardByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(f, w)
}

func (fc *flashcardController) post(w http.ResponseWriter, r *http.Request) {
	f, err := fc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse FlashCard object"))
		return
	}
	f, err = models.AddFlashCard(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(f, w)
}

func (fc *flashcardController) put(id int, w http.ResponseWriter, r *http.Request) {
	f, err := fc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse FlashCard object"))
		return
	}
	if id != f.ID {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ID of submitted flashcard must match ID in the URL"))
		return
	}
	f, err = models.UpdateFlashCard(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(f, w)
}

func (fc *flashcardController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveFlashCardByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (fc *flashcardController) parseRequest(r *http.Request) (models.FlashCard, error) {
	dec := json.NewDecoder(r.Body)
	var f models.FlashCard
	err := dec.Decode(&f)
	if err != nil {
		return models.FlashCard{}, err
	}
	return f, nil
}
