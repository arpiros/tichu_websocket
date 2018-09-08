package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestNewCardDeck(t *testing.T) {
	cardDeck := NewCardDeck()
	assert.Equal(t, len(cardDeck), TotalCardCount)
}