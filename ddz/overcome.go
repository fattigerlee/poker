package ddz

// 扑克牌列表中是否有比牌型更大的牌
func OvercomeCard(info CardTypeInfo, cards []*Card) bool {
	count := cardCount(cards)

	if info.CardType >= CardTypeDan && info.CardType < CardTypeZhaDan {
		// 炸弹
		if existZhaDan(count) {
			return true
		}

		// 火箭
		if existHuoJian(count) {
			return true
		}
	}

	switch info.CardType {
	case CardTypeDan:
		// 单
		return overcomeDan(count, info)
	case CardTypeDui:
		// 对
		return overcomeDui(cards, count, info)
	case CardTypeLianDui:
		// 连对
		return overcomeLianDui(cards, count, info)
	case CardTypeSanBuDai:
		// 三不带
		return overcomeSanBuDai(cards, count, info)
	case CardTypeSanDaiYi:
		// 三带一
		return overcomeSanDaiYi(cards, count, info)
	case CardTypeSanDaiEr:
		// 三带二
		return overcomeSanDaiEr(cards, count, info)
	case CardTypeShunZi:
		// 顺子
		return overcomeShunZi(cards, count, info)
	case CardTypeFeiJiBuDai:
		// 飞机不带
		return overcomeFeiJiBuDai(cards, count, info)
	case CardTypeFeiJiDaiYi:
		// 飞机带一
		return overcomeFeiJiDaiYi(cards, count, info)
	case CardTypeFeiJiDaiEr:
		// 飞机带二
		return overcomeFeiJiDaiEr(cards, count, info)
	case CardTypeZhaDan:
		// 炸弹
		return overcomeZhaDan(cards, count, info)
	case CardTypeHuoJian:
		// 火箭
		return overcomeHuoJian(cards, count)
	case CardTypeLianZha:
		// 连炸
		return overcomeLianZha(cards, count, info)
	}
	return false
}

func existZhaDan(count [18]int) bool {
	for i := 3; i <= 15; i++ {
		if count[i] == 4 {
			return true
		}
	}
	return false
}

func existHuoJian(count [18]int) bool {
	if count[16] == 1 && count[17] == 1 {
		return true
	}
	return false
}

func overcomeDan(count [18]int, info CardTypeInfo) bool {
	for i := info.MinValue + 1; i < 18; i++ {
		if count[i] >= 1 {
			return true
		}
	}
	return false
}

func overcomeDui(cards []*Card, count [18]int, info CardTypeInfo) bool {
	if len(cards) < 2 {
		return false
	}

	for i := info.MinValue + 1; i <= 15; i++ {
		if count[i] >= 2 {
			return true
		}
	}
	return false
}

func overcomeLianDui(cards []*Card, count [18]int, info CardTypeInfo) bool {
	valueRange := info.MaxValue - info.MinValue + 1

	if len(cards) < valueRange*2 {
		return false
	}

	for i := info.MinValue + 1; i <= 14-valueRange; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 2 {
				exist = false
				break
			}
		}
		if exist {
			return true
		}
	}
	return false
}

func overcomeSanBuDai(cards []*Card, count [18]int, info CardTypeInfo) bool {
	if len(cards) < 3 {
		return false
	}

	for i := info.MinValue + 1; i <= 15; i++ {
		if count[i] >= 3 {
			return true
		}
	}
	return false
}

func overcomeSanDaiYi(cards []*Card, count [18]int, info CardTypeInfo) bool {
	if len(cards) < 4 {
		return false
	}

	for i := info.MinValue + 1; i <= 15; i++ {
		if count[i] == 3 {
			for j := 3; j < 18; j++ {
				if j == i {
					continue
				}

				if count[j] >= 1 {
					return true
				}
			}
		}
	}
	return false
}

func overcomeSanDaiEr(cards []*Card, count [18]int, info CardTypeInfo) bool {
	if len(cards) < 5 {
		return false
	}

	for i := info.MinValue + 1; i <= 15; i++ {
		if count[i] == 3 {
			for j := 3; j < 18; j++ {
				if j == i {
					continue
				}

				if count[j] >= 2 {
					return true
				}
			}
		}
	}
	return false
}

func overcomeShunZi(cards []*Card, count [18]int, info CardTypeInfo) bool {
	valueRange := info.MaxValue - info.MinValue + 1

	if len(cards) < valueRange {
		return false
	}

	for i := info.MinValue + 1; i <= 14-valueRange; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 1 {
				exist = false
				break
			}
		}
		if exist {
			return true
		}
	}
	return false
}

func overcomeFeiJiBuDai(cards []*Card, count [18]int, info CardTypeInfo) bool {
	valueRange := info.MaxValue - info.MinValue + 1

	if len(cards) < valueRange*3 {
		return false
	}

	for i := info.MinValue + 1; i <= 14-valueRange; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 3 {
				exist = false
				break
			}
		}
		if exist {
			return true
		}
	}
	return false
}

func overcomeFeiJiDaiYi(cards []*Card, count [18]int, info CardTypeInfo) bool {
	valueRange := info.MaxValue - info.MinValue + 1

	if len(cards) < valueRange*4 {
		return false
	}

	for i := info.MinValue + 1; i <= 14-valueRange; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 3 {
				exist = false
				break
			}
		}
		if exist {
			return true
		}
	}
	return false
}

func overcomeFeiJiDaiEr(cards []*Card, count [18]int, info CardTypeInfo) bool {
	valueRange := info.MaxValue - info.MinValue + 1

	if len(cards) < valueRange*5 {
		return false
	}

	for i := info.MinValue + 1; i <= 14-valueRange; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 3 {
				exist = false
				break
			}
		}
		if exist {
			return true
		}
	}
	return false
}

func overcomeZhaDan(cards []*Card, count [18]int, info CardTypeInfo) bool {
	if existHuoJian(count) {
		return true
	}

	// 炸弹
	if len(cards) < 4 {
		return false
	}

	for i := info.MinValue + 1; i <= 15; i++ {
		if count[i] == 4 {
			return true
		}
	}

	// 连炸
	if len(cards) < 8 {
		return false
	}

	for i := 3; i < 14; i++ {
		if count[i] == 4 && count[i+1] == 4 {
			return true
		}
	}
	return false
}

func overcomeHuoJian(cards []*Card, count [18]int) bool {
	// 连炸
	if len(cards) < 8 {
		return false
	}

	for i := 3; i < 14; i++ {
		if count[i] == 4 && count[i+1] == 4 {
			return true
		}
	}
	return false
}

func overcomeLianZha(cards []*Card, count [18]int, info CardTypeInfo) bool {
	valueRange := info.MaxValue - info.MinValue + 1

	if len(cards) < valueRange*4 {
		return false
	}

	// 比连炸更大的连炸
	for i := info.MinValue + 1; i <= 14-valueRange; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] != 4 {
				exist = false
				break
			}
		}
		if exist {
			return true
		}
	}

	// 比连炸更多的连炸
	valueRange = valueRange + 1
	for i := 3; i <= 14-valueRange; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] != 4 {
				exist = false
				break
			}
		}
		if exist {
			return true
		}
	}
	return false
}
