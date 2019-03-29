package ddz

import (
	"sort"
)

// 扑克牌列表中找出比牌型更大的牌
func FindOvercomeCard(cardType Type, cards []*Card) (ret []*Card) {
	// 手牌排序
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Num < cards[j].Num
	})

	count := cardCount(cards)

	switch cardType.CardType {
	case CardTypeDan:
		return findBigDan(cards, count, cardType)
	case CardTypeDui:
		return findBigDui(cards, count, cardType)
	case CardTypeLianDui:
		return findBigLianDui(cards, count, cardType)
	//case CardTypeSanBuDai:
	//	// 三不带
	//	return overcomeSanBuDai(cards, count, cardType)
	//case CardTypeSanDaiYi:
	//	// 三带一
	//	return overcomeSanDaiYi(cards, count, cardType)
	//case CardTypeSanDaiEr:
	//	// 三带二
	//	return overcomeSanDaiEr(cards, count, cardType)
	case CardTypeShunZi:
		return findBigShunZi(cards, count, cardType)
	case CardTypeFeiJiBuDai:
		return findBigFeiJiBuDai(cards, count, cardType)
	case CardTypeFeiJiDaiYi:
		return findBigFeiJiDaiYi(cards, count, cardType)
		//case CardTypeFeiJiDaiEr:
		//	// 飞机带二
		//	return overcomeFeiJiDaiEr(cards, count, cardType)
		//case CardTypeZhaDan:
		//	// 炸弹
		//	return overcomeZhaDan(cards, count, cardType)
		//case CardTypeHuoJian:
		//	// 火箭
		//	return overcomeHuoJian(cards, count)
		//case CardTypeLianZha:
		//	// 连炸
		//	return overcomeLianZha(cards, count, cardType)
	}

	if cardType.CardType >= CardTypeDan && cardType.CardType < CardTypeZhaDan {
		// 炸弹
		if ret = findZhaDan(cards, count); ret != nil {
			return
		}

		// 火箭
		if ret = findHuoJian(cards, count); ret != nil {
			return
		}
	}
	return nil
}

func findZhaDan(cards []*Card, count [18]int) (ret []*Card) {
	for i := 3; i <= 15; i++ {
		if count[i] == 4 {
			for _, c := range cards {
				if int(c.Num) == i {
					ret = append(ret, c)
				}
			}
			return
		}
	}
	return
}

func findHuoJian(cards []*Card, count [18]int) (ret []*Card) {
	if count[16] == 1 && count[17] == 1 {
		for _, c := range cards {
			if int(c.Num) == 16 || int(c.Num) == 17 {
				ret = append(ret, c)
			}
		}
		return
	}
	return
}

func findBigDan(cards []*Card, count [18]int, cardType Type) (ret []*Card) {
	for i := cardType.MinValue + 1; i <= 17; i++ {
		if count[i] >= 1 {
			var times int
			for _, c := range cards {
				if times == 1 {
					break
				}

				if int(c.Num) == i {
					ret = append(ret, c)
					times++
				}
			}
			return
		}
	}
	return
}

func findBigDui(cards []*Card, count [18]int, cardType Type) (ret []*Card) {
	for i := cardType.MinValue + 1; i <= 15; i++ {
		if count[i] >= 2 {
			var times int
			for _, c := range cards {
				if times == 2 {
					break
				}

				if int(c.Num) == i {
					ret = append(ret, c)
					times++
				}
			}
			return
		}
	}
	return
}

func findBigLianDui(cards []*Card, count [18]int, cardType Type) (ret []*Card) {
	valueRange := cardType.MaxValue - cardType.MinValue + 1

	for i := cardType.MinValue + 1; i <= 14; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 2 {
				exist = false
				break
			}
		}

		if exist {
			for k := i; k < i+valueRange; k++ {
				var times int
				for _, c := range cards {
					if times == 2 {
						break
					}

					if int(c.Num) == k {
						ret = append(ret, c)
						times++
					}
				}
			}
			return
		}
	}
	return
}

