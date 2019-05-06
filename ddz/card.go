package ddz

import "fmt"

// 牌
type Card struct {
	Suit SuitType // 花色
	Num  NumType  // 牌值
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Suit, c.Num)
}

func NewCard(suit SuitType, num NumType) *Card {
	return &Card{
		Suit: suit,
		Num:  num,
	}
}
