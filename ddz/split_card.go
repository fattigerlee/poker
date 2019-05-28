package ddz

import (
	"sort"
)

// 拆牌(经典模式)
func SplitCardsJingDian(cards []*Card) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
	dictCards := convertToMap(cards)

	size := len(dictCards)
	count, value, _ := getCountValueLine(dictCards)

	// 火箭
	if retCards, retInfo := findHuoJian(size, dictCards, value); retInfo.CardType != CardTypeNone {
		retCardsList = append(retCardsList, retCards)
		retInfoList = append(retInfoList, retInfo)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 炸弹
	if cardsList, infoList := splitZhaDan(size, dictCards, value); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆2
	if retCards, retInfo := splitTwo(dictCards); retInfo.CardType != CardTypeNone {
		retCardsList = append(retCardsList, retCards)
		retInfoList = append(retInfoList, retInfo)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆飞机
	if cardsList, infoList := splitFeiJiBuDai(size, dictCards, count, value); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆顺子
	if cardsList, infoList := splitShunZi(size, dictCards, count); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆连对
	if cardsList, infoList := splitLianDui(size, dictCards, count, value); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆三不带
	if cardsList, infoList := splitSanBuDai(size, dictCards, value); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆对
	if cardsList, infoList := splitDui(size, dictCards, value); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆单
	if cardsList, infoList := splitDan(size, dictCards, value); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}
	return
}

// 拆牌(不洗牌)
func SplitCardsBuXiPai(cards []*Card) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
	dictCards := convertToMap(cards)

	size := len(dictCards)
	count, value, _ := getCountValueLine(dictCards)

	// 拆连炸
	if cardsList, infoList := splitLianZha(size, dictCards, count, value); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆火箭
	if retCards, retInfo := findHuoJian(size, dictCards, value); retInfo.CardType != CardTypeNone {
		retCardsList = append(retCardsList, retCards)
		retInfoList = append(retInfoList, retInfo)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆炸弹
	if cardsList, infoList := splitZhaDan(size, dictCards, value); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆2
	if retCards, retInfo := splitTwo(dictCards); retInfo.CardType != CardTypeNone {
		retCardsList = append(retCardsList, retCards)
		retInfoList = append(retInfoList, retInfo)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆飞机
	if cardsList, infoList := splitFeiJiBuDai(size, dictCards, count, value); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆顺子
	if cardsList, infoList := splitShunZi(size, dictCards, count); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆连对
	if cardsList, infoList := splitLianDui(size, dictCards, count, value); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆三不带
	if cardsList, infoList := splitSanBuDai(size, dictCards, value); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆对
	if cardsList, infoList := splitDui(size, dictCards, value); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}

	// 拆单
	if cardsList, infoList := splitDan(size, dictCards, value); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)

		size = len(dictCards)
		count, value, _ = getCountValueLine(dictCards)
	}
	return
}

