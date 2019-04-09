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
		if ret = findBigDan(cards, count, cardType); len(ret) != 0 {
			return
		}
	case CardTypeDui:
		if ret = findBigDui(cards, count, cardType); len(ret) != 0 {
			return
		}
	case CardTypeLianDui:
		if ret = findBigLianDui(cards, count, cardType); len(ret) != 0 {
			return
		}
	case CardTypeSanBuDai:
		if ret = findBigSanBuDai(cards, count, cardType); len(ret) != 0 {
			return
		}
	case CardTypeSanDaiYi:
		if ret = findBigSanDaiYi(cards, count, cardType); len(ret) != 0 {
			return
		}
	case CardTypeSanDaiEr:
		if ret = findBigSanDaiEr(cards, count, cardType); len(ret) != 0 {
			return
		}
	case CardTypeShunZi:
		if ret = findBigShunZi(cards, count, cardType); len(ret) != 0 {
			return
		}
	case CardTypeFeiJiBuDai:
		if ret = findBigFeiJiBuDai(cards, count, cardType); len(ret) != 0 {
			return
		}
	case CardTypeFeiJiDaiYi:
		if ret = findBigFeiJiDaiYi(cards, count, cardType); len(ret) != 0 {
			return
		}
	case CardTypeFeiJiDaiEr:
		if ret = findBigFeiJiDaiEr(cards, count, cardType); len(ret) != 0 {
			return
		}
	case CardTypeZhaDan:
		if ret = findBigZhaDan(cards, count, cardType); len(ret) != 0 {
			return
		}

		if ret = findHuoJian(cards, count); len(ret) != 0 {
			return
		}
	case CardTypeHuoJian:
		if ret = findBigHuoJian(cards, count, cardType); len(ret) != 0 {
			return
		}
	case CardTypeLianZha:
		if ret = findBigLianZha(cards, count, cardType); len(ret) != 0 {
			return
		}
	}

	if cardType.CardType >= CardTypeDan && cardType.CardType < CardTypeZhaDan {
		if ret = findZhaDan(cards, count); ret != nil {
			return
		}

		if ret = findHuoJian(cards, count); ret != nil {
			return
		}
	}
	return nil
}

func findZhaDan(cards []*Card, count [18]int) (ret []*Card) {
	for i := 3; i <= 15; i++ {
		if count[i] == 4 {
			for j := 0; j < len(cards); j++ {
				if int(cards[j].Num) == i {
					for k := 0; k < 4; k++ {
						ret = append(ret, cards[j+k])
					}
					return
				}
			}
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
			for j := 0; j < len(cards); j++ {
				if int(cards[j].Num) == i {
					ret = append(ret, cards[j])
					return
				}
			}
		}
	}
	return
}

func findBigDui(cards []*Card, count [18]int, cardType Type) (ret []*Card) {
	for i := cardType.MinValue + 1; i <= 15; i++ {
		if count[i] >= 2 {
			for j := 0; j < len(cards); j++ {
				if int(cards[j].Num) == i {
					for k := 0; k < 2; k++ {
						ret = append(ret, cards[j+k])
					}
					return
				}
			}
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
			for j := i; j < i+valueRange; j++ {
				for k := 0; k < len(cards); k++ {
					if int(cards[k].Num) == j {
						for m := 0; m < 2; m++ {
							ret = append(ret, cards[k+m])
						}
						break
					}
				}
			}
			return
		}
	}
	return
}

func findBigSanBuDai(cards []*Card, count [18]int, cardType Type) (ret []*Card) {
	for i := cardType.MinValue + 1; i <= 15; i++ {
		if count[i] >= 3 {
			for j := 0; j < len(cards); j++ {
				if int(cards[j].Num) == i {
					for k := 0; k < 3; k++ {
						ret = append(ret, cards[j+k])
					}
					return
				}
			}
		}
	}
	return
}

func findBigSanDaiYi(cards []*Card, count [18]int, cardType Type) (ret []*Card) {
	for i := cardType.MinValue + 1; i <= 15; i++ {
		if count[i] >= 3 {
			for j := 0; j < len(cards); j++ {
				if int(cards[j].Num) == i {
					for k := 0; k < 3; k++ {
						ret = append(ret, cards[j+k])
					}
					break
				}
			}

			for j := 3; j != i && j <= 15; j++ {
				if count[j] > 0 {
					for k := 0; k < len(cards); k++ {
						if int(cards[k].Num) == j {
							ret = append(ret, cards[k])
							return
						}
					}
				}
			}
		}
	}
	return
}

