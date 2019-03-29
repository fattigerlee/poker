package ddz

// 手牌是否大于上家牌型
func IsOvercomePreCardType(handCards []*Card, preCardType Type) bool {
	count := cardCount(handCards)

	if preCardType.CardType >= CardTypeDan && preCardType.CardType < CardTypeZhaDan {
		// 炸弹
		if hasZhaDan(count) {
			return true
		}

		// 火箭
		if hasHuoJian(count) {
			return true
		}
	}

	switch preCardType.CardType {
	case CardTypeDan:
		// 单
		return overcomeDan(count, preCardType)
	case CardTypeDui:
		// 对
		return overcomeDui(handCards, count, preCardType)
	case CardTypeLianDui:
		// 连对
		return overcomeLianDui(handCards, count, preCardType)
	case CardTypeSanBuDai:
		// 三不带
		return overcomeSanBuDai(handCards, count, preCardType)
	case CardTypeSanDaiYi:
		// 三带一
		return overcomeSanDaiYi(handCards, count, preCardType)
	case CardTypeSanDaiEr:
		// 三带二
		return overcomeSanDaiEr(handCards, count, preCardType)
	case CardTypeShunZi:
		// 顺子
		return overcomeShunZi(handCards, count, preCardType)
	case CardTypeFeiJiBuDai:
		// 飞机不带
		return overcomeFeiJiBuDai(handCards, count, preCardType)
	case CardTypeFeiJiDaiYi:
		// 飞机带一
		return overcomeFeiJiDaiYi(handCards, count, preCardType)
	case CardTypeFeiJiDaiEr:
		// 飞机带二
		return overcomeFeiJiDaiEr(handCards, count, preCardType)
	case CardTypeZhaDan:
		// 炸弹
		return overcomeZhaDan(handCards, count, preCardType)
	case CardTypeHuoJian:
		// 火箭
		return overcomeHuoJian(handCards, count)
	case CardTypeLianZha:
		// 连炸
		return overcomeLianZha(handCards, count, preCardType)
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

func overcomeDan(count [18]int, preCardType Type) bool {
	for i := preCardType.MinValue + 1; i <= 17; i++ {
		if count[i] >= 1 {
			return true
		}
	}
	return false
}

func overcomeDui(handCards []*Card, count [18]int, preCardType Type) bool {
	if len(handCards) < 2 {
		return false
	}

	for i := preCardType.MinValue + 1; i <= 15; i++ {
		if count[i] >= 2 {
			return true
		}
	}
	return false
}

func overcomeLianDui(handCards []*Card, count [18]int, preCardType Type) bool {
	valueRange := preCardType.MaxValue - preCardType.MinValue + 1

	if len(handCards) < valueRange*2 {
		return false
	}

	for i := preCardType.MinValue + 1; i <= 14-valueRange; i++ {
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

func overcomeSanBuDai(handCards []*Card, count [18]int, preCardType Type) bool {
	if len(handCards) < 3 {
		return false
	}

	for i := preCardType.MinValue + 1; i <= 15; i++ {
		if count[i] >= 3 {
			return true
		}
	}
	return false
}

func overcomeSanDaiYi(handCards []*Card, count [18]int, preCardType Type) bool {
	if len(handCards) < 4 {
		return false
	}

	for i := preCardType.MinValue + 1; i <= 15; i++ {
		if count[i] >= 3 {
			return true
		}
	}
	return false
}

func overcomeSanDaiEr(handCards []*Card, count [18]int, preCardType Type) bool {
	if len(handCards) < 5 {
		return false
	}

	for i := preCardType.MinValue + 1; i <= 15; i++ {
		if count[i] >= 3 {
			return true
		}
	}
	return false
}

func overcomeShunZi(handCards []*Card, count [18]int, preCardType Type) bool {
	valueRange := preCardType.MaxValue - preCardType.MinValue + 1

	if len(handCards) < valueRange {
		return false
	}

	for i := preCardType.MinValue + 1; i <= 14-valueRange; i++ {
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

func overcomeFeiJiBuDai(handCards []*Card, count [18]int, preCardType Type) bool {
	valueRange := preCardType.MaxValue - preCardType.MinValue + 1

	if len(handCards) < valueRange*3 {
		return false
	}

	for i := preCardType.MinValue + 1; i <= 14-valueRange; i++ {
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

func overcomeFeiJiDaiYi(handCards []*Card, count [18]int, preCardType Type) bool {
	valueRange := preCardType.MaxValue - preCardType.MinValue + 1

	if len(handCards) < valueRange*4 {
		return false
	}

	for i := preCardType.MinValue + 1; i <= 14-valueRange; i++ {
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

func overcomeFeiJiDaiEr(handCards []*Card, count [18]int, preCardType Type) bool {
	valueRange := preCardType.MaxValue - preCardType.MinValue + 1

	if len(handCards) < valueRange*5 {
		return false
	}

	for i := preCardType.MinValue + 1; i <= 14-valueRange; i++ {
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

func overcomeZhaDan(handCards []*Card, count [18]int, preCardType Type) bool {
	if hasHuoJian(count) {
		return true
	}

	// 炸弹
	if len(handCards) < 4 {
		return false
	}

	for i := preCardType.MinValue + 1; i <= 15; i++ {
		if count[i] == 4 {
			return true
		}
	}

	// 连炸
	if len(handCards) < 8 {
		return false
	}

	for i := 3; i < 14; i++ {
		if count[i] == 4 && count[i+1] == 4 {
			return true
		}
	}
	return false
}

func overcomeHuoJian(handCards []*Card, count [18]int) bool {
	// 连炸
	if len(handCards) < 8 {
		return false
	}

	for i := 3; i < 14; i++ {
		if count[i] == 4 && count[i+1] == 4 {
			return true
		}
	}
	return false
}

func overcomeLianZha(handCards []*Card, count [18]int, preCardType Type) bool {
	valueRange := preCardType.MaxValue - preCardType.MinValue + 1

	if len(handCards) < valueRange*4 {
		return false
	}

	for i := preCardType.MinValue + 1; i <= 14-valueRange; i++ {
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