// 拆牌(不洗牌+双王癞子模式)
func SplitCardsBuXiPaiLaiZi(cards []*Card, laiZiNums []NumType) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
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

	laiZiSize := len(laiZiCards) // 癞子牌数量

	if laiZiSize == 0 {
		return SplitCardsBuXiPai(cards)
	}

	// 癞子牌
	laiZiDictCards := convertToMap(laiZiCards)
	_, laiZiValue, _ := getCountValueLine(laiZiDictCards)

	// 非癞子牌拆牌
	retCardsList, retInfoList = SplitCardsBuXiPai(normalCards)

	// 拆癞子火箭
	if retCards, retInfo := findHuoJian(laiZiSize, laiZiDictCards, laiZiValue); retInfo.CardType != CardTypeNone {
		retCardsList = append(retCardsList, retCards)
		retInfoList = append(retInfoList, retInfo)

		laiZiSize = len(laiZiDictCards)
		if laiZiSize == 0 {
			return
		}
	}

	// 补四软炸(三不带补)
	for i := 0; i < len(retInfoList); i++ {
		if retInfoList[i].CardType == CardTypeSanBuDai {
			retInfoList[i].CardType = CardTypeRuanZhaDan4
			retCardsList[i] = append(retCardsList[i], findCardsLaiZi(laiZiDictCards, 1)...)

			laiZiSize = len(laiZiDictCards)
			if laiZiSize == 0 {
				return
			}
		}
	}

	// 补2
	for i := 0; i < len(retInfoList); i++ {
		if retInfoList[i].MinValue == NumTypeTwo && retInfoList[i].CardType != CardTypeZhaDan {
			retCardsList[i] = append(retCardsList[i], findCardsLaiZi(laiZiDictCards, 1)...)

			if len(retCardsList[i]) == 2 {
				retInfoList[i].CardType = CardTypeDui
			}

			if len(retCardsList[i]) == 3 {
				retInfoList[i].CardType = CardTypeSanBuDai
			}

			if len(retCardsList[i]) == 4 {
				retInfoList[i].CardType = CardTypeRuanZhaDan4
			}

			laiZiSize = len(laiZiDictCards)
			if laiZiSize == 0 {
				return
			}
		}
	}

	// 补三不带
	for i := 0; i < len(retInfoList); i++ {
		if retInfoList[i].CardType == CardTypeDui && retInfoList[i].MinValue > 10 {
			retInfoList[i].CardType = CardTypeSanBuDai
			retCardsList[i] = append(retCardsList[i], findCardsLaiZi(laiZiDictCards, 1)...)

			laiZiSize = len(laiZiDictCards)
			if laiZiSize == 0 {
				return
			}
		}
	}

	// 当做单牌用
	retCards := findCardsLaiZi(laiZiDictCards, 1)
	retCardsList = append(retCardsList, retCards)

	var retInfo CardTypeInfo
	retInfo.CardType = CardTypeDan
	retInfo.MinValue = int(laiZiCards[0].Num)
	retInfoList = append(retInfoList, retInfo)
	return
}

// 拆连炸
func splitLianZha(size int, dictCards dictMap, count countList, value valueList) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
	if size < 8 {
		return
	}

	for {
		var nums []int
		var exist bool

		for _, v := range value[4] {
			nums = nums[0:0]

			for i := v; i <= NumTypeAce; i++ {
				if count[i] != 4 {
					break
				}

				for j := 0; j < 4; j++ {
					nums = append(nums, i)
				}
			}

			if len(nums) < 8 {
				continue
			}

			exist = true
			var retCards []*Card
			var retInfo CardTypeInfo

			retCards = findCardsByNums(dictCards, nums)
			retCardsList = append(retCardsList, retCards)

			retInfo.CardType = CardTypeLianZha
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[len(nums)-1]
			retInfoList = append(retInfoList, retInfo)

			count, value, _ = getCountValueLine(dictCards)
			break
		}

		if !exist {
			break
		}
	}
	return
}

// 拆炸弹
func splitZhaDan(size int, dictCards dictMap, value valueList) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
	if size < 4 {
		return
	}

	var nums []int

	for _, v := range value[4] {
		nums = nums[0:0]
		for i := 0; i < 4; i++ {
			nums = append(nums, v)
		}

		var retCards []*Card
		var retInfo CardTypeInfo

		retCards = findCardsByNums(dictCards, nums)
		retCardsList = append(retCardsList, retCards)

		retInfo.CardType = CardTypeZhaDan
		retInfo.MinValue = v
		retInfoList = append(retInfoList, retInfo)
	}
	return
}

// 拆2
func splitTwo(dictCards dictMap) (retCards []*Card, retInfo CardTypeInfo) {
	var nums []int
	for c := range dictCards {
		if c.Num == NumTypeTwo {
			nums = append(nums, NumTypeTwo)
		}
	}

	switch len(nums) {
	case 0:
		return

	case 1:
		retInfo.CardType = CardTypeDan

	case 2:
		retInfo.CardType = CardTypeDui

	case 3:
		retInfo.CardType = CardTypeSanBuDai

	case 4:
		retInfo.CardType = CardTypeZhaDan
	}
	retInfo.MinValue = NumTypeTwo

	retCards = findCardsByNums(dictCards, nums)
	return
}

// 拆飞机不带
func splitFeiJiBuDai(size int, dictCards dictMap, count countList, value valueList) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
	if size < 6 {
		return
	}

	for {
		var nums []int
		var exist bool

		for _, v := range value[3] {
			nums = nums[0:0]

			for i := v; i <= NumTypeAce; i++ {
				if count[i] != 3 {
					break
				}

				for j := 0; j < 3; j++ {
					nums = append(nums, i)
				}
			}

			if len(nums) < 6 {
				continue
			}

			exist = true
			var retCards []*Card
			var retInfo CardTypeInfo

			retCards = findCardsByNums(dictCards, nums)
			retCardsList = append(retCardsList, retCards)

			retInfo.CardType = CardTypeFeiJiBuDai
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[len(nums)-1]
			retInfoList = append(retInfoList, retInfo)

			count, value, _ = getCountValueLine(dictCards)
			break
		}

		if !exist {
			break
		}
	}
	return
}

