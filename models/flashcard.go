package models

import (
	"errors"
	"fmt"
)

// FlashCard replresents a flash card
type FlashCard struct {
	ID       int
	Question string
	Hint     string
	Answer   string
}

var (
	cards  []*FlashCard
	nextID = 1
)

// GetFlashCards returns all flashcards
func GetFlashCards() []*FlashCard {
	return cards
}

// AddFlashCard adds a card to the flash cards deck
func AddFlashCard(card FlashCard) (FlashCard, error) {
	if card.ID != 0 {
		return FlashCard{}, errors.New("New FlashCard must not include ID or it must be set to zero")
	}
	card.ID = nextID
	nextID++
	cards = append(cards, &card)
	return card, nil
}

// GetFlashCardByID returns a flashcard with the given id
func GetFlashCardByID(id int) (FlashCard, error) {
	for _, c := range cards {
		if c.ID == id {
			return *c, nil
		}
	}
	return FlashCard{}, fmt.Errorf("Flashcard with ID '%v' not found", id)
}

// UpdateFlashCard updates the flashcard if found
func UpdateFlashCard(f FlashCard) (FlashCard, error) {
	for i, candidate := range cards {
		if candidate.ID == f.ID {
			cards[i] = &f
			return f, nil
		}
	}
	return FlashCard{}, fmt.Errorf("Flashcard with ID '%v' not found", f.ID)
}

// RemoveFlashCardByID removes the flashcard with the given id
func RemoveFlashCardByID(id int) error {
	for i, c := range cards {
		if c.ID == id {
			cards = append(cards[:i], cards[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Flashcard with ID '%v' not found", id)
}