func findBigSanDaiEr(cards []*Card, count [18]int, cardType Type) (ret []*Card) {
	for i := cardType.MinValue + 1; i <= 15; i++ {
		if count[i] >= 3 {
			for j := 0; j < len(cards); j++ {
				if int(cards[j].Num) == i {
					for k := 0; k < 3; k++ {
						ret = append(ret, cards[j+k])
					}
					break
				}
			}

			for j := 3; j != i && j <= 15; j++ {
				if count[j] > 1 {
					for k := 0; k < len(cards); k++ {
						if int(cards[k].Num) == j {
							for m := 0; m < 2; m++ {
								ret = append(ret, cards[k+m])
							}
							return
						}
					}
				}
			}
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
			for j := i; j < i+valueRange; j++ {
				for k := 0; k < len(cards); k++ {
					if int(cards[k].Num) == j {
						for m := 0; m < 1; m++ {
							ret = append(ret, cards[k+m])
						}
						break
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
			for j := i; j < i+valueRange; j++ {
				for k := 0; k < len(cards); k++ {
					if int(cards[k].Num) == j {
						for m := 0; m < 3; m++ {
							ret = append(ret, cards[k+m])
						}
						break
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
				for k := 0; k < len(cards); k++ {
					if int(cards[k].Num) == j {
						for m := 0; m < 3; m++ {
							ret = append(ret, cards[k+m])
						}
						break
					}
				}
			}

			var times int
			for j := 3; j < i; j++ {
				for k := 0; k < len(cards); k++ {
					if times == valueRange {
						return
					}

					if int(cards[k].Num) == j {
						ret = append(ret, cards[k])
						times++
					}
				}
			}

			for j := i + valueRange; j <= 14; j++ {
				for k := 0; k < len(cards); k++ {
					if times == valueRange {
						return
					}

					if int(cards[k].Num) == j {
						ret = append(ret, cards[k])
						times++
					}
				}
			}

			for j := i; j < i+valueRange; j++ {
				for k := 0; k < len(cards); k++ {
					if times == valueRange {
						return
					}

					var exist bool
					for m := 0; m < len(ret); m++ {
						if cards[k].Suit == ret[m].Suit && cards[k].Num == ret[m].Num {
							exist = true
							break
						}
					}

					if !exist {
						ret = append(ret, cards[k])
						times++
					}
				}
			}
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
			for j := i; j < i+valueRange; j++ {
				for k := 0; k < len(cards); k++ {
					if int(cards[k].Num) == j {
						for m := 0; m < 3; m++ {
							ret = append(ret, cards[k+m])
						}
						break
					}
				}
			}

			var times int
			for j := 3; j < i; j++ {
				for k := 0; k < len(cards); k++ {
					if times == valueRange {
						return
					}

					if int(cards[k].Num) == j && count[j] >= 2 {
						for m := 0; m < 2; m++ {
							ret = append(ret, cards[k+m])
						}
						times++
						break
					}
				}
			}

			for j := i + valueRange; j <= 15; j++ {
				for k := 0; k < len(cards); k++ {
					if times == valueRange {
						return
					}

					if int(cards[k].Num) == j && count[j] >= 2 {
						for m := 0; m < 2; m++ {
							ret = append(ret, cards[k+m])
						}
						times++
						break
					}
				}
			}
		}
	}
	return
}

func findBigZhaDan(cards []*Card, count [18]int, cardType Type) (ret []*Card) {
	for i := cardType.MinValue + 1; i <= 15; i++ {
		if count[i] == 4 {
			for j := 0; j < len(cards); j++ {
				if int(cards[j].Num) == i {
					for k := 0; k < 4; k++ {
						ret = append(ret, cards[j+k])
					}
					return
				}
			}
		}
	}
	return
}

func findBigHuoJian(cards []*Card, count [18]int, cardType Type) (ret []*Card) {
	for i := 3; i < 14; i++ {
		if count[i] == 4 && count[i+1] == 4 {
			for j := 0; j < len(cards); j++ {
				if int(cards[j].Num) == i {
					for k := 0; k < 8; k++ {
						ret = append(ret, cards[j+k])
					}
					return
				}
			}
		}
	}
	return
}

func findBigLianZha(cards []*Card, count [18]int, cardType Type) (ret []*Card) {
	valueRange := cardType.MaxValue - cardType.MinValue + 1

	for i := cardType.MinValue + 1; i <= 14; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 4 {
				exist = false
				break
			}
		}

		if exist {
			for j := i; j < i+valueRange; j++ {
				for k := 0; k < len(cards); k++ {
					if int(cards[k].Num) == j {
						for m := 0; m < 4; m++ {
							ret = append(ret, cards[k+m])
						}
						break
					}
				}
			}
			return
		}
	}
	return
}
