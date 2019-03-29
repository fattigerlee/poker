package main

import (
	"fmt"
	"github.com/fattigerlee/Poker/ddz"
)

func main() {
	// 对
	isDui()

	// 连对
	isLianDui()

	// 三不带
	isSanBuDai()

	// 三带一
	isSanDaiYi()

	// 三带二
	isSanDaiEr()

	// 四带单
	isSiDaiDan()

	// 四带对
	isSiDaiDui()

	// 顺子
	isShunZi()

	// 飞机不带
	isFeiJiBuDai()

	// 飞机带一
	isFeiJiDaiYi()

	// 飞机带二
	isFeiJiDaiEr()

	// 炸弹
	isZhaDan()

	// 火箭
	isHuoJian()

	// 连炸
	isLianZha()
}

func isDui() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
	}
	fmt.Println("对:", ddz.IsDui(cards))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
	}
	fmt.Println("对:", ddz.IsDui(cards))
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
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 6),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeClub, 6),
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
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
	}
	fmt.Println("顺子:", ddz.IsShunZi(cards))
}

func isFeiJiBuDai() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
	}
	fmt.Println("飞机不带:", ddz.IsFeiJiBuDai(cards))
}

func isFeiJiDaiYi() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
	}
	fmt.Println("飞机带一:", ddz.IsFeiJiDaiYi(cards))
}

func isFeiJiDaiEr() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
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

func isHuoJian() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 16),
		ddz.NewCard(ddz.SuitTypeJoker, 17),
	}
	fmt.Println("火箭:", ddz.IsHuoJian(cards))
}

func isLianZha() {
	cards := []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
	}
	fmt.Println("连炸:", ddz.IsLianZha(cards))
}
