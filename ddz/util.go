package ddz

import "fmt"

// 牌型
type CardTypeInfo struct {
	CardType CardType // 牌型
	MinValue int      // 最小值
	MaxValue int      // 最大值
}

func (c *CardTypeInfo) Reset() {
	c.CardType = CardTypeNone
	c.MinValue = 0
	c.MaxValue = 0
}

func (c CardTypeInfo) String() string {
	return fmt.Sprintf("牌型:%s 最小值:%d 最大值:%d", c.CardType, c.MinValue, c.MaxValue)
}

type dictMap map[*Card]bool // 字典,方便找牌
type countList [18]int      // 每张牌数量
type valueList [5][]int     // 单张,对子,三张,四张牌的牌值
type lineList []int         // 所有牌的唯一牌值

// 切片转map
func convertToMap(cards []*Card) (dictCards dictMap) {
	dictCards = map[*Card]bool{}
	for _, c := range cards {
		dictCards[c] = true
	}
	return
}

// count 每张牌数量
// value[1] 所有单张的牌值
// value[2] 所有对子的牌值
// value[3] 所有三张的牌值
// value[4] 所有四张的牌值
// line 所有牌的唯一牌值
func getCountValueLine(dictCards dictMap) (count countList, value valueList, line lineList) {
	// 牌值计数
	for c := range dictCards {
		count[int(c.Num)]++
	}

	for i := 3; i < 18; i++ {
		if count[i] > 0 {
			line = append(line, i)
		}

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
	return
}

// 找牌
func findCardsByNums(dictCards dictMap, nums []int) []*Card {
	var findCards []*Card

	for _, num := range nums {
		for c := range dictCards {
			if c.Num == NumType(num) {
				findCards = append(findCards, c)
				delete(dictCards, c)
				break
			}
		}
	}
	return findCards
}

// 找癞子
func findCardsLaiZi(dictCards dictMap, count int) []*Card {
	var findCards []*Card

	for i := 0; i < count; i++ {
		for c := range dictCards {
			findCards = append(findCards, c)
			delete(dictCards, c)
			break
		}
	}
	return findCards
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
