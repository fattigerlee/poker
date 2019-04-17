package ddz

// 牌型
type CardTypeInfo struct {
	CardType CardType // 牌型
	MinValue int      // 最小值
	MaxValue int      // 最大值
}

// 获取牌型
func GetCardType(cards []*Card) (info CardTypeInfo) {
	switch len(cards) {
	case 1:
		// 单
		return isDan(cards)
	case 2:
		// 对
		if info = isDui(cards); info.CardType != CardTypeNone {
			return
		}

		// 火箭
		if info = isHuoJian(cards); info.CardType != CardTypeNone {
			return
		}
	case 3:
		// 三不带
		if info = isSanBuDai(cards); info.CardType != CardTypeNone {
			return
		}
	case 4:
		// 三带一
		if info = isSanDaiYi(cards); info.CardType != CardTypeNone {
			return
		}

		// 炸弹
		if info = isZhaDan(cards); info.CardType != CardTypeNone {
			return
		}
	case 5:
		// 顺子
		if info = isShunZi(cards); info.CardType != CardTypeNone {
			return
		}

		// 三带二
		if info = isSanDaiEr(cards); info.CardType != CardTypeNone {
			return
		}
	case 6:
		// 顺子
		if info = isShunZi(cards); info.CardType != CardTypeNone {
			return
		}

		// 连对
		if info = isLianDui(cards); info.CardType != CardTypeNone {
			return
		}

		// 飞机不带
		if info = isFeiJiBuDai(cards); info.CardType != CardTypeNone {
			return
		}

		// 四带单
		if info = isSiDaiDan(cards); info.CardType != CardTypeNone {
			return
		}
	case 7:
		// 顺子
		if info = isShunZi(cards); info.CardType != CardTypeNone {
			return
		}
	case 8:
		// 连对
		if info = isLianDui(cards); info.CardType != CardTypeNone {
			return
		}

		// 飞机带一
		if info = isFeiJiDaiYi(cards); info.CardType != CardTypeNone {
			return
		}

		// 连炸
		if info = isLianZha(cards); info.CardType != CardTypeNone {
			return
		}

		// 顺子
		if info = isShunZi(cards); info.CardType != CardTypeNone {
			return
		}

		// 四带对
		if info = isSiDaiDui(cards); info.CardType != CardTypeNone {
			return
		}
	case 9:
		// 顺子
		if info = isShunZi(cards); info.CardType != CardTypeNone {
			return
		}

		// 飞机不带
		if info = isFeiJiBuDai(cards); info.CardType != CardTypeNone {
			return
		}
	case 10:
		// 飞机带二
		if info = isFeiJiDaiEr(cards); info.CardType != CardTypeNone {
			return
		}

		// 连对
		if info = isLianDui(cards); info.CardType != CardTypeNone {
			return
		}

		// 顺子
		if info = isShunZi(cards); info.CardType != CardTypeNone {
			return
		}
	case 11:
		// 顺子
		if info = isShunZi(cards); info.CardType != CardTypeNone {
			return
		}
	case 12:
		// 连对
		if info = isLianDui(cards); info.CardType != CardTypeNone {
			return
		}

		// 连炸
		if info = isLianZha(cards); info.CardType != CardTypeNone {
			return
		}

		// 飞机不带
		if info = isFeiJiBuDai(cards); info.CardType != CardTypeNone {
			return
		}

		// 飞机带一
		if info = isFeiJiDaiYi(cards); info.CardType != CardTypeNone {
			return
		}

		// 顺子
		if info = isShunZi(cards); info.CardType != CardTypeNone {
			return
		}
	case 14:
		// 连对
		if info = isLianDui(cards); info.CardType != CardTypeNone {
			return
		}
	case 15:
		// 飞机不带
		if info = isFeiJiBuDai(cards); info.CardType != CardTypeNone {
			return
		}

		// 飞机带二
		if info = isFeiJiDaiEr(cards); info.CardType != CardTypeNone {
			return
		}
	case 16:
		// 连对
		if info = isLianDui(cards); info.CardType != CardTypeNone {
			return
		}

		// 连炸
		if info = isLianZha(cards); info.CardType != CardTypeNone {
			return
		}

		// 飞机带一
		if info = isFeiJiDaiYi(cards); info.CardType != CardTypeNone {
			return
		}
	case 18:
		// 连对
		if info = isLianDui(cards); info.CardType != CardTypeNone {
			return
		}
	case 20:
		// 连对
		if info = isLianDui(cards); info.CardType != CardTypeNone {
			return
		}

		// 连炸
		if info = isLianZha(cards); info.CardType != CardTypeNone {
			return
		}

		// 飞机带一
		if info = isFeiJiDaiYi(cards); info.CardType != CardTypeNone {
			return
		}

		// 飞机带二
		if info = isFeiJiDaiEr(cards); info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 单
func isDan(cards []*Card) (info CardTypeInfo) {
	if len(cards) != 1 {
		return
	}

	info.CardType = CardTypeDan
	info.MinValue = int(cards[0].Num)
	return
}

// 对
func isDui(cards []*Card) (info CardTypeInfo) {
	if len(cards) != 2 {
		return
	}

	// 牌值相同
	if cards[0].Num != cards[1].Num {
		return
	}

	info.CardType = CardTypeDui
	info.MinValue = int(cards[0].Num)
	return
}

// 连对
func isLianDui(cards []*Card) (info CardTypeInfo) {
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
		info.CardType = CardTypeLianDui
		info.MinValue = minValue
		info.MaxValue = maxValue
		return
	}
	return
}

// 三不带
func isSanBuDai(cards []*Card) (info CardTypeInfo) {
	if len(cards) != 3 {
		return
	}

	count := cardCount(cards)
	for i := 3; i <= 15; i++ {
		if count[i] == 3 {
			info.CardType = CardTypeSanBuDai
			info.MinValue = i
			return
		}
	}
	return
}

// 三带一
func isSanDaiYi(cards []*Card) (info CardTypeInfo) {
	if len(cards) != 4 {
		return
	}

	count := cardCount(cards)
	for i := 3; i <= 15; i++ {
		if count[i] == 3 {
			info.CardType = CardTypeSanDaiYi
			info.MinValue = i
			return
		}
	}
	return
}

// 三带二
func isSanDaiEr(cards []*Card) (info CardTypeInfo) {
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
			info.CardType = CardTypeSanDaiEr
			info.MinValue = i
			return
		}
	}
	return
}

