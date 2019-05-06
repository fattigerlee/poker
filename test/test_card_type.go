package main

import (
	"fmt"
	"github.com/fattigerlee/poker/ddz"
)

func main() {
	isDan()
	isDui()
	isLianDui()
	isSanBuDai()
	isSanDaiYi()
	isSanDaiEr()
	isSiDaiDan()
	isSiDaiDui()
	isShunZi()
	isFeiJiBuDai()
	isFeiJiDaiYi()
	isFeiJiDaiEr()
	isZhaDan()
	isHuoJian()
	isLianZha()
}

func isDan() {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 16),
	}
	fmt.Println("单:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 17),
	}
	fmt.Println("单:", ddz.GetCardType(cards))
}

func isDui() {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
	}
	fmt.Println("对:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
	}
	fmt.Println("对:", ddz.GetCardType(cards))
}

func isLianDui() {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
	}
	fmt.Println("连对:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 12),
		ddz.NewCard(ddz.SuitTypeDiamond, 12),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeDiamond, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
	}
	fmt.Println("连对:", ddz.GetCardType(cards))
}

func isSanBuDai() {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
	}
	fmt.Println("三不带:", ddz.GetCardType(cards))
}

func isSanDaiYi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeClub, 5),
	}
	fmt.Println("三带一:", ddz.GetCardType(cards))
}

func isSanDaiEr() {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
	}
	fmt.Println("三带二:", ddz.GetCardType(cards))
}

func isSiDaiDan() {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards))
}

func isSiDaiDui() {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
	}
	fmt.Println("四带对:", ddz.GetCardType(cards))
}

func isShunZi() {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 11),
		ddz.NewCard(ddz.SuitTypeClub, 12),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
	}
	fmt.Println("顺子:", ddz.GetCardType(cards))
}

func isFeiJiBuDai() {
	var cards []*ddz.Card

	cards = []*ddz.Card{
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
	fmt.Println("飞机不带:", ddz.GetCardType(cards))
}

func isFeiJiDaiYi() {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
	}
	fmt.Println("飞机带一:", ddz.GetCardType(cards))
}

func isFeiJiDaiEr() {
	var cards []*ddz.Card

	cards = []*ddz.Card{
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
	fmt.Println("飞机带二:", ddz.GetCardType(cards))
}

func isZhaDan() {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
	}
	fmt.Println("炸弹:", ddz.GetCardType(cards))
}

func isLianZha() {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
	}
	fmt.Println("连炸:", ddz.GetCardType(cards))
}

func isHuoJian() {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 17),
		ddz.NewCard(ddz.SuitTypeJoker, 16),
	}
	fmt.Println("火箭:", ddz.GetCardType(cards))
}
