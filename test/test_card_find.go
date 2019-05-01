package main

import (
	"fmt"
	"github.com/fattigerlee/poker/ddz"
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
	findLianZha()
}

func findDan() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 10, false),
		ddz.NewCard(ddz.SuitTypeSpade, 10, false),
		ddz.NewCard(ddz.SuitTypeClub, 7, false),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeDan,
		MinValue: 15,
	}

	fmt.Println("单:", ddz.FindCards(info, cards))

	info = ddz.CardTypeInfo{
		CardType: ddz.CardTypeDan,
		MinValue: 3,
	}

	fmt.Println("单:", ddz.FindCards(info, cards))

	info = ddz.CardTypeInfo{
		CardType: ddz.CardTypeDan,
		MinValue: 8,
	}

	fmt.Println("单:", ddz.FindCards(info, cards))
}

func findDui() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 7, false),
		ddz.NewCard(ddz.SuitTypeSpade, 7, false),
		ddz.NewCard(ddz.SuitTypeHeart, 6, false),
		ddz.NewCard(ddz.SuitTypeSpade, 6, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeSpade, 4, false),
		ddz.NewCard(ddz.SuitTypeHeart, 5, false),
		ddz.NewCard(ddz.SuitTypeSpade, 5, false),
		ddz.NewCard(ddz.SuitTypeClub, 5, false),
		ddz.NewCard(ddz.SuitTypeSpade, 5, false),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeDui,
		MinValue: 3,
	}

	fmt.Println("对:", ddz.FindCards(info, cards))

	info = ddz.CardTypeInfo{
		CardType: ddz.CardTypeDui,
		MinValue: 6,
	}

	fmt.Println("对:", ddz.FindCards(info, cards))

	info = ddz.CardTypeInfo{
		CardType: ddz.CardTypeDui,
		MinValue: 8,
	}

	fmt.Println("对:", ddz.FindCards(info, cards))
}

func findLianDui() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 7, false),
		ddz.NewCard(ddz.SuitTypeSpade, 7, false),
		ddz.NewCard(ddz.SuitTypeHeart, 6, false),
		ddz.NewCard(ddz.SuitTypeSpade, 6, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeSpade, 4, false),
		ddz.NewCard(ddz.SuitTypeHeart, 5, false),
		ddz.NewCard(ddz.SuitTypeSpade, 5, false),
		ddz.NewCard(ddz.SuitTypeClub, 5, false),
		ddz.NewCard(ddz.SuitTypeSpade, 5, false),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeLianDui,
		MinValue: 3,
		MaxValue: 5,
	}

	fmt.Println("连对:", ddz.FindCards(info, cards))

	info = ddz.CardTypeInfo{
		CardType: ddz.CardTypeLianDui,
		MinValue: 4,
		MaxValue: 6,
	}

	fmt.Println("连对:", ddz.FindCards(info, cards))

	info = ddz.CardTypeInfo{
		CardType: ddz.CardTypeLianDui,
		MinValue: 5,
		MaxValue: 7,
	}

	fmt.Println("连对:", ddz.FindCards(info, cards))
}

func findSanBuDai() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 4, false),
		ddz.NewCard(ddz.SuitTypeSpade, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeSpade, 4, false),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeSanBuDai,
		MinValue: 3,
	}

	fmt.Println("三不带:", ddz.FindCards(info, cards))
}

func findSanDaiYi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 17, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
		ddz.NewCard(ddz.SuitTypeClub, 13, false),
		ddz.NewCard(ddz.SuitTypeClub, 10, false),
		ddz.NewCard(ddz.SuitTypeHeart, 13, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, false),
		ddz.NewCard(ddz.SuitTypeHeart, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 10, false),
		ddz.NewCard(ddz.SuitTypeClub, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 8, false),
		ddz.NewCard(ddz.SuitTypeHeart, 10, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, false),
		ddz.NewCard(ddz.SuitTypeSpade, 9, false),
		ddz.NewCard(ddz.SuitTypeHeart, 9, false),
		ddz.NewCard(ddz.SuitTypeClub, 9, false),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeSanDaiYi,
		MinValue: 3,
	}

	fmt.Println("三带一:", ddz.FindCards(info, cards))
}

func findSanDaiEr() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeClub, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeHeart, 6, false),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeSanDaiEr,
		MinValue: 8,
	}

	fmt.Println("三带二:", ddz.FindCards(info, cards))
}

func findShunZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 5, false),
		ddz.NewCard(ddz.SuitTypeSpade, 7, false),
		ddz.NewCard(ddz.SuitTypeHeart, 6, false),
		ddz.NewCard(ddz.SuitTypeSpade, 8, false),
		ddz.NewCard(ddz.SuitTypeClub, 9, false),
		ddz.NewCard(ddz.SuitTypeSpade, 10, false),
		ddz.NewCard(ddz.SuitTypeHeart, 11, false),
		ddz.NewCard(ddz.SuitTypeSpade, 12, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeShunZi,
		MinValue: 3,
		MaxValue: 13,
	}

	fmt.Println("顺子:", ddz.FindCards(info, cards))

	info = ddz.CardTypeInfo{
		CardType: ddz.CardTypeShunZi,
		MinValue: 6,
		MaxValue: 13,
	}

	fmt.Println("顺子:", ddz.FindCards(info, cards))

	info = ddz.CardTypeInfo{
		CardType: ddz.CardTypeShunZi,
		MinValue: 10,
		MaxValue: 14,
	}

	fmt.Println("顺子:", ddz.FindCards(info, cards))
}