// 拆顺子
func splitShunZi(size int, dictCards dictMap, count countList) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
	if size < 5 {
		return
	}

	var nums []int

	for {
		var exist bool

		// 取最小的5连
		for i := NumTypeThree; i <= NumTypeAce-4; i++ {
			nums = nums[0:0]

			for j := i; j <= i+4; j++ {
				if count[j] == 0 {
					break
				}
				nums = append(nums, j)
			}

			if len(nums) == 5 {
				exist = true

				var retCards []*Card
				var retInfo CardTypeInfo

				retCards = findCardsByNums(dictCards, nums)
				retCardsList = append(retCardsList, retCards)

				retInfo.CardType = CardTypeShunZi
				retInfo.MinValue = nums[0]
				retInfo.MaxValue = nums[len(nums)-1]
				retInfoList = append(retInfoList, retInfo)

				count, _, _ = getCountValueLine(dictCards)
				break
			}
		}

		if !exist {
			break
		}
	}

	if len(retInfoList) == 0 {
		return
	}

	// 剩余牌合并成新的顺子
	for i := 0; i < len(retInfoList); i++ {
		nums = nums[0:0]

		for j := retInfoList[i].MaxValue + 1; j <= NumTypeAce; j++ {
			if count[j] == 0 {
				break
			}
			nums = append(nums, j)
		}

		if len(nums) == 0 {
			continue
		}

		retCardsList[i] = append(retCardsList[i], findCardsByNums(dictCards, nums)...)
		retInfoList[i].MaxValue = nums[len(nums)-1]
		count, _, _ = getCountValueLine(dictCards)
	}

	// 顺子无缝合并
	for i := 0; i < len(retInfoList); {
		var exist bool
		for j := i + 1; j < len(retInfoList); j++ {
			if retInfoList[i].MaxValue+1 == retInfoList[j].MinValue {
				// 更长的顺子
				retInfoList[j].MinValue = retInfoList[i].MinValue
				cards := append(retCardsList[i], retCardsList[j]...)
				retCardsList[j] = cards

				exist = true
				break
			}
		}

		if !exist {
			i++
		} else {
			retInfoList = append(retInfoList[0:i], retInfoList[i+1:]...)
			retCardsList = append(retCardsList[0:i], retCardsList[i+1:]...)
		}
	}

	// 顺子完全重合,合并成连对
	for i := 0; i < len(retInfoList); {
		var exist bool
		for j := i + 1; j < len(retInfoList); j++ {
			if retInfoList[i] == retInfoList[j] {
				// 连对
				retInfoList[j].CardType = CardTypeLianDui
				cards := append(retCardsList[i], retCardsList[j]...)
				sort.Slice(cards, func(i, j int) bool {
					return cards[i].Num < cards[j].Num
				})
				retCardsList[j] = cards

				exist = true
				break
			}
		}

		if !exist {
			i++
		} else {
			retInfoList = append(retInfoList[0:i], retInfoList[i+1:]...)
			retCardsList = append(retCardsList[0:i], retCardsList[i+1:]...)
		}
	}

	// 超过8连的顺子,看能否和剩牌组成连对
	for i := 0; i < len(retInfoList); i++ {
		if retInfoList[i].CardType != CardTypeShunZi || len(retCardsList[i]) < 8 {
			continue
		}

		// 前面几张牌能否组成连对
		cards := retCardsList[i][0 : len(retCardsList[i])-5]
		for j := 0; j < len(cards)-2; j++ {
			nums = nums[0:0]

			for k := cards[j].Num; k <= cards[len(cards)-1].Num; k++ {
				if count[k] == 0 {
					break
				}
				nums = append(nums, int(k))
			}

			// 可以组成连对
			if len(nums) >= 3 {
				var retCards []*Card
				var retInfo CardTypeInfo

				// 顺子中找出连对使用的牌
				for _, num := range nums {
					for _, c := range retCardsList[i] {
						if int(c.Num) == num {
							retCards = append(retCards, c)
							break
						}
					}
				}

				// 剩牌中找出连对使用的牌
				retCards = append(retCards, findCardsByNums(dictCards, nums)...)
				count, _, _ = getCountValueLine(dictCards)

				sort.Slice(retCards, func(i, j int) bool {
					return retCards[i].Num < retCards[j].Num
				})

				retInfo.CardType = CardTypeLianDui
				retInfo.MinValue = nums[0]
				retInfo.MaxValue = nums[len(nums)-1]

				retCardsList = append(retCardsList, retCards)
				retInfoList = append(retInfoList, retInfo)

				// 修改顺子牌型,删除连对中使用的牌
				for _, num := range nums {
					for k, c := range retCardsList[i] {
						if int(c.Num) == num {
							retCardsList[i] = append(retCardsList[i][0:k], retCardsList[i][k+1:]...)
							break
						}
					}
				}

				// 顺子剩下的牌无法组成顺子,删除断牌,放入剩牌中
				for k := 0; j < len(retCardsList[i])-1; {
					var exist bool
					if retCardsList[i][k].Num+1 != retCardsList[i][k+1].Num {
						dictCards[retCardsList[i][j]] = true
						count, _, _ = getCountValueLine(dictCards)

						exist = true
					}

					if !exist {
						break
					}
					retCardsList[i] = append(retCardsList[i][0:k], retCardsList[i][k+1:]...)
				}

				// 确定顺子最终牌型
				retInfoList[i].MinValue = nums[len(nums)-1] + 1
			}
		}

		// 后面几张牌能否组成连对
		if len(retCardsList[i]) < 8 {
			continue
		}

		cards = retCardsList[i][5:len(retCardsList[i])]
		for j := 0; j < len(cards)-2; j++ {
			nums = nums[0:0]

			for k := cards[j].Num; k <= cards[len(cards)-1].Num; k++ {
				if count[k] == 0 {
					break
				}
				nums = append(nums, int(k))
			}

			// 可以组成连对
			if len(nums) >= 3 {
				var retCards []*Card
				var retInfo CardTypeInfo

				// 顺子中找出连对使用的牌
				for _, num := range nums {
					for _, c := range retCardsList[i] {
						if int(c.Num) == num {
							retCards = append(retCards, c)
							break
						}
					}
				}

				// 剩牌中找出连对使用的牌
				retCards = append(retCards, findCardsByNums(dictCards, nums)...)
				count, _, _ = getCountValueLine(dictCards)

				sort.Slice(retCards, func(i, j int) bool {
					return retCards[i].Num < retCards[j].Num
				})

				retInfo.CardType = CardTypeLianDui
				retInfo.MinValue = nums[0]
				retInfo.MaxValue = nums[len(nums)-1]

				retCardsList = append(retCardsList, retCards)
				retInfoList = append(retInfoList, retInfo)

				// 修改顺子牌型,删除连对中使用的牌
				for _, num := range nums {
					for k, c := range retCardsList[i] {
						if int(c.Num) == num {
							retCardsList[i] = append(retCardsList[i][0:k], retCardsList[i][k+1:]...)
							break
						}
					}
				}

				// 顺子剩下的牌无法组成顺子,删除断牌,放入剩牌中
				for k := len(retCardsList[i]) - 1; k > 0; {
					var exist bool
					if retCardsList[i][k].Num-1 != retCardsList[i][k-1].Num {
						dictCards[retCardsList[i][k]] = true
						count, _, _ = getCountValueLine(dictCards)

						exist = true
					}

					if !exist {
						break
					}

					retCardsList[i] = retCardsList[i][:k]
					k--
				}

				// 确定顺子最终牌型
				retInfoList[i].MaxValue = nums[0] - 1
			}
		}
	}

	// 超过6连的顺子,看能否和剩牌组成三张
	for i := 0; i < len(retInfoList); i++ {
		if retInfoList[i].CardType != CardTypeShunZi || len(retCardsList[i]) < 6 {
			continue
		}

		var nums []int

		// 第一张牌和剩牌能否组成三张
		card := retCardsList[i][0]
		if count[card.Num] == 2 {
			for j := 0; j < 2; j++ {
				nums = append(nums, int(card.Num))
			}

			var retCards []*Card
			var retInfo CardTypeInfo

			// 顺子中找出三张使用的牌
			for _, c := range retCardsList[i] {
				if c.Num == card.Num {
					retCards = append(retCards, c)
					break
				}
			}

			// 剩牌中找出三张使用的牌
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
			count, _, _ = getCountValueLine(dictCards)

			retInfo.CardType = CardTypeSanBuDai
			retInfo.MinValue = nums[0]

			retCardsList = append(retCardsList, retCards)
			retInfoList = append(retInfoList, retInfo)

			// 修改顺子牌型,删除三张中使用的牌
			for k, c := range retCardsList[i] {
				if c.Num == retCardsList[i][0].Num {
					retCardsList[i] = append(retCardsList[i][0:k], retCardsList[i][k+1:]...)
					break
				}
			}

			// 确定顺子最终牌型
			retInfoList[i].MinValue = retInfoList[i].MinValue + 1
		}

		// 最后一张牌和剩牌能否组成三张
		if len(retCardsList[i]) < 6 {
			continue
		}

		card = retCardsList[i][len(retCardsList[i])-1]
		if count[card.Num] == 2 {
			for j := 0; j < 2; j++ {
				nums = append(nums, int(card.Num))
			}

			var retCards []*Card
			var retInfo CardTypeInfo

			// 顺子中找出三张使用的牌
			for _, c := range retCardsList[i] {
				if c.Num == card.Num {
					retCards = append(retCards, c)
					break
				}
			}

			// 剩牌中找出三张使用的牌
			retCards = append(retCards, findCardsByNums(dictCards, nums)...)
			count, _, _ = getCountValueLine(dictCards)

			retInfo.CardType = CardTypeSanBuDai
			retInfo.MinValue = nums[0]

			retCardsList = append(retCardsList, retCards)
			retInfoList = append(retInfoList, retInfo)

			// 修改顺子牌型,删除三张中使用的牌
			for k, c := range retCardsList[i] {
				if c.Num == card.Num {
					retCardsList[i] = append(retCardsList[i][0:k], retCardsList[i][k+1:]...)
					break
				}
			}

			// 确定顺子最终牌型
			retInfoList[i].MaxValue = retInfoList[i].MaxValue - 1
		}
	}
	return
}

