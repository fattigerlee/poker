package main

import (
	"fmt"
	"github.com/fattigerlee/poker/ddz"
)

func main() {
	laiZiNums := []ddz.NumType{9, 10}

	findDanLaiZi(laiZiNums)
	findDuiLaiZi(laiZiNums)

	findSanBuDaiLaiZi(laiZiNums)
	findSanDaiYiLaiZi(laiZiNums)

	findShunZiLaiZi(laiZiNums)

	findRuanZhaDan4(laiZiNums)
	findLaiZiZhaDan4(laiZiNums)
	findChunLaiZiZhaDan(laiZiNums)

	findHuoJianLaiZi(laiZiNums)
	findRuanLianZha(laiZiNums)
}

// 单张
func findDanLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeDan,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("单张:", retCards, retInfo)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeDan,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("单张:", retCards, retInfo)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 13),
		ddz.NewCard(ddz.SuitTypeClub, 13),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeDan,
		MinValue: 14,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("单张:", retCards, retInfo)
}

// 对子
func findDuiLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeDui,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("对子:", retCards, retInfo)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeDui,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("对子:", retCards, retInfo)
}

// 三不带
func findSanBuDaiLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeSanBuDai,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("三不带:", retCards, retInfo)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeSanBuDai,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("三不带:", retCards, retInfo)
}

// 三带一
func findSanDaiYiLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeClub, 11),
		ddz.NewCard(ddz.SuitTypeSpade, 11),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeSanDaiYi,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("三带一:", retCards, retInfo)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeClub, 11),
		ddz.NewCard(ddz.SuitTypeSpade, 11),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeSanDaiYi,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("三带一:", retCards, retInfo)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeClub, 11),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeSanDaiYi,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("三带一:", retCards, retInfo)
}

// 顺子
func findShunZiLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeClub, 11),
		ddz.NewCard(ddz.SuitTypeSpade, 11),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeShunZi,
		MinValue: 3,
		MaxValue: 7,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("顺子:", retCards, retInfo)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 11),
		ddz.NewCard(ddz.SuitTypeHeart, 12),
		ddz.NewCard(ddz.SuitTypeClub, 13),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeShunZi,
		MinValue: 3,
		MaxValue: 7,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("顺子:", retCards, retInfo)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeClub, 11),
		ddz.NewCard(ddz.SuitTypeSpade, 11),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeShunZi,
		MinValue: 3,
		MaxValue: 7,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("顺子:", retCards, retInfo)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeClub, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 11),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeShunZi,
		MinValue: 3,
		MaxValue: 7,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("顺子:", retCards, retInfo)
}

// 四软炸
func findRuanZhaDan4(laiZiNums []ddz.NumType) {
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
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeFeiJiBuDai,
		MinValue: 12,
		MaxValue: 14,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("四软炸:", retCards, retInfo)
}

// 四癞子炸
func findLaiZiZhaDan4(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeFeiJiBuDai,
		MinValue: 12,
		MaxValue: 13,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("四癞子炸:", retCards, retInfo)
}

// 四纯癞子炸
func findChunLaiZiZhaDan(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeLaiZiZhaDan4,
		MinValue: 8,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("四纯癞子炸:", retCards, retInfo)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeJoker, 16),
		ddz.NewCard(ddz.SuitTypeJoker, 17),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeZhaDan,
		MinValue: 12,
		MaxValue: 0,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("四纯癞子炸:", retCards, retInfo)
}

// 火箭
func findHuoJianLaiZi(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeHuoJian,
		MinValue: 0,
		MaxValue: 0,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("火箭:", retCards, retInfo)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 6),
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeHuoJian,
		MinValue: 0,
		MaxValue: 0,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("火箭:", retCards, retInfo)
}

// 软连炸
func findRuanLianZha(laiZiNums []ddz.NumType) {
	var cards []*ddz.Card
	var info *ddz.CardTypeInfo
	var retCards []*ddz.Card
	var retInfo ddz.CardTypeInfo

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeRuanLianZha,
		MinValue: 3,
		MaxValue: 4,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("软连炸:", retCards, retInfo)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 6),
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeClub, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
	}

	info = &ddz.CardTypeInfo{
		CardType: ddz.CardTypeRuanLianZha,
		MinValue: 11,
		MaxValue: 12,
	}
	retCards, retInfo = ddz.FindCards(info, cards, laiZiNums...)
	fmt.Println("软连炸:", retCards, retInfo)
}
