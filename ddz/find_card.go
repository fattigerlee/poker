package ddz

// 找出比牌型更大的牌
func FindCards(info *CardTypeInfo, cards []*Card, laiZiNums ...NumType) (retCards []*Card, retInfo CardTypeInfo) {
	if len(laiZiNums) == 0 {
		// 无癞子找牌
		return findCards(info, cards)
	}

	newCards := getCardsWithoutLaiZi(cards, laiZiNums...) // 无癞子牌数组
	laiZiCount := len(cards) - len(newCards)              // 癞子数量

	if laiZiCount == 0 {
		// 无癞子找牌
		return findCards(info, cards)
	} else {
		// 癞子找牌
		return findCardsLaiZi(info, cards, newCards, laiZiCount)
	}
}

// 无癞子找牌
func findCards(info *CardTypeInfo, cards []*Card) (retCards []*Card, retInfo CardTypeInfo) {
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
	//case CardTypeFeiJiDaiYi:
	//	if ret = findBigFeiJiDaiYi(cards, count, info); len(ret) != 0 {
	//		return
	//	}
	//case CardTypeFeiJiDaiEr:
	//	if ret = findBigFeiJiDaiEr(cards, count, info); len(ret) != 0 {
	//		return
	//	}
	case CardTypeZhaDan:
		if retCards, retInfo = findBigZhaDan(size, info, dictCards, count, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeHuoJian:
		if retCards, retInfo = findBigHuoJian(size, dictCards, count, value); retInfo.CardType != CardTypeNone {
			return
		}
	case CardTypeLianZha:
		if retCards, retInfo = findBigLianZha(size, info, dictCards, count); retInfo.CardType != CardTypeNone {
			return
		}
	}

	// 炸弹
	if retCards, retInfo = findZhaDan(size, dictCards, value); retInfo.CardType != CardTypeNone {
		return
	}
	return
}

// 带单张
func withDan(cards map[*Card]bool, value valueList) (ret []*Card) {
	var nums []int

	// 找单张
	if huoJian(value) {
		// 有火箭
		for _, v := range value[1] {
			if isJoker(v) {
				continue
			}
			nums = append(nums, v)
			ret = append(ret, findCardsByNums(cards, nums)...)
			return
		}
	} else {
		for _, v := range value[1] {
			nums = append(nums, v)
			ret = append(ret, findCardsByNums(cards, nums)...)
			return
		}
	}

	// 拆对子
	for _, v := range value[2] {
		nums = append(nums, v)
		ret = append(ret, findCardsByNums(cards, nums)...)
		return
	}

	// 拆三张
	for _, v := range value[3] {
		nums = append(nums, v)
		ret = append(ret, findCardsByNums(cards, nums)...)
		return
	}
	return
}

// 带对子
func withDui(cards map[*Card]bool, value valueList) (ret []*Card) {
	var nums []int

	// 找对子
	for _, v := range value[2] {
		for i := 0; i < 2; i++ {
			nums = append(nums, v)
		}
		ret = append(ret, findCardsByNums(cards, nums)...)
		return
	}

	// 拆三张
	for _, v := range value[3] {
		for i := 0; i < 2; i++ {
			nums = append(nums, v)
		}
		ret = append(ret, findCardsByNums(cards, nums)...)
		return
	}
	return
}

