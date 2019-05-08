package main

import (
	"fmt"
	"github.com/fattigerlee/poker/ddz"
)

func main() {
	findDan()
	findDui()

	findSanBuDai()
	findSanDaiYi()
	findSanDaiEr()

	findSiDaiDan()
	findSiDaiDui()

	findShunZi()
	findLianDui()
	findFeiJiBuDai()

	findFeiJiDaiYi()
	findFeiJiDaiEr()

	findZhaDan()
	findHuoJian()
	findLianZha()
}

func findDan() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeClub, 7),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeDan,
		MinValue: 15,
	}

	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("单:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeDan,
		MinValue: 3,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("单:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeDan,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("单:", retCards, retInfo)
}

func findDui() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
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

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeDui,
		MinValue: 3,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("对:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeDui,
		MinValue: 6,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("对:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeDui,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("对:", retCards, retInfo)
}

func findSanBuDai() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeSanBuDai,
		MinValue: 3,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("三不带:", retCards, retInfo)
}

func findSanDaiYi() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 17),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeSanDaiYi,
		MinValue: 3,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("三带一:", retCards, retInfo)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeSanDaiYi,
		MinValue: 3,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("三带一:", retCards, retInfo)
}

func findSanDaiEr() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
		ddz.NewCard(ddz.SuitTypeClub, 12),
		ddz.NewCard(ddz.SuitTypeHeart, 12),
		ddz.NewCard(ddz.SuitTypeDiamond, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeSanDaiEr,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("三带二:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeSanDaiEr,
		MinValue: 14,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("三带二:", retCards, retInfo)
}

func findSiDaiDan() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
		ddz.NewCard(ddz.SuitTypeClub, 12),
		ddz.NewCard(ddz.SuitTypeHeart, 12),
		ddz.NewCard(ddz.SuitTypeDiamond, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeSiDaiDan,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("四带单:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeSiDaiDan,
		MinValue: 14,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("四带单:", retCards, retInfo)
}

func findSiDaiDui() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
		ddz.NewCard(ddz.SuitTypeClub, 12),
		ddz.NewCard(ddz.SuitTypeHeart, 12),
		ddz.NewCard(ddz.SuitTypeDiamond, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeSiDaiDui,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("四带对:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeSiDaiDui,
		MinValue: 14,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("四带对:", retCards, retInfo)
}

func findShunZi() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
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
		ddz.NewCard(ddz.SuitTypeJoker, 16),
		ddz.NewCard(ddz.SuitTypeJoker, 17),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeShunZi,
		MinValue: 3,
		MaxValue: 13,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("顺子:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeShunZi,
		MinValue: 6,
		MaxValue: 13,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("顺子:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeShunZi,
		MinValue: 10,
		MaxValue: 14,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("顺子:", retCards, retInfo)
}

func findLianDui() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeLianDui,
		MinValue: 3,
		MaxValue: 5,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("连对:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeLianDui,
		MinValue: 4,
		MaxValue: 6,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("连对:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeLianDui,
		MinValue: 5,
		MaxValue: 8,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("连对:", retCards, retInfo)
}

func findFeiJiBuDai() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
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

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeFeiJiBuDai,
		MinValue: 12,
		MaxValue: 13,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("飞机不带:", retCards, retInfo)
}

func findFeiJiDaiYi() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeSpade, 11),
		ddz.NewCard(ddz.SuitTypeClub, 11),
		ddz.NewCard(ddz.SuitTypeHeart, 12),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
		ddz.NewCard(ddz.SuitTypeClub, 12),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeDiamond, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeFeiJiDaiYi,
		MinValue: 12,
		MaxValue: 13,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("飞机带一:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeFeiJiDaiYi,
		MinValue: 9,
		MaxValue: 10,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("飞机带一:", retCards, retInfo)
}

func findFeiJiDaiEr() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeDiamond, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeDiamond, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 15),
		ddz.NewCard(ddz.SuitTypeClub, 15),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeFeiJiDaiEr,
		MinValue: 7,
		MaxValue: 8,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("飞机带二:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeFeiJiDaiEr,
		MinValue: 7,
		MaxValue: 9,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("飞机带二:", retCards, retInfo)
}

func findZhaDan() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
		ddz.NewCard(ddz.SuitTypeClub, 12),
		ddz.NewCard(ddz.SuitTypeJoker, 16),
		ddz.NewCard(ddz.SuitTypeJoker, 17),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeZhaDan,
		MinValue: 12,
		MaxValue: 0,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("炸弹:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeZhaDan,
		MinValue: 14,
		MaxValue: 0,
	}
	retCards, retInfo = ddz.FindCardsJingDian(info, cards)
	fmt.Println("炸弹:", retCards, retInfo)
}

func findHuoJian() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeDiamond, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
		ddz.NewCard(ddz.SuitTypeClub, 14),
		ddz.NewCard(ddz.SuitTypeDiamond, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
		ddz.NewCard(ddz.SuitTypeClub, 12),
		ddz.NewCard(ddz.SuitTypeSpade, 11),
		ddz.NewCard(ddz.SuitTypeClub, 11),
		ddz.NewCard(ddz.SuitTypeJoker, 16),
		ddz.NewCard(ddz.SuitTypeJoker, 17),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeHuoJian,
		MinValue: 0,
		MaxValue: 0,
	}
	retCards, retInfo = ddz.FindCardsBuXiPai(info, cards)
	fmt.Println("火箭:", retCards, retInfo)
}

func findLianZha() {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeDiamond, 10),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeZhaDan,
		MinValue: 12,
		MaxValue: 0,
	}
	retCards, retInfo = ddz.FindCardsBuXiPai(info, cards)
	fmt.Println("连炸:", retCards, retInfo)

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeLianZha,
		MinValue: 13,
		MaxValue: 14,
	}
	retCards, retInfo = ddz.FindCardsBuXiPai(info, cards)
	fmt.Println("连炸:", retCards, retInfo)
}
