package main

import (
	"fmt"
	"github.com/fattigerlee/poker/ddz"
)

func main() {
	splitJingDian()
	splitBuXiPai()
}

func splitJingDian() {
	var cards []*ddz.Card
	var infoList []ddz.CardTypeInfo
	var cardsList [][]*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeJoker, 16),
		ddz.NewCard(ddz.SuitTypeJoker, 17),
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeDiamond, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)1:", cardsList, infoList)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)2:", cardsList, infoList)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeHeart, 12),
		ddz.NewCard(ddz.SuitTypeDiamond, 13),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)3:", cardsList, infoList)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeHeart, 12),
		ddz.NewCard(ddz.SuitTypeDiamond, 13),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)4:", cardsList, infoList)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeHeart, 12),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeDiamond, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)5:", cardsList, infoList)
	fmt.Println("手数(经典模式)5:", ddz.GetMinCount(infoList))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeClub, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeClub, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeDiamond, 11),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeDiamond, 15),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)6:", cardsList, infoList)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeDiamond, 11),
		ddz.NewCard(ddz.SuitTypeSpade, 12),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)7:", cardsList, infoList)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)8:", cardsList, infoList)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)9:", cardsList, infoList)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeDiamond, 12),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)10:", cardsList, infoList)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeHeart, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeDiamond, 12),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)11:", cardsList, infoList)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 11),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeDiamond, 12),
		ddz.NewCard(ddz.SuitTypeHeart, 12),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)12:", cardsList, infoList)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 7),
		ddz.NewCard(ddz.SuitTypeSpade, 8),
		ddz.NewCard(ddz.SuitTypeHeart, 8),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeHeart, 12),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)13:", cardsList, infoList)

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeDiamond, 9),
		ddz.NewCard(ddz.SuitTypeSpade, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeHeart, 12),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)14:", cardsList, infoList)
	fmt.Println("手数(经典模式)14:", ddz.GetMinCount(infoList))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 9),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeHeart, 11),
		ddz.NewCard(ddz.SuitTypeHeart, 12),
		ddz.NewCard(ddz.SuitTypeHeart, 13),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeDiamond, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)15:", cardsList, infoList)
	fmt.Println("手数(经典模式)15:", ddz.GetMinCount(infoList))

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeDiamond, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 8),
		ddz.NewCard(ddz.SuitTypeHeart, 14),
		ddz.NewCard(ddz.SuitTypeDiamond, 14),
		ddz.NewCard(ddz.SuitTypeSpade, 14),
	}

	cardsList, infoList = ddz.SplitCardsJingDian(cards)
	fmt.Println("拆牌(经典模式)16:", cardsList, infoList)
	fmt.Println("手数(经典模式)16:", ddz.GetMinCount(infoList))
}

func splitBuXiPai() {
	var cards []*ddz.Card
	var infoList []ddz.CardTypeInfo
	var cardsList [][]*ddz.Card

	cards = []*ddz.Card{
		ddz.NewCard(ddz.SuitTypeClub, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 3),
		ddz.NewCard(ddz.SuitTypeDiamond, 3),
		ddz.NewCard(ddz.SuitTypeSpade, 3),
		ddz.NewCard(ddz.SuitTypeHeart, 4),
		ddz.NewCard(ddz.SuitTypeDiamond, 4),
		ddz.NewCard(ddz.SuitTypeClub, 4),
		ddz.NewCard(ddz.SuitTypeSpade, 4),
		ddz.NewCard(ddz.SuitTypeHeart, 5),
		ddz.NewCard(ddz.SuitTypeDiamond, 5),
		ddz.NewCard(ddz.SuitTypeSpade, 5),
		ddz.NewCard(ddz.SuitTypeHeart, 6),
		ddz.NewCard(ddz.SuitTypeDiamond, 6),
		ddz.NewCard(ddz.SuitTypeSpade, 6),
		ddz.NewCard(ddz.SuitTypeHeart, 10),
		ddz.NewCard(ddz.SuitTypeDiamond, 10),
		ddz.NewCard(ddz.SuitTypeSpade, 10),
	}

	cardsList, infoList = ddz.SplitCardsBuXiPai(cards)
	fmt.Println("拆牌(不洗牌模式):", cardsList, infoList)
}