func findBigShunZi(cards []*Card, count [18]int, cardType Type) (ret []*Card) {
	valueRange := cardType.MaxValue - cardType.MinValue + 1

	for i := cardType.MinValue + 1; i <= 14; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 1 {
				exist = false
				break
			}
		}

		if exist {
			for k := i; k < i+valueRange; k++ {
				var times int
				for _, c := range cards {
					if times == 1 {
						break
					}

					if int(c.Num) == k {
						ret = append(ret, c)
						times++
					}
				}
			}
			return
		}
	}
	return
}

func findBigFeiJiBuDai(cards []*Card, count [18]int, cardType Type) (ret []*Card) {
	valueRange := cardType.MaxValue - cardType.MinValue + 1

	for i := cardType.MinValue + 1; i <= 14; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 3 {
				exist = false
				break
			}
		}

		if exist {
			for k := i; k < i+valueRange; k++ {
				var times int
				for _, c := range cards {
					if times == 3 {
						break
					}

					if int(c.Num) == k {
						ret = append(ret, c)
						times++
					}
				}
			}
			return
		}
	}
	return
}

func findBigFeiJiDaiYi(cards []*Card, count [18]int, cardType Type) (ret []*Card) {
	valueRange := cardType.MaxValue - cardType.MinValue + 1

	for i := cardType.MinValue + 1; i <= 14; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 3 {
				exist = false
				break
			}
		}

		if exist {
			for j := i; j < i+valueRange; j++ {
				var times int
				for _, c := range cards {
					if times == 3 {
						break
					}

					if int(c.Num) == j {
						ret = append(ret, c)
						times++
					}
				}
			}

			var times int
			for j := 3; j < i; j++ {
				for _, c := range cards {
					if times == valueRange {
						return
					}

					if int(c.Num) == j {
						ret = append(ret, c)
						times++
					}
				}
			}

			for j := i + valueRange; j <= 14; j++ {
				for _, c := range cards {
					if times == valueRange {
						return
					}

					if int(c.Num) == j {
						ret = append(ret, c)
						times++
					}
				}
			}

			for j := i; j < i+valueRange; j++ {
				for _, c := range cards {
					if times == valueRange {
						return
					}

					var exist bool
					for _, rc := range ret {
						if rc.Suit == c.Suit && rc.Num == c.Num {
							exist = true
							break
						}
					}

					if !exist {
						ret = append(ret, c)
						times++
					}
				}
			}
			//
			//for j := 0; j < valueRange; j++ {
			//	for _, c := range cards {
			//		for _, rc := range ret {
			//			if c.Suit == rc.Suit && c.Num == rc.Num {
			//				continue
			//			}
			//
			//		}
			//	}
			//}
			//
			//var totalTimes int
			//for _, rc := range ret {
			//	for _, c := range cards {
			//		if c.Suit == rc.Suit && c.Num == rc.Num {
			//			continue
			//		}
			//	}
			//}
			//
			//for j := 3; j <= 14; j++ {
			//	var times int
			//	for _, c := range cards {
			//		if totalTimes == valueRange {
			//			return
			//		}
			//
			//		if times == 1 {
			//			break
			//		}
			//
			//		if int(c.Num) == j && int(c.Num) != 3 {
			//			ret = append(ret, c)
			//			times++
			//			totalTimes++
			//		}
			//	}
			//}
		}
	}
	return
}

func findBigFeiJiDaiEr(cards []*Card, count [18]int, cardType Type) (ret []*Card) {
	valueRange := cardType.MaxValue - cardType.MinValue + 1

	for i := cardType.MinValue + 1; i <= 14; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 3 {
				exist = false
				break
			}
		}

		if exist {
			for k := i; k < i+valueRange; k++ {
				var times int
				for _, c := range cards {
					if times == 3 {
						break
					}

					if int(c.Num) == k {
						ret = append(ret, c)
						times++
					}
				}
			}
			return
		}
	}
	return
}
