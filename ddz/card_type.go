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

// 获取牌型
func GetCardType(cards []*Card) (list []*CardTypeInfo) {
	newCards := getNoLaiZiCards(cards)       // 无癞子牌数组
	laiZiCount := len(cards) - len(newCards) // 癞子数量

	if len(cards)-len(newCards) == 0 {
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

// 无癞子算法
func analysis(cards []*Card) (list []*CardTypeInfo) {
	// 排序
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Num < cards[j].Num
	})

	// 统计牌数量
	count := cardCount(cards)

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
	if info := isLianDui(cards, count); info.CardType != CardTypeNone {
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
func isLianDui(cards []*Card, count [18]int) (info CardTypeInfo) {
	minValue := getMinValue(count, 2)
	maxValue := getMaxValue(count, 2)
	valueRange := maxValue - minValue + 1

	size := len(cards)
	if size < 6 || size != valueRange*2 {
		return
	}

	exist := true
	for i := minValue; i <= maxValue; i++ {
		if count[i] != 2 {
			exist = false
			break
		}
	}

	if exist {
		info.CardType = CardTypeLianDui
		info.MinValue = minValue
		info.MaxValue = maxValue
		return
	}
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
	// 排序
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Num < cards[j].Num
	})

	sort.Slice(newCards, func(i, j int) bool {
		return newCards[i].Num < newCards[j].Num
	})

	// 统计牌数量
	cardsCount := cardCount(cards)
	newCardsCount := cardCount(newCards)

	// 癞子单
	if info := isDan(cards); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 癞子对
	if info := isDuiLaiZi(cards, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 癞子连对
	if info := isLianDuiLaiZi(cards, cardsCount, newCards, newCardsCount, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 癞子三不带
	if info := isSanBuDaiLaiZi(cards, cardsCount, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 四软炸
	if info := isRuanZhaDan4(cards, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 五软炸
	if info := isRuanZhaDan5(cards, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 六软炸
	if info := isRuanZhaDan6(cards, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 七软炸
	if info := isRuanZhaDan7(cards, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 八软炸
	if info := isRuanZhaDan8(cards, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 九软炸
	if info := isRuanZhaDan9(cards, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 十软炸
	if info := isRuanZhaDan10(cards, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 十一软炸
	if info := isRuanZhaDan11(cards, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 十二软炸
	if info := isRuanZhaDan12(cards, laiZiCount); info.CardType != CardTypeNone {
		list = append(list, &info)
	}
	return
}

// 癞子对
func isDuiLaiZi(cards []*Card, laiZiCount int) (info CardTypeInfo) {
	if len(cards) != 2 {
		return
	}

	switch laiZiCount {
	case 1:
		info.CardType = CardTypeDui
		info.MinValue = int(cards[0].Num)

	case 2:
		return isDui(cards)
	}
	return
}

// 癞子连对
func isLianDuiLaiZi(cards []*Card, cardsCount [18]int, newCards []*Card, newCardsCount [18]int, laiZiCount int) (info CardTypeInfo) {
	size := len(cards)
	if size < 6 {
		return
	}

	if info = isLianDui(cards, cardsCount); info.CardType != CardTypeNone {
		// 连对
		return info
	} else {
		minValue := int(newCards[0].Num)               // 最小值
		maxValue := int(newCards[len(newCards)-1].Num) // 最大值
		valueRange := maxValue - minValue + 1          // 连对数量
		if size != valueRange*2 {
			return
		}

		// 癞子代替需要的牌
		for i := minValue; i <= maxValue; i++ {
			for j := newCardsCount[i]; j < 2; j++ {
				newCards = append(newCards, NewCard(SuitTypeNone, NumType(i), false))
				newCardsCount[i]++
				laiZiCount--
			}
		}

		if laiZiCount != 0 {
			return
		}
		return isLianDui(newCards, newCardsCount)
	}
}

// 癞子三不带
func isSanBuDaiLaiZi(cards []*Card, cardsCount [18]int, laiZiCount int) (info CardTypeInfo) {
	if len(cards) != 3 {
		return
	}

	switch laiZiCount {
	case 1:
		if cards[0].Num == cards[1].Num {
			info.CardType = CardTypeSanBuDai
			info.MinValue = int(cards[0].Num)
		}

	case 2:
		info.CardType = CardTypeSanBuDai
		info.MinValue = int(cards[0].Num)

	case 3:
		return isSanBuDai(cards, cardsCount)
	}
	return
}

// 癞子三带一
func isSanDaiYiLaiZi(cards []*Card, laiZiCount int) (info CardTypeInfo) {
	if len(cards) != 2 {
		return
	}

	switch laiZiCount {
	case 1:
		info.CardType = CardTypeDui
		info.MinValue = int(cards[0].Num)

	case 2:
		return isDui(cards)
	}
	return
}

// 癞子三带二
func isSanDaiErLaiZi(cards []*Card, laiZiCount int) (info CardTypeInfo) {
	if len(cards) != 2 {
		return
	}

	switch laiZiCount {
	case 1:
		info.CardType = CardTypeDui
		info.MinValue = int(cards[0].Num)

	case 2:
		return isDui(cards)
	}
	return
}

// 四软炸
func isRuanZhaDan4(cards []*Card, laiZiCount int) (info CardTypeInfo) {
	if len(cards) != 4 {
		return
	}

	switch laiZiCount {
	case 1:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num {
			info.CardType = CardTypeRuanZhaDan4
			info.MinValue = int(cards[0].Num)
		}

	case 2:
		if cards[0].Num == cards[1].Num {
			info.CardType = CardTypeRuanZhaDan4
			info.MinValue = int(cards[0].Num)
		}

	case 3:
		info.CardType = CardTypeRuanZhaDan4
		info.MinValue = int(cards[0].Num)
	}
	return
}

// 五软炸
func isRuanZhaDan5(cards []*Card, laiZiCount int) (info CardTypeInfo) {
	if len(cards) != 5 {
		return
	}

	switch laiZiCount {
	case 1:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num && cards[0].Num == cards[3].Num {
			info.CardType = CardTypeRuanZhaDan5
			info.MinValue = int(cards[0].Num)
		}

	case 2:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num {
			info.CardType = CardTypeRuanZhaDan5
			info.MinValue = int(cards[0].Num)
		}

	case 3:
		if cards[0].Num == cards[1].Num {
			info.CardType = CardTypeRuanZhaDan5
			info.MinValue = int(cards[0].Num)
		}

	case 4:
		info.CardType = CardTypeRuanZhaDan5
		info.MinValue = int(cards[0].Num)
	}
	return
}

// 六软炸
func isRuanZhaDan6(cards []*Card, laiZiCount int) (info CardTypeInfo) {
	if len(cards) != 6 {
		return
	}

	switch laiZiCount {
	case 2:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num && cards[0].Num == cards[3].Num {
			info.CardType = CardTypeRuanZhaDan6
			info.MinValue = int(cards[0].Num)
		}

	case 3:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num {
			info.CardType = CardTypeRuanZhaDan6
			info.MinValue = int(cards[0].Num)
		}

	case 4:
		if cards[0].Num == cards[1].Num {
			info.CardType = CardTypeRuanZhaDan6
			info.MinValue = int(cards[0].Num)
		}

	case 5:
		info.CardType = CardTypeRuanZhaDan6
		info.MinValue = int(cards[0].Num)
	}
	return
}

// 七软炸
func isRuanZhaDan7(cards []*Card, laiZiCount int) (info CardTypeInfo) {
	if len(cards) != 7 {
		return
	}

	switch laiZiCount {
	case 3:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num && cards[0].Num == cards[3].Num {
			info.CardType = CardTypeRuanZhaDan7
			info.MinValue = int(cards[0].Num)
		}

	case 4:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num {
			info.CardType = CardTypeRuanZhaDan7
			info.MinValue = int(cards[0].Num)
		}

	case 5:
		if cards[0].Num == cards[1].Num {
			info.CardType = CardTypeRuanZhaDan7
			info.MinValue = int(cards[0].Num)
		}

	case 6:
		info.CardType = CardTypeRuanZhaDan7
		info.MinValue = int(cards[0].Num)
	}
	return
}

// 八软炸
func isRuanZhaDan8(cards []*Card, laiZiCount int) (info CardTypeInfo) {
	if len(cards) != 8 {
		return
	}

	switch laiZiCount {
	case 4:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num && cards[0].Num == cards[3].Num {
			info.CardType = CardTypeRuanZhaDan8
			info.MinValue = int(cards[0].Num)
		}

	case 5:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num {
			info.CardType = CardTypeRuanZhaDan8
			info.MinValue = int(cards[0].Num)
		}

	case 6:
		if cards[0].Num == cards[1].Num {
			info.CardType = CardTypeRuanZhaDan8
			info.MinValue = int(cards[0].Num)
		}

	case 7:
		info.CardType = CardTypeRuanZhaDan8
		info.MinValue = int(cards[0].Num)
	}
	return
}

// 九软炸
func isRuanZhaDan9(cards []*Card, laiZiCount int) (info CardTypeInfo) {
	if len(cards) != 9 {
		return
	}

	switch laiZiCount {
	case 5:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num && cards[0].Num == cards[3].Num {
			info.CardType = CardTypeRuanZhaDan9
			info.MinValue = int(cards[0].Num)
		}

	case 6:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num {
			info.CardType = CardTypeRuanZhaDan9
			info.MinValue = int(cards[0].Num)
		}

	case 7:
		if cards[0].Num == cards[1].Num {
			info.CardType = CardTypeRuanZhaDan9
			info.MinValue = int(cards[0].Num)
		}

	case 8:
		info.CardType = CardTypeRuanZhaDan9
		info.MinValue = int(cards[0].Num)
	}
	return
}

// 10软炸
func isRuanZhaDan10(cards []*Card, laiZiCount int) (info CardTypeInfo) {
	if len(cards) != 10 {
		return
	}

	switch laiZiCount {
	case 6:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num && cards[0].Num == cards[3].Num {
			info.CardType = CardTypeRuanZhaDan10
			info.MinValue = int(cards[0].Num)
		}

	case 7:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num {
			info.CardType = CardTypeRuanZhaDan10
			info.MinValue = int(cards[0].Num)
		}

	case 8:
		if cards[0].Num == cards[1].Num {
			info.CardType = CardTypeRuanZhaDan10
			info.MinValue = int(cards[0].Num)
		}
	}
	return
}

// 11软炸
func isRuanZhaDan11(cards []*Card, laiZiCount int) (info CardTypeInfo) {
	if len(cards) != 11 {
		return
	}

	switch laiZiCount {
	case 7:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num && cards[0].Num == cards[3].Num {
			info.CardType = CardTypeRuanZhaDan11
			info.MinValue = int(cards[0].Num)
		}

	case 8:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num {
			info.CardType = CardTypeRuanZhaDan11
			info.MinValue = int(cards[0].Num)
		}
	}
	return
}

// 十二软炸
func isRuanZhaDan12(cards []*Card, laiZiCount int) (info CardTypeInfo) {
	if len(cards) != 12 {
		return
	}

	switch laiZiCount {
	case 8:
		if cards[0].Num == cards[1].Num && cards[0].Num == cards[2].Num && cards[0].Num == cards[3].Num {
			info.CardType = CardTypeRuanZhaDan12
			info.MinValue = int(cards[0].Num)
		}
	}
	return
}
