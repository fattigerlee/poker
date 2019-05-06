package ddz

// 拆牌
func SplitCards(cards []*Card) []*CardTypeInfo {
	return nil
}

//  无癞子拆牌
func splitCards(cards []*Card) []*CardTypeInfo {
	//size := len(cards)
	//myCards := convertToMap(cards)
	//count, value, line := getCountValueLine(myCards)

	//size := len(cards)             // 牌总数量
	//cardsCount := cardCount(cards) // 每张牌数量
	//var value valueList            // 所有单张,对子,三张,四张的牌值
	//var line []int                 // 连牌
	//for i := 3; i < 18; i++ {
	//	if cardsCount[i] > 0 {
	//		line = append(line, i)
	//	}
	//
	//	switch cardsCount[i] {
	//	case 1:
	//		value[1] = append(value[1], i)
	//	case 2:
	//		value[2] = append(value[2], i)
	//	case 3:
	//		value[3] = append(value[3], i)
	//	case 4:
	//		value[4] = append(value[4], i)
	//	}
	//}
	return nil
}

// 连炸
func splitLianZha(size int, value valueList, cards map[*Card]bool) (ret []*CardTypeInfo) {
	if size < 8 {
		return
	}

	for i := 0; i < len(value[4]); i++ {

	}
	return
}

// 癞子拆牌
func splitCardsLaiZi() {

}
