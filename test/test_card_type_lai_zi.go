package main

import (
	"fmt"
	"github.com/fattigerlee/poker/ddz"
)

func main() {
	isDanLaiZi()
	isDuiLaiZi()

	isSanBuDaiLaiZi()
	isSanDaiYiLaiZi()
	isSanDaiErLaiZi()

	isSiDaiDanLaiZi()
	isSiDaiDuiLaiZi()

	isShunZiLaiZi()
	isLianDuiLaiZi()

	isFeiJiBuDaiLaiZi()
	isFeiJiDaiYiLaiZi()
	isFeiJiDaiErLaiZi()

	isRuanZhaDan4()
	isRuanZhaDan5()
	isRuanLianZhaDan()
}

func isDanLaiZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
	}
	fmt.Println("单:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3, true),
	}
	fmt.Println("单:", ddz.GetCardType(cards))
}

func isDuiLaiZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, true),
	}
	fmt.Println("对:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, true),
	}
	fmt.Println("对:", ddz.GetCardType(cards))
}

func isSanBuDaiLaiZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, true),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
	}
	fmt.Println("三不带:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, true),
		ddz.NewCard(ddz.SuitTypeSpade, 3, true),
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
	}
	fmt.Println("三不带:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, true),
		ddz.NewCard(ddz.SuitTypeSpade, 3, true),
		ddz.NewCard(ddz.SuitTypeClub, 3, true),
	}
	fmt.Println("三不带:", ddz.GetCardType(cards))
}

func isSanDaiYiLaiZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, false),
	}
	fmt.Println("三带一:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeSpade, 8, true),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
	}
	fmt.Println("三带一:", ddz.GetCardType(cards))
}

func isSanDaiErLaiZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
	}
	fmt.Println("三带二:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 8, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
	}
	fmt.Println("三带二:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, true),
		ddz.NewCard(ddz.SuitTypeHeart, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, true),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
	}
	fmt.Println("三带二:", ddz.GetCardType(cards))
}

func isSiDaiDanLaiZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, false),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, false),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 5, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 4, false),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 5, false),
		ddz.NewCard(ddz.SuitTypeHeart, 4, false),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 5, false),
		ddz.NewCard(ddz.SuitTypeHeart, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 16, false),
		ddz.NewCard(ddz.SuitTypeJoker, 17, false),
		ddz.NewCard(ddz.SuitTypeHeart, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards))
}

func isSiDaiDuiLaiZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeHeart, 8, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, true),
		ddz.NewCard(ddz.SuitTypeHeart, 6, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 6, false),
	}
	fmt.Println("四带对:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeHeart, 8, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, true),
		ddz.NewCard(ddz.SuitTypeHeart, 6, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 6, false),
	}
	fmt.Println("四带对:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 7, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, false),
		ddz.NewCard(ddz.SuitTypeHeart, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeHeart, 10, true),
		ddz.NewCard(ddz.SuitTypeSpade, 10, true),
	}
	fmt.Println("四带对:", ddz.GetCardType(cards))
}

func isShunZiLaiZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 5, false),
		ddz.NewCard(ddz.SuitTypeClub, 6, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, true),
	}
	fmt.Println("顺子:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 7, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeHeart, 9, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 10, true),
		ddz.NewCard(ddz.SuitTypeClub, 10, true),
		ddz.NewCard(ddz.SuitTypeSpade, 10, true),
		ddz.NewCard(ddz.SuitTypeHeart, 10, true),
	}
	fmt.Println("顺子:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 11, false),
		ddz.NewCard(ddz.SuitTypeClub, 12, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeHeart, 9, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 10, true),
		ddz.NewCard(ddz.SuitTypeClub, 10, true),
		ddz.NewCard(ddz.SuitTypeSpade, 10, true),
		ddz.NewCard(ddz.SuitTypeHeart, 10, true),
	}
	fmt.Println("顺子:", ddz.GetCardType(cards))
}

func isLianDuiLaiZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 5, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
	}
	fmt.Println("连对:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 5, true),
		ddz.NewCard(ddz.SuitTypeClub, 5, true),
	}
	fmt.Println("连对:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 5, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
	}
	fmt.Println("连对:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 7, false),
		ddz.NewCard(ddz.SuitTypeClub, 7, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeHeart, 9, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 10, true),
		ddz.NewCard(ddz.SuitTypeClub, 10, true),
		ddz.NewCard(ddz.SuitTypeSpade, 10, true),
		ddz.NewCard(ddz.SuitTypeHeart, 10, true),
	}
	fmt.Println("连对:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 7, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeHeart, 9, true),
	}
	fmt.Println("连对:", ddz.GetCardType(cards))
}

func isFeiJiBuDaiLaiZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeSpade, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
	}
	fmt.Println("飞机不带:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 5, false),
		ddz.NewCard(ddz.SuitTypeClub, 6, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeHeart, 9, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 10, true),
		ddz.NewCard(ddz.SuitTypeClub, 10, true),
		ddz.NewCard(ddz.SuitTypeSpade, 10, true),
		ddz.NewCard(ddz.SuitTypeHeart, 10, true),
	}
	fmt.Println("飞机不带:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 5, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeHeart, 9, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 10, true),
		ddz.NewCard(ddz.SuitTypeClub, 10, true),
		ddz.NewCard(ddz.SuitTypeSpade, 10, true),
		ddz.NewCard(ddz.SuitTypeHeart, 10, true),
	}
	fmt.Println("飞机不带:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 5, false),
		ddz.NewCard(ddz.SuitTypeClub, 5, false),
		ddz.NewCard(ddz.SuitTypeSpade, 5, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeSpade, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
	}
	fmt.Println("飞机不带:", ddz.GetCardType(cards))
}

func isFeiJiDaiYiLaiZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeSpade, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 5, false),
		ddz.NewCard(ddz.SuitTypeClub, 6, false),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
	}
	fmt.Println("飞机带一:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 5, false),
		ddz.NewCard(ddz.SuitTypeSpade, 5, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeSpade, 5, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 6, false),
		ddz.NewCard(ddz.SuitTypeSpade, 6, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 6, false),
		ddz.NewCard(ddz.SuitTypeSpade, 7, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 17, false),
		ddz.NewCard(ddz.SuitTypeClub, 10, true),
		ddz.NewCard(ddz.SuitTypeSpade, 10, true),
		ddz.NewCard(ddz.SuitTypeClub, 10, true),
		ddz.NewCard(ddz.SuitTypeSpade, 10, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
	}
	fmt.Println("飞机带一:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 7, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 7, false),
		ddz.NewCard(ddz.SuitTypeSpade, 8, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, false),
		ddz.NewCard(ddz.SuitTypeClub, 10, true),
		ddz.NewCard(ddz.SuitTypeSpade, 10, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 10, true),
		ddz.NewCard(ddz.SuitTypeHeart, 10, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeHeart, 9, true),
	}
	fmt.Println("飞机带一:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 7, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, false),
		ddz.NewCard(ddz.SuitTypeClub, 10, true),
		ddz.NewCard(ddz.SuitTypeSpade, 10, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 10, true),
		ddz.NewCard(ddz.SuitTypeHeart, 10, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
	}
	fmt.Println("飞机带一:", ddz.GetCardType(cards))
}

func isFeiJiDaiErLaiZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 7, false),
		ddz.NewCard(ddz.SuitTypeSpade, 8, false),
		ddz.NewCard(ddz.SuitTypeClub, 10, true),
		ddz.NewCard(ddz.SuitTypeSpade, 10, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 10, true),
		ddz.NewCard(ddz.SuitTypeHeart, 10, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
		ddz.NewCard(ddz.SuitTypeHeart, 9, true),
	}
	fmt.Println("飞机带二:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 7, false),
		ddz.NewCard(ddz.SuitTypeSpade, 8, false),
		ddz.NewCard(ddz.SuitTypeClub, 10, true),
		ddz.NewCard(ddz.SuitTypeSpade, 10, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 10, true),
		ddz.NewCard(ddz.SuitTypeHeart, 10, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 9, true),
	}
	fmt.Println("飞机带二:", ddz.GetCardType(cards))
}

func isRuanZhaDan4() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
	}
	fmt.Println("四软炸:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 8, true),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
	}
	fmt.Println("四软炸:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, true),
		ddz.NewCard(ddz.SuitTypeSpade, 8, true),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
	}
	fmt.Println("四软炸:", ddz.GetCardType(cards))
}

func isRuanZhaDan5() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
	}
	fmt.Println("五软炸:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
	}
	fmt.Println("五软炸:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
	}
	fmt.Println("五软炸:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 8, true),
		ddz.NewCard(ddz.SuitTypeSpade, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 9, true),
		ddz.NewCard(ddz.SuitTypeClub, 8, true),
	}
	fmt.Println("五软炸:", ddz.GetCardType(cards))
}

func isRuanLianZhaDan() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeSpade, 4, false),
		ddz.NewCard(ddz.SuitTypeJoker, 16, true),
		ddz.NewCard(ddz.SuitTypeJoker, 17, true),
	}
	fmt.Println("软连炸:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 3, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 4, false),
		ddz.NewCard(ddz.SuitTypeClub, 4, false),
		ddz.NewCard(ddz.SuitTypeSpade, 4, false),
		ddz.NewCard(ddz.SuitTypeJoker, 16, true),
		ddz.NewCard(ddz.SuitTypeJoker, 17, true),
	}
	fmt.Println("软连炸:", ddz.GetCardType(cards))
}
