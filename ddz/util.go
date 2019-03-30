package ddz

type Type struct {
	CardType CardType // 牌型
	MinValue int      // 最小值
	MaxValue int      // 最大值
}

// 单
func isDan(cards []*Card) (cardType Type) {
	if len(cards) != 1 {
		return
	}

	cardType.CardType = CardTypeDan
	cardType.MinValue = int(cards[0].Num)
	return
}

// 对
func isDui(cards []*Card) (cardType Type) {
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
func isLianDui(cards []*Card) (cardType Type) {
	count := cardCount(cards)
	minValue := getMinValue(count, 2)
	maxValue := getMaxValue(count, 2)
	valueRange := maxValue - minValue + 1

	if len(cards) != valueRange*2 {
		return
	}

	exist := true
	for i := minValue; i <= maxValue; i++ {
		if count[i] != 2 {
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
	return
}

// 三不带
func isSanBuDai(cards []*Card) (cardType Type) {
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
func isSanDaiYi(cards []*Card) (cardType Type) {
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
func isSanDaiEr(cards []*Card) (cardType Type) {
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
func isSiDaiDan(cards []*Card) (cardType Type) {
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
func isSiDaiDui(cards []*Card) (cardType Type) {
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
func isShunZi(cards []*Card) (cardType Type) {
	if len(cards) < 5 {
		return
	}

	count := cardCount(cards)
	minValue := getMinValue(count, 1)
	maxValue := getMaxValue(count, 1)
	valueRange := maxValue - minValue + 1

	if len(cards) != valueRange {
		return
	}

	exist := true
	for i := minValue; i <= maxValue; i++ {
		if count[i] != 1 {
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
	return
}

// 飞机不带
func isFeiJiBuDai(cards []*Card) (cardType Type) {
	if len(cards) < 6 {
		return
	}

	count := cardCount(cards)
	minValue := getMinValue(count, 3)
	maxValue := getMaxValue(count, 3)
	valueRange := maxValue - minValue + 1

	if len(cards) != valueRange*3 {
		return
	}

	exist := true
	for i := minValue; i <= maxValue; i++ {
		if count[i] != 3 {
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
	return
}

// 飞机带一
func isFeiJiDaiYi(cards []*Card) (cardType Type) {
	if len(cards) < 8 {
		return
	}

	// 连炸不是飞机带一
	if isLianZha(cards).CardType != CardTypeNone {
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

	exist := true
	for i := minValue; i <= maxValue; i++ {
		if count[i] < 3 {
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
	return
}

// 飞机带二
func isFeiJiDaiEr(cards []*Card) (cardType Type) {
	if len(cards) < 10 {
		return
	}

	count := cardCount(cards)

	var times int
	for i := 3; i <= 15; i++ {
		if count[i] == 2 {
			times++
		}
	}

	minValue := getMinValue(count, 3)
	maxValue := getMaxValue(count, 3)
	valueRange := maxValue - minValue + 1

	if len(cards) != valueRange*5 || times != valueRange {
		return
	}

	exist := true
	for i := minValue; i <= maxValue; i++ {
		if count[i] != 3 {
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
	return
}

// 炸弹
func isZhaDan(cards []*Card) (cardType Type) {
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
func isHuoJian(cards []*Card) (cardType Type) {
	if len(cards) != 2 {
		return
	}

	count := cardCount(cards)
	if count[16] == 1 && count[17] == 1 {
		cardType.CardType = CardTypeHuoJian
		return
	}
	return
}

// 连炸
func isLianZha(cards []*Card) (cardType Type) {
	if len(cards) < 8 {
		return
	}

	count := cardCount(cards)
	minValue := getMinValue(count, 4)
	maxValue := getMaxValue(count, 4)
	valueRange := maxValue - minValue + 1

	if len(cards) != valueRange*4 {
		return
	}

	exist := true
	for i := minValue; i <= maxValue; i++ {
		if count[i] != 4 {
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

// 获取最小的牌
func getMinValue(count [18]int, repeatCount int) int {
	for i := 3; i <= 14; i++ {
		if count[i] == repeatCount {
			return i
		}
	}
	return 0
}

// 获取最大的牌
func getMaxValue(count [18]int, repeatCount int) int {
	for i := 14; i >= 3; i-- {
		if count[i] == repeatCount {
			return i
		}
	}
	return 0
}