// 拆连对
func splitLianDui(size int, dictCards dictMap, count countList, value valueList) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
	if size < 6 {
		return
	}

	for {
		var nums []int
		var exist bool

		for _, v := range value[2] {
			nums = nums[0:0]

			for i := v; i <= NumTypeAce; i++ {
				if count[i] != 2 {
					break
				}

				for j := 0; j < 2; j++ {
					nums = append(nums, i)
				}
			}

			if len(nums) < 6 {
				continue
			}

			exist = true
			var retCards []*Card
			var retInfo CardTypeInfo

			retCards = findCardsByNums(dictCards, nums)
			retCardsList = append(retCardsList, retCards)

			retInfo.CardType = CardTypeLianDui
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[len(nums)-1]
			retInfoList = append(retInfoList, retInfo)

			count, value, _ = getCountValueLine(dictCards)
			break
		}

		if !exist {
			break
		}
	}
	return
}

// 拆三不带
func splitSanBuDai(size int, dictCards dictMap, value valueList) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
	if size < 3 {
		return
	}

	var nums []int

	for _, v := range value[3] {
		nums = nums[0:0]

		for i := 0; i < 3; i++ {
			nums = append(nums, v)
		}

		var retCards []*Card
		var retInfo CardTypeInfo

		retCards = findCardsByNums(dictCards, nums)
		retCardsList = append(retCardsList, retCards)

		retInfo.CardType = CardTypeSanBuDai
		retInfo.MinValue = v
		retInfoList = append(retInfoList, retInfo)
	}
	return
}