// 四带单(炸弹带两张单)
func isSiDaiDan(cards []*Card) (info CardTypeInfo) {
	if len(cards) != 6 {
		return
	}

	count := cardCount(cards)
	for i := 3; i <= 15; i++ {
		if count[i] == 4 {
			info.CardType = CardTypeSiDaiDan
			info.MinValue = i
			return
		}
	}
	return
}

// 四带对(炸弹带两对)
func isSiDaiDui(cards []*Card) (info CardTypeInfo) {
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
			info.CardType = CardTypeSiDaiDui
			info.MinValue = i
			return
		}
	}
	return
}

// 顺子
func isShunZi(cards []*Card) (info CardTypeInfo) {
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
		info.CardType = CardTypeShunZi
		info.MinValue = minValue
		info.MaxValue = maxValue
		return
	}
	return
}

// 飞机不带
func isFeiJiBuDai(cards []*Card) (info CardTypeInfo) {
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
		info.CardType = CardTypeFeiJiBuDai
		info.MinValue = minValue
		info.MaxValue = maxValue
		return
	}
	return
}

// 飞机带一
func isFeiJiDaiYi(cards []*Card) (info CardTypeInfo) {
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
		info.CardType = CardTypeFeiJiDaiYi
		info.MinValue = minValue
		info.MaxValue = maxValue
		return
	}
	return
}

// 飞机带二
func isFeiJiDaiEr(cards []*Card) (info CardTypeInfo) {
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
		info.CardType = CardTypeFeiJiDaiEr
		info.MinValue = minValue
		info.MaxValue = maxValue
		return
	}
	return
}

// 炸弹
func isZhaDan(cards []*Card) (info CardTypeInfo) {
	if len(cards) != 4 {
		return
	}

	count := cardCount(cards)
	for i := 3; i <= 15; i++ {
		if count[i] == 4 {
			info.CardType = CardTypeZhaDan
			info.MinValue = i
			return
		}
	}
	return
}

// 火箭
func isHuoJian(cards []*Card) (info CardTypeInfo) {
	if len(cards) != 2 {
		return
	}

	count := cardCount(cards)
	if count[16] == 1 && count[17] == 1 {
		info.CardType = CardTypeHuoJian
		return
	}
	return
}

// 连炸
func isLianZha(cards []*Card) (info CardTypeInfo) {
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
		info.CardType = CardTypeLianZha
		info.MinValue = minValue
		info.MaxValue = maxValue
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
