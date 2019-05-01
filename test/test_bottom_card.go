package main

import (
	"fmt"
	"github.com/fattigerlee/poker/ddz"
)

func main() {
	huoJian()
	oneJoker()
	tongHuaShun()
	tongHua()
	shunZi()
	sanBuDai()
}

func huoJian() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 16, false),
		ddz.NewCard(ddz.SuitTypeJoker, 17, false),
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
	}
	fmt.Println("火箭:", ddz.GetBottomCardType(cards))
}

func oneJoker() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 16, false),
		ddz.NewCard(ddz.SuitTypeSpade, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
	}
	fmt.Println("单王:", ddz.GetBottomCardType(cards))
}

func tongHuaShun() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 4, false),
		ddz.NewCard(ddz.SuitTypeHeart, 5, false),
	}
	fmt.Println("同花顺:", ddz.GetBottomCardType(cards))
}

func tongHua() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 4, false),
		ddz.NewCard(ddz.SuitTypeHeart, 7, false),
	}
	fmt.Println("同花:", ddz.GetBottomCardType(cards))
}

func shunZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3, false),
		ddz.NewCard(ddz.SuitTypeHeart, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 5, false),
	}
	fmt.Println("顺子:", ddz.GetBottomCardType(cards))
}

func sanBuDai() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 4, false),
		ddz.NewCard(ddz.SuitTypeHeart, 4, false),
		ddz.NewCard(ddz.SuitTypeDiamond, 4, false),
	}
	fmt.Println("三不带:", ddz.GetBottomCardType(cards))
}