// 拆对子
func splitDui(size int, dictCards dictMap, value valueList) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
	if size < 2 {
		return
	}

	var nums []int

	for _, v := range value[2] {
		nums = nums[0:0]

		for i := 0; i < 2; i++ {
			nums = append(nums, v)
		}

		var retCards []*Card
		var retInfo CardTypeInfo

		retCards = findCardsByNums(dictCards, nums)
		retCardsList = append(retCardsList, retCards)

		retInfo.CardType = CardTypeDui
		retInfo.MinValue = v
		retInfoList = append(retInfoList, retInfo)
	}
	return
}

// 拆单
func splitDan(size int, dictCards dictMap, value valueList) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
	if size < 1 {
		return
	}

	var nums []int

	for _, v := range value[1] {
		nums = nums[0:0]

		nums = append(nums, v)

		var retCards []*Card
		var retInfo CardTypeInfo

		retCards = findCardsByNums(dictCards, nums)
		retCardsList = append(retCardsList, retCards)

		retInfo.CardType = CardTypeDan
		retInfo.MinValue = v
		retInfoList = append(retInfoList, retInfo)
	}
	return
}

// 拆四软炸
func splitRuanZhaDan(size int, laiZiSize int, dictCards dictMap, count countList, value valueList, laiZiDictCards dictMap) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
	if size < 4 || laiZiSize == 0 {
		return
	}

	var nums []int

	// 补三张
	for _, v := range value[3] {
		for i := 0; i < 3; i++ {
			nums = append(nums, v)
		}

		var retCards []*Card
		var retInfo CardTypeInfo

		retCards = findCardsByNums(dictCards, nums)                       // 非癞子牌
		retCards = append(retCards, findCardsLaiZi(laiZiDictCards, 1)...) // 癞子牌
		retCardsList = append(retCardsList, retCards)

		retInfo.CardType = CardTypeRuanZhaDan4
		retInfo.MinValue = nums[0]
		retInfoList = append(retInfoList, retInfo)

		count, value, _ = getCountValueLine(dictCards)
		laiZiSize = len(laiZiDictCards)

		if laiZiSize == 0 {
			return
		}
	}
	return
}

