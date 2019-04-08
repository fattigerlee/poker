package ddz

// 底牌牌型判定
func GetBottomCardType(cards []*Card) BottomCardType {
	if wangZha(cards) {
		return BottomCardTypeWangZha
	}

	if danWang(cards) {
		return BottomCardTypeDanWang
	}

	if tongHuaShun(cards) {
		return BottomCardTypeTongHuaShun
	}

	if tongHua(cards) {
		return BottomCardTypeTongHua
	}

	if shunZi(cards) {
		return BottomCardTypeShunZi
	}

	if sanBuDai(cards) {
		return BottomCardTypeSanBuDai
	}
	return BottomCardTypeNone
}

// 王炸
func wangZha(cards []*Card) bool {
	count := cardCount(cards)
	if count[16] == 1 && count[17] == 1 {
		return true
	}
	return false
}

// 单王
func danWang(cards []*Card) bool {
	count := cardCount(cards)

	if count[16] == 1 || count[17] == 1 {
		return true
	}
	return false
}

// 同花顺
func tongHuaShun(cards []*Card) bool {
	if !tongHua(cards) {
		return false
	}

	if !shunZi(cards) {
		return false
	}
	return true
}

// 同花
func tongHua(cards []*Card) bool {
	for i := 0; i < len(cards)-1; i++ {
		if cards[i].Suit != cards[i+1].Suit {
			return false
		}
	}
	return true
}

// 顺子
func shunZi(cards []*Card) bool {
	count := cardCount(cards)
	minValue := getMinValue(count, 1)
	maxValue := getMaxValue(count, 1)
	valueRange := maxValue - minValue + 1

	if len(cards) != valueRange {
		return false
	}

	exist := true
	for i := minValue; i <= maxValue; i++ {
		if count[i] != 1 {
			exist = false
			break
		}
	}

	if exist {
		return true
	}
	return false
}

// 三不带
func sanBuDai(cards []*Card) bool {
	count := cardCount(cards)
	for i := 3; i <= 15; i++ {
		if count[i] == 3 {
			return true
		}
	}
	return false
}
