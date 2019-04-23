package main

import (
	"fmt"
	"github.com/fattigerlee/poker/ddz"
)

func main() {
	isDuiLaiZi()
	isLianDuiLaiZi()
	isSanBuDaiLaiZi()
	isRuanZhaDan4()
	isRuanZhaDan5()
}

func isDuiLaiZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 3, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
	}
	fmt.Println("对:", ddz.GetCardType(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 8, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 3, false),
	}
	fmt.Println("对:", ddz.GetCardType(cards))
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
		ddz.NewCard(ddz.SuitTypeDiamond, 3, true),
		ddz.NewCard(ddz.SuitTypeClub, 3, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, true),
		ddz.NewCard(ddz.SuitTypeClub, 4, true),
		ddz.NewCard(ddz.SuitTypeDiamond, 5, true),
		ddz.NewCard(ddz.SuitTypeClub, 5, true),
	}
	fmt.Println("连对:", ddz.GetCardType(cards))
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
