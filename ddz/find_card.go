package ddz

// 找出比牌型更大的牌

// 找牌(经典模式)
func FindCardsJingDian(info *CardTypeInfo, cards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	size := len(cards)
	dictCards := convertToMap(cards)
	count, value, _ := getCountValueLine(dictCards)

	switch info.CardType {
	case CardTypeDan:
		if retCards, retInfo = findBigDan(size, info, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeDui:
		if retCards, retInfo = findBigDui(size, info, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSanBuDai:
		if retCards, retInfo = findBigSanBuDai(size, info, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSanDaiYi:
		if retCards, retInfo = findBigSanDaiYi(size, info, cards, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSanDaiEr:
		if retCards, retInfo = findBigSanDaiEr(size, info, cards, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSiDaiDan:
		if retCards, retInfo = findBigSiDaiDan(size, info, cards, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSiDaiDui:
		if retCards, retInfo = findBigSiDaiDui(size, info, cards, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeShunZi:
		if retCards, retInfo = findBigShunZi(size, info, dictCards, count); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeLianDui:
		if retCards, retInfo = findBigLianDui(size, info, dictCards, count); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeFeiJiBuDai:
		if retCards, retInfo = findBigFeiJiBuDai(size, info, dictCards, count); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeFeiJiDaiYi:
		if retCards, retInfo = findBigFeiJiDaiYi(size, info, cards, dictCards, count, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeFeiJiDaiEr:
		if retCards, retInfo = findBigFeiJiDaiEr(size, info, cards, dictCards, count, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeZhaDan:
		if retCards, retInfo = findBigZhaDan(size, info, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeZhaDan {
		// 炸弹
		if retCards, retInfo = findZhaDan(size, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeHuoJian {
		// 火箭
		if retCards, retInfo = findHuoJian(size, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 找牌(不洗牌模式)
func FindCardsBuXiPai(info *CardTypeInfo, cards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	size := len(cards)
	dictCards := convertToMap(cards)
	count, value, _ := getCountValueLine(dictCards)

	switch info.CardType {
	case CardTypeDan:
		if retCards, retInfo = findBigDan(size, info, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeDui:
		if retCards, retInfo = findBigDui(size, info, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSanBuDai:
		if retCards, retInfo = findBigSanBuDai(size, info, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSanDaiYi:
		if retCards, retInfo = findBigSanDaiYi(size, info, cards, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSanDaiEr:
		if retCards, retInfo = findBigSanDaiEr(size, info, cards, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSiDaiDan:
		if retCards, retInfo = findBigSiDaiDan(size, info, cards, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSiDaiDui:
		if retCards, retInfo = findBigSiDaiDui(size, info, cards, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeShunZi:
		if retCards, retInfo = findBigShunZi(size, info, dictCards, count); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeLianDui:
		if retCards, retInfo = findBigLianDui(size, info, dictCards, count); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeFeiJiBuDai:
		if retCards, retInfo = findBigFeiJiBuDai(size, info, dictCards, count); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeFeiJiDaiYi:
		if retCards, retInfo = findBigFeiJiDaiYi(size, info, cards, dictCards, count, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeFeiJiDaiEr:
		if retCards, retInfo = findBigFeiJiDaiEr(size, info, cards, dictCards, count, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeZhaDan:
		if retCards, retInfo = findBigZhaDan(size, info, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeHuoJian:
		if retCards, retInfo = findLianZha(size, dictCards, count, value, 2); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeLianZha:
		if retCards, retInfo = findBigLianZha(size, info, dictCards, count); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeZhaDan {
		// 炸弹
		if retCards, retInfo = findZhaDan(size, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeHuoJian {
		// 火箭
		if retCards, retInfo = findHuoJian(size, dictCards, value); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeLianZha {
		// 连炸
		if retCards, retInfo = findLianZha(size, dictCards, count, value, 2); retInfo.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 找牌(不洗牌+双王癞子模式)
func FindCardsBuXiPaiLaiZi(info *CardTypeInfo, cards []*Card, laiZiNums ...NumType) (retCards []*Card, retInfo CardTypeInfo) {
	var normalCards []*Card // 普通牌
	var laiZiCards []*Card  // 癞子牌
	for _, c := range cards {
		var exist bool
		for _, num := range laiZiNums {
			if c.Num == num {
				exist = true
				break
			}
		}

		if exist {
			laiZiCards = append(laiZiCards, c)
		} else {
			normalCards = append(normalCards, c)
		}
	}

	size := len(cards)           // 手牌数量
	laiZiSize := len(laiZiCards) // 癞子牌数量

	normalDictCards := convertToMap(normalCards)
	normalCount, normalValue, _ := getCountValueLine(normalDictCards)

	switch info.CardType {
	case CardTypeDan:
		if retCards, retInfo = findBigDanLaiZi(size, laiZiSize, info, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeDui:
		if retCards, retInfo = findBigDuiLaiZi(size, laiZiSize, info, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeSanBuDai:
		if retCards, retInfo = findBigSanBuDaiLaiZi(size, laiZiSize, info, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeSanDaiYi:
		if retCards, retInfo = findBigSanDaiYiLaiZi(size, laiZiSize, info, normalCards, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeSanDaiEr:
		if retCards, retInfo = findBigSanDaiErLaiZi(size, laiZiSize, info, normalCards, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeSiDaiDan:
		if retCards, retInfo = findBigSiDaiDanLaiZi(size, laiZiSize, info, normalCards, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeSiDaiDui:
		if retCards, retInfo = findBigSiDaiDuiLaiZi(size, laiZiSize, info, normalCards, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeShunZi:
		if retCards, retInfo = findBigShunZiLaiZi(size, laiZiSize, info, normalDictCards, normalCount, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeLianDui:
		if retCards, retInfo = findBigLianDuiLaiZi(size, laiZiSize, info, normalDictCards, normalCount, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeFeiJiBuDai:
		if retCards, retInfo = findBigFeiJiBuDaiLaiZi(size, laiZiSize, info, normalDictCards, normalCount, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeFeiJiDaiYi:
		if retCards, retInfo = findBigFeiJiDaiYiLaiZi(size, laiZiSize, info, normalCards, normalDictCards, normalCount, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeFeiJiDaiEr:
		if retCards, retInfo = findBigFeiJiDaiErLaiZi(size, laiZiSize, info, normalCards, normalDictCards, normalCount, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeRuanZhaDan4, CardTypeZhaDan:
		// 更大的硬炸弹
		if retCards, retInfo = findBigZhaDan(size, info, normalDictCards, normalValue); retInfo.CardType != CardTypeNone {
			return
		}

		// 更大的四软炸
		if retCards, retInfo = findBigRuanZhaDan4(size, laiZiSize, info, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

		// 火箭癞子
		laiZiDictCards := convertToMap(laiZiCards)
		_, laiZiValue, _ := getCountValueLine(laiZiDictCards)
		if retCards, retInfo = findHuoJian(size, laiZiDictCards, laiZiValue); retInfo.CardType != CardTypeNone {
			return
		}

		// 硬连炸
		if retCards, retInfo = findLianZha(size, normalDictCards, normalCount, normalValue, 2); retInfo.CardType != CardTypeNone {
			return
		}

		// 软连炸
		if retCards, retInfo = findRuanLianZha(size, laiZiSize, normalDictCards, normalCount, laiZiCards, 2); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeHuoJian:
		// 硬连炸
		if retCards, retInfo = findLianZha(size, normalDictCards, normalCount, normalValue, 2); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeRuanZhaDan5:
		// 更大的五软炸
		if retCards, retInfo = findBigRuanZhaDan5(size, laiZiSize, info, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

		// 硬连炸
		if retCards, retInfo = findLianZha(size, normalDictCards, normalCount, normalValue, 2); retInfo.CardType != CardTypeNone {
			return
		}

		// 软连炸
		if retCards, retInfo = findRuanLianZha(size, laiZiSize, normalDictCards, normalCount, laiZiCards, 2); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeRuanZhaDan6:
		// 硬连炸
		if retCards, retInfo = findLianZha(size, normalDictCards, normalCount, normalValue, 3); retInfo.CardType != CardTypeNone {
			return
		}

	case CardTypeRuanLianZha, CardTypeLianZha:
		// 更大的硬连炸
		if retCards, retInfo = findBigLianZha(size, info, normalDictCards, normalCount); retInfo.CardType != CardTypeNone {
			return
		}

		// 更大的软连炸
		if retCards, retInfo = findBigRuanLianZha(size, laiZiSize, info, normalDictCards, normalCount, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	}

	// 普通牌型
	if info.CardType <= CardTypeFeiJiDaiEr {
		// 硬炸弹
		if retCards, retInfo = findZhaDan(size, normalDictCards, normalValue); retInfo.CardType != CardTypeNone {
			return
		}

		// 四软炸
		if retCards, retInfo = findRuanZhaDan4(size, laiZiSize, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}

		// 火箭癞子
		laiZiDictCards := convertToMap(laiZiCards)
		_, laiZiValue, _ := getCountValueLine(laiZiDictCards)
		if retCards, retInfo = findHuoJian(size, laiZiDictCards, laiZiValue); retInfo.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 找牌(癞子模式)

// 找牌(天地癞子模式)
func FindCardsTianDiLaiZi(info *CardTypeInfo, cards []*Card, laiZiNums ...NumType) (retCards []*Card, retInfo CardTypeInfo) {
	var normalCards []*Card // 普通牌
	var laiZiCards []*Card  // 癞子牌
	for _, c := range cards {
		var exist bool
		for _, num := range laiZiNums {
			if c.Num == num {
				exist = true
				break
			}
		}

		if exist {
			laiZiCards = append(laiZiCards, c)
		} else {
			normalCards = append(normalCards, c)
		}
	}

	size := len(cards)           // 手牌数量
	laiZiSize := len(laiZiCards) // 癞子牌数量

	if laiZiSize == 0 {
		return FindCardsJingDian(info, cards)
	}

	normalDictCards := convertToMap(normalCards)
	normalCount, normalValue, _ := getCountValueLine(normalDictCards)

	switch info.CardType {
	case CardTypeDan:
		if retCards, retInfo = findBigDanLaiZi(size, laiZiSize, info, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeDui:
		if retCards, retInfo = findBigDuiLaiZi(size, laiZiSize, info, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSanBuDai:
		if retCards, retInfo = findBigSanBuDaiLaiZi(size, laiZiSize, info, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSanDaiYi:
		if retCards, retInfo = findBigSanDaiYiLaiZi(size, laiZiSize, info, normalCards, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSanDaiEr:
		if retCards, retInfo = findBigSanDaiErLaiZi(size, laiZiSize, info, normalCards, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSiDaiDan:
		if retCards, retInfo = findBigSiDaiDanLaiZi(size, laiZiSize, info, normalCards, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeSiDaiDui:
		if retCards, retInfo = findBigSiDaiDuiLaiZi(size, laiZiSize, info, normalCards, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeShunZi:
		if retCards, retInfo = findBigShunZiLaiZi(size, laiZiSize, info, normalDictCards, normalCount, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeLianDui:
		if retCards, retInfo = findBigLianDuiLaiZi(size, laiZiSize, info, normalDictCards, normalCount, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeFeiJiBuDai:
		if retCards, retInfo = findBigFeiJiBuDaiLaiZi(size, laiZiSize, info, normalDictCards, normalCount, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeFeiJiDaiYi:
		if retCards, retInfo = findBigFeiJiDaiYiLaiZi(size, laiZiSize, info, normalCards, normalDictCards, normalCount, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeFeiJiDaiEr:
		if retCards, retInfo = findBigFeiJiDaiErLaiZi(size, laiZiSize, info, normalCards, normalDictCards, normalCount, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeRuanZhaDan4:
		// 更大的四软炸
		if retCards, retInfo = findBigRuanZhaDan4(size, laiZiSize, info, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeZhaDan:
		// 更大的硬炸弹
		if retCards, retInfo = findBigZhaDan(size, info, normalDictCards, normalValue); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeChunLaiZiZhaDan:
		// 更大的四纯癞子炸
		if retCards, retInfo = findBigChunLaiZiZhaDan(laiZiSize, info, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeRuanZhaDan5:
		// 更大的五软炸
		if retCards, retInfo = findBigRuanZhaDan5(size, laiZiSize, info, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeRuanZhaDan6:
		// 更大的六软炸
		if retCards, retInfo = findBigRuanZhaDan6(size, laiZiSize, info, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeRuanZhaDan7:
		// 更大的七软炸
		if retCards, retInfo = findBigRuanZhaDan7(size, laiZiSize, info, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeRuanZhaDan8:
		// 更大的八软炸
		if retCards, retInfo = findBigRuanZhaDan8(size, laiZiSize, info, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeRuanZhaDan4 {
		// 四软炸
		if retCards, retInfo = findRuanZhaDan4(size, laiZiSize, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeZhaDan {
		// 硬炸弹
		if retCards, retInfo = findZhaDan(size, normalDictCards, normalValue); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeChunLaiZiZhaDan {
		// 四纯癞子炸,四癞子炸
		if retCards, retInfo = findLaiZiZhaDan4(laiZiSize, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeRuanZhaDan5 {
		// 五软炸
		if retCards, retInfo = findRuanZhaDan5(size, laiZiSize, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeLaiZiZhaDan5 {
		// 五癞子炸
		if retCards, retInfo = findLaiZiZhaDan(CardTypeLaiZiZhaDan5, laiZiSize, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeRuanZhaDan6 {
		// 六软炸
		if retCards, retInfo = findRuanZhaDan6(size, laiZiSize, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeLaiZiZhaDan6 {
		// 六癞子炸
		if retCards, retInfo = findLaiZiZhaDan(CardTypeLaiZiZhaDan6, laiZiSize, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeRuanZhaDan7 {
		// 七软炸
		if retCards, retInfo = findRuanZhaDan7(size, laiZiSize, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeRuanZhaDan8 {
		// 八软炸
		if retCards, retInfo = findRuanZhaDan8(size, laiZiSize, normalDictCards, normalValue, laiZiCards); retInfo.CardType != CardTypeNone {
			return
		}
	}

	if info.CardType < CardTypeHuoJian {
		// 火箭
		if retCards, retInfo = findHuoJian(size, normalDictCards, normalValue); retInfo.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 带单张
func withDan(dictCards dictMap, value valueList) (ret []*Card) {
	var nums []int

	// 找单张
	if huoJian(value) {
		// 有火箭
		for _, v := range value[1] {
			if isJoker(v) {
				continue
			}
			nums = append(nums, v)
			ret = append(ret, findCardsByNums(dictCards, nums)...)
			return
		}
	} else {
		for _, v := range value[1] {
			nums = append(nums, v)
			ret = append(ret, findCardsByNums(dictCards, nums)...)
			return
		}
	}

	// 拆对子
	for _, v := range value[2] {
		nums = append(nums, v)
		ret = append(ret, findCardsByNums(dictCards, nums)...)
		return
	}

	// 拆三张
	for _, v := range value[3] {
		nums = append(nums, v)
		ret = append(ret, findCardsByNums(dictCards, nums)...)
		return
	}
	return
}

// 带对子
func withDui(dictCards dictMap, value valueList) (ret []*Card) {
	var nums []int

	// 找对子
	for _, v := range value[2] {
		for i := 0; i < 2; i++ {
			nums = append(nums, v)
		}
		ret = append(ret, findCardsByNums(dictCards, nums)...)
		return
	}

	// 拆三张
	for _, v := range value[3] {
		for i := 0; i < 2; i++ {
			nums = append(nums, v)
		}
		ret = append(ret, findCardsByNums(dictCards, nums)...)
		return
	}
	return
}

// 单张
func findBigDan(size int, info *CardTypeInfo, dictCards dictMap, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 1 {
		return
	}

	var nums []int

	// 手上有火箭
	if huoJian(value) {
		// 更大的单
		for _, v := range value[1] {
			if isJoker(v) {
				continue
			}

			if info.MinValue < v {
				nums = append(nums, v)
				retCards = findCardsByNums(dictCards, nums)
				retInfo.CardType = CardTypeDan
				retInfo.MinValue = v
				return
			}
		}
	} else {
		// 更大的单
		for _, v := range value[1] {
			if info.MinValue < v {
				nums = append(nums, v)
				retCards = findCardsByNums(dictCards, nums)
				retInfo.CardType = CardTypeDan
				retInfo.MinValue = v
				return
			}
		}
	}

	// 拆对子
	for _, v := range value[2] {
		if info.MinValue < v {
			nums = append(nums, v)
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeDan
			retInfo.MinValue = v
			return
		}
	}

	// 拆三张
	for _, v := range value[3] {
		if info.MinValue < v {
			nums = append(nums, v)
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeDan
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 对子
func findBigDui(size int, info *CardTypeInfo, dictCards dictMap, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 2 {
		return
	}

	var nums []int

	// 更大的对
	for _, v := range value[2] {
		if info.MinValue < v {
			for i := 0; i < 2; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeDui
			retInfo.MinValue = v
			return
		}
	}

	// 拆三张
	for _, v := range value[3] {
		if info.MinValue < v {
			for i := 0; i < 2; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeDui
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 三不带
func findBigSanBuDai(size int, info *CardTypeInfo, dictCards dictMap, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 3 {
		return
	}

	var nums []int

	// 更大的三张
	for _, v := range value[3] {
		if info.MinValue < v {
			for i := 0; i < 3; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeSanBuDai
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 三带一
func findBigSanDaiYi(size int, info *CardTypeInfo, cards []*Card, dictCards dictMap, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 4 {
		return
	}

	var nums []int

	// 更大的三张
	for _, v := range value[3] {
		if info.MinValue < v {
			for i := 0; i < 3; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeSanDaiYi
			retInfo.MinValue = v
			break
		}
	}

	if retInfo.CardType == CardTypeNone {
		return
	}

	_, value, _ = getCountValueLine(dictCards)
	retCards = append(retCards, withDan(dictCards, value)...)

	// 没有合适的牌
	if len(retCards) != 4 {
		retCards = retCards[0:0]
		retInfo.Reset()
		dictCards = convertToMap(cards)
		_, value, _ = getCountValueLine(dictCards)
	}
	return
}

// 三带二
func findBigSanDaiEr(size int, info *CardTypeInfo, cards []*Card, dictCards dictMap, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 5 {
		return
	}

	var nums []int

	// 更大的三张
	for _, v := range value[3] {
		if info.MinValue < v {
			for i := 0; i < 3; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeSanDaiEr
			retInfo.MinValue = v
			break
		}
	}

	if retInfo.CardType == CardTypeNone {
		return
	}

	_, value, _ = getCountValueLine(dictCards)
	retCards = append(retCards, withDui(dictCards, value)...)

	// 没有合适的牌
	if len(retCards) != 5 {
		retCards = retCards[0:0]
		retInfo.Reset()
		dictCards = convertToMap(cards)
		_, value, _ = getCountValueLine(dictCards)
	}
	return
}

// 四带单
func findBigSiDaiDan(size int, info *CardTypeInfo, cards []*Card, dictCards dictMap, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 6 {
		return
	}

	var nums []int

	// 更大的四张
	for _, v := range value[4] {
		if info.MinValue < v {
			for i := 0; i < 4; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeSiDaiDan
			retInfo.MinValue = v
			break
		}
	}

	if retInfo.CardType == CardTypeNone {
		return
	}

	for i := 0; i < 2; i++ {
		_, value, _ = getCountValueLine(dictCards)

		ret := withDan(dictCards, value)
		if len(ret) == 0 {
			break
		}
		retCards = append(retCards, ret...)
	}

	// 没有合适的牌
	if len(retCards) != 6 {
		retCards = retCards[0:0]
		retInfo.Reset()
		dictCards = convertToMap(cards)
		_, value, _ = getCountValueLine(dictCards)
	}
	return
}

// 四带对
func findBigSiDaiDui(size int, info *CardTypeInfo, cards []*Card, dictCards dictMap, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 8 {
		return
	}

	var nums []int

	// 更大的四张
	for _, v := range value[4] {
		if info.MinValue < v {
			for i := 0; i < 4; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeSiDaiDui
			retInfo.MinValue = v
			break
		}
	}

	if retInfo.CardType == CardTypeNone {
		return
	}

	for i := 0; i < 2; i++ {
		_, value, _ = getCountValueLine(dictCards)

		ret := withDui(dictCards, value)
		if len(ret) == 0 {
			break
		}
		retCards = append(retCards, ret...)
	}

	// 没有合适的牌
	if len(retCards) != 8 {
		retCards = retCards[0:0]
		retInfo.Reset()
		dictCards = convertToMap(cards)
		_, value, _ = getCountValueLine(dictCards)
	}
	return
}

// 顺子
func findBigShunZi(size int, info *CardTypeInfo, dictCards dictMap, count countList) (retCards []*Card, retInfo CardTypeInfo) {
	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange {
		return
	}

	var nums []int

	// 更大的顺子
	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		for j := i; j < i+valueRange; j++ {
			switch count[j] {
			case 1, 2, 3:
				nums = append(nums, j)
			}
		}

		if len(nums) == valueRange {
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeShunZi
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[0] + valueRange - 1
			break
		}
		nums = nums[0:0]
	}
	return
}

// 连对
func findBigLianDui(size int, info *CardTypeInfo, dictCards dictMap, count countList) (retCards []*Card, retInfo CardTypeInfo) {
	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange*2 {
		return
	}

	var nums []int

	// 更大的连对
	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		for j := i; j < i+valueRange; j++ {
			switch count[j] {
			case 2, 3:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
			}
		}

		if len(nums) == valueRange*2 {
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeLianDui
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[0] + valueRange - 1
			break
		}
		nums = nums[0:0]
	}
	return
}

// 飞机不带
func findBigFeiJiBuDai(size int, info *CardTypeInfo, dictCards dictMap, count countList) (retCards []*Card, retInfo CardTypeInfo) {
	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange*3 {
		return
	}

	var nums []int

	// 更大的飞机不带
	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		for j := i; j < i+valueRange; j++ {
			if count[j] == 3 {
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
			}
		}

		if len(nums) == valueRange*3 {
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeFeiJiBuDai
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[0] + valueRange - 1
			break
		}
		nums = nums[0:0]
	}
	return
}

// 飞机带一
func findBigFeiJiDaiYi(size int, info *CardTypeInfo, cards []*Card, dictCards dictMap, count countList, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange*4 {
		return
	}

	var nums []int

	// 更大的飞机带一
	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		for j := i; j < i+valueRange; j++ {
			if count[j] == 3 {
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
			}
		}

		if len(nums) == valueRange*3 {
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeFeiJiDaiYi
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[0] + valueRange - 1
			break
		}
		nums = nums[0:0]
	}

	if retInfo.CardType == CardTypeNone {
		return
	}

	for i := 0; i < valueRange; i++ {
		_, value, _ = getCountValueLine(dictCards)

		ret := withDan(dictCards, value)
		if len(ret) == 0 {
			break
		}
		retCards = append(retCards, ret...)
	}

	// 没有合适的牌
	if len(retCards) != valueRange*4 {
		retCards = retCards[0:0]
		retInfo.Reset()
		dictCards = convertToMap(cards)
		_, value, _ = getCountValueLine(dictCards)
	}
	return
}

// 飞机带二
func findBigFeiJiDaiEr(size int, info *CardTypeInfo, cards []*Card, dictCards dictMap, count countList, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange*5 {
		return
	}

	var nums []int

	// 更大的飞机带二
	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		for j := i; j < i+valueRange; j++ {
			if count[j] == 3 {
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
			}
		}

		if len(nums) == valueRange*3 {
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeFeiJiDaiEr
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[0] + valueRange - 1
			break
		}
		nums = nums[0:0]
	}

	if retInfo.CardType == CardTypeNone {
		return
	}

	for i := 0; i < valueRange; i++ {
		_, value, _ = getCountValueLine(dictCards)

		ret := withDui(dictCards, value)
		if len(ret) == 0 {
			break
		}
		retCards = append(retCards, ret...)
	}

	// 没有合适的牌
	if len(retCards) != valueRange*5 {
		retCards = retCards[0:0]
		retInfo.Reset()
		dictCards = convertToMap(cards)
		_, value, _ = getCountValueLine(dictCards)
	}
	return
}

// 炸弹
func findZhaDan(size int, dictCards dictMap, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 4 {
		return
	}

	var nums []int

	// 炸弹
	for _, v := range value[4] {
		for i := 0; i < 4; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(dictCards, nums)
		retInfo.CardType = CardTypeZhaDan
		retInfo.MinValue = v
	}
	return
}

// 更大的炸弹
func findBigZhaDan(size int, info *CardTypeInfo, dictCards dictMap, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 4 {
		return
	}

	var nums []int

	// 更大的炸弹
	for _, v := range value[4] {
		if v <= info.MinValue {
			continue
		}

		for i := 0; i < 4; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(dictCards, nums)
		retInfo.CardType = CardTypeZhaDan
		retInfo.MinValue = v
	}
	return
}

// 火箭
func findHuoJian(size int, dictCards dictMap, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 2 {
		return
	}

	var nums []int

	if huoJian(value) {
		nums = append(nums, NumTypeSmallJoker, NumTypeBigJoker)
		retCards = findCardsByNums(dictCards, nums)
		retInfo.CardType = CardTypeHuoJian
	}
	return
}

// 连炸
func findLianZha(size int, dictCards dictMap, count countList, value valueList, valueRange int) (retCards []*Card, retInfo CardTypeInfo) {
	if size < valueRange*4 {
		return
	}

	var nums []int

	for _, v := range value[4] {
		for i := v; i < v+valueRange; i++ {
			if v > NumTypeAce {
				break
			}

			if count[i] == 4 {
				for j := 0; j < 4; j++ {
					nums = append(nums, i)
				}
			}
		}

		if len(nums) == valueRange*4 {
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeLianZha
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[0] + valueRange - 1
			break
		}
		nums = nums[0:0]
	}
	return
}

// 更大的连炸
func findBigLianZha(size int, info *CardTypeInfo, dictCards dictMap, count countList) (retCards []*Card, retInfo CardTypeInfo) {
	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange*4 {
		return
	}

	var nums []int

	// 更大的连炸
	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		for j := i; j < i+valueRange; j++ {
			if count[j] == 4 {
				for k := 0; k < 4; k++ {
					nums = append(nums, j)
				}
			}
		}

		if len(nums) == valueRange*4 {
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeLianZha
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[0] + valueRange - 1
			break
		}
		nums = nums[0:0]
	}

	if retInfo.CardType != CardTypeNone {
		return
	}

	// 更多的连炸
	valueRange += 1
	for i := NumTypeThree; i <= NumTypeAce-valueRange+1; i++ {
		for j := i; j < i+valueRange; j++ {
			if count[j] == 4 {
				for k := 0; k < 4; k++ {
					nums = append(nums, j)
				}
			}
		}

		if len(nums) == valueRange*4 {
			retCards = findCardsByNums(dictCards, nums)
			retInfo.CardType = CardTypeLianZha
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[0] + valueRange - 1
			break
		}
		nums = nums[0:0]
	}
	return
}

// 癞子单张
func findBigDanLaiZi(size int, laiZiSize int, info *CardTypeInfo, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	// 非癞子牌里有大牌
	if retCards, retInfo = findBigDan(len(normalDictCards), info, normalDictCards, normalValue); retInfo.CardType != CardTypeNone {
		return
	}

	// 手牌大于三张,不拆癞子
	if size > 3 {
		return
	}

	// 纯癞子
	if laiZiSize < 1 {
		return
	}

	var nums []int

	laiZiDictCards := convertToMap(laiZiCards)
	_, laiZiValue, _ := getCountValueLine(laiZiDictCards)

	// 手上有火箭
	if huoJian(laiZiValue) {
		for _, v := range laiZiValue[1] {
			if isJoker(v) {
				continue
			}

			if info.MinValue < v {
				nums = append(nums, v)
				retCards = findCardsByNums(laiZiDictCards, nums)
				retInfo.CardType = CardTypeDan
				retInfo.MinValue = v
				return
			}
		}
	} else {
		for _, v := range laiZiValue[1] {
			if info.MinValue < v {
				nums = append(nums, v)
				retCards = findCardsByNums(laiZiDictCards, nums)
				retInfo.CardType = CardTypeDan
				retInfo.MinValue = v
				return
			}
		}
	}

	// 拆对子
	for _, v := range laiZiValue[2] {
		if info.MinValue < v {
			nums = append(nums, v)
			retCards = findCardsByNums(laiZiDictCards, nums)
			retInfo.CardType = CardTypeDan
			retInfo.MinValue = v
			return
		}
	}

	// 拆三张
	for _, v := range laiZiValue[3] {
		if info.MinValue < v {
			nums = append(nums, v)
			retCards = findCardsByNums(laiZiDictCards, nums)
			retInfo.CardType = CardTypeDan
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 癞子对子
func findBigDuiLaiZi(size int, laiZiSize int, info *CardTypeInfo, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	// 非癞子牌里有大牌
	if retCards, retInfo = findBigDui(len(normalDictCards), info, normalDictCards, normalValue); retInfo.CardType != CardTypeNone {
		return
	}

	if size < 2 {
		return
	}

	var nums []int

	// 一个癞子
	if laiZiSize < 1 {
		return
	}

	// 补一个癞子
	for _, v := range normalValue[1] {
		if isJoker(v) {
			continue
		}

		if info.MinValue < v {
			if laiZiSize < 1 {
				return
			}

			nums = append(nums, v)
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeDui
			retInfo.MinValue = v
			return
		}
	}

	// 手牌大于三张,不拆癞子
	if size > 3 {
		return
	}

	// 纯癞子
	if laiZiSize < 2 {
		return
	}

	laiZiDictCards := convertToMap(laiZiCards)
	_, laiZiValue, _ := getCountValueLine(laiZiDictCards)

	for _, v := range laiZiValue[2] {
		if info.MinValue < v {
			for i := 0; i < 2; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(laiZiDictCards, nums)
			retInfo.CardType = CardTypeDui
			retInfo.MinValue = v
			return
		}
	}

	// 拆三张
	for _, v := range laiZiValue[3] {
		if info.MinValue < v {
			for i := 0; i < 2; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(laiZiDictCards, nums)
			retInfo.CardType = CardTypeDui
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 癞子三不带
func findBigSanBuDaiLaiZi(size int, laiZiSize int, info *CardTypeInfo, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	// 非癞子牌里有大牌
	if retCards, retInfo = findBigSanBuDai(len(normalDictCards), info, normalDictCards, normalValue); retInfo.CardType != CardTypeNone {
		return
	}

	if size < 3 {
		return
	}

	var nums []int

	// 补一个癞子
	for _, v := range normalValue[2] {
		if info.MinValue < v {
			if laiZiSize < 1 {
				return
			}

			for i := 0; i < 2; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeSanBuDai
			retInfo.MinValue = v
			return
		}
	}

	// 补两个癞子
	for _, v := range normalValue[1] {
		if info.MinValue < v {
			if laiZiSize < 2 {
				return
			}

			nums = append(nums, v)
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeSanBuDai
			retInfo.MinValue = v
			return
		}
	}

	// 手牌大于三张,不拆癞子
	if size > 3 {
		return
	}

	// 纯癞子
	if laiZiSize < 3 {
		return
	}

	laiZiDictCards := convertToMap(laiZiCards)
	_, laiZiValue, _ := getCountValueLine(laiZiDictCards)

	for _, v := range laiZiValue[3] {
		if info.MinValue < v {
			for i := 0; i < 3; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(laiZiDictCards, nums)
			retInfo.CardType = CardTypeSanBuDai
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 癞子三带一
func findBigSanDaiYiLaiZi(size int, laiZiSize int, info *CardTypeInfo, normalCards []*Card, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	// 非癞子牌里有大牌
	if retCards, retInfo = findBigSanDaiYi(len(normalDictCards), info, normalCards, normalDictCards, normalValue); retInfo.CardType != CardTypeNone {
		return
	}

	if size < 4 {
		return
	}

	var nums []int

	// 补一个癞子
	for _, v := range normalValue[2] {
		if info.MinValue < v {
			if laiZiSize < 1 {
				return
			}

			for i := 0; i < 2; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeSanDaiYi
			retInfo.MinValue = v
			break
		}
	}

	if retInfo.CardType != CardTypeNone {
		_, normalValue, _ = getCountValueLine(normalDictCards)
		retCards = append(retCards, withDan(normalDictCards, normalValue)...)

		// 没有合适的牌
		if len(retCards) != 4 {
			retCards = retCards[0:0]
			retInfo.Reset()
			normalDictCards = convertToMap(normalCards)
			_, normalValue, _ = getCountValueLine(normalDictCards)
		}
	}

	// 补两个癞子
	for _, v := range normalValue[1] {
		if info.MinValue < v {
			if laiZiSize < 2 {
				return
			}

			nums = append(nums, v)
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeSanDaiYi
			retInfo.MinValue = v
			return
		}
	}

	if retInfo.CardType != CardTypeNone {
		_, normalValue, _ = getCountValueLine(normalDictCards)
		retCards = append(retCards, withDan(normalDictCards, normalValue)...)

		// 没有合适的牌
		if len(retCards) != 4 {
			retCards = retCards[0:0]
			retInfo.Reset()
			normalDictCards = convertToMap(normalCards)
			_, normalValue, _ = getCountValueLine(normalDictCards)
		}
	}
	return
}

// 癞子三带二
func findBigSanDaiErLaiZi(size int, laiZiSize int, info *CardTypeInfo, normalCards []*Card, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	// 非癞子牌里有大牌
	if retCards, retInfo = findBigSanDaiEr(len(normalDictCards), info, normalCards, normalDictCards, normalValue); retInfo.CardType != CardTypeNone {
		return
	}

	if size < 5 {
		return
	}

	var nums []int

	// 补一个癞子
	for _, v := range normalValue[2] {
		if info.MinValue < v {
			if laiZiSize < 1 {
				return
			}

			for i := 0; i < 2; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeSanDaiEr
			retInfo.MinValue = v
			break
		}
	}

	if retInfo.CardType != CardTypeNone {
		_, normalValue, _ = getCountValueLine(normalDictCards)
		retCards = append(retCards, withDui(normalDictCards, normalValue)...)

		// 没有合适的牌
		if len(retCards) != 5 {
			retCards = retCards[0:0]
			retInfo.Reset()
			normalDictCards = convertToMap(normalCards)
			_, normalValue, _ = getCountValueLine(normalDictCards)
		}
	}

	// 补两个癞子
	for _, v := range normalValue[1] {
		if info.MinValue < v {
			if laiZiSize < 2 {
				return
			}

			nums = append(nums, v)
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeSanDaiEr
			retInfo.MinValue = v
			return
		}
	}

	if retInfo.CardType != CardTypeNone {
		_, normalValue, _ = getCountValueLine(normalDictCards)
		retCards = append(retCards, withDan(normalDictCards, normalValue)...)

		// 没有合适的牌
		if len(retCards) != 5 {
			retCards = retCards[0:0]
			retInfo.Reset()
			normalDictCards = convertToMap(normalCards)
			_, normalValue, _ = getCountValueLine(normalDictCards)
		}
	}
	return
}

// 癞子四带单
func findBigSiDaiDanLaiZi(size int, laiZiSize int, info *CardTypeInfo, normalCards []*Card, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	// 非癞子牌里有大牌
	if retCards, retInfo = findBigSiDaiDan(len(normalDictCards), info, normalCards, normalDictCards, normalValue); retInfo.CardType != CardTypeNone {
		return
	}

	if size < 6 {
		return
	}

	var nums []int

	// 补一个癞子
	for _, v := range normalValue[3] {
		if info.MinValue < v {
			if laiZiSize < 1 {
				return
			}

			for i := 0; i < 3; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeSiDaiDan
			retInfo.MinValue = v
			break
		}
	}

	if retInfo.CardType != CardTypeNone {
		for i := 0; i < 2; i++ {
			_, normalValue, _ = getCountValueLine(normalDictCards)

			ret := withDan(normalDictCards, normalValue)
			if len(ret) == 0 {
				break
			}
			retCards = append(retCards, ret...)
		}

		// 没有合适的牌
		if len(retCards) != 6 {
			retCards = retCards[0:0]
			retInfo.Reset()
			normalDictCards = convertToMap(normalCards)
			_, normalValue, _ = getCountValueLine(normalDictCards)
		}
	}

	// 补两个癞子
	for _, v := range normalValue[2] {
		if info.MinValue < v {
			if laiZiSize < 2 {
				return
			}

			for i := 0; i < 2; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeSiDaiDan
			retInfo.MinValue = v
			break
		}
	}

	if retInfo.CardType != CardTypeNone {
		for i := 0; i < 2; i++ {
			_, normalValue, _ = getCountValueLine(normalDictCards)

			ret := withDan(normalDictCards, normalValue)
			if len(ret) == 0 {
				break
			}
			retCards = append(retCards, ret...)
		}

		// 没有合适的牌
		if len(retCards) != 6 {
			retCards = retCards[0:0]
			retInfo.Reset()
			normalDictCards = convertToMap(normalCards)
			_, normalValue, _ = getCountValueLine(normalDictCards)
		}
	}
	return
}

// 癞子四带对
func findBigSiDaiDuiLaiZi(size int, laiZiSize int, info *CardTypeInfo, normalCards []*Card, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	// 非癞子牌里有大牌
	if retCards, retInfo = findBigSiDaiDui(len(normalDictCards), info, normalCards, normalDictCards, normalValue); retInfo.CardType != CardTypeNone {
		return
	}

	if size < 8 {
		return
	}

	var nums []int

	// 补一个癞子
	for _, v := range normalValue[3] {
		if info.MinValue < v {
			if laiZiSize < 1 {
				return
			}

			for i := 0; i < 3; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeSiDaiDui
			retInfo.MinValue = v
			break
		}
	}

	if retInfo.CardType != CardTypeNone {
		for i := 0; i < 2; i++ {
			_, normalValue, _ = getCountValueLine(normalDictCards)

			ret := withDui(normalDictCards, normalValue)
			if len(ret) == 0 {
				break
			}
			retCards = append(retCards, ret...)
		}

		// 没有合适的牌
		if len(retCards) != 8 {
			retCards = retCards[0:0]
			retInfo.Reset()
			normalDictCards = convertToMap(normalCards)
			_, normalValue, _ = getCountValueLine(normalDictCards)
		}
	}

	// 补两个癞子
	for _, v := range normalValue[2] {
		if info.MinValue < v {
			if laiZiSize < 2 {
				return
			}

			for i := 0; i < 2; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeSiDaiDui
			retInfo.MinValue = v
			break
		}
	}

	if retInfo.CardType != CardTypeNone {
		for i := 0; i < 2; i++ {
			_, normalValue, _ = getCountValueLine(normalDictCards)

			ret := withDui(normalDictCards, normalValue)
			if len(ret) == 0 {
				break
			}
			retCards = append(retCards, ret...)
		}

		// 没有合适的牌
		if len(retCards) != 8 {
			retCards = retCards[0:0]
			retInfo.Reset()
			normalDictCards = convertToMap(normalCards)
			_, normalValue, _ = getCountValueLine(normalDictCards)
		}
	}
	return
}

// 癞子顺子
func findBigShunZiLaiZi(size int, laiZiSize int, info *CardTypeInfo, normalDictCards dictMap, normalCount countList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	// 非癞子牌里有大牌
	if retCards, retInfo = findBigShunZi(len(normalDictCards), info, normalDictCards, normalCount); retInfo.CardType != CardTypeNone {
		return
	}

	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange {
		return
	}

	var nums []int
	var line []int

	// 补一个癞子
	if laiZiSize < 1 {
		return
	}

	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			line = append(line, j)

			switch normalCount[j] {
			case 0:
				needLaiZi++
			case 1, 2, 3:
				nums = append(nums, j)
			}

			if needLaiZi > 1 {
				break
			}
		}

		if needLaiZi == 1 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeShunZi
			retInfo.MinValue = line[0]
			retInfo.MaxValue = line[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
		line = line[0:0]
	}

	// 补两个癞子
	if laiZiSize < 2 {
		return
	}

	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			line = append(line, j)

			switch normalCount[j] {
			case 0:
				needLaiZi++
			case 1, 2, 3:
				nums = append(nums, j)
			}

			if needLaiZi > 2 {
				break
			}
		}

		if needLaiZi == 2 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeShunZi
			retInfo.MinValue = line[0]
			retInfo.MaxValue = line[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
		line = line[0:0]
	}
	return
}

// 癞子连对
func findBigLianDuiLaiZi(size int, laiZiSize int, info *CardTypeInfo, normalDictCards dictMap, normalCount countList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	// 非癞子牌里有大牌
	if retCards, retInfo = findBigLianDui(len(normalDictCards), info, normalDictCards, normalCount); retInfo.CardType != CardTypeNone {
		return
	}

	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange*2 {
		return
	}

	var nums []int
	var line []int

	// 补一个癞子
	if laiZiSize < 1 {
		return
	}

	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			line = append(line, j)

			switch normalCount[j] {
			case 0:
				needLaiZi += 2
			case 1:
				nums = append(nums, j)
				needLaiZi++
			case 2, 3:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
			}

			if needLaiZi > 1 {
				break
			}
		}

		if needLaiZi == 1 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeLianDui
			retInfo.MinValue = line[0]
			retInfo.MaxValue = line[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
		line = line[0:0]
	}

	// 补两个癞子
	if laiZiSize < 2 {
		return
	}

	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			line = append(line, j)

			switch normalCount[j] {
			case 0:
				needLaiZi += 2
			case 1:
				nums = append(nums, j)
				needLaiZi++
			case 2, 3:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
			}

			if needLaiZi > 2 {
				break
			}
		}

		if needLaiZi == 2 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeLianDui
			retInfo.MinValue = line[0]
			retInfo.MaxValue = line[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
		line = line[0:0]
	}
	return
}

// 癞子飞机不带
func findBigFeiJiBuDaiLaiZi(size int, laiZiSize int, info *CardTypeInfo, normalDictCards dictMap, normalCount countList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	// 非癞子牌里有大牌
	if retCards, retInfo = findBigFeiJiBuDai(len(normalDictCards), info, normalDictCards, normalCount); retInfo.CardType != CardTypeNone {
		return
	}

	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange*3 {
		return
	}

	var nums []int
	var line []int

	// 补一个癞子
	if laiZiSize < 1 {
		return
	}

	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			line = append(line, j)

			switch normalCount[j] {
			case 0:
				needLaiZi += 3
			case 1:
				nums = append(nums, j)
				needLaiZi += 2
			case 2:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
				needLaiZi++
			case 3:
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
			}

			if needLaiZi > 1 {
				break
			}
		}

		if needLaiZi == 1 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeFeiJiBuDai
			retInfo.MinValue = line[0]
			retInfo.MaxValue = line[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
		line = line[0:0]
	}

	// 补两个癞子
	if laiZiSize < 2 {
		return
	}

	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			line = append(line, j)

			switch normalCount[j] {
			case 0:
				needLaiZi += 3
			case 1:
				nums = append(nums, j)
				needLaiZi += 2
			case 2:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
				needLaiZi++
			case 3:
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
			}

			if needLaiZi > 2 {
				break
			}
		}

		if needLaiZi == 2 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeFeiJiBuDai
			retInfo.MinValue = line[0]
			retInfo.MaxValue = line[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
		line = line[0:0]
	}
	return
}

// 癞子飞机带一
func findBigFeiJiDaiYiLaiZi(size int, laiZiSize int, info *CardTypeInfo, normalCards []*Card, normalDictCards dictMap, normalCount countList, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	// 非癞子牌里有大牌
	if retCards, retInfo = findBigFeiJiDaiYi(len(normalDictCards), info, normalCards, normalDictCards, normalCount, normalValue); retInfo.CardType != CardTypeNone {
		return
	}

	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange*4 {
		return
	}

	var nums []int
	var line []int

	// 补一个癞子
	if laiZiSize < 1 {
		return
	}

	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			line = append(line, j)

			switch normalCount[j] {
			case 0:
				needLaiZi += 3
			case 1:
				nums = append(nums, j)
				needLaiZi += 2
			case 2:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
				needLaiZi++
			case 3:
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
			}

			if needLaiZi > 1 {
				break
			}
		}

		if needLaiZi == 1 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeFeiJiDaiYi
			retInfo.MinValue = line[0]
			retInfo.MaxValue = line[0] + valueRange - 1
			break
		}
		nums = nums[0:0]
		line = line[0:0]
	}

	if retInfo.CardType != CardTypeNone {
		for i := 0; i < valueRange; i++ {
			_, normalValue, _ = getCountValueLine(normalDictCards)

			ret := withDan(normalDictCards, normalValue)
			if len(ret) == 0 {
				break
			}
			retCards = append(retCards, ret...)
		}

		// 没有合适的牌
		if len(retCards) != valueRange*4 {
			retCards = retCards[0:0]
			retInfo.Reset()
			normalDictCards = convertToMap(normalCards)
			_, normalValue, _ = getCountValueLine(normalDictCards)
		}
	}

	// 补两个癞子
	if laiZiSize < 2 {
		return
	}

	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			line = append(line, j)

			switch normalCount[j] {
			case 0:
				needLaiZi += 3
			case 1:
				nums = append(nums, j)
				needLaiZi += 2
			case 2:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
				needLaiZi++
			case 3:
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
			}

			if needLaiZi > 2 {
				break
			}
		}

		if needLaiZi == 2 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeFeiJiDaiYi
			retInfo.MinValue = line[0]
			retInfo.MaxValue = line[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
		line = line[0:0]
	}

	if retInfo.CardType != CardTypeNone {
		for i := 0; i < valueRange; i++ {
			_, normalValue, _ = getCountValueLine(normalDictCards)

			ret := withDan(normalDictCards, normalValue)
			if len(ret) == 0 {
				break
			}
			retCards = append(retCards, ret...)
		}

		// 没有合适的牌
		if len(retCards) != valueRange*4 {
			retCards = retCards[0:0]
			retInfo.Reset()
			normalDictCards = convertToMap(normalCards)
			_, normalValue, _ = getCountValueLine(normalDictCards)
		}
	}
	return
}

// 癞子飞机带二
func findBigFeiJiDaiErLaiZi(size int, laiZiSize int, info *CardTypeInfo, normalCards []*Card, normalDictCards dictMap, normalCount countList, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	// 非癞子牌里有大牌
	if retCards, retInfo = findBigFeiJiDaiEr(len(normalDictCards), info, normalCards, normalDictCards, normalCount, normalValue); retInfo.CardType != CardTypeNone {
		return
	}

	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange*5 {
		return
	}

	var nums []int
	var line []int

	// 补一个癞子
	if laiZiSize < 1 {
		return
	}

	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			line = append(line, j)

			switch normalCount[j] {
			case 0:
				needLaiZi += 3
			case 1:
				nums = append(nums, j)
				needLaiZi += 2
			case 2:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
				needLaiZi++
			case 3:
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
			}

			if needLaiZi > 1 {
				break
			}
		}

		if needLaiZi == 1 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeFeiJiDaiEr
			retInfo.MinValue = line[0]
			retInfo.MaxValue = line[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
		line = line[0:0]
	}

	if retInfo.CardType != CardTypeNone {
		for i := 0; i < valueRange; i++ {
			_, normalValue, _ = getCountValueLine(normalDictCards)

			ret := withDui(normalDictCards, normalValue)
			if len(ret) == 0 {
				break
			}
			retCards = append(retCards, ret...)
		}

		// 没有合适的牌
		if len(retCards) != valueRange*5 {
			retCards = retCards[0:0]
			retInfo.Reset()
			normalDictCards = convertToMap(normalCards)
			_, normalValue, _ = getCountValueLine(normalDictCards)
		}
	}

	// 补两个癞子
	if laiZiSize < 2 {
		return
	}

	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			line = append(line, j)

			switch normalCount[j] {
			case 0:
				needLaiZi += 3
			case 1:
				nums = append(nums, j)
				needLaiZi += 2
			case 2:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
				needLaiZi++
			case 3:
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
			}

			if needLaiZi > 2 {
				break
			}
		}

		if needLaiZi == 2 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeFeiJiDaiEr
			retInfo.MinValue = line[0]
			retInfo.MaxValue = line[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
		line = line[0:0]
	}

	if retInfo.CardType != CardTypeNone {
		for i := 0; i < valueRange; i++ {
			_, normalValue, _ = getCountValueLine(normalDictCards)

			ret := withDui(normalDictCards, normalValue)
			if len(ret) == 0 {
				break
			}
			retCards = append(retCards, ret...)
		}

		// 没有合适的牌
		if len(retCards) != valueRange*5 {
			retCards = retCards[0:0]
			retInfo.Reset()
			normalDictCards = convertToMap(normalCards)
			_, normalValue, _ = getCountValueLine(normalDictCards)
		}
	}
	return
}

// 四软炸
func findRuanZhaDan4(size int, laiZiSize int, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 4 {
		return
	}

	var nums []int

	if laiZiSize < 1 {
		return
	}

	for _, v := range normalValue[3] {
		for i := 0; i < 3; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0])
		retInfo.CardType = CardTypeRuanZhaDan4
		retInfo.MinValue = v
		return
	}

	if laiZiSize < 2 {
		return
	}

	for _, v := range normalValue[2] {
		for i := 0; i < 2; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:2]...)
		retInfo.CardType = CardTypeRuanZhaDan4
		retInfo.MinValue = v
		return
	}

	if laiZiSize < 3 {
		return
	}

	for _, v := range normalValue[1] {
		if isJoker(v) {
			continue
		}

		nums = append(nums, v)
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:3]...)
		retInfo.CardType = CardTypeRuanZhaDan4
		retInfo.MinValue = v
		return
	}
	return
}

// 更大的四软炸
func findBigRuanZhaDan4(size int, laiZiSize int, info *CardTypeInfo, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 4 {
		return
	}

	var nums []int

	// 一个癞子
	if laiZiSize < 1 {
		return
	}

	for _, v := range normalValue[3] {
		if info.MinValue < v {
			for i := 0; i < 3; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeRuanZhaDan4
			retInfo.MinValue = v
			return
		}
	}

	// 两个癞子
	if laiZiSize < 2 {
		return
	}

	for _, v := range normalValue[2] {
		if info.MinValue < v {
			for i := 0; i < 2; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeRuanZhaDan4
			retInfo.MinValue = v
			return
		}
	}

	// 三个癞子
	if laiZiSize < 3 {
		return
	}

	for _, v := range normalValue[1] {
		if info.MinValue < v {
			nums = append(nums, v)
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:3]...)
			retInfo.CardType = CardTypeRuanZhaDan4
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 五软炸
func findRuanZhaDan5(size int, laiZiSize int, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 5 {
		return
	}

	var nums []int

	if laiZiSize < 1 {
		return
	}

	for _, v := range normalValue[4] {
		for i := 0; i < 4; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0])
		retInfo.CardType = CardTypeRuanZhaDan5
		retInfo.MinValue = v
		return
	}

	if laiZiSize < 2 {
		return
	}

	for _, v := range normalValue[3] {
		for i := 0; i < 3; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:2]...)
		retInfo.CardType = CardTypeRuanZhaDan5
		retInfo.MinValue = v
		return
	}

	if laiZiSize < 3 {
		return
	}

	for _, v := range normalValue[2] {
		for i := 0; i < 2; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:3]...)
		retInfo.CardType = CardTypeRuanZhaDan5
		retInfo.MinValue = v
		return
	}

	if laiZiSize < 4 {
		return
	}

	for _, v := range normalValue[1] {
		if isJoker(v) {
			continue
		}

		nums = append(nums, v)
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:4]...)
		retInfo.CardType = CardTypeRuanZhaDan5
		retInfo.MinValue = v
		return
	}
	return
}

// 更大的五软炸
func findBigRuanZhaDan5(size int, laiZiSize int, info *CardTypeInfo, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 5 {
		return
	}

	var nums []int

	if laiZiSize < 1 {
		return
	}

	for _, v := range normalValue[4] {
		if info.MinValue < v {
			for i := 0; i < 4; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeRuanZhaDan5
			retInfo.MinValue = v
			return
		}
	}

	if laiZiSize < 2 {
		return
	}

	for _, v := range normalValue[3] {
		if info.MinValue < v {
			for i := 0; i < 3; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeRuanZhaDan5
			retInfo.MinValue = v
			return
		}
	}

	if laiZiSize < 3 {
		return
	}

	for _, v := range normalValue[2] {
		if info.MinValue < v {
			for i := 0; i < 2; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:3]...)
			retInfo.CardType = CardTypeRuanZhaDan5
			retInfo.MinValue = v
			return
		}
	}

	if laiZiSize < 4 {
		return
	}

	for _, v := range normalValue[1] {
		if info.MinValue < v {
			nums = append(nums, v)
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:4]...)
			retInfo.CardType = CardTypeRuanZhaDan5
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 六软炸
func findRuanZhaDan6(size int, laiZiSize int, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 6 {
		return
	}

	var nums []int

	if laiZiSize < 2 {
		return
	}

	for _, v := range normalValue[4] {
		for i := 0; i < 4; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:2]...)
		retInfo.CardType = CardTypeRuanZhaDan6
		retInfo.MinValue = v
		return
	}

	if laiZiSize < 3 {
		return
	}

	for _, v := range normalValue[3] {
		for i := 0; i < 3; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:3]...)
		retInfo.CardType = CardTypeRuanZhaDan6
		retInfo.MinValue = v
		return
	}

	if laiZiSize < 4 {
		return
	}

	for _, v := range normalValue[2] {
		for i := 0; i < 2; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:4]...)
		retInfo.CardType = CardTypeRuanZhaDan6
		retInfo.MinValue = v
		return
	}

	if laiZiSize < 5 {
		return
	}

	for _, v := range normalValue[1] {
		if isJoker(v) {
			continue
		}

		nums = append(nums, v)
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:5]...)
		retInfo.CardType = CardTypeRuanZhaDan6
		retInfo.MinValue = v
		return
	}
	return
}

// 更大的六软炸
func findBigRuanZhaDan6(size int, laiZiSize int, info *CardTypeInfo, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 6 {
		return
	}

	var nums []int

	if laiZiSize < 2 {
		return
	}

	for _, v := range normalValue[4] {
		if info.MinValue < v {
			for i := 0; i < 4; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeRuanZhaDan6
			retInfo.MinValue = v
			return
		}
	}

	if laiZiSize < 3 {
		return
	}

	for _, v := range normalValue[3] {
		if info.MinValue < v {
			for i := 0; i < 3; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:3]...)
			retInfo.CardType = CardTypeRuanZhaDan6
			retInfo.MinValue = v
			return
		}
	}

	if laiZiSize < 4 {
		return
	}

	for _, v := range normalValue[2] {
		if info.MinValue < v {
			for i := 0; i < 2; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:4]...)
			retInfo.CardType = CardTypeRuanZhaDan6
			retInfo.MinValue = v
			return
		}
	}

	if laiZiSize < 5 {
		return
	}

	for _, v := range normalValue[1] {
		if info.MinValue < v {
			nums = append(nums, v)
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:5]...)
			retInfo.CardType = CardTypeRuanZhaDan6
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 七软炸
func findRuanZhaDan7(size int, laiZiSize int, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 7 {
		return
	}

	var nums []int

	if laiZiSize < 3 {
		return
	}

	for _, v := range normalValue[4] {
		for i := 0; i < 4; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:3]...)
		retInfo.CardType = CardTypeRuanZhaDan7
		retInfo.MinValue = v
		return
	}

	if laiZiSize < 4 {
		return
	}

	for _, v := range normalValue[3] {
		for i := 0; i < 3; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:4]...)
		retInfo.CardType = CardTypeRuanZhaDan7
		retInfo.MinValue = v
		return
	}

	if laiZiSize < 5 {
		return
	}

	for _, v := range normalValue[2] {
		for i := 0; i < 2; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:5]...)
		retInfo.CardType = CardTypeRuanZhaDan7
		retInfo.MinValue = v
		return
	}

	if laiZiSize < 6 {
		return
	}

	for _, v := range normalValue[1] {
		if isJoker(v) {
			continue
		}

		nums = append(nums, v)
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:6]...)
		retInfo.CardType = CardTypeRuanZhaDan7
		retInfo.MinValue = v
		return
	}
	return
}

// 更大的七软炸
func findBigRuanZhaDan7(size int, laiZiSize int, info *CardTypeInfo, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 7 {
		return
	}

	var nums []int

	if laiZiSize < 3 {
		return
	}

	for _, v := range normalValue[4] {
		if info.MinValue < v {
			for i := 0; i < 4; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:3]...)
			retInfo.CardType = CardTypeRuanZhaDan7
			retInfo.MinValue = v
			return
		}
	}

	if laiZiSize < 4 {
		return
	}

	for _, v := range normalValue[3] {
		if info.MinValue < v {
			for i := 0; i < 3; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:4]...)
			retInfo.CardType = CardTypeRuanZhaDan7
			retInfo.MinValue = v
			return
		}
	}

	if laiZiSize < 5 {
		return
	}

	for _, v := range normalValue[2] {
		if info.MinValue < v {
			for i := 0; i < 2; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:5]...)
			retInfo.CardType = CardTypeRuanZhaDan7
			retInfo.MinValue = v
			return
		}
	}

	if laiZiSize < 6 {
		return
	}

	for _, v := range normalValue[1] {
		if info.MinValue < v {
			nums = append(nums, v)
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:6]...)
			retInfo.CardType = CardTypeRuanZhaDan7
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 八软炸
func findRuanZhaDan8(size int, laiZiSize int, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 8 {
		return
	}

	var nums []int

	if laiZiSize < 4 {
		return
	}

	for _, v := range normalValue[4] {
		for i := 0; i < 4; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:4]...)
		retInfo.CardType = CardTypeRuanZhaDan8
		retInfo.MinValue = v
		return
	}

	if laiZiSize < 5 {
		return
	}

	for _, v := range normalValue[3] {
		for i := 0; i < 3; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:5]...)
		retInfo.CardType = CardTypeRuanZhaDan8
		retInfo.MinValue = v
		return
	}

	if laiZiSize < 6 {
		return
	}

	for _, v := range normalValue[2] {
		for i := 0; i < 2; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:6]...)
		retInfo.CardType = CardTypeRuanZhaDan8
		retInfo.MinValue = v
		return
	}

	if laiZiSize < 7 {
		return
	}

	for _, v := range normalValue[1] {
		if isJoker(v) {
			continue
		}

		nums = append(nums, v)
		retCards = findCardsByNums(normalDictCards, nums)
		retCards = append(retCards, laiZiCards[0:7]...)
		retInfo.CardType = CardTypeRuanZhaDan8
		retInfo.MinValue = v
		return
	}
	return
}

// 更大的八软炸
func findBigRuanZhaDan8(size int, laiZiSize int, info *CardTypeInfo, normalDictCards dictMap, normalValue valueList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 8 {
		return
	}

	var nums []int

	if laiZiSize < 4 {
		return
	}

	for _, v := range normalValue[4] {
		if info.MinValue < v {
			for i := 0; i < 4; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:4]...)
			retInfo.CardType = CardTypeRuanZhaDan8
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 癞子炸
func findLaiZiZhaDan(cardType CardType, laiZiSize int, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	switch cardType {
	case CardTypeLaiZiZhaDan5:
		if laiZiSize < 5 {
			return
		}
		retCards = laiZiCards[0:5]
		retInfo.CardType = CardTypeLaiZiZhaDan5

	case CardTypeLaiZiZhaDan6:
		if laiZiSize < 6 {
			return
		}
		retCards = laiZiCards[0:6]
		retInfo.CardType = CardTypeLaiZiZhaDan6
	}
	return
}

// 四纯癞子炸,四癞子炸
func findLaiZiZhaDan4(laiZiSize int, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	if laiZiSize < 4 {
		return
	}

	var nums []int

	laiZiDictCards := convertToMap(laiZiCards)
	_, laiZiValue, _ := getCountValueLine(laiZiDictCards)

	// 优先四纯癞子炸
	for _, v := range laiZiValue[4] {
		for i := 0; i < 4; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(laiZiDictCards, nums)
		retInfo.CardType = CardTypeChunLaiZiZhaDan
		retInfo.MinValue = v
		return
	}

	// 四癞子炸
	retCards = laiZiCards[0:4]
	retInfo.CardType = CardTypeLaiZiZhaDan4
	return
}

// 更大的四纯癞子炸
func findBigChunLaiZiZhaDan(laiZiSize int, info *CardTypeInfo, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	if laiZiSize < 4 {
		return
	}

	var nums []int

	laiZiDictCards := convertToMap(laiZiCards)
	_, laiZiValue, _ := getCountValueLine(laiZiDictCards)

	for _, v := range laiZiValue[4] {
		if info.MinValue < v {
			for i := 0; i < 4; i++ {
				nums = append(nums, v)
			}
			retCards = findCardsByNums(laiZiDictCards, nums)
			retInfo.CardType = CardTypeChunLaiZiZhaDan
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 软连炸
func findRuanLianZha(size int, laiZiSize int, normalDictCards dictMap, normalCount countList, laiZiCards []*Card, valueRange int) (retCards []*Card, retInfo CardTypeInfo) {
	if size < valueRange*4 {
		return
	}

	var nums []int

	// 补一个癞子
	if laiZiSize < 1 {
		return
	}

	for i := NumTypeThree; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			switch normalCount[j] {
			case 0:
				needLaiZi += 4
			case 1:
				nums = append(nums, j)
				needLaiZi += 3
			case 2:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
				needLaiZi += 2
			case 3:
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
				needLaiZi++
			case 4:
				for k := 0; k < 4; k++ {
					nums = append(nums, j)
				}
			}

			if needLaiZi > 1 {
				break
			}
		}

		if needLaiZi == 1 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeRuanLianZha
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
	}

	// 补两个癞子
	if laiZiSize < 2 {
		return
	}

	for i := NumTypeThree; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			switch normalCount[j] {
			case 0:
				needLaiZi += 4
			case 1:
				nums = append(nums, j)
				needLaiZi += 3
			case 2:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
				needLaiZi += 2
			case 3:
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
				needLaiZi++
			case 4:
				for k := 0; k < 4; k++ {
					nums = append(nums, j)
				}
			}

			if needLaiZi > 2 {
				break
			}
		}

		if needLaiZi == 2 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeRuanLianZha
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
	}
	return
}

// 更大的软连炸
func findBigRuanLianZha(size int, laiZiSize int, info *CardTypeInfo, normalDictCards dictMap, normalCount countList, laiZiCards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange*4 {
		return
	}

	var nums []int

	// 补一个癞子
	if laiZiSize < 1 {
		return
	}

	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			switch normalCount[j] {
			case 0:
				needLaiZi += 4
			case 1:
				nums = append(nums, j)
				needLaiZi += 3
			case 2:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
				needLaiZi += 2
			case 3:
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
				needLaiZi++
			case 4:
				for k := 0; k < 4; k++ {
					nums = append(nums, j)
				}
			}

			if needLaiZi > 1 {
				break
			}
		}

		if needLaiZi == 1 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeRuanLianZha
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
	}

	// 补两个癞子
	if laiZiSize < 2 {
		return
	}

	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			switch normalCount[j] {
			case 0:
				needLaiZi += 4
			case 1:
				nums = append(nums, j)
				needLaiZi += 3
			case 2:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
				needLaiZi += 2
			case 3:
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
				needLaiZi++
			case 4:
				for k := 0; k < 4; k++ {
					nums = append(nums, j)
				}
			}

			if needLaiZi > 2 {
				break
			}
		}

		if needLaiZi == 2 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeRuanLianZha
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
	}

	// 更多的软连炸
	valueRange++
	if size < valueRange*4 {
		return
	}

	// 补一个癞子
	if laiZiSize < 1 {
		return
	}

	for i := NumTypeThree; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			switch normalCount[j] {
			case 0:
				needLaiZi += 4
			case 1:
				nums = append(nums, j)
				needLaiZi += 3
			case 2:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
				needLaiZi += 2
			case 3:
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
				needLaiZi++
			case 4:
				for k := 0; k < 4; k++ {
					nums = append(nums, j)
				}
			}

			if needLaiZi > 1 {
				break
			}
		}

		if needLaiZi == 1 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0])
			retInfo.CardType = CardTypeRuanLianZha
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
	}

	// 补两个癞子
	if laiZiSize < 2 {
		return
	}

	for i := NumTypeThree; i <= NumTypeAce-valueRange+1; i++ {
		var needLaiZi int
		for j := i; j < i+valueRange; j++ {
			switch normalCount[j] {
			case 0:
				needLaiZi += 4
			case 1:
				nums = append(nums, j)
				needLaiZi += 3
			case 2:
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
				needLaiZi += 2
			case 3:
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
				needLaiZi++
			case 4:
				for k := 0; k < 4; k++ {
					nums = append(nums, j)
				}
			}

			if needLaiZi > 2 {
				break
			}
		}

		if needLaiZi == 2 {
			retCards = findCardsByNums(normalDictCards, nums)
			retCards = append(retCards, laiZiCards[0:2]...)
			retInfo.CardType = CardTypeRuanLianZha
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[0] + valueRange - 1
			return
		}
		nums = nums[0:0]
	}

	return
}
