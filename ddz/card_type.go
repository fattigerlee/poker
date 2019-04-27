package ddz

import (
	"fmt"
	"sort"
)

// 牌型
type CardTypeInfo struct {
	CardType CardType // 牌型
	MinValue int      // 最小值
	MaxValue int      // 最大值
}

func (c *CardTypeInfo) String() string {
	return fmt.Sprintf("牌型:%s 最小值:%d 最大值:%d", c.CardType, c.MinValue, c.MaxValue)
}

type valueList [5][]int // 单张,对子,三张,四张牌的牌值
type countList [18]int  // 所有牌的数量

// 获取牌型
func GetCardType(cards []*Card) (list []*CardTypeInfo) {
	newCards := getNoLaiZiCards(cards)       // 无癞子牌数组
	laiZiCount := len(cards) - len(newCards) // 癞子数量

	if laiZiCount == 0 {
		// 无癞子算法
		return analysis(cards)
	} else {
		// 癞子算法
		return analysisLaiZi(cards, newCards, laiZiCount)
	}
}

// 牌值计数
func cardCount(cards []*Card) countList {
	var count [18]int
	for _, c := range cards {
		count[int(c.Num)]++
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

// 获取非癞子数组
func getNoLaiZiCards(cards []*Card) []*Card {
	var newCards []*Card

	for _, c := range cards {
		if !c.LaiZi {
			newCards = append(newCards, c)
		}
	}
	return newCards
}

// 判断牌值是否是大小王
func isJoker(value int) bool {
	if value == NumTypeSmallJoker || value == NumTypeBigJoker {
		return true
	}
	return false
}

// 判断牌值是否是大小王和2
func isJokerAndTwo(value int) bool {
	if value == NumTypeSmallJoker || value == NumTypeBigJoker || value == NumTypeTwo {
		return true
	}
	return false
}

// 无癞子算法
func analysis(cards []*Card) (list []*CardTypeInfo) {
	// 排序
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Num < cards[j].Num
	})

	size := len(cards)        // 牌总数量
	count := cardCount(cards) // 每张牌数量
	var value valueList       // 所有单张,对子,三张,四张的牌值
	for i := 3; i < 18; i++ {
		switch count[i] {
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

	switch len(cards) {
	case 1:
		// 单
		if info := isDan(cards); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

	case 2:
		// 对
		if info := isDui(cards); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

		// 火箭
		if info := isHuoJian(cards, count); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

	case 3:
		// 三不带
		if info := isSanBuDai(cards, count); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

	case 4:
		// 三带一
		if info := isSanDaiYi(cards, count); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

		// 炸弹
		if info := isZhaDan(cards, count); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

	case 5:
		// 三带二
		if info := isSanDaiEr(cards, count); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

	case 6:
		// 四带单
		if info := isSiDaiDan(cards, count); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

	case 8:
		// 四带对
		if info := isSiDaiDui(cards, count); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}
	}

	// 顺子
	if info := isShunZi(cards, count); info.CardType != CardTypeNone {
		list = append(list, &info)
		return
	}

	// 连对
	if info := isLianDui(size, value); info.CardType != CardTypeNone {
		list = append(list, &info)
		return
	}

	// 飞机不带
	if info := isFeiJiBuDai(cards, count); info.CardType != CardTypeNone {
		list = append(list, &info)
		return
	}

	// 飞机带一
	if info := isFeiJiDaiYi(cards, count); info.CardType != CardTypeNone {
		list = append(list, &info)
		return
	}

	// 飞机带二
	if info := isFeiJiDaiEr(cards, count); info.CardType != CardTypeNone {
		list = append(list, &info)
		return
	}

	// 连炸
	if info := isLianZha(cards, count); info.CardType != CardTypeNone {
		list = append(list, &info)
		return
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
func isLianDui(size int, value valueList) (info CardTypeInfo) {
	if size < 6 || size%2 != 0 {
		return
	}

	// 只能有对子
	if len(value[1]) != 0 || len(value[3]) != 0 || len(value[4]) != 0 {
		return
	}

	// 对子必须连续
	minValue := value[2][0]
	maxValue := value[2][len(value[2])-1]
	valueRange := maxValue - minValue + 1

	if valueRange != size/2 {
		return
	}

	info.CardType = CardTypeLianDui
	info.MinValue = minValue
	info.MaxValue = maxValue

	//minValue := getMinValue(count, 2)
	//maxValue := getMaxValue(count, 2)
	//valueRange := maxValue - minValue + 1
	//
	//size := len(cards)
	//if size < 6 || size != valueRange*2 {
	//	return
	//}
	//
	//exist := true
	//for i := minValue; i <= maxValue; i++ {
	//	if count[i] != 2 {
	//		exist = false
	//		break
	//	}
	//}
	//
	//if exist {
	//	info.CardType = CardTypeLianDui
	//	info.MinValue = minValue
	//	info.MaxValue = maxValue
	//	return
	//}
	return
}

// 三不带
func isSanBuDai(cards []*Card, count [18]int) (info CardTypeInfo) {
	if len(cards) != 3 {
		return
	}

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
func isSanDaiYi(cards []*Card, count [18]int) (info CardTypeInfo) {
	if len(cards) != 4 {
		return
	}

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
func isSanDaiEr(cards []*Card, count [18]int) (info CardTypeInfo) {
	if len(cards) != 5 {
		return
	}

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
func isSiDaiDan(cards []*Card, count [18]int) (info CardTypeInfo) {
	if len(cards) != 6 {
		return
	}

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
func isSiDaiDui(cards []*Card, count [18]int) (info CardTypeInfo) {
	if len(cards) != 8 {
		return
	}

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
func isShunZi(cards []*Card, count [18]int) (info CardTypeInfo) {
	if len(cards) < 5 {
		return
	}

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
func isFeiJiBuDai(cards []*Card, count [18]int) (info CardTypeInfo) {
	if len(cards) < 6 {
		return
	}

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
func isFeiJiDaiYi(cards []*Card, count [18]int) (info CardTypeInfo) {
	if len(cards) < 8 {
		return
	}

	// 连炸不是飞机带一
	if isLianZha(cards, count).CardType != CardTypeNone {
		return
	}

	var minValue int
	var maxValue int
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
func isFeiJiDaiEr(cards []*Card, count [18]int) (info CardTypeInfo) {
	if len(cards) < 10 {
		return
	}

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
func isZhaDan(cards []*Card, count [18]int) (info CardTypeInfo) {
	if len(cards) != 4 {
		return
	}

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
func isHuoJian(cards []*Card, count [18]int) (info CardTypeInfo) {
	if len(cards) != 2 {
		return
	}

	if count[16] == 1 && count[17] == 1 {
		info.CardType = CardTypeHuoJian
		return
	}
	return
}

// 连炸
func isLianZha(cards []*Card, count [18]int) (info CardTypeInfo) {
	if len(cards) < 8 {
		return
	}

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

// 癞子算法
func analysisLaiZi(cards []*Card, newCards []*Card, laiZiCount int) (list []*CardTypeInfo) {
	size := len(cards)             // 牌总数量
	cardsCount := cardCount(cards) // 每张牌数量
	var value valueList            // 所有单张,对子,三张,四张的牌值
	for i := 3; i < 18; i++ {
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

	newCardsCount := cardCount(newCards) // 非癞子每张牌数量
	var newValue valueList               // 所有非癞子单张,对子,三张,四张的牌值
	var newLine []int                    // 非癞子连牌
	for i := 3; i < 18; i++ {
		if newCardsCount[i] > 0 {
			newLine = append(newLine, i)
		}

		switch newCardsCount[i] {
		case 1:
			newValue[1] = append(newValue[1], i)
		case 2:
			newValue[2] = append(newValue[2], i)
		case 3:
			newValue[3] = append(newValue[3], i)
		case 4:
			newValue[4] = append(newValue[4], i)
		}
	}

	// 癞子单
	if info := isDanLaiZi(size, value); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 癞子对
	if info := isDuiLaiZi(size, value, newLine); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 癞子三不带
	if info := isSanBuDaiLaiZi(size, value, newLine); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 癞子三带一
	if infoList := isSanDaiYiLaiZi(size, newCardsCount, newLine, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子三带二
	if infoList := isSanDaiErLaiZi(size, newCardsCount, newLine, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子四带单
	if infoList := isSiDaiDanLaiZi(size, newCardsCount, newLine, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子四带对
	if infoList := isSiDaiDuiLaiZi(size, newCardsCount, newLine, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子顺子
	if infoList := isLianLaiZi(CardTypeShunZi, size, newValue, newCardsCount, newLine, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子连对
	if infoList := isLianLaiZi(CardTypeLianDui, size, newValue, newCardsCount, newLine, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子飞机不带
	if infoList := isLianLaiZi(CardTypeFeiJiBuDai, size, newValue, newCardsCount, newLine, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子飞机带一
	if infoList := isFeiJiDaiYiLaiZi(size, newCardsCount, newLine, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子飞机带二
	if infoList := isFeiJiDaiErLaiZi(size, newCardsCount, newLine, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 四纯癞子炸
	if info := isChunLaiZiZhaDan(size, value, newLine); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 软炸
	if info := isRuanZhaDan(size, newLine); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 癞子炸
	if info := isLaiZiZhaDan(size, newLine); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 软连炸
	if info := isRuanLianZha(size, newCardsCount, newLine, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}
	return
}

// 癞子单
func isDanLaiZi(size int, value valueList) (info CardTypeInfo) {
	if size != 1 {
		return
	}

	info.CardType = CardTypeDan
	info.MinValue = value[1][0]
	return
}

// 癞子对
func isDuiLaiZi(size int, value valueList, newLine []int) (info CardTypeInfo) {
	if size != 2 {
		return
	}

	// 非癞子牌有多少种牌值
	switch len(newLine) {
	case 0:
		if len(value[2]) != 1 {
			return
		}

		info.CardType = CardTypeDui
		info.MinValue = value[2][0]

	case 1:
		if isJoker(newLine[0]) {
			return
		}

		info.CardType = CardTypeDui
		info.MinValue = newLine[0]
	}
	return
}

// 癞子三不带
func isSanBuDaiLaiZi(size int, value valueList, newLine []int) (info CardTypeInfo) {
	if size != 3 {
		return
	}

	// 非癞子牌有多少种牌值
	switch len(newLine) {
	case 0:
		if len(value[3]) != 1 {
			return
		}

		info.CardType = CardTypeSanBuDai
		info.MinValue = value[3][0]

	case 1:
		if isJoker(newLine[0]) {
			return
		}

		info.CardType = CardTypeSanBuDai
		info.MinValue = newLine[0]
	}
	return
}

// 癞子三带一
func isSanDaiYiLaiZi(size int, newCardsCount countList, newLine []int, laiZiCount int) (list []*CardTypeInfo) {
	if size != 4 {
		return
	}

	var min CardTypeInfo // 最小值
	var max CardTypeInfo // 最大值

	// 非癞子牌有多少种牌值
	switch len(newLine) {
	case 2:
		var line []int
		for _, v := range newLine {
			if isJoker(v) {
				continue
			}
			line = append(line, v)
		}

		if len(line) == 0 {
			return
		}

		min = minSanDaiYiLaiZi(newCardsCount, line, laiZiCount)
		max = maxSanDaiYiLaiZi(newCardsCount, line, laiZiCount)

	default:
		return
	}

	if min.CardType == CardTypeNone && max.CardType == CardTypeNone {
		return
	}

	if min.MinValue == max.MinValue {
		list = append(list, &min)
	} else {
		list = append(list, &min, &max)
	}
	return
}

// 最小癞子三带一
func minSanDaiYiLaiZi(newCardsCount countList, line []int, laiZiCount int) (info CardTypeInfo) {
	for i := 0; i < len(line); i++ {
		info = sanDaiYiLaiZi(line[i], newCardsCount, line, laiZiCount)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 最大癞子三带一
func maxSanDaiYiLaiZi(newCardsCount countList, line []int, laiZiCount int) (info CardTypeInfo) {
	for i := len(line) - 1; i >= 0; i-- {
		info = sanDaiYiLaiZi(line[i], newCardsCount, line, laiZiCount)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 癞子三带一结果
func sanDaiYiLaiZi(value int, newCardsCount countList, line []int, laiZiCount int) (info CardTypeInfo) {
	var laiZi int

	// 补三张
	switch newCardsCount[value] {
	case 0:
		laiZi += 3
	case 1:
		laiZi += 2
	case 2:
		laiZi += 1
	}

	if laiZi > laiZiCount {
		return
	}

	info.CardType = CardTypeSanDaiYi
	info.MinValue = value
	return
}

// 癞子三带二
func isSanDaiErLaiZi(size int, newCardsCount countList, newLine []int, laiZiCount int) (list []*CardTypeInfo) {
	if size != 5 {
		return
	}

	var min CardTypeInfo // 最小值
	var max CardTypeInfo // 最大值

	// 非癞子牌有多少种牌值
	switch len(newLine) {
	case 2:
		for _, v := range newLine {
			if isJoker(v) {
				return
			}
		}

		min = minSanDaiErLaiZi(newCardsCount, newLine, laiZiCount)
		max = maxSanDaiErLaiZi(newCardsCount, newLine, laiZiCount)

	default:
		return
	}

	if min.CardType == CardTypeNone && max.CardType == CardTypeNone {
		return
	}

	if min.MinValue == max.MinValue {
		list = append(list, &min)
	} else {
		list = append(list, &min, &max)
	}
	return
}

// 最小癞子三带二
func minSanDaiErLaiZi(newCardsCount countList, line []int, laiZiCount int) (info CardTypeInfo) {
	for i := 0; i < len(line); i++ {
		info = sanDaiErLaiZi(line[i], newCardsCount, line, laiZiCount)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 最大癞子三带二
func maxSanDaiErLaiZi(newCardsCount countList, line []int, laiZiCount int) (info CardTypeInfo) {
	for i := len(line) - 1; i >= 0; i-- {
		info = sanDaiErLaiZi(line[i], newCardsCount, line, laiZiCount)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 癞子三带二结果
func sanDaiErLaiZi(value int, newCardsCount countList, line []int, laiZiCount int) (info CardTypeInfo) {
	var laiZi int

	// 补三张
	switch newCardsCount[value] {
	case 0:
		laiZi += 3
	case 1:
		laiZi += 2
	case 2:
		laiZi += 1
	}

	if laiZi > laiZiCount {
		return
	}

	// 补对子
	for _, v := range line {
		if v == value {
			continue
		}

		switch newCardsCount[v] {
		case 0:
			laiZi += 2
		case 1:
			laiZi += 1
		}

		if laiZi > laiZiCount {
			return
		}
	}

	info.CardType = CardTypeSanDaiEr
	info.MinValue = value
	return
}

// 癞子四带单
func isSiDaiDanLaiZi(size int, newCardsCount countList, newLine []int, laiZiCount int) (list []*CardTypeInfo) {
	if size != 6 {
		return
	}

	var min CardTypeInfo // 最小值
	var max CardTypeInfo // 最大值

	// 非癞子牌有多少种牌值
	switch len(newLine) {
	case 0, 1:
		return

	default:
		var line []int
		for _, v := range newLine {
			if isJoker(v) {
				continue
			}
			line = append(line, v)
		}

		min = minSiDaiDanLaiZi(newCardsCount, line, laiZiCount)
		max = maxSiDaiDanLaiZi(newCardsCount, line, laiZiCount)
	}

	if min.CardType == CardTypeNone && max.CardType == CardTypeNone {
		return
	}

	if min.MinValue == max.MinValue {
		list = append(list, &min)
	} else {
		list = append(list, &min, &max)
	}
	return
}

// 最小癞子四带单
func minSiDaiDanLaiZi(newCardsCount countList, line []int, laiZiCount int) (info CardTypeInfo) {
	// 最小值为3
	info = siDaiDanLaiZi(NumTypeThree, newCardsCount, laiZiCount)
	if info.CardType != CardTypeNone {
		return
	}

	for i := 0; i < len(line); i++ {
		if line[i] == NumTypeThree {
			continue
		}

		info = siDaiDanLaiZi(line[i], newCardsCount, laiZiCount)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 最大癞子四带单
func maxSiDaiDanLaiZi(newCardsCount countList, line []int, laiZiCount int) (info CardTypeInfo) {
	// 最大值为2
	info = siDaiDanLaiZi(NumTypeTwo, newCardsCount, laiZiCount)
	if info.CardType != CardTypeNone {
		return
	}

	for i := len(line) - 1; i >= 0; i-- {
		if line[i] == NumTypeTwo {
			continue
		}

		info = siDaiDanLaiZi(line[i], newCardsCount, laiZiCount)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 癞子四带单结果
func siDaiDanLaiZi(value int, newCardsCount countList, laiZiCount int) (info CardTypeInfo) {
	var laiZi int

	// 补四张
	switch newCardsCount[value] {
	case 0:
		laiZi += 4
	case 1:
		laiZi += 3
	case 2:
		laiZi += 2
	case 3:
		laiZi += 1
	}

	if laiZi > laiZiCount {
		return
	}

	info.CardType = CardTypeSiDaiDan
	info.MinValue = value
	return
}

// 癞子四带对
func isSiDaiDuiLaiZi(size int, newCardsCount countList, newLine []int, laiZiCount int) (list []*CardTypeInfo) {
	if size != 8 {
		return
	}

	var min CardTypeInfo // 最小值
	var max CardTypeInfo // 最大值

	// 非癞子牌有多少种牌值
	switch len(newLine) {
	case 0, 1:
		return

	default:
		for _, v := range newLine {
			if isJoker(v) {
				return
			}
		}

		min = minSiDaiDuiLaiZi(newCardsCount, newLine, laiZiCount)
		max = maxSiDaiDuiLaiZi(newCardsCount, newLine, laiZiCount)
	}

	if min.CardType == CardTypeNone && max.CardType == CardTypeNone {
		return
	}

	if min.MinValue == max.MinValue {
		list = append(list, &min)
	} else {
		list = append(list, &min, &max)
	}
	return
}

// 最小癞子四带对
func minSiDaiDuiLaiZi(newCardsCount countList, line []int, laiZiCount int) (info CardTypeInfo) {
	// 最小值为3
	info = siDaiDuiLaiZi(NumTypeThree, newCardsCount, line, laiZiCount)
	if info.CardType != CardTypeNone {
		return
	}

	for i := 0; i < len(line); i++ {
		if line[i] == NumTypeThree {
			continue
		}

		info = siDaiDuiLaiZi(line[i], newCardsCount, line, laiZiCount)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 最大癞子四带对
func maxSiDaiDuiLaiZi(newCardsCount countList, line []int, laiZiCount int) (info CardTypeInfo) {
	// 最大值为2
	info = siDaiDuiLaiZi(NumTypeTwo, newCardsCount, line, laiZiCount)
	if info.CardType != CardTypeNone {
		return
	}

	for i := len(line) - 1; i >= 0; i-- {
		if line[i] == NumTypeTwo {
			continue
		}

		info = siDaiDuiLaiZi(line[i], newCardsCount, line, laiZiCount)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 癞子四带对结果
func siDaiDuiLaiZi(value int, newCardsCount countList, line []int, laiZiCount int) (info CardTypeInfo) {
	var laiZi int

	// 补四张
	switch newCardsCount[value] {
	case 0:
		laiZi += 4
	case 1:
		laiZi += 3
	case 2:
		laiZi += 2
	case 3:
		laiZi += 1
	}

	if laiZi > laiZiCount {
		return
	}

	// 补对子
	for _, v := range line {
		if v == value {
			continue
		}

		switch newCardsCount[v] {
		case 0:
			laiZi += 2
		case 1:
			laiZi += 1
		}

		if laiZi > laiZiCount {
			return
		}
	}

	info.CardType = CardTypeSiDaiDui
	info.MinValue = value
	return
}

// 癞子连牌(顺子,连对,飞机不带)
func isLianLaiZi(cardType CardType, size int, newValue valueList, newCardsCount countList, newLine []int, laiZiCount int) (list []*CardTypeInfo) {
	var min CardTypeInfo // 最小值
	var max CardTypeInfo // 最大值

	switch cardType {
	case CardTypeShunZi:
		// 顺子
		if size < 5 || size > 12 {
			return
		}

		// 非癞子牌只能有单张
		if len(newValue[2]) != 0 || len(newValue[3]) != 0 || len(newValue[4]) != 0 {
			return
		}

	case CardTypeLianDui:
		// 连对
		if size < 6 || size%2 != 0 {
			return
		}

		// 非癞子牌只能有单张和对子
		if len(newValue[3]) != 0 || len(newValue[4]) != 0 {
			return
		}

	case CardTypeFeiJiBuDai:
		// 飞机不带
		if size < 6 || size%3 != 0 {
			return
		}

		// 非癞子牌只能有单张和对子和三张
		if len(newValue[4]) != 0 {
			return
		}
	}

	// 非癞子牌有多少种牌值
	switch len(newLine) {
	case 0, 1:
		return

	default:
		for _, v := range newLine {
			if isJokerAndTwo(v) {
				return
			}
		}

		// 最小值
		min = lianLaiZi(ResultTypeMin, cardType, newCardsCount, newLine, laiZiCount)

		// 最大值
		max = lianLaiZi(ResultTypeMax, cardType, newCardsCount, newLine, laiZiCount)
	}

	if min.CardType == CardTypeNone && max.CardType == CardTypeNone {
		return
	}

	if min.MinValue == max.MinValue {
		list = append(list, &min)
	} else {
		list = append(list, &min, &max)
	}
	return
}

// 连癞子
func lianLaiZi(resultType ResultType, cardType CardType, newCardsCount countList, line []int, laiZiCount int) (info CardTypeInfo) {
	var laiZi int

	// 补牌
	for i := line[0]; i <= line[len(line)-1]; i++ {
		switch newCardsCount[i] {
		case 0:
			switch cardType {
			case CardTypeShunZi:
				laiZi += 1
			case CardTypeLianDui:
				laiZi += 2
			case CardTypeFeiJiBuDai:
				laiZi += 3
			}
		case 1:
			switch cardType {
			case CardTypeLianDui:
				laiZi += 1
			case CardTypeFeiJiBuDai:
				laiZi += 2
			}
		case 2:
			switch cardType {
			case CardTypeFeiJiBuDai:
				laiZi += 1
			}
		}
	}

	if laiZi > laiZiCount {
		return
	}

	// 处理多余的癞子
	var laiZiGroup int
	limitLaiZi := laiZiCount - laiZi
	switch cardType {
	case CardTypeShunZi:
		laiZiGroup = limitLaiZi

	case CardTypeLianDui:
		laiZiGroup = limitLaiZi / 2

	case CardTypeFeiJiBuDai:
		laiZiGroup = limitLaiZi / 3
	}

	// 结果处理
	var minValue int
	var maxValue int
	switch resultType {
	case ResultTypeMin:
		if line[0]-laiZiGroup < NumTypeThree {
			minValue = NumTypeThree
			maxValue = line[len(line)-1] + laiZiGroup - (line[0] - NumTypeThree)
		} else {
			minValue = line[0] - laiZiGroup
			maxValue = line[len(line)-1]
		}

	case ResultTypeMax:
		if line[len(line)-1]+laiZiGroup > NumTypeAce {
			minValue = line[0] - (laiZiGroup - (NumTypeAce - line[len(line)-1]))
			maxValue = NumTypeAce
		} else {
			minValue = line[0]
			maxValue = line[len(line)-1] + laiZiGroup
		}
	}

	switch cardType {
	case CardTypeShunZi:
		info.CardType = CardTypeShunZi
		info.MinValue = minValue
		info.MaxValue = maxValue
	case CardTypeLianDui:
		info.CardType = CardTypeLianDui
		info.MinValue = minValue
		info.MaxValue = maxValue
	case CardTypeFeiJiBuDai:
		info.CardType = CardTypeFeiJiBuDai
		info.MinValue = minValue
		info.MaxValue = maxValue
	}
	return
}

// 癞子飞机带一
func isFeiJiDaiYiLaiZi(size int, newCardsCount countList, newLine []int, laiZiCount int) (list []*CardTypeInfo) {
	if size < 8 || size%4 != 0 {
		return
	}

	var min CardTypeInfo   // 最小值
	var max CardTypeInfo   // 最大值
	valueRange := size / 4 // 几飞机

	switch len(newLine) {
	case 0, 1:
		return

	default:
		var line []int // 去除2和大小王的值
		for _, v := range newLine {
			if isJokerAndTwo(v) {
				continue
			}
			line = append(line, v)
		}

		// 最小值
		min = minFeiJiDaiYiLaiZi(newCardsCount, line, valueRange, laiZiCount)

		// 最大值
		max = maxFeiJiDaiYiLaiZi(newCardsCount, line, valueRange, laiZiCount)
	}

	if min.CardType == CardTypeNone && max.CardType == CardTypeNone {
		return
	}

	if min.MinValue == max.MinValue {
		list = append(list, &min)
	} else {
		list = append(list, &min, &max)
	}
	return
}

// 最小癞子飞机带一
func minFeiJiDaiYiLaiZi(newCardsCount countList, line []int, valueRange int, laiZiCount int) (info CardTypeInfo) {
	// 最小值为3的飞机
	info = feiJiDaiYiLaiZi(newCardsCount, NumTypeThree, NumTypeThree+valueRange-1, laiZiCount)
	if info.CardType != CardTypeNone {
		return
	}

	// 最小值不为3的飞机
	for i := 0; i < len(line); i++ {
		min := line[i] - valueRange + 1
		if min <= NumTypeThree {
			continue
		}

		info = feiJiDaiYiLaiZi(newCardsCount, min, line[i], laiZiCount)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 最大癞子飞机带一
func maxFeiJiDaiYiLaiZi(newCardsCount countList, line []int, valueRange int, laiZiCount int) (info CardTypeInfo) {
	// 最大值为A的飞机
	info = feiJiDaiYiLaiZi(newCardsCount, NumTypeAce-valueRange+1, NumTypeAce, laiZiCount)
	if info.CardType != CardTypeNone {
		return
	}

	// 最大值不为A的飞机
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] == NumTypeAce {
			continue
		}

		min := line[i] - valueRange + 1
		if min <= NumTypeThree {
			info = feiJiDaiYiLaiZi(newCardsCount, NumTypeThree, NumTypeThree+valueRange-1, laiZiCount)
			if info.CardType != CardTypeNone {
				return
			}
		} else {
			info = feiJiDaiYiLaiZi(newCardsCount, min, line[i], laiZiCount)
			if info.CardType != CardTypeNone {
				return
			}
		}
	}
	return
}

// 癞子飞机带一结果
func feiJiDaiYiLaiZi(newCardsCount countList, min int, max int, laiZiCount int) (info CardTypeInfo) {
	var laiZi int

	// 补飞机
	for i := min; i <= max; i++ {
		switch newCardsCount[i] {
		case 0:
			laiZi += 3
		case 1:
			laiZi += 2
		case 2:
			laiZi += 1
		}
	}

	if laiZi > laiZiCount {
		return
	}
	info.CardType = CardTypeFeiJiDaiYi
	info.MinValue = min
	info.MaxValue = max
	return
}

// 癞子飞机带二
func isFeiJiDaiErLaiZi(size int, newCardsCount countList, newLine []int, laiZiCount int) (list []*CardTypeInfo) {
	if size < 10 || size%5 != 0 {
		return
	}

	var min CardTypeInfo   // 最小值
	var max CardTypeInfo   // 最大值
	valueRange := size / 5 // 几飞机

	switch len(newLine) {
	case 0, 1:
		return

	default:
		// 不能有大小王
		for _, v := range newLine {
			if isJoker(v) {
				return
			}
		}

		var line []int // 去除2
		for _, v := range newLine {
			if isJokerAndTwo(v) {
				continue
			}
			line = append(line, v)
		}

		// 最小值
		min = minFeiJiDaiErLaiZi(newCardsCount, line, newLine, valueRange, laiZiCount)

		// 最大值
		max = maxFeiJiDaiErLaiZi(newCardsCount, line, newLine, valueRange, laiZiCount)
	}

	if min.CardType == CardTypeNone && max.CardType == CardTypeNone {
		return
	}

	if min.MinValue == max.MinValue {
		list = append(list, &min)
	} else {
		list = append(list, &min, &max)
	}
	return
}

// 最小癞子飞机带二
func minFeiJiDaiErLaiZi(newCardsCount countList, line []int, newLine []int, valueRange int, laiZiCount int) (info CardTypeInfo) {
	// 最小值为3的飞机
	info = feiJiDaiErLaiZi(newCardsCount, newLine, NumTypeThree, NumTypeThree+valueRange-1, laiZiCount)
	if info.CardType != CardTypeNone {
		return
	}

	// 最小值不为3的飞机
	for i := 0; i < len(line); i++ {
		min := line[i] - valueRange + 1
		if min <= NumTypeThree {
			continue
		}

		info = feiJiDaiErLaiZi(newCardsCount, newLine, min, line[i], laiZiCount)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 最大癞子飞机带二
func maxFeiJiDaiErLaiZi(newCardsCount countList, line []int, newLine []int, valueRange int, laiZiCount int) (info CardTypeInfo) {
	// 最大值为A的飞机
	info = feiJiDaiErLaiZi(newCardsCount, newLine, NumTypeAce-valueRange+1, NumTypeAce, laiZiCount)
	if info.CardType != CardTypeNone {
		return
	}

	// 最大值不为A的飞机
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] == NumTypeAce {
			continue
		}

		min := line[i] - valueRange + 1
		if min <= NumTypeThree {
			info = feiJiDaiErLaiZi(newCardsCount, newLine, NumTypeThree, NumTypeThree+valueRange-1, laiZiCount)
			if info.CardType != CardTypeNone {
				return
			}
		} else {
			info = feiJiDaiErLaiZi(newCardsCount, newLine, min, line[i], laiZiCount)
			if info.CardType != CardTypeNone {
				return
			}
		}
	}
	return
}

// 癞子飞机带二结果
func feiJiDaiErLaiZi(newCardsCount countList, newLine []int, min int, max int, laiZiCount int) (info CardTypeInfo) {
	var laiZi int

	// 补飞机
	for i := min; i <= max; i++ {
		switch newCardsCount[i] {
		case 0:
			laiZi += 3
		case 1:
			laiZi += 2
		case 2:
			laiZi += 1
		}

		if laiZi > laiZiCount {
			return
		}
	}

	// 补对子
	for _, v := range newLine {
		var exist bool
		for i := min; i <= max; i++ {
			if v == i {
				exist = true
				break
			}
		}

		if exist {
			continue
		}

		switch newCardsCount[v] {
		case 0:
			laiZi += 2
		case 1:
			laiZi += 1
		}

		if laiZi > laiZiCount {
			return
		}
	}

	info.CardType = CardTypeFeiJiDaiEr
	info.MinValue = min
	info.MaxValue = max
	return
}

// 软炸
func isRuanZhaDan(size int, newLine []int) (info CardTypeInfo) {
	if len(newLine) != 1 {
		return
	}

	if isJoker(newLine[0]) {
		return
	}

	switch size {
	case 4:
		info.CardType = CardTypeRuanZhaDan4

	case 5:
		info.CardType = CardTypeRuanZhaDan5

	case 6:
		info.CardType = CardTypeRuanZhaDan6

	case 7:
		info.CardType = CardTypeRuanZhaDan7

	case 8:
		info.CardType = CardTypeRuanZhaDan8

	case 9:
		info.CardType = CardTypeRuanZhaDan9

	case 10:
		info.CardType = CardTypeRuanZhaDan10

	case 11:
		info.CardType = CardTypeRuanZhaDan11

	case 12:
		info.CardType = CardTypeRuanZhaDan12
	}

	if info.CardType == CardTypeNone {
		return
	}

	info.MinValue = newLine[0]
	return
}

// 四纯癞子炸
func isChunLaiZiZhaDan(size int, value valueList, newLine []int) (info CardTypeInfo) {
	if size != 4 {
		return
	}

	if len(newLine) != 0 {
		return
	}

	if len(value[4]) != 1 {
		return
	}

	info.CardType = CardTypeChunLaiZiZhaDan
	info.MinValue = value[4][0]
	return
}

// 癞子炸
func isLaiZiZhaDan(size int, newLine []int) (info CardTypeInfo) {
	if len(newLine) != 0 {
		return
	}

	switch size {
	case 4:
		info.CardType = CardTypeLaiZiZhaDan4

	case 5:
		info.CardType = CardTypeLaiZiZhaDan5

	case 6:
		info.CardType = CardTypeLaiZiZhaDan6

	case 7:
		info.CardType = CardTypeLaiZiZhaDan7

	case 8:
		info.CardType = CardTypeLaiZiZhaDan8
	}

	if info.CardType == CardTypeNone {
		return
	}
	return
}

// 软连炸
func isRuanLianZha(size int, newCardsCount countList, newLine []int, laiZiCount int) (info CardTypeInfo) {
	if size < 8 || size%4 != 0 {
		return
	}

	// 不能有2和大小王
	for _, v := range newLine {
		if isJokerAndTwo(v) {
			return
		}
	}

	var laiZi int

	// 补牌
	for i := newLine[0]; i <= newLine[len(newLine)-1]; i++ {
		switch newCardsCount[i] {
		case 0:
			laiZi += 4
		case 1:
			laiZi += 3
		case 2:
			laiZi += 2
		case 3:
			laiZi += 1
		}

		if laiZi > laiZiCount {
			return
		}
	}

	info.CardType = CardTypeRuanLianZha
	info.MinValue = newLine[0]
	info.MaxValue = newLine[len(newLine)-1]
	return
}
