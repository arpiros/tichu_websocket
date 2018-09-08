package models

import "math/rand"

const (
	CardColerNone  = iota
	CardColorRed
	CardColorBlue
	CardColorGreen
	CardColorBlock
	CardColorCount
)

const (
	CardTypeNone    = iota
	CardTypeSparrow
	CardTypeDog
	CardTypeDragon
	CardTypePhoenix
	CardTypeCount
)

const (
	StartCardNumber = 2
	EndCardNumber   = 13
	TotalCardCount  = 52
)

type Card struct {
	Number   int
	Color    int
	CardType int
}

func NewCardDeck() []*Card {
	var newDeck []*Card
	for cType := 0; cType < CardTypeCount; cType++ {
		switch cType {
		case CardTypeNone:
			for color := CardColorRed; color < CardColorCount; color++ {
				for num := StartCardNumber; num < EndCardNumber+1; num++ {
					newDeck = append(newDeck, &Card{
						Number:   num,
						Color:    color,
						CardType: cType,
					})
				}
			}
		default:
			newDeck = append(newDeck, &Card{
				Number:   0,
				Color:    CardColerNone,
				CardType: cType,
			})
		}
	}

	shuffleCardDeck(newDeck)

	return newDeck
}

func shuffleCardDeck(deck []*Card) {
	for key := range deck {
		deck[rand.Intn(TotalCardCount)], deck[key] = deck[key], deck[rand.Intn(TotalCardCount)]
	}
}
