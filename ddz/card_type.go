package ddz

import "sort"

// 获取牌型
func GetCardType(cards []*Card, laiZiNums ...NumType) (list []*CardTypeInfo) {
	if len(laiZiNums) == 0 {
		// 无癞子算法
		return analysis(cards)
	}

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

	if len(laiZiCards) == 0 {
		// 无癞子算法
		return analysis(cards)
	} else {
		// 癞子算法
		return analysisLaiZi(cards, normalCards, laiZiCards)
	}
}

// 无癞子算法
func analysis(cards []*Card) (list []*CardTypeInfo) {
	size := len(cards)
	dictCards := convertToMap(cards)
	_, value, line := getCountValueLine(dictCards)

	switch size {
	case 1:
		// 单
		if info := isDan(size, value); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

	case 2:
		// 对
		if info := isDui(size, value); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

		// 火箭
		if info := isHuoJian(size, value); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

	case 3:
		// 三不带
		if info := isSanBuDai(size, value); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

	case 4:
		// 三带一
		if info := isSanDaiYi(size, value); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

		// 炸弹
		if info := isZhaDan(size, value); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

	case 5:
		// 三带二
		if info := isSanDaiEr(size, value); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

	case 6:
		// 四带单
		if info := isSiDaiDan(size, value); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}

	case 8:
		// 四带对
		if info := isSiDaiDui(size, value); info.CardType != CardTypeNone {
			list = append(list, &info)
			return
		}
	}

	// 顺子
	if info := isShunZi(size, value, line); info.CardType != CardTypeNone {
		list = append(list, &info)
		return
	}

	// 连对
	if info := isLianDui(size, value, line); info.CardType != CardTypeNone {
		list = append(list, &info)
		return
	}

	// 飞机不带
	if info := isFeiJiBuDai(size, value, line); info.CardType != CardTypeNone {
		list = append(list, &info)
		return
	}

	// 飞机带一
	if info := isFeiJiDaiYi(size, value, line); info.CardType != CardTypeNone {
		list = append(list, &info)
		return
	}

	// 飞机带二
	if info := isFeiJiDaiEr(size, value, line); info.CardType != CardTypeNone {
		list = append(list, &info)
		return
	}

	// 连炸
	if info := isLianZha(size, value, line); info.CardType != CardTypeNone {
		list = append(list, &info)
		return
	}
	return
}

// 单
func isDan(size int, value valueList) (info CardTypeInfo) {
	if size != 1 {
		return
	}

	info.CardType = CardTypeDan
	info.MinValue = value[1][0]
	return
}

// 对
func isDui(size int, value valueList) (info CardTypeInfo) {
	if size != 2 {
		return
	}

	// 牌值相同
	if len(value[2]) == 1 {
		info.CardType = CardTypeDui
		info.MinValue = value[2][0]
	}
	return
}

// 三不带
func isSanBuDai(size int, value valueList) (info CardTypeInfo) {
	if size != 3 {
		return
	}

	if len(value[3]) == 1 {
		info.CardType = CardTypeSanBuDai
		info.MinValue = value[3][0]
	}
	return
}

// 三带一
func isSanDaiYi(size int, value valueList) (info CardTypeInfo) {
	if size != 4 {
		return
	}

	if len(value[3]) == 1 && len(value[1]) == 1 {
		info.CardType = CardTypeSanDaiYi
		info.MinValue = value[3][0]
	}
	return
}

// 三带二
func isSanDaiEr(size int, value valueList) (info CardTypeInfo) {
	if size != 5 {
		return
	}

	if len(value[3]) == 1 && len(value[2]) == 1 {
		info.CardType = CardTypeSanDaiEr
		info.MinValue = value[3][0]
	}
	return
}

// 四带单(炸弹带两张单)
func isSiDaiDan(size int, value valueList) (info CardTypeInfo) {
	if size != 6 {
		return
	}

	if len(value[4]) == 1 && (len(value[1]) == 2 || len(value[2]) == 1) {
		info.CardType = CardTypeSiDaiDan
		info.MinValue = value[4][0]
	}
	return
}

