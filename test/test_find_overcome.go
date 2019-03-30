package main

import (
	"fmt"
	"github.com/fattigerlee/Poker/ddz"
)

func main() {
	findDan()
	findDui()
	findLianDui()
	findSanBuDai()
	findSanDaiYi()
	findSanDaiEr()
	findShunZi()
	findFeiJiBuDai()
	findFeiJiDaiYi()
	findFeiJiDaiEr()
	findZhaDan()
	findHuoJian()
}

func findDan() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
	}

	cardType := ddz.Type{
		CardType: ddz.CardTypeDan,
		MinValue: 6,
	}

	fmt.Println(ddz.FindOvercomeCard(cardType, cards))
}

func findDui() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
	}

	cardType := ddz.Type{
		CardType: ddz.CardTypeDui,
		MinValue: 3,
	}

	fmt.Println(ddz.FindOvercomeCard(cardType, cards))
}

func findLianDui() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
	}

	cardType := ddz.Type{
		CardType: ddz.CardTypeLianDui,
		MinValue: 3,
		MaxValue: 5,
	}

	fmt.Println(ddz.FindOvercomeCard(cardType, cards))
}

func findSanBuDai() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
	}

	cardType := ddz.Type{
		CardType: ddz.CardTypeSanBuDai,
		MinValue: 3,
	}

	fmt.Println(ddz.FindOvercomeCard(cardType, cards))
}

func findSanDaiYi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeClub, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
	}

	cardType := ddz.Type{
		CardType: ddz.CardTypeSanDaiYi,
		MinValue: 3,
	}

	fmt.Println(ddz.FindOvercomeCard(cardType, cards))
}

func findSanDaiEr() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 15),
		ddz.NewCard(ddz.SuitTypeSpade, 15),
		ddz.NewCard(ddz.SuitTypeClub, 15),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
	}

	cardType := ddz.Type{
		CardType: ddz.CardTypeSanDaiEr,
		MinValue: 3,
	}

	fmt.Println(ddz.FindOvercomeCard(cardType, cards))
}

func findShunZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
	}

	cardType := ddz.Type{
		CardType: ddz.CardTypeShunZi,
		MinValue: 3,
		MaxValue: 13,
	}

	fmt.Println(ddz.FindOvercomeCard(cardType, cards))
}

func findFeiJiBuDai() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
	}

	cardType := ddz.Type{
		CardType: ddz.CardTypeFeiJiBuDai,
		MinValue: 12,
		MaxValue: 13,
	}

	fmt.Println(ddz.FindOvercomeCard(cardType, cards))
}

func findFeiJiDaiYi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeDiamond, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
	}

	cardType := ddz.Type{
		CardType: ddz.CardTypeFeiJiDaiYi,
		MinValue: 12,
		MaxValue: 13,
	}

	fmt.Println(ddz.FindOvercomeCard(cardType, cards))
}

func findFeiJiDaiEr() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeDiamond, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeDiamond, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
		ddz.NewCard(ddz.SuitTypeClub, 12),
		ddz.NewCard(ddz.SuitTypeSpade, 15),
		ddz.NewCard(ddz.SuitTypeClub, 15),
	}

	cardType := ddz.Type{
		CardType: ddz.CardTypeFeiJiDaiEr,
		MinValue: 12,
		MaxValue: 13,
	}

	fmt.Println(ddz.FindOvercomeCard(cardType, cards))
}

func findZhaDan() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeDiamond, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeDiamond, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
		ddz.NewCard(ddz.SuitTypeClub, 12),
		ddz.NewCard(ddz.SuitTypeJoker, 16),
		ddz.NewCard(ddz.SuitTypeJoker, 17),
	}

	cardType := ddz.Type{
		CardType: ddz.CardTypeZhaDan,
		MinValue: 12,
		MaxValue: 0,
	}

	fmt.Println(ddz.FindOvercomeCard(cardType, cards))

	cardType = ddz.Type{
		CardType: ddz.CardTypeZhaDan,
		MinValue: 14,
		MaxValue: 0,
	}

	fmt.Println(ddz.FindOvercomeCard(cardType, cards))
}

func findHuoJian() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeDiamond, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeDiamond, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
		ddz.NewCard(ddz.SuitTypeClub, 12),
		ddz.NewCard(ddz.SuitTypeSpade, 11),
		ddz.NewCard(ddz.SuitTypeClub, 11),
	}

	cardType := ddz.Type{
		CardType: ddz.CardTypeHuoJian,
		MinValue: 0,
		MaxValue: 0,
	}

	fmt.Println(ddz.FindOvercomeCard(cardType, cards))
}
