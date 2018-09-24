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
	StartCardNumber = 2
	EndCardNumber   = 14
	TotalCardCount  = 56
)

const BoomFourCard = 4

type Card struct {
	Number   int
	Color    int
	CardType int
}

type CardList []*Card

func (c CardList) Len() int           { return len(c) }
func (c CardList) Less(i, j int) bool { return c[i].Number < c[j].Number }
func (c CardList) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func NewCardDeck() CardList {
	var newDeck CardList
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

func shuffleCardDeck(deck CardList) {
	for key := range deck {
		randNum := rand.Intn(TotalCardCount)
		deck[randNum], deck[key] = deck[key], deck[randNum]
	}
}

func IsBoom(cards CardList) bool {
	baseCard := cards[0]
	if baseCard.CardType != CardTypeNone {
		return false
	}

	switch {
	case len(cards) == BoomFourCard:
		for _, card := range cards {
			if card.CardType != CardTypeNone {
				return false
			}

			if baseCard.Number != card.Number {
				return false
			}
		}
	case len(cards) > BoomFourCard:
		for _, card := range cards {
			if card.CardType != CardTypeNone {
				return false
			}

			if baseCard.Color != card.Color {
				return false
			}
		}

		sort.Sort(cards)
		for i := 0; i < len(cards)-1; i++ {
			if cards[i+1].Number-cards[i].Number != 1 {
				return false
			}
		}

	default:
		return false
	}

	return true
}

func GetLargestNumber(cards CardList) int {
	return 0
}

func IsSingle(cards CardList) bool {
	return len(cards) == 1
}

func IsPair(cards CardList) bool {
	if len(cards) != 2 {
		return false
	}

	for _, v := range cards {
		if v.CardType == CardTypePhoenix {
			return true
		}

		if v.Number == 0 {
			return false
		}
	}

	return cards[0].Number == cards[1].Number
}

func IsTriple(cards CardList) bool {
	if len(cards) != 3 {
		return false
	}

	cardCounter := make(map[int]int)
	for _, v := range cards {
		if v.CardType == CardTypePhoenix {
			continue
		}

		if v.Number == 0 {
			return false
		}

		cardCounter[v.Number]++
	}

	if len(cardCounter) == 1 {
		return true
	}

	return false
}

func IsFullHouse(cards CardList) bool {
	if len(cards) != 5 {
		return false
	}

	cardCounter := make(map[int]int)
	for _, v := range cards {
		if v.CardType == CardTypePhoenix {
			continue
		}

		if v.Number == 0 {
			return false
		}

		cardCounter[v.Number]++

		if cardCounter[v.Number] > 3 {
			return false
		}
	}

	if len(cardCounter) != 2 {
		return false
	}

	return true
}

//
//func IsStrait(cards CardList) bool {
//	if len(cards) < 5 {
//		return false
//	}
//
//	sort.Sort(cards)
//	for i := 0; i < len(cards)-1; i++ {
//		if cards[i+1].Number-cards[i].Number != 1 {
//			return false
//		}
//	}
//
//	return true
//}
//
//func IsStraitPair(cards CardList) bool {
//	if len(cards) < 4 && len(cards)%2 != 0 {
//		return false
//	}
//
//	cardCounter := make(map[int]int)
//	var numbers []int
//	for _, v := range cards {
//		cardCounter[v.Number]++
//		if cardCounter[v.Number] == 2 {
//			numbers = append(numbers, v.Number)
//		} else if cardCounter[v.Number] > 2 {
//			return false
//		}
//	}
//
//	if len(numbers)*2 != len(cards) {
//		return false
//	}
//
//	sort.Ints(numbers)
//	for i := 0; i < len(numbers)-1; i++ {
//		if numbers[i+1]-numbers[i] != 1 {
//			return false
//		}
//	}
//
//	return true
//}