func findFeiJiBuDai() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeClub, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
		ddz.NewCard(ddz.SuitTypeHeart, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
		ddz.NewCard(ddz.SuitTypeClub, 9, false),
		ddz.NewCard(ddz.SuitTypeSpade, 10, false),
		ddz.NewCard(ddz.SuitTypeHeart, 11, false),
		ddz.NewCard(ddz.SuitTypeSpade, 12, false),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeFeiJiBuDai,
		MinValue: 12,
		MaxValue: 13,
	}

	fmt.Println("飞机不带:", ddz.FindCards(info, cards))
}

func findFeiJiDaiYi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeClub, 13, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 13, false),
		ddz.NewCard(ddz.SuitTypeHeart, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
		ddz.NewCard(ddz.SuitTypeClub, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 10, false),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeFeiJiDaiYi,
		MinValue: 12,
		MaxValue: 13,
	}

	fmt.Println("飞机带一:", ddz.FindCards(info, cards))
}

func findFeiJiDaiEr() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeClub, 13, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 13, false),
		ddz.NewCard(ddz.SuitTypeHeart, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
		ddz.NewCard(ddz.SuitTypeClub, 14, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 12, false),
		ddz.NewCard(ddz.SuitTypeClub, 10, false),
		ddz.NewCard(ddz.SuitTypeSpade, 15, false),
		ddz.NewCard(ddz.SuitTypeClub, 15, false),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeFeiJiDaiEr,
		MinValue: 12,
		MaxValue: 13,
	}

	fmt.Println("飞机带二:", ddz.FindCards(info, cards))
}

func findZhaDan() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 10, false),
		ddz.NewCard(ddz.SuitTypeSpade, 10, false),
		ddz.NewCard(ddz.SuitTypeClub, 10, false),
		ddz.NewCard(ddz.SuitTypeHeart, 10, false),
		ddz.NewCard(ddz.SuitTypeHeart, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeClub, 13, false),
		ddz.NewCard(ddz.SuitTypeHeart, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
		ddz.NewCard(ddz.SuitTypeClub, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 12, false),
		ddz.NewCard(ddz.SuitTypeClub, 12, false),
		ddz.NewCard(ddz.SuitTypeJoker, 16, false),
		ddz.NewCard(ddz.SuitTypeJoker, 17, false),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeZhaDan,
		MinValue: 12,
		MaxValue: 0,
	}

	fmt.Println("炸弹:", ddz.FindCards(info, cards))

	info = ddz.CardTypeInfo{
		CardType: ddz.CardTypeZhaDan,
		MinValue: 14,
		MaxValue: 0,
	}

	fmt.Println("炸弹:", ddz.FindCards(info, cards))
}

func findHuoJian() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeClub, 13, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 12, false),
		ddz.NewCard(ddz.SuitTypeHeart, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
		ddz.NewCard(ddz.SuitTypeClub, 11, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 12, false),
		ddz.NewCard(ddz.SuitTypeClub, 12, false),
		ddz.NewCard(ddz.SuitTypeSpade, 11, false),
		ddz.NewCard(ddz.SuitTypeClub, 11, false),
		ddz.NewCard(ddz.SuitTypeJoker, 16, false),
		ddz.NewCard(ddz.SuitTypeJoker, 17, false),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeHuoJian,
		MinValue: 0,
		MaxValue: 0,
	}

	fmt.Println("火箭:", ddz.FindCards(info, cards))
}

func findLianZha() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 7, false),
		ddz.NewCard(ddz.SuitTypeSpade, 7, false),
		ddz.NewCard(ddz.SuitTypeClub, 7, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 7, false),
		ddz.NewCard(ddz.SuitTypeHeart, 8, false),
		ddz.NewCard(ddz.SuitTypeSpade, 8, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, false),
		ddz.NewCard(ddz.SuitTypeSpade, 9, false),
		ddz.NewCard(ddz.SuitTypeClub, 9, false),
		ddz.NewCard(ddz.SuitTypeSpade, 9, false),
		ddz.NewCard(ddz.SuitTypeClub, 9, false),
		ddz.NewCard(ddz.SuitTypeHeart, 10, false),
		ddz.NewCard(ddz.SuitTypeSpade, 10, false),
		ddz.NewCard(ddz.SuitTypeClub, 10, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 10, false),
	}

	info := ddz.CardTypeInfo{
		CardType: ddz.CardTypeZhaDan,
		MinValue: 12,
		MaxValue: 0,
	}

	fmt.Println("连炸:", ddz.FindCards(info, cards))

	info = ddz.CardTypeInfo{
		CardType: ddz.CardTypeLianZha,
		MinValue: 13,
		MaxValue: 14,
	}

	fmt.Println("连炸:", ddz.FindCards(info, cards))
}
