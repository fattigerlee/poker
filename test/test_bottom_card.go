package main

import (
	"fmt"
	"github.com/fattigerlee/poker/ddz"
)

func main() {
	wangZha()
	danWang()
	tongHuaShun()
	tongHua()
	shunZi()
	sanBuDai()
}

func wangZha() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 16),
		ddz.NewCard(ddz.SuitTypeJoker, 17),
		ddz.NewCard(ddz.SuitTypeHeart, 3),
	}
	fmt.Println("王炸:", ddz.GetBottomCardType(cards))
}

func danWang() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 16),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 3),
	}
	fmt.Println("单王:", ddz.GetBottomCardType(cards))
}

func tongHuaShun() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
	}
	fmt.Println("同花顺:", ddz.GetBottomCardType(cards))
}

func tongHua() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
	}
	fmt.Println("同花:", ddz.GetBottomCardType(cards))
}

func shunZi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
	}
	fmt.Println("顺子:", ddz.GetBottomCardType(cards))
}

func sanBuDai() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
	}
	fmt.Println("三不带:", ddz.GetBottomCardType(cards))
}
