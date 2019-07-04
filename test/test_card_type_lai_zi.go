package main

import (
	"fmt"
	"github.com/fattigerlee/poker/ddz"
)

func main() {
	laiZiNums := []ddz.NumType{16, 17}

	isDanLaiZi(laiZiNums)
	isDuiLaiZi(laiZiNums)

	isSanBuDaiLaiZi(laiZiNums)
	isSanDaiYiLaiZi(laiZiNums)
	isSanDaiErLaiZi(laiZiNums)

	isSiDaiDanLaiZi(laiZiNums)
	isSiDaiDuiLaiZi(laiZiNums)

	isShunZiLaiZi(laiZiNums)
	isLianDuiLaiZi(laiZiNums)

	isFeiJiBuDaiLaiZi(laiZiNums)
	isFeiJiDaiYiLaiZi(laiZiNums)
	isFeiJiDaiErLaiZi(laiZiNums)

	isLaiZiHuoJian(laiZiNums)

	isLaiZiZhaDan4(laiZiNums)
	isChunLaiZiZhaDan(laiZiNums)

	isRuanZhaDan4(laiZiNums)
	isRuanZhaDan5(laiZiNums)
	isRuanLianZhaDan(laiZiNums)
}

func isDanLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3),
	}
	fmt.Println("单:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 9),
	}
	fmt.Println("单:", ddz.GetCardType(cards))
}

func isDuiLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
	}
	fmt.Println("对:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
	}
	fmt.Println("对:", ddz.GetCardType(cards, laiZiNums...))
}

func isSanBuDaiLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("三不带:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("三不带:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("三不带:", ddz.GetCardType(cards, laiZiNums...))
}

func isSanDaiYiLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("三带一:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("三带一:", ddz.GetCardType(cards, laiZiNums...))
}

func isSanDaiErLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("三带二:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("三带二:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("三带二:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
	}
	fmt.Println("三带二:", ddz.GetCardType(cards, laiZiNums...))
}

func isSiDaiDanLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 16),
		ddz.NewCard(ddz.SuitTypeJoker, 17),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("四带单:", ddz.GetCardType(cards, laiZiNums...))
}

func isSiDaiDuiLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 6),
	}
	fmt.Println("四带对:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 6),
	}
	fmt.Println("四带对:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
	}
	fmt.Println("四带对:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeJoker, 16),
	}
	fmt.Println("四带对:", ddz.GetCardType(cards, laiZiNums...))
}

func isShunZiLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeClub, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
	}
	fmt.Println("顺子:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 10),
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
	}
	fmt.Println("顺子:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 11),
		ddz.NewCard(ddz.SuitTypeClub, 12),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 10),
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
	}
	fmt.Println("顺子:", ddz.GetCardType(cards, laiZiNums...))
}

func isLianDuiLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeClub, 8),
	}
	fmt.Println("连对:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
	}
	fmt.Println("连对:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeClub, 8),
	}
	fmt.Println("连对:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 10),
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
	}
	fmt.Println("连对:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
	}
	fmt.Println("连对:", ddz.GetCardType(cards, laiZiNums...))
}

func isFeiJiBuDaiLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
	}
	fmt.Println("飞机不带:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeClub, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 10),
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
	}
	fmt.Println("飞机不带:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 10),
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
	}
	fmt.Println("飞机不带:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
	}
	fmt.Println("飞机不带:", ddz.GetCardType(cards, laiZiNums...))
}

func isFeiJiDaiYiLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeClub, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
	}
	fmt.Println("飞机带一:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeClub, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 17),
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
	}
	fmt.Println("飞机带一:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeDiamond, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
	}
	fmt.Println("飞机带一:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeDiamond, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
	}
	fmt.Println("飞机带一:", ddz.GetCardType(cards, laiZiNums...))
}

func isFeiJiDaiErLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeDiamond, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
	}
	fmt.Println("飞机带二:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeDiamond, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
	}
	fmt.Println("飞机带二:", ddz.GetCardType(cards, laiZiNums...))
}

func isLaiZiHuoJian(laiZiNums []ddz.NumType) {
	laiZiNums = []ddz.NumType{16, 17}

	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 17),
		ddz.NewCard(ddz.SuitTypeJoker, 16),
	}
	fmt.Println("火箭:", ddz.GetCardType(cards, laiZiNums...))
}

func isLaiZiZhaDan4(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeClub, 10),
	}
	fmt.Println("四癞子炸:", ddz.GetCardType(cards, laiZiNums...))
}

func isChunLaiZiZhaDan(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
	}
	fmt.Println("四纯癞子炸:", ddz.GetCardType(cards, laiZiNums...))
}

func isRuanZhaDan4(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 8),
	}
	fmt.Println("四软炸:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeClub, 8),
	}
	fmt.Println("四软炸:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeClub, 8),
	}
	fmt.Println("四软炸:", ddz.GetCardType(cards, laiZiNums...))
}

func isRuanZhaDan5(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeClub, 8),
	}
	fmt.Println("五软炸:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeClub, 8),
	}
	fmt.Println("五软炸:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeClub, 8),
	}
	fmt.Println("五软炸:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeClub, 8),
	}
	fmt.Println("五软炸:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeDiamond, 11),
		ddz.NewCard(ddz.SuitTypeSpade, 11),
		ddz.NewCard(ddz.SuitTypeJoker, 17),
		ddz.NewCard(ddz.SuitTypeJoker, 16),
	}
	fmt.Println("五软炸:", ddz.GetCardType(cards, laiZiNums...))
}

func isRuanLianZhaDan(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeJoker, 16),
		ddz.NewCard(ddz.SuitTypeJoker, 17),
	}
	fmt.Println("软连炸:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeJoker, 16),
		ddz.NewCard(ddz.SuitTypeJoker, 17),
	}
	fmt.Println("软连炸:", ddz.GetCardType(cards, laiZiNums...))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeDiamond, 14),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeClub, 15),
		ddz.NewCard(ddz.SuitTypeSpade, 15),
		ddz.NewCard(ddz.SuitTypeHeart, 15),
		ddz.NewCard(ddz.SuitTypeJoker, 16),
	}
	fmt.Println("软连炸:", ddz.GetCardType(cards, laiZiNums...))
}
