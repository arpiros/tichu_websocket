package models

import (
	"math/rand"
	"sort"
)

const (
	CardColorRed = iota
	CardColorGreen
	CardColorBlock
	CardColorBlue
	CardColorNone
)

const (
	CardTypeNone = iota
	CardTypeMahjong
	CardTypePhoenix
	CardTypeDrache
	CardTypeDashund
	CardTypeCount
)

const (
	StartCardNumber = 1
	EndCardNumber   = 13
	TotalCardCount  = 56
)

const BoomFourCard = 4

type Card struct {
	Number   int
	Color    int
	CardType int
}

func NewCardDeck() []*Card {
	var newDeck []*Card
	for cardType := 0; cardType < CardTypeCount; cardType++ {
		switch cardType {
		case CardTypeNone:
			for color := CardColorRed; color < CardColorNone; color++ {
				for num := StartCardNumber; num < EndCardNumber+1; num++ {
					newDeck = append(newDeck, &Card{
						Number:   num,
						Color:    color,
						CardType: cardType,
					})
				}
			}
		default:
			newDeck = append(newDeck, &Card{
				Number:   0,
				Color:    CardColorNone,
				CardType: cardType,
			})
		}
	}

	shuffleCardDeck(newDeck)

	return newDeck
}

func shuffleCardDeck(deck []*Card) {
	for key := range deck {
		randNum := rand.Intn(TotalCardCount)
		deck[randNum], deck[key] = deck[key], deck[randNum]
	}
}

func IsBoom(cards []Card) bool {
	baseCard := cards[0]
	if baseCard.CardType != CardTypeNone {
		return false
	}

	switch {
	case len(cards) == BoomFourCard:
		for _, card := range cards {
			if baseCard.Number != card.Number {
				return false
			}
		}
	case len(cards) > BoomFourCard:
		var numbers []int
		for _, card := range cards {
			if baseCard.Color != card.Color {
				return false
			}

			numbers = append(numbers, card.Number)

			sort.Ints(numbers)

			for i := 0; i < len(numbers)-1; i++ {
				if numbers[i+1]-numbers[i] != 1 {
					return false
				}
			}
		}

	default:
		return false
	}

	return true
}
