package ddz

// 拆牌(经典模式)
func SplitCardsJingDian(cards []*Card) (retCardsList [][]*Card, retInfoList []CardTypeInfo) {
	size := len(cards)
	dictCards := convertToMap(cards)
	_, value, _ := getCountValueLine(dictCards)

	// 火箭
	if retCards, retInfo := findHuoJian(size, dictCards, value); retInfo.CardType != CardTypeNone {
		retCardsList = append(retCardsList, retCards)
		retInfoList = append(retInfoList, retInfo)
	}

	// 炸弹
	if cardsList, infoList := splitZhaDan(size, dictCards, value); len(infoList) != 0 {
		retCardsList = append(retCardsList, cardsList...)
		retInfoList = append(retInfoList, infoList...)
	}

	return
}

// 拆牌(不洗牌)
func SplitCardsBuXiPai() {

}

// 拆牌(不洗牌+双王癞子模式)
func SplitCardsBuXiPaiLaiZi() {

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
