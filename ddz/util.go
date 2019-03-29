package ddz

type Type struct {
	CardType CardType // 牌型
	MinValue int      // 最小值
	MaxValue int      // 最大值
}

// 单
func IsDan(cards []*Card) (cardType Type) {
	if len(cards) != 1 {
		return
	}

	cardType.CardType = CardTypeDan
	cardType.MinValue = int(cards[0].Num)
	return
}

// 对
func IsDui(cards []*Card) (cardType Type) {
	if len(cards) != 2 {
		return
	}

	// 牌值相同
	if cards[0].Num != cards[1].Num {
		return
	}

	cardType.CardType = CardTypeDui
	cardType.MinValue = int(cards[0].Num)
	return
}

// 连对
func IsLianDui(cards []*Card) (cardType Type) {
	var minValue int
	var maxValue int
	count := cardCount(cards)
	for i := 3; i <= 14; i++ {
		if count[i] == 2 {
			minValue = i
			break
		}
	}

	for i := 14; i >= 3; i-- {
		if count[i] == 2 {
			maxValue = i
			break
		}
	}
	valueRange := maxValue - minValue + 1

	if len(cards) != valueRange*2 {
		return
	}

	for i := minValue; i <= 14-valueRange; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] != 2 {
				exist = false
				break
			}
		}
		if exist {
			cardType.CardType = CardTypeLianDui
			cardType.MinValue = minValue
			cardType.MaxValue = maxValue
			return
		}
	}
	return
}

// 三不带
func IsSanBuDai(cards []*Card) (cardType Type) {
	if len(cards) != 3 {
		return
	}

	count := cardCount(cards)
	for i := 3; i <= 15; i++ {
		if count[i] == 3 {
			cardType.CardType = CardTypeSanBuDai
			cardType.MinValue = i
			return
		}
	}
	return
}

// 三带一
func IsSanDaiYi(cards []*Card) (cardType Type) {
	if len(cards) != 4 {
		return
	}

	count := cardCount(cards)
	for i := 3; i <= 15; i++ {
		if count[i] == 3 {
			cardType.CardType = CardTypeSanDaiYi
			cardType.MinValue = i
			return
		}
	}
	return
}

// 三带二
func IsSanDaiEr(cards []*Card) (cardType Type) {
	if len(cards) != 5 {
		return
	}

	count := cardCount(cards)

	var exist bool
	for i := 3; i <= 15; i++ {
		if count[i] == 2 {
			exist = true
			break
		}
	}

	if !exist {
		return
	}

	for i := 3; i <= 15; i++ {
		if count[i] == 3 {
			cardType.CardType = CardTypeSanDaiEr
			cardType.MinValue = i
			return
		}
	}
	return
}

// 四带单(炸弹带两张单)
func IsSiDaiDan(cards []*Card) (cardType Type) {
	if len(cards) != 6 {
		return
	}

	count := cardCount(cards)
	for i := 3; i <= 15; i++ {
		if count[i] == 4 {
			cardType.CardType = CardTypeSiDaiDan
			cardType.MinValue = i
			return
		}
	}
	return
}

// 四带对(炸弹带两对)
func IsSiDaiDui(cards []*Card) (cardType Type) {
	if len(cards) != 8 {
		return
	}

	count := cardCount(cards)

	var exist int
	for i := 3; i <= 15; i++ {
		if count[i] == 2 {
			exist++
		}
	}

	if exist != 2 {
		return
	}

	for i := 3; i <= 15; i++ {
		if count[i] == 4 {
			cardType.CardType = CardTypeSiDaiDui
			cardType.MinValue = i
			return
		}
	}
	return
}

// 顺子
func IsShunZi(cards []*Card) (cardType Type) {
	if len(cards) < 5 {
		return
	}

	var minValue int
	var maxValue int
	count := cardCount(cards)
	for i := 3; i <= 14; i++ {
		if count[i] == 1 {
			minValue = i
			break
		}
	}

	for i := 14; i >= 3; i-- {
		if count[i] == 1 {
			maxValue = i
			break
		}
	}
	valueRange := maxValue - minValue + 1

	if len(cards) != valueRange {
		return
	}

	for i := minValue; i <= 14-valueRange; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] != 1 {
				exist = false
				break
			}
		}
		if exist {
			cardType.CardType = CardTypeShunZi
			cardType.MinValue = minValue
			cardType.MaxValue = maxValue
			return
		}
	}
	return
}

