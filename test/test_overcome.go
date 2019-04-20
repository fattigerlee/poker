package main

import (
	"fmt"
	"github.com/fattigerlee/poker/ddz"
)

func main() {
	overcomeSanDaiEr()
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
