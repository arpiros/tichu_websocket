package models

const (
	CardColer_None  = iota
	CardColor_Red
	CardColor_Blue
	CardColor_Green
	CardColor_Block
)

const (
	CardType_None    = iota
	CardType_SPARROW
	CardType_Dog
	CardType_Dragon
	CardType_Phoenix
)

type Card struct {
	Number int
	Color  int
	cType  int
}

func NewCardDeck() []*Card {
	return nil
}