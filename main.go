package main

import (
	"fmt"
	"github.com/fattigerlee/Poker/ddz"
)

func main() {
	typeTest()
	findTest()
}

func typeTest() {
	isDan()
	isHuoJian()
	isLianDui()
	isShunZi()
	isFeiJiBuDai()
	isFeiJiDaiYi()
	isFeiJiDaiEr()
	isLianZha()
}

func isDan() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 16),
	}
	fmt.Println(ddz.GetType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 17),
	}
	fmt.Println(ddz.GetType(cards))
}

func isDui() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
	}
	fmt.Println("对:", ddz.IsDui(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
	}
	fmt.Println("对:", ddz.IsDui(cards))
}

func isHuoJian() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 17),
		ddz.NewCard(ddz.SuitTypeJoker, 16),
	}
	fmt.Println(ddz.GetType(cards))
}

func isLianDui() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
	}
	fmt.Println("连对:", ddz.IsLianDui(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 12),
		ddz.NewCard(ddz.SuitTypeDiamond, 12),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeDiamond, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
	}
	fmt.Println("连对:", ddz.IsLianDui(cards))
}

func isSanBuDai() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
	}
	fmt.Println("三不带:", ddz.IsSanBuDai(cards))
}

func isSanDaiYi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeClub, 5),
	}
	fmt.Println("三带一:", ddz.IsSanDaiYi(cards))
}

func isSanDaiEr() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
	}
	fmt.Println("三带二:", ddz.IsSanDaiEr(cards))
}

func isSiDaiDan() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
	}
	fmt.Println("四带单:", ddz.IsSiDaiDan(cards))
}

func isSiDaiDui() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
	}
	fmt.Println("四带对:", ddz.IsSiDaiDui(cards))
}

func isShunZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 11),
		ddz.NewCard(ddz.SuitTypeClub, 12),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
	}
	fmt.Println("顺子:", ddz.IsShunZi(cards))
}

func isFeiJiBuDai() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 12),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
		ddz.NewCard(ddz.SuitTypeClub, 12),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
	}
	fmt.Println("飞机不带:", ddz.IsFeiJiBuDai(cards))
}

func isFeiJiDaiYi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
	}
	fmt.Println("飞机带一:", ddz.IsFeiJiDaiYi(cards))
}

func isFeiJiDaiEr() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeClub, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
	}
	fmt.Println("飞机带二:", ddz.IsFeiJiDaiEr(cards))
}

func isZhaDan() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
	}
	fmt.Println("炸弹:", ddz.IsZhaDan(cards))
}

func isLianZha() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
	}
	fmt.Println("连炸:", ddz.IsLianZha(cards))
}

func findTest() {
	findDan()
	findDui()
	findLianDui()
	findShunZi()
	findFeiJiBuDai()
	findFeiJiDaiYi()
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