// 飞机不带
func IsFeiJiBuDai(cards []*Card) (cardType Type) {
	if len(cards) < 6 {
		return
	}

	var minValue int
	var maxValue int
	count := cardCount(cards)
	for i := 3; i <= 14; i++ {
		if count[i] == 3 {
			minValue = i
			break
		}
	}

	for i := 14; i >= 3; i-- {
		if count[i] == 3 {
			maxValue = i
			break
		}
	}
	valueRange := maxValue - minValue + 1

	if len(cards) != valueRange*3 {
		return
	}

	for i := minValue; i <= 14-valueRange; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] != 3 {
				exist = false
				break
			}
		}
		if exist {
			cardType.CardType = CardTypeFeiJiBuDai
			cardType.MinValue = minValue
			cardType.MaxValue = maxValue
			return
		}
	}
	return
}

// 飞机带一
func IsFeiJiDaiYi(cards []*Card) (cardType Type) {
	if len(cards) < 8 {
		return
	}

	// 连炸不是飞机带一
	if IsLianZha(cards).CardType != CardTypeNone {
		return
	}

	var minValue int
	var maxValue int
	count := cardCount(cards)
	for i := 3; i <= 14; i++ {
		if count[i] >= 3 {
			minValue = i
			break
		}
	}

	for i := 14; i >= 3; i-- {
		if count[i] >= 3 {
			maxValue = i
			break
		}
	}
	valueRange := maxValue - minValue + 1

	if len(cards) != valueRange*4 {
		return
	}

	for i := minValue; i <= 14-valueRange; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 3 {
				exist = false
				break
			}
		}
		if exist {
			cardType.CardType = CardTypeFeiJiDaiYi
			cardType.MinValue = minValue
			cardType.MaxValue = maxValue
			return
		}
	}
	return
}

// 飞机带二
func IsFeiJiDaiEr(cards []*Card) (cardType Type) {
	if len(cards) < 10 {
		return
	}

	count := cardCount(cards)

	var exist int
	for i := 3; i <= 15; i++ {
		if count[i] == 2 {
			exist++
		}
	}

	var minValue int
	var maxValue int
	for i := 3; i <= 14; i++ {
		if count[i] == 3 {
			minValue = i
			break
		}
	}

	for i := 14; i >= 3; i-- {
		if count[i] == 3 {
			maxValue = i
			break
		}
	}
	valueRange := maxValue - minValue + 1

	if len(cards) != valueRange*5 || exist != valueRange {
		return
	}

	for i := minValue; i <= 14-valueRange; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] != 3 {
				exist = false
				break
			}
		}
		if exist {
			cardType.CardType = CardTypeFeiJiDaiEr
			cardType.MinValue = minValue
			cardType.MaxValue = maxValue
			return
		}
	}
	return
}

// 炸弹
func IsZhaDan(cards []*Card) (cardType Type) {
	if len(cards) != 4 {
		return
	}

	count := cardCount(cards)
	for i := 3; i <= 15; i++ {
		if count[i] == 4 {
			cardType.CardType = CardTypeZhaDan
			cardType.MinValue = i
			return
		}
	}
	return
}

// 火箭
func IsHuoJian(cards []*Card) (cardType Type) {
	if len(cards) != 2 {
		return
	}

	if cards[0].Num == 16 && cards[1].Num == 17 {
		cardType.CardType = CardTypeHuoJian
		return
	}
	return
}

// 连炸
func IsLianZha(cards []*Card) (cardType Type) {
	if len(cards) < 8 {
		return
	}

	var minValue int
	var maxValue int
	count := cardCount(cards)
	for i := 3; i <= 14; i++ {
		if count[i] == 4 {
			minValue = i
			break
		}
	}

	for i := 14; i >= 3; i-- {
		if count[i] == 4 {
			maxValue = i
			break
		}
	}
	valueRange := maxValue - minValue + 1

	if len(cards) != valueRange*4 {
		return
	}

	for i := minValue; i <= 14-valueRange; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] != 4 {
				exist = false
				break
			}
		}
		if exist {
			cardType.CardType = CardTypeLianZha
			cardType.MinValue = minValue
			cardType.MaxValue = maxValue
			return
		}
	}
	return
}

// 牌值计数
func cardCount(cards []*Card) [18]int {
	var count [18]int
	for _, card := range cards {
		count[int(card.Num)]++
	}
	return count
}