// 单张
func findBigDan(size int, info *CardTypeInfo, dictCards map[*Card]bool, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
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
				retCards = append(retCards, findCardsByNums(dictCards, nums)...)
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
				retCards = append(retCards, findCardsByNums(dictCards, nums)...)
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
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
			retInfo.CardType = CardTypeDan
			retInfo.MinValue = v
			return
		}
	}

	// 拆三张
	for _, v := range value[3] {
		if info.MinValue < v {
			nums = append(nums)
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
			retInfo.CardType = CardTypeDan
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 对子
func findBigDui(size int, info *CardTypeInfo, dictCards map[*Card]bool, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
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
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
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
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
			retInfo.CardType = CardTypeDui
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 三不带
func findBigSanBuDai(size int, info *CardTypeInfo, dictCards map[*Card]bool, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
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
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
			retInfo.CardType = CardTypeSanBuDai
			retInfo.MinValue = v
			return
		}
	}
	return
}

// 三带一
func findBigSanDaiYi(size int, info *CardTypeInfo, cards []*Card, dictCards map[*Card]bool, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
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
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
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
func findBigSanDaiEr(size int, info *CardTypeInfo, cards []*Card, dictCards map[*Card]bool, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
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
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
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
func findBigSiDaiDan(size int, info *CardTypeInfo, cards []*Card, dictCards map[*Card]bool, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
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
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
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
func findBigSiDaiDui(size int, info *CardTypeInfo, cards []*Card, dictCards map[*Card]bool, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
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
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
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
func findBigShunZi(size int, info *CardTypeInfo, dictCards map[*Card]bool, count countList) (retCards []*Card, retInfo CardTypeInfo) {
	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange {
		return
	}

	var nums []int

	// 更大的顺子
	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		for j := i; j < i+valueRange; j++ {
			if count[j] > 0 && count[j] < 4 {
				nums = append(nums, j)
			}
		}

		if len(nums) == valueRange {
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
			retInfo.CardType = CardTypeShunZi
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[len(nums)-1]
			break
		}
		nums = nums[0:0]
	}
	return
}

// 连对
func findBigLianDui(size int, info *CardTypeInfo, dictCards map[*Card]bool, count countList) (retCards []*Card, retInfo CardTypeInfo) {
	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange*2 {
		return
	}

	var nums []int

	// 更大的连对
	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		for j := i; j < i+valueRange; j++ {
			if count[j] > 1 && count[j] < 4 {
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}
			}
		}

		if len(nums) == valueRange*2 {
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
			retInfo.CardType = CardTypeLianDui
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[len(nums)-1]
			break
		}
		nums = nums[0:0]
	}
	return
}

// 飞机不带
func findBigFeiJiBuDai(size int, info *CardTypeInfo, dictCards map[*Card]bool, count countList) (retCards []*Card, retInfo CardTypeInfo) {
	valueRange := info.MaxValue - info.MinValue + 1
	if size < valueRange*3 {
		return
	}

	var nums []int

	// 更大的飞机不带
	for i := info.MinValue + 1; i <= NumTypeAce-valueRange+1; i++ {
		for j := i; j < i+valueRange; j++ {
			if count[j] > 1 && count[j] < 4 {
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
			}
		}

		if len(nums) == valueRange*3 {
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
			retInfo.CardType = CardTypeFeiJiBuDai
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[len(nums)-1]
			break
		}
		nums = nums[0:0]
	}
	return
}

// 飞机带一

// 飞机带二

// 炸弹
func findZhaDan(size int, dictCards map[*Card]bool, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	var nums []int

	// 炸弹
	for _, v := range value[4] {
		for i := 0; i < 4; i++ {
			nums = append(nums, v)
		}
		retCards = findCardsByNums(dictCards, nums)
		retInfo.CardType = CardTypeZhaDan
		retInfo.MinValue = v
		return
	}

	// 火箭
	retCards, retInfo = findHuoJian(size, dictCards, value)
	return
}

func findBigZhaDan(size int, info *CardTypeInfo, dictCards map[*Card]bool, count countList, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
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
		return
	}

	// 火箭
	if retCards, retInfo = findHuoJian(size, dictCards, value); retInfo.CardType != CardTypeNone {
		return
	}

	// 连炸
	if retCards, retInfo = findLianZha(size, dictCards, count, value); retInfo.CardType != CardTypeNone {
		return
	}
	return
}

// 火箭
func findHuoJian(size int, dictCards map[*Card]bool, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
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

func findBigHuoJian(size int, dictCards map[*Card]bool, count countList, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	return findLianZha(size, dictCards, count, value)
}

// 连炸
func findLianZha(size int, dictCards map[*Card]bool, count countList, value valueList) (retCards []*Card, retInfo CardTypeInfo) {
	if size < 8 {
		return
	}

	var nums []int

	valueRange := 2
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
			retInfo.MaxValue = nums[len(nums)-1]
			break
		}
		nums = nums[0:0]
	}
	return
}

func findBigLianZha(size int, info *CardTypeInfo, dictCards map[*Card]bool, count countList) (retCards []*Card, retInfo CardTypeInfo) {
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
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
			retInfo.CardType = CardTypeLianZha
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[len(nums)-1]
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
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
			retInfo.CardType = CardTypeLianZha
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[len(nums)-1]
			break
		}
		nums = nums[0:0]
	}
	return
}

// 癞子找牌
func findCardsLaiZi(info *CardTypeInfo, cards []*Card, newCards []*Card, laiZiCount int) (retCards []*Card, retInfo CardTypeInfo) {
	return
}
