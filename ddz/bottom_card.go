package ddz

// 底牌牌型判定
func GetBottomCardType(cards []*Card) BottomCardType {
	cardsCount := cardCount(cards) // 每张牌数量
	var value valueList            // 所有单张,对子,三张,四张的牌值
	var line []int                 // 连牌
	for i := 3; i < 18; i++ {
		if cardsCount[i] > 0 {
			line = append(line, i)
		}

		switch cardsCount[i] {
		case 1:
			value[1] = append(value[1], i)
		case 2:
			value[2] = append(value[2], i)
		case 3:
			value[3] = append(value[3], i)
		case 4:
			value[4] = append(value[4], i)
		}
	}

	if huoJian(value) {
		return BottomCardTypeHuoJian
	}

	if oneJoker(value) {
		return BottomCardTypeOneJoker
	}

	if tongHuaShun(cards, value, line) {
		return BottomCardTypeTongHuaShun
	}

	if tongHua(cards) {
		return BottomCardTypeTongHua
	}

	if shunZi(value, line) {
		return BottomCardTypeShunZi
	}

	if sanBuDai(value) {
		return BottomCardTypeSanBuDai
	}
	return BottomCardTypeNone
}

// 火箭
func huoJian(value valueList) bool {
	var smallJoker bool
	var bigJoker bool
	for _, v := range value[1] {
		if v == NumTypeSmallJoker {
			smallJoker = true
		}

		if v == NumTypeBigJoker {
			bigJoker = true
		}
	}

	if smallJoker == true && bigJoker == true {
		return true
	}
	return false
}

// 单王
func oneJoker(value valueList) bool {
	var joker int
	for _, v := range value[1] {
		if v == NumTypeSmallJoker || v == NumTypeBigJoker {
			joker++
		}
	}

	if joker == 1 {
		return true
	}
	return false
}

// 同花顺
func tongHuaShun(cards []*Card, value valueList, line []int) bool {
	if !tongHua(cards) {
		return false
	}

	if !shunZi(value, line) {
		return false
	}
	return true
}

// 同花
func tongHua(cards []*Card) bool {
	if len(cards) == 0 {
		return false
	}

	suit := cards[0].Suit
	for _, c := range cards {
		if suit != c.Suit {
			return false
		}
	}
	return true
}

// 顺子
func shunZi(value valueList, line []int) bool {
	// 只能有单张
	if len(value[2]) != 0 || len(value[3]) != 0 || len(value[4]) != 0 {
		return false
	}

	// 不能有2和大小王
	for _, v := range line {
		if isJokerAndTwo(v) {
			return false
		}
	}

	// 必须连续
	valueRange := line[len(line)-1] - line[0] + 1
	if valueRange != len(line) {
		return false
	}
	return true
}

// 三不带
func sanBuDai(value valueList) bool {
	// 只能有三张
	if len(value[1]) != 0 || len(value[2]) != 0 || len(value[4]) != 0 {
		return false
	}

	if len(value[3]) != 1 {
		return false
	}
	return true
}