// 四带对(炸弹带两对)
func isSiDaiDui(size int, value valueList) (info CardTypeInfo) {
	if size != 8 {
		return
	}

	if len(value[4]) == 1 && len(value[2]) == 2 {
		info.CardType = CardTypeSiDaiDui
		info.MinValue = value[4][0]
	}
	return
}

// 顺子
func isShunZi(size int, value valueList, line lineList) (info CardTypeInfo) {
	if size < 5 || size > 12 {
		return
	}

	// 只能有单张
	if len(value[2]) != 0 || len(value[3]) != 0 || len(value[4]) != 0 {
		return
	}

	// 不能有2和大小王
	for _, v := range line {
		if isJokerAndTwo(v) {
			return
		}
	}

	// 必须连续
	valueRange := line[len(line)-1] - line[0] + 1
	if size != valueRange {
		return
	}

	info.CardType = CardTypeShunZi
	info.MinValue = line[0]
	info.MaxValue = line[len(line)-1]
	return
}

// 连对
func isLianDui(size int, value valueList, line lineList) (info CardTypeInfo) {
	if size < 6 || size%2 != 0 {
		return
	}

	// 只能有对子
	if len(value[1]) != 0 || len(value[3]) != 0 || len(value[4]) != 0 {
		return
	}

	// 不能有2和大小王
	for _, v := range line {
		if isJokerAndTwo(v) {
			return
		}
	}

	// 必须连续
	valueRange := line[len(line)-1] - line[0] + 1
	if size != valueRange*2 {
		return
	}

	info.CardType = CardTypeLianDui
	info.MinValue = line[0]
	info.MaxValue = line[len(line)-1]
	return
}

// 飞机不带
func isFeiJiBuDai(size int, value valueList, line lineList) (info CardTypeInfo) {
	if size < 6 || size%3 != 0 {
		return
	}

	// 只能有三张
	if len(value[1]) != 0 || len(value[2]) != 0 || len(value[4]) != 0 {
		return
	}

	// 不能有2和大小王
	for _, v := range line {
		if isJokerAndTwo(v) {
			return
		}
	}

	// 必须连续
	valueRange := line[len(line)-1] - line[0] + 1
	if size != valueRange*3 {
		return
	}

	info.CardType = CardTypeFeiJiBuDai
	info.MinValue = line[0]
	info.MaxValue = line[len(line)-1]
	return
}

// 飞机带一
func isFeiJiDaiYi(size int, value valueList, line lineList) (info CardTypeInfo) {
	if size < 8 || size%4 != 0 {
		return
	}

	// 连炸不是飞机带一
	info = isLianZha(size, value, line)
	if info.CardType != CardTypeNone {
		return
	}

	// 统计所有的三张和四张
	var normalLine []int
	for _, v := range value[3] {
		if isJokerAndTwo(v) {
			continue
		}
		normalLine = append(normalLine, v)
	}

	for _, v := range value[4] {
		if isJokerAndTwo(v) {
			continue
		}
		normalLine = append(normalLine, v)
	}

	if size != len(normalLine)*4 {
		return
	}

	// 排序
	sort.Slice(normalLine, func(i, j int) bool {
		return normalLine[i] < normalLine[j]
	})

	// 必须连续
	valueRange := normalLine[len(normalLine)-1] - normalLine[0] + 1
	if size != valueRange*4 {
		return
	}

	info.CardType = CardTypeFeiJiDaiYi
	info.MinValue = normalLine[0]
	info.MaxValue = normalLine[len(normalLine)-1]
	return
}

