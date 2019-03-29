package ddz

// 扑克牌列表中是否有比牌型更大的牌
func IsOvercomeCardType(cardType Type, cards []*Card) bool {
	count := cardCount(cards)

	if cardType.CardType >= CardTypeDan && cardType.CardType < CardTypeZhaDan {
		// 炸弹
		if hasZhaDan(count) {
			return true
		}

		// 火箭
		if hasHuoJian(count) {
			return true
		}
	}

	switch cardType.CardType {
	case CardTypeDan:
		// 单
		return overcomeDan(count, cardType)
	case CardTypeDui:
		// 对
		return overcomeDui(cards, count, cardType)
	case CardTypeLianDui:
		// 连对
		return overcomeLianDui(cards, count, cardType)
	case CardTypeSanBuDai:
		// 三不带
		return overcomeSanBuDai(cards, count, cardType)
	case CardTypeSanDaiYi:
		// 三带一
		return overcomeSanDaiYi(cards, count, cardType)
	case CardTypeSanDaiEr:
		// 三带二
		return overcomeSanDaiEr(cards, count, cardType)
	case CardTypeShunZi:
		// 顺子
		return overcomeShunZi(cards, count, cardType)
	case CardTypeFeiJiBuDai:
		// 飞机不带
		return overcomeFeiJiBuDai(cards, count, cardType)
	case CardTypeFeiJiDaiYi:
		// 飞机带一
		return overcomeFeiJiDaiYi(cards, count, cardType)
	case CardTypeFeiJiDaiEr:
		// 飞机带二
		return overcomeFeiJiDaiEr(cards, count, cardType)
	case CardTypeZhaDan:
		// 炸弹
		return overcomeZhaDan(cards, count, cardType)
	case CardTypeHuoJian:
		// 火箭
		return overcomeHuoJian(cards, count)
	case CardTypeLianZha:
		// 连炸
		return overcomeLianZha(cards, count, cardType)
	}
	return false
}

func hasZhaDan(count [18]int) bool {
	for i := 3; i <= 15; i++ {
		if count[i] == 4 {
			return true
		}
	}
	return false
}

func hasHuoJian(count [18]int) bool {
	if count[16] == 1 && count[17] == 1 {
		return true
	}
	return false
}

func overcomeDan(count [18]int, cardType Type) bool {
	for i := cardType.MinValue + 1; i <= 17; i++ {
		if count[i] >= 1 {
			return true
		}
	}
	return false
}

func overcomeDui(cards []*Card, count [18]int, cardType Type) bool {
	if len(cards) < 2 {
		return false
	}

	for i := cardType.MinValue + 1; i <= 15; i++ {
		if count[i] >= 2 {
			return true
		}
	}
	return false
}

func overcomeLianDui(cards []*Card, count [18]int, cardType Type) bool {
	valueRange := cardType.MaxValue - cardType.MinValue + 1

	if len(cards) < valueRange*2 {
		return false
	}

	for i := cardType.MinValue + 1; i <= 14-valueRange; i++ {
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

func overcomeSanBuDai(cards []*Card, count [18]int, cardType Type) bool {
	if len(cards) < 3 {
		return false
	}

	for i := cardType.MinValue + 1; i <= 15; i++ {
		if count[i] >= 3 {
			return true
		}
	}
	return false
}

func overcomeSanDaiYi(cards []*Card, count [18]int, cardType Type) bool {
	if len(cards) < 4 {
		return false
	}

	for i := cardType.MinValue + 1; i <= 15; i++ {
		if count[i] >= 3 {
			return true
		}
	}
	return false
}

func overcomeSanDaiEr(cards []*Card, count [18]int, cardType Type) bool {
	if len(cards) < 5 {
		return false
	}

	for i := cardType.MinValue + 1; i <= 15; i++ {
		if count[i] >= 3 {
			return true
		}
	}
	return false
}

func overcomeShunZi(cards []*Card, count [18]int, cardType Type) bool {
	valueRange := cardType.MaxValue - cardType.MinValue + 1

	if len(cards) < valueRange {
		return false
	}

	for i := cardType.MinValue + 1; i <= 14-valueRange; i++ {
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

func overcomeFeiJiBuDai(cards []*Card, count [18]int, cardType Type) bool {
	valueRange := cardType.MaxValue - cardType.MinValue + 1

	if len(cards) < valueRange*3 {
		return false
	}

	for i := cardType.MinValue + 1; i <= 14-valueRange; i++ {
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

func overcomeFeiJiDaiYi(cards []*Card, count [18]int, cardType Type) bool {
	valueRange := cardType.MaxValue - cardType.MinValue + 1

	if len(cards) < valueRange*4 {
		return false
	}

	for i := cardType.MinValue + 1; i <= 14-valueRange; i++ {
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

func overcomeFeiJiDaiEr(cards []*Card, count [18]int, cardType Type) bool {
	valueRange := cardType.MaxValue - cardType.MinValue + 1

	if len(cards) < valueRange*5 {
		return false
	}

	for i := cardType.MinValue + 1; i <= 14-valueRange; i++ {
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

func overcomeZhaDan(cards []*Card, count [18]int, cardType Type) bool {
	if hasHuoJian(count) {
		return true
	}

	// 炸弹
	if len(cards) < 4 {
		return false
	}

	for i := cardType.MinValue + 1; i <= 15; i++ {
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

func overcomeLianZha(cards []*Card, count [18]int, cardType Type) bool {
	valueRange := cardType.MaxValue - cardType.MinValue + 1

	if len(cards) < valueRange*4 {
		return false
	}

	for i := cardType.MinValue + 1; i <= 14-valueRange; i++ {
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