// 拆癞子飞机
func splitFeiJiBuDaiLaiZi(size int, laiZiSize int, dictCards dictMap, count countList, value valueList, laiZiDictCards dictMap) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
	if size < 6 {
		return
	}

	if laiZiSize == 0 {
		return splitFeiJiBuDai(size, dictCards, count, value)
	}

	for i := NumTypeThree; i <= NumTypeAce; i++ {
		var nums []int
		var needLaiZi int

		for j := i; j <= NumTypeAce; j++ {
			i = j

			if count[j] == 3 {
				for k := 0; k < 3; k++ {
					nums = append(nums, j)
				}
				continue
			}

			if count[j] == 2 {
				needLaiZi++
				for k := 0; k < 2; k++ {
					nums = append(nums, j)
				}

				if needLaiZi == laiZiSize {
					break
				}
				continue
			}
			break
		}

		// 可以补出飞机
		if len(nums) > 0 && nums[len(nums)-1]-nums[0]+1 >= 2 {
			var retCards []*Card
			var retInfo CardTypeInfo

			retCards = findCardsByNums(dictCards, nums)                               // 非癞子牌
			retCards = append(retCards, findCardsLaiZi(laiZiDictCards, needLaiZi)...) // 癞子牌
			retCardsList = append(retCardsList, retCards)

			retInfo.CardType = CardTypeFeiJiBuDai
			retInfo.MinValue = nums[0]
			retInfo.MaxValue = nums[len(nums)-1]
			retInfoList = append(retInfoList, retInfo)

			count, value, _ = getCountValueLine(dictCards)
			laiZiSize = len(laiZiDictCards)
		}
	}
	return
}

// 拆癞子三不带
func splitSanBuDaiLaiZi(size int, laiZiSize int, dictCards dictMap, count countList, value valueList, laiZiDictCards dictMap) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
	if size < 3 || laiZiSize == 0 {
		return
	}

	var nums []int

	// 补三张
	for _, v := range value[3] {
		for i := 0; i < 3; i++ {
			nums = append(nums, v)
		}

		var retCards []*Card
		var retInfo CardTypeInfo

		retCards = findCardsByNums(dictCards, nums)                       // 非癞子牌
		retCards = append(retCards, findCardsLaiZi(laiZiDictCards, 1)...) // 癞子牌
		retCardsList = append(retCardsList, retCards)

		retInfo.CardType = CardTypeRuanZhaDan4
		retInfo.MinValue = nums[0]
		retInfoList = append(retInfoList, retInfo)

		count, value, _ = getCountValueLine(dictCards)
		laiZiSize = len(laiZiDictCards)

		if laiZiSize == 0 {
			return
		}
	}
	return
}

// 拆癞子对
