package main

import (
	"fmt"
	"github.com/fattigerlee/poker/ddz"
)

func main() {
	splitJingDian()
}

func splitJingDian() {
	var cards []*ddz.Card
	var infoList []ddz.CardTypeInfo
	var cardsList [][]*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 16),
		ddz.NewCard(ddz.SuitTypeJoker, 17),
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式):", cardsList, infoList)
}
