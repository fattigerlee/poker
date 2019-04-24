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
func cardCount(cards []*Card) [18]int {
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

// 判断牌中是否有2和大小王
func existTwoAndJoker(value valueList) bool {
	for _, vs := range value {
		for _, v := range vs {
			if v == 15 || v == 16 || v == 17 {
				return true
			}
		}
	}
	return false
}

// 获取连牌缺失的牌数量
func getLineGap(value []int) int {
	var gap int
	for i := 0; i < len(value)-1; i++ {
		gap += value[i+1] - value[i] - 1
	}
	return gap
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
	if info := isDuiLaiZi(size, value, newValue, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 癞子三不带
	if info := isSanBuDaiLaiZi(size, value, newValue, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 癞子三带一
	if infoList := isSanDaiYiLaiZi(size, newValue, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子三带二
	if infoList := isSanDaiErLaiZi(size, newValue, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子四带单
	if infoList := isSiDaiDanLaiZi(size, newValue, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子四带对
	if infoList := isSiDaiDuiLaiZi(size, newValue, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子顺子
	if infoList := isShunZiLaiZi(size, newValue, newLine, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子连对
	if infoList := isLianDuiLaiZi(size, newValue, newLine, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子飞机不带
	if infoList := isFeiJiBuDaiLaiZi(size, newValue, newLine, laiZiCount); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子飞机带一

	// 癞子飞机带二

	// 四软炸
	if info := isRuanZhaDan4(size, newValue, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 五软炸
	if info := isRuanZhaDan5(size, newValue, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 六软炸
	if info := isRuanZhaDan6(size, newValue, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 七软炸
	if info := isRuanZhaDan7(size, newValue, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 八软炸
	if info := isRuanZhaDan8(size, newValue, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 九软炸
	if info := isRuanZhaDan9(size, newValue, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 十软炸
	if info := isRuanZhaDan10(size, newValue, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 十一软炸
	if info := isRuanZhaDan11(size, newValue, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 十二软炸
	if info := isRuanZhaDan12(size, newValue, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}
	return
}

// 癞子单
func isDanLaiZi(size int, value valueList) (info CardTypeInfo) {
	if size != 1 {
		return
	}

	// 癞子牌为1单张
	if len(value[1]) == 1 {
		info.CardType = CardTypeDan
		info.MinValue = value[1][0]
	}
	return
}

// 癞子对
func isDuiLaiZi(size int, value valueList, newValue valueList, laiZiCount int) (info CardTypeInfo) {
	if size != 2 {
		return
	}

	switch laiZiCount {
	case 1:
		// 非癞子牌为1单张
		if len(newValue[1]) == 1 {
			info.CardType = CardTypeDui
			info.MinValue = newValue[1][0]
		}

	case 2:
		// 癞子牌为1对子
		if len(value[2]) == 1 {
			info.CardType = CardTypeDui
			info.MinValue = value[2][0]
		}
	}
	return
}

// 癞子三不带
func isSanBuDaiLaiZi(size int, value valueList, newValue valueList, laiZiCount int) (info CardTypeInfo) {
	if size != 3 {
		return
	}

	switch laiZiCount {
	case 1:
		// 非癞子牌为1对子
		if len(newValue[2]) == 1 {
			info.CardType = CardTypeSanBuDai
			info.MinValue = newValue[2][0]
		}

	case 2:
		// 非癞子牌为1单张
		if len(newValue[1]) == 1 {
			info.CardType = CardTypeSanBuDai
			info.MinValue = newValue[1][0]
		}

	case 3:
		// 癞子牌为1三张
		if len(value[3]) == 1 {
			info.CardType = CardTypeSanBuDai
			info.MinValue = value[3][0]
		}
	}
	return
}

// 癞子三带一
func isSanDaiYiLaiZi(size int, newValue valueList, laiZiCount int) (list []*CardTypeInfo) {
	if size != 4 {
		return
	}

	switch laiZiCount {
	case 1:
		// 非癞子牌为1对子
		if len(newValue[2]) == 1 {
			info := &CardTypeInfo{}
			info.CardType = CardTypeSanDaiYi
			info.MinValue = newValue[2][0]
			list = append(list, info)
		}

	case 2:
		// 非癞子牌为2单张
		if len(newValue[1]) == 2 {
			// 不能是大小王
			for _, v := range newValue[1] {
				if v != 16 && v != 17 {
					info := &CardTypeInfo{}
					info.CardType = CardTypeSanDaiYi
					info.MinValue = newValue[1][0]
					list = append(list, info)
				}
			}
		}
	}
	return
}

// 癞子三带二
func isSanDaiErLaiZi(size int, newValue valueList, laiZiCount int) (list []*CardTypeInfo) {
	if size != 5 {
		return
	}

	switch laiZiCount {
	case 1:
		// 非癞子牌为2对子
		if len(newValue[2]) == 2 {
			for _, v := range newValue[2] {
				info := &CardTypeInfo{}
				info.CardType = CardTypeSanDaiEr
				info.MinValue = v
				list = append(list, info)
			}
		}

	case 2:
		// 非癞子牌为1对子
		if len(newValue[2]) == 1 {
			info := &CardTypeInfo{}
			info.CardType = CardTypeSanDaiEr
			info.MinValue = newValue[2][0]
			list = append(list, info)
		}

	case 3:
		// 非癞子牌为2单张
		if len(newValue[1]) == 2 {
			// 不能是大小王
			for _, v := range newValue[1] {
				if v != 16 && v != 17 {
					info := &CardTypeInfo{}
					info.CardType = CardTypeSanDaiEr
					info.MinValue = v
					list = append(list, info)
				}
			}
		}
	}
	return
}

// 癞子四带单
func isSiDaiDanLaiZi(size int, newValue valueList, laiZiCount int) (list []*CardTypeInfo) {
	if size != 6 {
		return
	}

	switch laiZiCount {
	case 1:
		// 非癞子牌为1三张
		if len(newValue[3]) == 1 {
			info := &CardTypeInfo{}
			info.CardType = CardTypeSiDaiDan
			info.MinValue = newValue[3][0]
			list = append(list, info)
		}

		// 非癞子牌为1四张
		if len(newValue[4]) == 1 {
			info := &CardTypeInfo{}
			info.CardType = CardTypeSiDaiDan
			info.MinValue = newValue[4][0]
			list = append(list, info)
		}

	case 2:
		// 非癞子牌为1三张
		if len(newValue[3]) == 1 {
			info := &CardTypeInfo{}
			info.CardType = CardTypeSiDaiDan
			info.MinValue = newValue[3][0]
			list = append(list, info)
		}

		// 非癞子牌为2对子
		if len(newValue[2]) == 2 {
			for _, v := range newValue[2] {
				info := &CardTypeInfo{}
				info.CardType = CardTypeSiDaiDan
				info.MinValue = v
				list = append(list, info)
			}
		}

		// 非癞子牌为1对子
		if len(newValue[2]) == 1 {
			info := &CardTypeInfo{}
			info.CardType = CardTypeSiDaiDan
			info.MinValue = newValue[2][0]
			list = append(list, info)
		}

	case 3:
		// 非癞子牌为1对子+1单张
		if len(newValue[2]) == 1 {
			info := &CardTypeInfo{}
			info.CardType = CardTypeSiDaiDan
			info.MinValue = newValue[2][0]
			list = append(list, info)

			if newValue[1][0] != 16 && newValue[1][0] != 17 {
				info := &CardTypeInfo{}
				info.CardType = CardTypeSiDaiDan
				info.MinValue = newValue[1][0]
				list = append(list, info)
			}
		}

		// 非癞子牌为3单张
		if len(newValue[1]) == 3 {
			for _, v := range newValue[1] {
				if v != 16 && v != 17 {
					info := &CardTypeInfo{}
					info.CardType = CardTypeSiDaiDan
					info.MinValue = v
					list = append(list, info)
				}
			}
		}
	}
	return
}

// 癞子四带对
func isSiDaiDuiLaiZi(size int, newValue valueList, laiZiCount int) (list []*CardTypeInfo) {
	if size != 8 {
		return
	}

	switch laiZiCount {
	case 1:
		// 非癞子牌为1三张+2对子
		if len(newValue[3]) == 1 && len(newValue[2]) == 2 {
			info := &CardTypeInfo{}
			info.CardType = CardTypeSiDaiDui
			info.MinValue = newValue[3][0]
			list = append(list, info)
		}

		// 非癞子牌为1四张
		if len(newValue[4]) == 1 {
			info := &CardTypeInfo{}
			info.CardType = CardTypeSiDaiDui
			info.MinValue = newValue[4][0]
			list = append(list, info)
		}

	case 2:
		// 非癞子牌为1三张
		if len(newValue[3]) == 1 {
			info := &CardTypeInfo{}
			info.CardType = CardTypeSiDaiDui
			info.MinValue = newValue[3][0]
			list = append(list, info)
		}

		// 非癞子牌为2对子
		if len(newValue[2]) == 2 {
			for _, v := range newValue[2] {
				info := &CardTypeInfo{}
				info.CardType = CardTypeSiDaiDui
				info.MinValue = v
				list = append(list, info)
			}
		}

		// 非癞子牌为1对子
		if len(newValue[2]) == 1 {
			info := &CardTypeInfo{}
			info.CardType = CardTypeSiDaiDui
			info.MinValue = newValue[2][0]
			list = append(list, info)
		}

	case 3:
		// 非癞子牌为1对子+1单张
		if len(newValue[2]) == 1 && len(newValue[1]) == 1 {
			// 单张不能是大小王
			if newValue[1][0] == 16 || newValue[1][0] == 17 {
				return
			}

			info1 := &CardTypeInfo{}
			info1.CardType = CardTypeSiDaiDui
			info1.MinValue = newValue[2][0]

			info2 := &CardTypeInfo{}
			info2.CardType = CardTypeSiDaiDui
			info2.MinValue = newValue[1][0]
			list = append(list, info1, info2)
		}

		// 非癞子牌为3单张
		if len(newValue[1]) == 3 {
			for _, v := range newValue[1] {
				if v != 16 && v != 17 {
					info := &CardTypeInfo{}
					info.CardType = CardTypeSiDaiDui
					info.MinValue = v
					list = append(list, info)
				}
			}
		}
	}
	return
}

// 癞子顺子
func isShunZiLaiZi(size int, newValue valueList, newLine []int, laiZiCount int) (list []*CardTypeInfo) {
	if size < 5 {
		return
	}

	// 非癞子牌只能有单张
	if len(newValue[2]) != 0 || len(newValue[3]) != 0 || len(newValue[4]) != 0 {
		return
	}

	// 不能有2和大小王
	if existTwoAndJoker(newValue) {
		return
	}

	// 获取连牌缺失的值
	gap := getLineGap(newLine)

	if laiZiCount < gap {
		return
	}
	laiZiCount -= gap

	// 补其他牌
	minValue := newLine[0] - laiZiCount
	if minValue < 3 {
		minValue = 3
	}

	maxValue := newLine[len(newLine)-1] + laiZiCount
	if maxValue > 14 {
		maxValue = 14
	}

	valueRange := size - 1

	for i := minValue; i <= maxValue-valueRange; i++ {
		info := &CardTypeInfo{}
		info.CardType = CardTypeShunZi
		info.MinValue = i
		info.MaxValue = i + valueRange
		list = append(list, info)
	}
	return
}

// 癞子连对
func isLianDuiLaiZi(size int, newValue valueList, newLine []int, laiZiCount int) (list []*CardTypeInfo) {
	if size < 6 || size%2 != 0 {
		return
	}

	// 非癞子牌只能有单张和对子
	if len(newValue[3]) != 0 || len(newValue[4]) != 0 {
		return
	}

	// 不能有2和大小王
	if existTwoAndJoker(newValue) {
		return
	}

	// 配单张
	if laiZiCount < len(newValue[1]) {
		return
	}
	laiZiCount -= len(newValue[1])

	// 配对子
	if laiZiCount%2 != 0 {
		return
	}

	// 获取连牌缺失的值
	gap := getLineGap(newLine)

	if laiZiCount < gap*2 {
		return
	}
	laiZiCount -= gap * 2

	// 补其他牌
	minValue := newLine[0] - laiZiCount/2
	if minValue < 3 {
		minValue = 3
	}

	maxValue := newLine[len(newLine)-1] + laiZiCount/2
	if maxValue > 14 {
		maxValue = 14
	}

	valueRange := size/2 - 1

	for i := minValue; i <= maxValue-valueRange; i++ {
		info := &CardTypeInfo{}
		info.CardType = CardTypeLianDui
		info.MinValue = i
		info.MaxValue = i + valueRange
		list = append(list, info)
	}
	return
}

// 癞子飞机不带
func isFeiJiBuDaiLaiZi(size int, newValue valueList, newLine []int, laiZiCount int) (list []*CardTypeInfo) {
	if size < 6 || size%3 != 0 {
		return
	}

	// 非癞子牌不能有四张
	if len(newValue[4]) != 0 {
		return
	}

	// 不能有2和大小王
	if existTwoAndJoker(newValue) {
		return
	}

	// 配单张
	if laiZiCount < len(newValue[1]) {
		return
	}
	laiZiCount -= len(newValue[1]) * 2

	// 配对子
	if laiZiCount < len(newValue[2]) {
		return
	}
	laiZiCount -= len(newValue[2])

	if laiZiCount%3 != 0 {
		return
	}

	// 获取连牌缺失的值
	gap := getLineGap(newLine)

	if laiZiCount < gap*3 {
		return
	}
	laiZiCount -= gap * 3

	// 补其他牌
	minValue := newLine[0] - laiZiCount/3
	if minValue < 3 {
		minValue = 3
	}

	maxValue := newLine[len(newLine)-1] + laiZiCount/2
	if maxValue > 14 {
		maxValue = 14
	}

	valueRange := size/3 - 1

	for i := minValue; i <= maxValue-valueRange; i++ {
		info := &CardTypeInfo{}
		info.CardType = CardTypeFeiJiBuDai
		info.MinValue = i
		info.MaxValue = i + valueRange
		list = append(list, info)
	}
	return
}

// 四软炸
func isRuanZhaDan4(size int, newValue valueList, laiZiCount int) (info CardTypeInfo) {
	if size != 4 {
		return
	}

	switch laiZiCount {
	case 1:
		// 非癞子牌为1三张
		if len(newValue[3]) == 1 {
			info.CardType = CardTypeRuanZhaDan4
			info.MinValue = newValue[3][0]
		}

	case 2:
		// 非癞子牌为1对子
		if len(newValue[2]) == 1 {
			info.CardType = CardTypeRuanZhaDan4
			info.MinValue = newValue[2][0]
		}

	case 3:
		// 非癞子牌为1单张
		if len(newValue[1]) == 1 {
			info.CardType = CardTypeRuanZhaDan4
			info.MinValue = newValue[1][0]
		}
	}
	return
}

// 五软炸
func isRuanZhaDan5(size int, newValue valueList, laiZiCount int) (info CardTypeInfo) {
	if size != 5 {
		return
	}

	switch laiZiCount {
	case 1:
		// 非癞子牌为1四张
		if len(newValue[4]) == 1 {
			info.CardType = CardTypeRuanZhaDan5
			info.MinValue = newValue[4][0]
		}

	case 2:
		// 非癞子牌为1三张
		if len(newValue[3]) == 1 {
			info.CardType = CardTypeRuanZhaDan5
			info.MinValue = newValue[3][0]
		}

	case 3:
		// 非癞子牌为1对子
		if len(newValue[2]) == 1 {
			info.CardType = CardTypeRuanZhaDan5
			info.MinValue = newValue[2][0]
		}

	case 4:
		// 非癞子牌为1单张
		if len(newValue[1]) == 1 {
			info.CardType = CardTypeRuanZhaDan5
			info.MinValue = newValue[1][0]
		}
	}
	return
}

// 六软炸
func isRuanZhaDan6(size int, newValue valueList, laiZiCount int) (info CardTypeInfo) {
	if size != 6 {
		return
	}

	switch laiZiCount {
	case 2:
		// 非癞子牌为1四张
		if len(newValue[4]) == 1 {
			info.CardType = CardTypeRuanZhaDan6
			info.MinValue = newValue[4][0]
		}

	case 3:
		// 非癞子牌为1三张
		if len(newValue[3]) == 1 {
			info.CardType = CardTypeRuanZhaDan6
			info.MinValue = newValue[3][0]
		}

	case 4:
		// 非癞子牌为1对子
		if len(newValue[2]) == 1 {
			info.CardType = CardTypeRuanZhaDan6
			info.MinValue = newValue[2][0]
		}

	case 5:
		// 非癞子牌为1单张
		if len(newValue[1]) == 1 {
			info.CardType = CardTypeRuanZhaDan6
			info.MinValue = newValue[1][0]
		}
	}
	return
}

// 七软炸
func isRuanZhaDan7(size int, newValue valueList, laiZiCount int) (info CardTypeInfo) {
	if size != 7 {
		return
	}

	switch laiZiCount {
	case 3:
		// 非癞子牌为1四张
		if len(newValue[4]) == 1 {
			info.CardType = CardTypeRuanZhaDan7
			info.MinValue = newValue[4][0]
		}

	case 4:
		// 非癞子牌为1三张
		if len(newValue[3]) == 1 {
			info.CardType = CardTypeRuanZhaDan7
			info.MinValue = newValue[3][0]
		}

	case 5:
		// 非癞子牌为1对子
		if len(newValue[2]) == 1 {
			info.CardType = CardTypeRuanZhaDan7
			info.MinValue = newValue[2][0]
		}

	case 6:
		// 非癞子牌为1单张
		if len(newValue[1]) == 1 {
			info.CardType = CardTypeRuanZhaDan7
			info.MinValue = newValue[1][0]
		}
	}
	return
}

// 八软炸
func isRuanZhaDan8(size int, newValue valueList, laiZiCount int) (info CardTypeInfo) {
	if size != 8 {
		return
	}

	switch laiZiCount {
	case 4:
		// 非癞子牌为1四张
		if len(newValue[4]) == 1 {
			info.CardType = CardTypeRuanZhaDan8
			info.MinValue = newValue[4][0]
		}

	case 5:
		// 非癞子牌为1三张
		if len(newValue[3]) == 1 {
			info.CardType = CardTypeRuanZhaDan8
			info.MinValue = newValue[3][0]
		}

	case 6:
		// 非癞子牌为1对子
		if len(newValue[2]) == 1 {
			info.CardType = CardTypeRuanZhaDan8
			info.MinValue = newValue[2][0]
		}

	case 7:
		// 非癞子牌为1单张
		if len(newValue[1]) == 1 {
			info.CardType = CardTypeRuanZhaDan8
			info.MinValue = newValue[1][0]
		}
	}
	return
}

// 九软炸
func isRuanZhaDan9(size int, newValue valueList, laiZiCount int) (info CardTypeInfo) {
	if size != 9 {
		return
	}

	switch laiZiCount {
	case 5:
		// 非癞子牌为1四张
		if len(newValue[4]) == 1 {
			info.CardType = CardTypeRuanZhaDan9
			info.MinValue = newValue[4][0]
		}

	case 6:
		// 非癞子牌为1三张
		if len(newValue[3]) == 1 {
			info.CardType = CardTypeRuanZhaDan9
			info.MinValue = newValue[3][0]
		}

	case 7:
		// 非癞子牌为1对子
		if len(newValue[2]) == 1 {
			info.CardType = CardTypeRuanZhaDan9
			info.MinValue = newValue[2][0]
		}

	case 8:
		// 非癞子牌为1单张
		if len(newValue[1]) == 1 {
			info.CardType = CardTypeRuanZhaDan9
			info.MinValue = newValue[1][0]
		}
	}
	return
}

// 10软炸
func isRuanZhaDan10(size int, newValue valueList, laiZiCount int) (info CardTypeInfo) {
	if size != 10 {
		return
	}

	switch laiZiCount {
	case 6:
		// 非癞子牌为1四张
		if len(newValue[4]) == 1 {
			info.CardType = CardTypeRuanZhaDan10
			info.MinValue = newValue[4][0]
		}

	case 7:
		// 非癞子牌为1三张
		if len(newValue[3]) == 1 {
			info.CardType = CardTypeRuanZhaDan10
			info.MinValue = newValue[3][0]
		}

	case 8:
		// 非癞子牌为1对子
		if len(newValue[2]) == 1 {
			info.CardType = CardTypeRuanZhaDan10
			info.MinValue = newValue[2][0]
		}
	}
	return
}

// 11软炸
func isRuanZhaDan11(size int, newValue valueList, laiZiCount int) (info CardTypeInfo) {
	if size != 11 {
		return
	}

	switch laiZiCount {
	case 7:
		// 非癞子牌为1四张
		if len(newValue[4]) == 1 {
			info.CardType = CardTypeRuanZhaDan11
			info.MinValue = newValue[4][0]
		}

	case 8:
		// 非癞子牌为1三张
		if len(newValue[3]) == 1 {
			info.CardType = CardTypeRuanZhaDan11
			info.MinValue = newValue[3][0]
		}
	}
	return
}

// 十二软炸
func isRuanZhaDan12(size int, newValue valueList, laiZiCount int) (info CardTypeInfo) {
	if size != 12 {
		return
	}

	switch laiZiCount {
	case 8:
		// 非癞子牌为1四张
		if len(newValue[4]) == 1 {
			info.CardType = CardTypeRuanZhaDan12
			info.MinValue = newValue[4][0]
		}
	}
	return
}
