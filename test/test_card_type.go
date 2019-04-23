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
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 16, false),
	}
	fmt.Println("单:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 17, false),
	}
	fmt.Println("单:", ddz.GetCardType(cards))
}

func isDui() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
	}
	fmt.Println("对:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
	}
	fmt.Println("对:", ddz.GetCardType(cards))
}

func isLianDui() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 5, false),
		ddz.NewCard(ddz.SuitTypeClub, 5, false),
	}
	fmt.Println("连对:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 12, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 12, false),
		ddz.NewCard(ddz.SuitTypeClub, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 13, false),
		ddz.NewCard(ddz.SuitTypeClub, 13, false),
	}
	fmt.Println("连对:", ddz.GetCardType(cards))
}

func isSanBuDai() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
	}
	fmt.Println("三不带:", ddz.GetCardType(cards))
}

func isSanDaiYi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 5, false),
	}
	fmt.Println("三带一:", ddz.GetCardType(cards))
}

func isSanDaiEr() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 5, false),
		ddz.NewCard(ddz.SuitTypeSpade, 5, false),
	}
	fmt.Println("三带二:", ddz.GetCardType(cards))
}

func isSiDaiDan() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 5, false),
		ddz.NewCard(ddz.SuitTypeSpade, 6, false),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards))
}

func isSiDaiDui() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 5, false),
		ddz.NewCard(ddz.SuitTypeHeart, 5, false),
		ddz.NewCard(ddz.SuitTypeSpade, 7, false),
		ddz.NewCard(ddz.SuitTypeHeart, 7, false),
	}
	fmt.Println("四带对:", ddz.GetCardType(cards))
}

func isShunZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 10, false),
		ddz.NewCard(ddz.SuitTypeSpade, 11, false),
		ddz.NewCard(ddz.SuitTypeClub, 12, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
	}
	fmt.Println("顺子:", ddz.GetCardType(cards))
}

func isFeiJiBuDai() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 12, false),
		ddz.NewCard(ddz.SuitTypeSpade, 12, false),
		ddz.NewCard(ddz.SuitTypeClub, 12, false),
		ddz.NewCard(ddz.SuitTypeHeart, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeClub, 13, false),
		ddz.NewCard(ddz.SuitTypeHeart, 14, false),
		ddz.NewCard(ddz.SuitTypeClub, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
	}
	fmt.Println("飞机不带:", ddz.GetCardType(cards))
}

func isFeiJiDaiYi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
		ddz.NewCard(ddz.SuitTypeClub, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
		ddz.NewCard(ddz.SuitTypeHeart, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeClub, 13, false),
		ddz.NewCard(ddz.SuitTypeHeart, 5, false),
	}
	fmt.Println("飞机带一:", ddz.GetCardType(cards))
}

func isFeiJiDaiEr() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeClub, 13, false),
		ddz.NewCard(ddz.SuitTypeHeart, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
		ddz.NewCard(ddz.SuitTypeClub, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 5, false),
		ddz.NewCard(ddz.SuitTypeHeart, 5, false),
		ddz.NewCard(ddz.SuitTypeClub, 6, false),
		ddz.NewCard(ddz.SuitTypeSpade, 6, false),
	}
	fmt.Println("飞机带二:", ddz.GetCardType(cards))
}

func isZhaDan() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 5, false),
		ddz.NewCard(ddz.SuitTypeSpade, 5, false),
		ddz.NewCard(ddz.SuitTypeClub, 5, false),
		ddz.NewCard(ddz.SuitTypeSpade, 5, false),
	}
	fmt.Println("炸弹:", ddz.GetCardType(cards))
}

func isLianZha() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeClub, 13, false),
		ddz.NewCard(ddz.SuitTypeSpade, 13, false),
		ddz.NewCard(ddz.SuitTypeHeart, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
		ddz.NewCard(ddz.SuitTypeClub, 14, false),
		ddz.NewCard(ddz.SuitTypeSpade, 14, false),
	}
	fmt.Println("连炸:", ddz.GetCardType(cards))
}

func isHuoJian() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 17, false),
		ddz.NewCard(ddz.SuitTypeJoker, 16, false),
	}
	fmt.Println("火箭:", ddz.GetCardType(cards))
}