// 飞机带二
func isFeiJiDaiEr(size int, value valueList, line []int) (info CardTypeInfo) {
	if size < 10 || size%5 != 0 {
		return
	}

	// 不能有单张和四张
	if len(value[1]) != 0 || len(value[4]) != 0 {
		return
	}

	// 三张不能有2
	for _, v := range value[3] {
		if isJokerAndTwo(v) {
			return
		}
	}

	// 必须连续
	valueRange := value[3][len(value[3])-1] - value[3][0] + 1
	if size != valueRange*5 {
		return
	}

	info.CardType = CardTypeFeiJiDaiEr
	info.MinValue = value[3][0]
	info.MaxValue = value[3][len(value[3])-1]
	return
}

// 炸弹
func isZhaDan(size int, value valueList) (info CardTypeInfo) {
	if size != 4 {
		return
	}

	if len(value[4]) == 1 {
		info.CardType = CardTypeZhaDan
		info.MinValue = value[4][0]
	}
	return
}

// 火箭
func isHuoJian(size int, value valueList) (info CardTypeInfo) {
	if size != 2 {
		return
	}

	if huoJian(value) {
		info.CardType = CardTypeHuoJian
	}
	return
}

// 连炸
func isLianZha(size int, value valueList, line lineList) (info CardTypeInfo) {
	if size < 8 || size%4 != 0 {
		return
	}

	// 只能有四张
	if len(value[1]) != 0 || len(value[2]) != 0 || len(value[3]) != 0 {
		return
	}

	// 不能有2和大小王
	for _, v := range line {
		if isJokerAndTwo(v) {
			return
		}
	}

	// 必须连续
	valueRange := line[len(line)-1] - line[0] + 1
	if valueRange != size/4 {
		return
	}

	info.CardType = CardTypeLianZha
	info.MinValue = line[0]
	info.MaxValue = line[len(line)-1]
	return
}

