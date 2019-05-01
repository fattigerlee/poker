package ddz

import "sort"

// 找出比牌型更大的牌
func FindCards(info CardTypeInfo, cards []*Card) (ret []*Card) {
	newCards := getNoLaiZiCards(cards)       // 无癞子牌数组
	laiZiCount := len(cards) - len(newCards) // 癞子数量

	if laiZiCount == 0 {
		// 无癞子找牌
		return findCards(&info, cards)
	} else {
		// 癞子找牌
		return findCardsLaiZi(&info, cards, newCards, laiZiCount)
	}
}

// 无癞子找牌
func findCards(info *CardTypeInfo, cards []*Card) (ret []*Card) {
	// 排序
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Num < cards[j].Num
	})

	size := len(cards)
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

	switch info.CardType {
	case CardTypeDan:
		return findBigDan(size, info, cards, value)

	case CardTypeDui:
		if ret = findBigDui(size, info, cards, value); len(ret) != 0 {
			return
		}
		//case CardTypeLianDui:
		//	if ret = findBigLianDui(cards, count, info); len(ret) != 0 {
		//		return
		//	}
		//case CardTypeSanBuDai:
		//	if ret = findBigSanBuDai(cards, count, info); len(ret) != 0 {
		//		return
		//	}
		//case CardTypeSanDaiYi:
		//	if ret = findBigSanDaiYi(cards, count, info); len(ret) != 0 {
		//		return
		//	}
		//case CardTypeSanDaiEr:
		//	if ret = findBigSanDaiEr(cards, count, info); len(ret) != 0 {
		//		return
		//	}
		//case CardTypeShunZi:
		//	if ret = findBigShunZi(cards, count, info); len(ret) != 0 {
		//		return
		//	}
		//case CardTypeFeiJiBuDai:
		//	if ret = findBigFeiJiBuDai(cards, count, info); len(ret) != 0 {
		//		return
		//	}
		//case CardTypeFeiJiDaiYi:
		//	if ret = findBigFeiJiDaiYi(cards, count, info); len(ret) != 0 {
		//		return
		//	}
		//case CardTypeFeiJiDaiEr:
		//	if ret = findBigFeiJiDaiEr(cards, count, info); len(ret) != 0 {
		//		return
		//	}
		//case CardTypeZhaDan:
		//	if ret = findBigZhaDan(cards, count, info); len(ret) != 0 {
		//		return
		//	}
		//
		//	if ret = findHuoJian(cards, count); len(ret) != 0 {
		//		return
		//	}
		//
		//	if ret = findLianZha(cards, count); len(ret) != 0 {
		//		return
		//	}
		//case CardTypeHuoJian:
		//	if ret = findBigHuoJian(cards, count); len(ret) != 0 {
		//		return
		//	}
		//case CardTypeLianZha:
		//	if ret = findBigLianZha(cards, count, info); len(ret) != 0 {
		//		return
		//	}
	}
	return
}

// 取牌
func getCards(cards []*Card, num int, count int) (ret []*Card) {
	for i := 0; i < len(cards); i++ {
		if cards[i].Num == NumType(num) {
			ret = append(ret, cards[i:i+count]...)
			return
		}
	}
	return
}

// 找炸弹
func findZhaDan(cards []*Card, value valueList) (ret []*Card) {
	// 炸弹
	if len(value[4]) != 0 {
		return getCards(cards, value[4][0], 4)
	}

	// 火箭
	if ret = findHuoJian(cards, value); len(ret) != 0 {
		return
	}
	return
}

// 找火箭
func findHuoJian(cards []*Card, value valueList) (ret []*Card) {
	if huoJian(value) {
		ret = append(ret, getCards(cards, NumTypeSmallJoker, 1)...)
		ret = append(ret, getCards(cards, NumTypeBigJoker, 1)...)
	}
	return
}

// 找单
func findBigDan(size int, info *CardTypeInfo, cards []*Card, value valueList) (ret []*Card) {
	if size < 1 {
		return
	}

	// 手上有火箭
	if huoJian(value) {
		// 更大的单
		for _, v := range value[1] {
			if isJoker(v) {
				continue
			}

			if info.MinValue < v {
				return getCards(cards, v, 1)
			}
		}
	} else {
		// 更大的单
		for _, v := range value[1] {
			if info.MinValue < v {
				return getCards(cards, v, 1)
			}
		}
	}

	// 找炸弹
	if ret = findZhaDan(cards, value); len(ret) != 0 {
		return
	}

	// 拆对子
	for _, v := range value[2] {
		if info.MinValue < v {
			return getCards(cards, v, 1)
		}
	}

	// 拆三张
	for _, v := range value[3] {
		if info.MinValue < v {
			return getCards(cards, v, 1)
		}
	}
	return
}

// 找对
func findBigDui(size int, info *CardTypeInfo, cards []*Card, value valueList) (ret []*Card) {
	if size < 2 {
		return
	}

	// 更大的对
	for _, v := range value[2] {
		if info.MinValue < v {
			return getCards(cards, v, 2)
		}
	}

	// 找炸弹
	if ret = findZhaDan(cards, value); len(ret) != 0 {
		return
	}

	// 拆三张
	for _, v := range value[3] {
		if info.MinValue < v {
			return getCards(cards, v, 2)
		}
	}
	return
}

// 找三不带
func findBigSanBuDai(size int, info *CardTypeInfo, cards []*Card, value valueList) (ret []*Card) {
	if size < 3 {
		return
	}

	// 更大的三不带
	for _, v := range value[3] {
		if info.MinValue < v {
			return getCards(cards, v, 3)
		}
	}

	// 找炸弹
	if ret = findZhaDan(cards, value); len(ret) != 0 {
		return
	}
	return
}

// 找三带一
func findBigSanDaiYi(size int, info *CardTypeInfo, cards []*Card, value valueList) (ret []*Card) {
	if size < 4 {
		return
	}

	// 更大的三不带
	for _, v := range value[3] {
		if info.MinValue < v {
			ret = getCards(cards, v, 3)
			break
		}
	}

	if len(ret) != 0 {
		// 找单张
		for _, v := range value[1] {

		}
	}

	// 找炸弹
	if ret = findZhaDan(cards, value); len(ret) != 0 {
		return
	}
	return
}

// 癞子找牌
func findCardsLaiZi(info *CardTypeInfo, cards []*Card, newCards []*Card, laiZiCount int) (ret []*Card) {
	return
}
