package main

import (
	"fmt"
	"github.com/fattigerlee/poker/ddz"
)

func main() {
	overcomeSanDaiEr()

	overcomeLianZha()
}

func overcomeSanDaiEr() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeSanDaiEr,
		MinValue: 8,
	}

	fmt.Println(ddz.OvercomeCard(info, cards))
}

func overcomeLianZha() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeLianZha,
		MinValue: 12,
		MaxValue: 13,
	}

	fmt.Println("超过连炸:", ddz.OvercomeCard(info, cards))
}