// 癞子算法
func analysisLaiZi(cards []*Card, normalCards []*Card, laiZiCards []*Card) (list []*CardTypeInfo) {
	size := len(cards)
	laiZiSize := len(laiZiCards)

	dictCards := convertToMap(cards)
	_, value, _ := getCountValueLine(dictCards)

	normalDictCards := convertToMap(normalCards)
	normalCount, normalValue, normalLine := getCountValueLine(normalDictCards)

	// 癞子单
	if info := isDanLaiZi(size, value); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 癞子对
	if info := isDuiLaiZi(size, value, normalLine); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 癞子三不带
	if info := isSanBuDaiLaiZi(size, value, normalLine); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 癞子三带一
	if infoList := isSanDaiYiLaiZi(size, normalCount, normalLine, laiZiSize); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子三带二
	if infoList := isSanDaiErLaiZi(size, normalCount, normalLine, laiZiSize); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子四带单
	if infoList := isSiDaiDanLaiZi(size, normalCount, normalLine, laiZiSize); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子四带对
	if infoList := isSiDaiDuiLaiZi(size, normalCount, normalLine, laiZiSize); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子顺子
	if infoList := isLianLaiZi(CardTypeShunZi, size, normalValue, normalCount, normalLine, laiZiSize); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子连对
	if infoList := isLianLaiZi(CardTypeLianDui, size, normalValue, normalCount, normalLine, laiZiSize); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子飞机不带
	if infoList := isLianLaiZi(CardTypeFeiJiBuDai, size, normalValue, normalCount, normalLine, laiZiSize); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子飞机带一
	if infoList := isFeiJiDaiYiLaiZi(size, normalCount, normalLine, laiZiSize); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 癞子飞机带二
	if infoList := isFeiJiDaiErLaiZi(size, normalCount, normalLine, laiZiSize); len(infoList) != 0 {
		list = append(list, infoList...)
	}

	// 四纯癞子炸
	if info := isChunLaiZiZhaDan(size, laiZiSize, laiZiCards); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 软炸
	if info := isRuanZhaDan(size, normalLine); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 癞子炸
	if info := isLaiZiZhaDan(size, laiZiSize, laiZiCards); info.CardType != CardTypeNone {
		list = append(list, &info)
	}

	// 软连炸
	if info := isRuanLianZha(size, normalCount, normalLine, laiZiSize); info.CardType != CardTypeNone {
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
func isDuiLaiZi(size int, value valueList, normalLine lineList) (info CardTypeInfo) {
	if size != 2 {
		return
	}

	// 非癞子牌有多少种牌值
	switch len(normalLine) {
	case 0:
		if len(value[2]) != 1 {
			return
		}

		info.CardType = CardTypeDui
		info.MinValue = value[2][0]

	case 1:
		if isJoker(normalLine[0]) {
			return
		}

		info.CardType = CardTypeDui
		info.MinValue = normalLine[0]
	}
	return
}

// 癞子三不带
func isSanBuDaiLaiZi(size int, value valueList, normalLine lineList) (info CardTypeInfo) {
	if size != 3 {
		return
	}

	// 非癞子牌有多少种牌值
	switch len(normalLine) {
	case 0:
		if len(value[3]) != 1 {
			return
		}

		info.CardType = CardTypeSanBuDai
		info.MinValue = value[3][0]

	case 1:
		if isJoker(normalLine[0]) {
			return
		}

		info.CardType = CardTypeSanBuDai
		info.MinValue = normalLine[0]
	}
	return
}

// 癞子三带一
func isSanDaiYiLaiZi(size int, normalCount countList, normalLine lineList, laiZiSize int) (list []*CardTypeInfo) {
	if size != 4 {
		return
	}

	var min CardTypeInfo // 最小值
	var max CardTypeInfo // 最大值

	// 非癞子牌有多少种牌值
	switch len(normalLine) {
	case 2:
		var line []int
		for _, v := range normalLine {
			if isJoker(v) {
				continue
			}
			line = append(line, v)
		}

		if len(line) == 0 {
			return
		}

		min = minSanDaiYiLaiZi(normalCount, line, laiZiSize)
		max = maxSanDaiYiLaiZi(normalCount, line, laiZiSize)

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
func minSanDaiYiLaiZi(normalCount countList, line lineList, laiZiSize int) (info CardTypeInfo) {
	for i := 0; i < len(line); i++ {
		info = sanDaiYiLaiZi(line[i], normalCount, laiZiSize)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 最大癞子三带一
func maxSanDaiYiLaiZi(normalCount countList, line lineList, laiZiSize int) (info CardTypeInfo) {
	for i := len(line) - 1; i >= 0; i-- {
		info = sanDaiYiLaiZi(line[i], normalCount, laiZiSize)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 癞子三带一结果
func sanDaiYiLaiZi(value int, normalCount countList, laiZiSize int) (info CardTypeInfo) {
	var laiZi int

	// 补三张
	switch normalCount[value] {
	case 0:
		laiZi += 3
	case 1:
		laiZi += 2
	case 2:
		laiZi += 1
	}

	if laiZi > laiZiSize {
		return
	}

	info.CardType = CardTypeSanDaiYi
	info.MinValue = value
	return
}

// 癞子三带二
func isSanDaiErLaiZi(size int, normalCount countList, normalLine lineList, laiZiSize int) (list []*CardTypeInfo) {
	if size != 5 {
		return
	}

	var min CardTypeInfo // 最小值
	var max CardTypeInfo // 最大值

	// 非癞子牌有多少种牌值
	switch len(normalLine) {
	case 2:
		for _, v := range normalLine {
			if isJoker(v) {
				return
			}
		}

		min = minSanDaiErLaiZi(normalCount, normalLine, laiZiSize)
		max = maxSanDaiErLaiZi(normalCount, normalLine, laiZiSize)

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
func minSanDaiErLaiZi(normalCount countList, line lineList, laiZiSize int) (info CardTypeInfo) {
	for i := 0; i < len(line); i++ {
		info = sanDaiErLaiZi(line[i], normalCount, line, laiZiSize)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 最大癞子三带二
func maxSanDaiErLaiZi(normalCount countList, line lineList, laiZiSize int) (info CardTypeInfo) {
	for i := len(line) - 1; i >= 0; i-- {
		info = sanDaiErLaiZi(line[i], normalCount, line, laiZiSize)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 癞子三带二结果
func sanDaiErLaiZi(value int, normalCount countList, line lineList, laiZiSize int) (info CardTypeInfo) {
	var laiZi int

	// 补三张
	switch normalCount[value] {
	case 0:
		laiZi += 3
	case 1:
		laiZi += 2
	case 2:
		laiZi += 1
	}

	if laiZi > laiZiSize {
		return
	}

	// 补对子
	for _, v := range line {
		if v == value {
			continue
		}

		switch normalCount[v] {
		case 0:
			laiZi += 2
		case 1:
			laiZi += 1
		}

		if laiZi > laiZiSize {
			return
		}
	}

	info.CardType = CardTypeSanDaiEr
	info.MinValue = value
	return
}

// 癞子四带单
func isSiDaiDanLaiZi(size int, normalCount countList, normalLine lineList, laiZiSize int) (list []*CardTypeInfo) {
	if size != 6 {
		return
	}

	var min CardTypeInfo // 最小值
	var max CardTypeInfo // 最大值

	// 非癞子牌有多少种牌值
	switch len(normalLine) {
	case 0, 1:
		return

	default:
		var line []int
		for _, v := range normalLine {
			if isJoker(v) {
				continue
			}
			line = append(line, v)
		}

		min = minSiDaiDanLaiZi(normalCount, line, laiZiSize)
		max = maxSiDaiDanLaiZi(normalCount, line, laiZiSize)
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
func minSiDaiDanLaiZi(normalCount countList, line lineList, laiZiSize int) (info CardTypeInfo) {
	// 最小值为3
	info = siDaiDanLaiZi(NumTypeThree, normalCount, laiZiSize)
	if info.CardType != CardTypeNone {
		return
	}

	for i := 0; i < len(line); i++ {
		if line[i] == NumTypeThree {
			continue
		}

		info = siDaiDanLaiZi(line[i], normalCount, laiZiSize)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 最大癞子四带单
func maxSiDaiDanLaiZi(normalCount countList, line lineList, laiZiSize int) (info CardTypeInfo) {
	// 最大值为2
	info = siDaiDanLaiZi(NumTypeTwo, normalCount, laiZiSize)
	if info.CardType != CardTypeNone {
		return
	}

	for i := len(line) - 1; i >= 0; i-- {
		if line[i] == NumTypeTwo {
			continue
		}

		info = siDaiDanLaiZi(line[i], normalCount, laiZiSize)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 癞子四带单结果
func siDaiDanLaiZi(value int, normalCount countList, laiZiSize int) (info CardTypeInfo) {
	var laiZi int

	// 补四张
	switch normalCount[value] {
	case 0:
		laiZi += 4
	case 1:
		laiZi += 3
	case 2:
		laiZi += 2
	case 3:
		laiZi += 1
	}

	if laiZi > laiZiSize {
		return
	}

	info.CardType = CardTypeSiDaiDan
	info.MinValue = value
	return
}

// 癞子四带对
func isSiDaiDuiLaiZi(size int, normalCount countList, normalLine lineList, laiZiSize int) (list []*CardTypeInfo) {
	if size != 8 {
		return
	}

	var min CardTypeInfo // 最小值
	var max CardTypeInfo // 最大值

	// 非癞子牌有多少种牌值
	switch len(normalLine) {
	case 0, 1:
		return

	default:
		for _, v := range normalLine {
			if isJoker(v) {
				return
			}
		}

		min = minSiDaiDuiLaiZi(normalCount, normalLine, laiZiSize)
		max = maxSiDaiDuiLaiZi(normalCount, normalLine, laiZiSize)
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
func minSiDaiDuiLaiZi(normalCount countList, line lineList, laiZiSize int) (info CardTypeInfo) {
	// 最小值为3
	info = siDaiDuiLaiZi(NumTypeThree, normalCount, line, laiZiSize)
	if info.CardType != CardTypeNone {
		return
	}

	for i := 0; i < len(line); i++ {
		if line[i] == NumTypeThree {
			continue
		}

		info = siDaiDuiLaiZi(line[i], normalCount, line, laiZiSize)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 最大癞子四带对
func maxSiDaiDuiLaiZi(normalCount countList, line lineList, laiZiSize int) (info CardTypeInfo) {
	// 最大值为2
	info = siDaiDuiLaiZi(NumTypeTwo, normalCount, line, laiZiSize)
	if info.CardType != CardTypeNone {
		return
	}

	for i := len(line) - 1; i >= 0; i-- {
		if line[i] == NumTypeTwo {
			continue
		}

		info = siDaiDuiLaiZi(line[i], normalCount, line, laiZiSize)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 癞子四带对结果
func siDaiDuiLaiZi(value int, normalCount countList, line lineList, laiZiSize int) (info CardTypeInfo) {
	var laiZi int

	// 补四张
	switch normalCount[value] {
	case 0:
		laiZi += 4
	case 1:
		laiZi += 3
	case 2:
		laiZi += 2
	case 3:
		laiZi += 1
	}

	if laiZi > laiZiSize {
		return
	}

	// 补对子
	for _, v := range line {
		if v == value {
			continue
		}

		switch normalCount[v] {
		case 0:
			laiZi += 2
		case 1:
			laiZi += 1
		}

		if laiZi > laiZiSize {
			return
		}
	}

	info.CardType = CardTypeSiDaiDui
	info.MinValue = value
	return
}

// 癞子连牌(顺子,连对,飞机不带)
func isLianLaiZi(cardType CardType, size int, normalValue valueList, normalCount countList, normalLine lineList, laiZiSize int) (list []*CardTypeInfo) {
	var min CardTypeInfo // 最小值
	var max CardTypeInfo // 最大值

	switch cardType {
	case CardTypeShunZi:
		// 顺子
		if size < 5 || size > 12 {
			return
		}

		// 非癞子牌只能有单张
		if len(normalValue[2]) != 0 || len(normalValue[3]) != 0 || len(normalValue[4]) != 0 {
			return
		}

	case CardTypeLianDui:
		// 连对
		if size < 6 || size%2 != 0 {
			return
		}

		// 非癞子牌只能有单张和对子
		if len(normalValue[3]) != 0 || len(normalValue[4]) != 0 {
			return
		}

	case CardTypeFeiJiBuDai:
		// 飞机不带
		if size < 6 || size%3 != 0 {
			return
		}

		// 非癞子牌只能有单张和对子和三张
		if len(normalValue[4]) != 0 {
			return
		}
	}

	// 非癞子牌有多少种牌值
	switch len(normalLine) {
	case 0, 1:
		return

	default:
		for _, v := range normalLine {
			if isJokerAndTwo(v) {
				return
			}
		}

		// 最小值
		min = lianLaiZi(ResultTypeMin, cardType, normalCount, normalLine, laiZiSize)

		// 最大值
		max = lianLaiZi(ResultTypeMax, cardType, normalCount, normalLine, laiZiSize)
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
func lianLaiZi(resultType ResultType, cardType CardType, normalCount countList, line lineList, laiZiSize int) (info CardTypeInfo) {
	var laiZi int

	// 补牌
	for i := line[0]; i <= line[len(line)-1]; i++ {
		switch normalCount[i] {
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

	if laiZi > laiZiSize {
		return
	}

	// 处理多余的癞子
	var laiZiGroup int
	limitLaiZi := laiZiSize - laiZi
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
func isFeiJiDaiYiLaiZi(size int, normalCount countList, normalLine lineList, laiZiSize int) (list []*CardTypeInfo) {
	if size < 8 || size%4 != 0 {
		return
	}

	var min CardTypeInfo   // 最小值
	var max CardTypeInfo   // 最大值
	valueRange := size / 4 // 几飞机

	switch len(normalLine) {
	case 0, 1:
		return

	default:
		var line []int // 去除2和大小王的值
		for _, v := range normalLine {
			if isJokerAndTwo(v) {
				continue
			}
			line = append(line, v)
		}

		// 最小值
		min = minFeiJiDaiYiLaiZi(normalCount, line, valueRange, laiZiSize)

		// 最大值
		max = maxFeiJiDaiYiLaiZi(normalCount, line, valueRange, laiZiSize)
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
func minFeiJiDaiYiLaiZi(normalCount countList, line lineList, valueRange int, laiZiSize int) (info CardTypeInfo) {
	// 最小值为3的飞机
	info = feiJiDaiYiLaiZi(normalCount, NumTypeThree, NumTypeThree+valueRange-1, laiZiSize)
	if info.CardType != CardTypeNone {
		return
	}

	// 最小值不为3的飞机
	for i := 0; i < len(line); i++ {
		min := line[i] - valueRange + 1
		if min <= NumTypeThree {
			continue
		}

		info = feiJiDaiYiLaiZi(normalCount, min, line[i], laiZiSize)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 最大癞子飞机带一
func maxFeiJiDaiYiLaiZi(normalCount countList, line lineList, valueRange int, laiZiSize int) (info CardTypeInfo) {
	// 最大值为A的飞机
	info = feiJiDaiYiLaiZi(normalCount, NumTypeAce-valueRange+1, NumTypeAce, laiZiSize)
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
			info = feiJiDaiYiLaiZi(normalCount, NumTypeThree, NumTypeThree+valueRange-1, laiZiSize)
			if info.CardType != CardTypeNone {
				return
			}
		} else {
			info = feiJiDaiYiLaiZi(normalCount, min, line[i], laiZiSize)
			if info.CardType != CardTypeNone {
				return
			}
		}
	}
	return
}

// 癞子飞机带一结果
func feiJiDaiYiLaiZi(normalCount countList, min int, max int, laiZiSize int) (info CardTypeInfo) {
	var laiZi int

	// 补飞机
	for i := min; i <= max; i++ {
		switch normalCount[i] {
		case 0:
			laiZi += 3
		case 1:
			laiZi += 2
		case 2:
			laiZi += 1
		}
	}

	if laiZi > laiZiSize {
		return
	}
	info.CardType = CardTypeFeiJiDaiYi
	info.MinValue = min
	info.MaxValue = max
	return
}

// 癞子飞机带二
func isFeiJiDaiErLaiZi(size int, normalCount countList, normalLine lineList, laiZiSize int) (list []*CardTypeInfo) {
	if size < 10 || size%5 != 0 {
		return
	}

	var min CardTypeInfo   // 最小值
	var max CardTypeInfo   // 最大值
	valueRange := size / 5 // 几飞机

	switch len(normalLine) {
	case 0, 1:
		return

	default:
		// 不能有大小王
		for _, v := range normalLine {
			if isJoker(v) {
				return
			}
		}

		var line []int // 去除2
		for _, v := range normalLine {
			if isJokerAndTwo(v) {
				continue
			}
			line = append(line, v)
		}

		// 最小值
		min = minFeiJiDaiErLaiZi(normalCount, line, normalLine, valueRange, laiZiSize)

		// 最大值
		max = maxFeiJiDaiErLaiZi(normalCount, line, normalLine, valueRange, laiZiSize)
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
func minFeiJiDaiErLaiZi(normalCount countList, line lineList, normalLine lineList, valueRange int, laiZiSize int) (info CardTypeInfo) {
	// 最小值为3的飞机
	info = feiJiDaiErLaiZi(normalCount, normalLine, NumTypeThree, NumTypeThree+valueRange-1, laiZiSize)
	if info.CardType != CardTypeNone {
		return
	}

	// 最小值不为3的飞机
	for i := 0; i < len(line); i++ {
		min := line[i] - valueRange + 1
		if min <= NumTypeThree {
			continue
		}

		info = feiJiDaiErLaiZi(normalCount, normalLine, min, line[i], laiZiSize)
		if info.CardType != CardTypeNone {
			return
		}
	}
	return
}

// 最大癞子飞机带二
func maxFeiJiDaiErLaiZi(normalCount countList, line lineList, normalLine lineList, valueRange int, laiZiSize int) (info CardTypeInfo) {
	// 最大值为A的飞机
	info = feiJiDaiErLaiZi(normalCount, normalLine, NumTypeAce-valueRange+1, NumTypeAce, laiZiSize)
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
			info = feiJiDaiErLaiZi(normalCount, normalLine, NumTypeThree, NumTypeThree+valueRange-1, laiZiSize)
			if info.CardType != CardTypeNone {
				return
			}
		} else {
			info = feiJiDaiErLaiZi(normalCount, normalLine, min, line[i], laiZiSize)
			if info.CardType != CardTypeNone {
				return
			}
		}
	}
	return
}

// 癞子飞机带二结果
func feiJiDaiErLaiZi(normalCount countList, normalLine lineList, min int, max int, laiZiSize int) (info CardTypeInfo) {
	var laiZi int

	// 补飞机
	for i := min; i <= max; i++ {
		switch normalCount[i] {
		case 0:
			laiZi += 3
		case 1:
			laiZi += 2
		case 2:
			laiZi += 1
		}

		if laiZi > laiZiSize {
			return
		}
	}

	// 补对子
	for _, v := range normalLine {
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

		switch normalCount[v] {
		case 0:
			laiZi += 2
		case 1:
			laiZi += 1
		}

		if laiZi > laiZiSize {
			return
		}
	}

	info.CardType = CardTypeFeiJiDaiEr
	info.MinValue = min
	info.MaxValue = max
	return
}

// 软炸
func isRuanZhaDan(size int, normalLine lineList) (info CardTypeInfo) {
	if len(normalLine) != 1 {
		return
	}

	if isJoker(normalLine[0]) {
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

	info.MinValue = normalLine[0]
	return
}

// 四纯癞子炸
func isChunLaiZiZhaDan(size int, laiZiSize int, laiZiCards []*Card) (info CardTypeInfo) {
	if size != 4 || laiZiSize != 4 {
		return
	}

	laiZiDictCards := convertToMap(laiZiCards)
	_, _, laiZiLine := getCountValueLine(laiZiDictCards)

	if len(laiZiLine) == 1 {
		info.CardType = CardTypeChunLaiZiZhaDan
		info.MinValue = laiZiLine[0]
	}
	return
}

// 癞子炸
func isLaiZiZhaDan(size int, laiZiSize int, laiZiCards []*Card) (info CardTypeInfo) {
	if size != laiZiSize {
		return
	}

	switch size {
	case 4:
		laiZiDictCards := convertToMap(laiZiCards)
		_, _, laiZiLine := getCountValueLine(laiZiDictCards)

		if len(laiZiLine) != 1 {
			info.CardType = CardTypeLaiZiZhaDan4
		}

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
func isRuanLianZha(size int, normalCount countList, normalLine lineList, laiZiSize int) (info CardTypeInfo) {
	if size < 8 || size%4 != 0 {
		return
	}

	// 不能有2和大小王
	for _, v := range normalLine {
		if isJokerAndTwo(v) {
			return
		}
	}

	var laiZi int

	// 补牌
	for i := normalLine[0]; i <= normalLine[len(normalLine)-1]; i++ {
		switch normalCount[i] {
		case 0:
			laiZi += 4
		case 1:
			laiZi += 3
		case 2:
			laiZi += 2
		case 3:
			laiZi += 1
		}

		if laiZi > laiZiSize {
			return
		}
	}

	info.CardType = CardTypeRuanLianZha
	info.MinValue = normalLine[0]
	info.MaxValue = normalLine[len(normalLine)-1]
	return
}
