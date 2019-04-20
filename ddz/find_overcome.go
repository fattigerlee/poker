package ddz

import "sort"

// 扑克牌列表中找出比牌型更大的牌
func FindOvercomeCard(info CardTypeInfo, cards []*Card) (ret []*Card) {
	// 手牌排序
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Num < cards[j].Num
	})

	count := cardCount(cards)

	switch info.CardType {
	case CardTypeDan:
		if ret = findBigDan(cards, count, info); len(ret) != 0 {
			return
		}
	case CardTypeDui:
		if ret = findBigDui(cards, count, info); len(ret) != 0 {
			return
		}
	case CardTypeLianDui:
		if ret = findBigLianDui(cards, count, info); len(ret) != 0 {
			return
		}
	case CardTypeSanBuDai:
		if ret = findBigSanBuDai(cards, count, info); len(ret) != 0 {
			return
		}
	case CardTypeSanDaiYi:
		if ret = findBigSanDaiYi(cards, count, info); len(ret) != 0 {
			return
		}
	case CardTypeSanDaiEr:
		if ret = findBigSanDaiEr(cards, count, info); len(ret) != 0 {
			return
		}
	case CardTypeShunZi:
		if ret = findBigShunZi(cards, count, info); len(ret) != 0 {
			return
		}
	case CardTypeFeiJiBuDai:
		if ret = findBigFeiJiBuDai(cards, count, info); len(ret) != 0 {
			return
		}
	case CardTypeFeiJiDaiYi:
		if ret = findBigFeiJiDaiYi(cards, count, info); len(ret) != 0 {
			return
		}
	case CardTypeFeiJiDaiEr:
		if ret = findBigFeiJiDaiEr(cards, count, info); len(ret) != 0 {
			return
		}
	case CardTypeZhaDan:
		if ret = findBigZhaDan(cards, count, info); len(ret) != 0 {
			return
		}

		if ret = findHuoJian(cards, count); len(ret) != 0 {
			return
		}

		if ret = findLianZha(cards, count); len(ret) != 0 {
			return
		}
	case CardTypeHuoJian:
		if ret = findBigHuoJian(cards, count); len(ret) != 0 {
			return
		}
	case CardTypeLianZha:
		if ret = findBigLianZha(cards, count, info); len(ret) != 0 {
			return
		}
	}

	if info.CardType >= CardTypeDan && info.CardType < CardTypeZhaDan {
		if ret = findZhaDan(cards, count); ret != nil {
			return
		}

		if ret = findHuoJian(cards, count); ret != nil {
			return
		}
	}
	return nil
}

func findZhaDan(cards []*Card, count [18]int) (ret []*Card) {
	var zhaDan int
	for i := 3; i <= 15; i++ {
		if count[i] == 4 {
			zhaDan = i
			break
		}
	}

	if zhaDan == 0 {
		return
	}

	for i := 0; i < len(cards); i++ {
		if int(cards[i].Num) == zhaDan {
			for k := 0; k < 4; k++ {
				ret = append(ret, cards[i+k])
			}
			break
		}
	}
	return
}

func findHuoJian(cards []*Card, count [18]int) (ret []*Card) {
	var huoJian int
	if count[16] == 1 && count[17] == 1 {
		huoJian = 1
	}

	if huoJian == 0 {
		return
	}

	for i := 0; i < len(cards); i++ {
		if int(cards[i].Num) == 16 {
			for k := 0; k < 1; k++ {
				ret = append(ret, cards[i+k])
			}
			break
		}
	}

	for i := 0; i < len(cards); i++ {
		if int(cards[i].Num) == 17 {
			for k := 0; k < 1; k++ {
				ret = append(ret, cards[i+k])
			}
			break
		}
	}
	return
}

func findLianZha(cards []*Card, count [18]int) (ret []*Card) {
	var zhaDanList []int
	for i := 3; i < 14; i++ {
		if count[i] == 4 && count[i+1] == 4 {
			zhaDanList = append(zhaDanList, i, i+1)
			break
		}
	}

	if len(zhaDanList) == 0 {
		return
	}

	for _, zhaDan := range zhaDanList {
		for i := 0; i < len(cards); i++ {
			if int(cards[i].Num) == zhaDan {
				for k := 0; k < 4; k++ {
					ret = append(ret, cards[i+k])
				}
				break
			}
		}
	}
	return
}

func findBigDan(cards []*Card, count [18]int, info CardTypeInfo) (ret []*Card) {
	var dan int
	for i := info.MinValue + 1; i <= 17; i++ {
		if count[i] >= 1 {
			dan = i
			break
		}
	}

	if dan == 0 {
		return
	}

	for i := 0; i < len(cards); i++ {
		if int(cards[i].Num) == dan {
			for k := 0; k < 1; k++ {
				ret = append(ret, cards[i+k])
			}
			break
		}
	}
	return
}

func findBigDui(cards []*Card, count [18]int, info CardTypeInfo) (ret []*Card) {
	var dui int
	for i := info.MinValue + 1; i <= 15; i++ {
		if count[i] >= 2 {
			dui = i
			break
		}
	}

	if dui == 0 {
		return
	}

	for i := 0; i < len(cards); i++ {
		if int(cards[i].Num) == dui {
			for k := 0; k < 2; k++ {
				ret = append(ret, cards[i+k])
			}
			break
		}
	}
	return
}

func findBigLianDui(cards []*Card, count [18]int, info CardTypeInfo) (ret []*Card) {
	valueRange := info.MaxValue - info.MinValue + 1

	var duiList []int
	for i := info.MinValue + 1; i <= 14; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 2 {
				exist = false
				break
			}
		}

		if exist {
			for j := i; j < i+valueRange; j++ {
				duiList = append(duiList, j)
			}
			break
		}
	}

	if len(duiList) == 0 {
		return
	}

	for _, dui := range duiList {
		for i := 0; i < len(cards); i++ {
			if int(cards[i].Num) == dui {
				for k := 0; k < 2; k++ {
					ret = append(ret, cards[i+k])
				}
				break
			}
		}
	}
	return
}

func findBigSanBuDai(cards []*Card, count [18]int, info CardTypeInfo) (ret []*Card) {
	var san int
	for i := info.MinValue + 1; i <= 15; i++ {
		if count[i] >= 3 {
			san = i
			break
		}
	}

	if san == 0 {
		return
	}

	for i := 0; i < len(cards); i++ {
		if int(cards[i].Num) == san {
			for k := 0; k < 3; k++ {
				ret = append(ret, cards[i+k])
			}
			break
		}
	}
	return
}

func findBigSanDaiYi(cards []*Card, count [18]int, info CardTypeInfo) (ret []*Card) {
	var san int
	var dan int
	for i := info.MinValue + 1; i <= 15; i++ {
		if count[i] >= 3 {
			san = i
			break
		}
	}

	if san == 0 {
		return
	}

	for i := 3; i <= 15; i++ {
		if i != san && count[i] >= 1 {
			dan = i
			break
		}
	}

	if dan == 0 {
		return
	}

	for i := 0; i < len(cards); i++ {
		if int(cards[i].Num) == san {
			for k := 0; k < 3; k++ {
				ret = append(ret, cards[i+k])
			}
			break
		}
	}

	for i := 0; i < len(cards); i++ {
		if int(cards[i].Num) == dan {
			for k := 0; k < 1; k++ {
				ret = append(ret, cards[i+k])
			}
			break
		}
	}
	return
}

func findBigSanDaiEr(cards []*Card, count [18]int, info CardTypeInfo) (ret []*Card) {
	var san int
	var dui int
	for i := info.MinValue + 1; i <= 15; i++ {
		if count[i] >= 3 {
			san = i
			break
		}
	}

	if san == 0 {
		return
	}

	for i := 3; i <= 15; i++ {
		if i != san && count[i] >= 2 {
			dui = i
			break
		}
	}

	if dui == 0 {
		return
	}

	for i := 0; i < len(cards); i++ {
		if int(cards[i].Num) == san {
			for k := 0; k < 3; k++ {
				ret = append(ret, cards[i+k])
			}
			break
		}
	}

	for i := 0; i < len(cards); i++ {
		if int(cards[i].Num) == dui {
			for k := 0; k < 2; k++ {
				ret = append(ret, cards[i+k])
			}
			break
		}
	}
	return
}

func findBigShunZi(cards []*Card, count [18]int, info CardTypeInfo) (ret []*Card) {
	valueRange := info.MaxValue - info.MinValue + 1

	var danList []int
	for i := info.MinValue + 1; i <= 14; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 1 {
				exist = false
				break
			}
		}

		if exist {
			for j := i; j < i+valueRange; j++ {
				danList = append(danList, j)
			}
			break
		}
	}

	if len(danList) == 0 {
		return
	}

	for _, dan := range danList {
		for i := 0; i < len(cards); i++ {
			if int(cards[i].Num) == dan {
				for k := 0; k < 1; k++ {
					ret = append(ret, cards[i+k])
				}
				break
			}
		}
	}
	return
}

func findBigFeiJiBuDai(cards []*Card, count [18]int, info CardTypeInfo) (ret []*Card) {
	valueRange := info.MaxValue - info.MinValue + 1

	var sanList []int
	for i := info.MinValue + 1; i <= 14; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 3 {
				exist = false
				break
			}
		}

		if exist {
			for j := i; j < i+valueRange; j++ {
				sanList = append(sanList, j)
			}
			break
		}
	}

	if len(sanList) == 0 {
		return
	}

	for _, san := range sanList {
		for i := 0; i < len(cards); i++ {
			if int(cards[i].Num) == san {
				for k := 0; k < 3; k++ {
					ret = append(ret, cards[i+k])
				}
				break
			}
		}
	}
	return
}

func findBigFeiJiDaiYi(cards []*Card, count [18]int, info CardTypeInfo) (ret []*Card) {
	valueRange := info.MaxValue - info.MinValue + 1

	var sanList []int
	for i := info.MinValue + 1; i <= 14; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 3 {
				exist = false
				break
			}
		}

		if exist {
			for j := i; j < i+valueRange; j++ {
				sanList = append(sanList, j)
			}
			break
		}
	}

	if len(sanList) == 0 {
		return
	}

	var danList []int
	for i := 3; i <= 15; i++ {
		if count[i] > 0 && count[i] < 3 {
			for j := 0; j < count[i]; j++ {
				danList = append(danList, i)

				if len(danList) == len(sanList) {
					break
				}
			}

			if len(danList) == len(sanList) {
				break
			}
		}
	}

	if len(danList) != len(sanList) {
		return
	}

	for _, san := range sanList {
		for i := 0; i < len(cards); i++ {
			if int(cards[i].Num) == san {
				for k := 0; k < 3; k++ {
					ret = append(ret, cards[i+k])
				}
				break
			}
		}
	}

	for _, dan := range danList {
		for i := 0; i < len(cards); i++ {
			if int(cards[i].Num) == dan {
				ret = append(ret, cards[i])
			}
		}
	}
	return
}

func findBigFeiJiDaiEr(cards []*Card, count [18]int, info CardTypeInfo) (ret []*Card) {
	valueRange := info.MaxValue - info.MinValue + 1

	var sanList []int
	for i := info.MinValue + 1; i <= 14; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 3 {
				exist = false
				break
			}
		}

		if exist {
			for j := i; j < i+valueRange; j++ {
				sanList = append(sanList, j)
			}
			break
		}
	}

	if len(sanList) == 0 {
		return
	}

	var duiList []int
	for i := 3; i <= 15; i++ {
		var exist bool
		for _, san := range sanList {
			if san == i {
				exist = true
				break
			}
		}

		if exist {
			continue
		}

		if count[i] >= 2 && count[i] < 4 {
			duiList = append(duiList, i)

			if len(duiList) == len(sanList) {
				break
			}
		}
	}

	if len(duiList) != len(sanList) {
		return
	}

	for _, san := range sanList {
		for i := 0; i < len(cards); i++ {
			if int(cards[i].Num) == san {
				for k := 0; k < 3; k++ {
					ret = append(ret, cards[i+k])
				}
				break
			}
		}
	}

	for _, dui := range duiList {
		for i := 0; i < len(cards); i++ {
			if int(cards[i].Num) == dui {
				for k := 0; k < 2; k++ {
					ret = append(ret, cards[i+k])
				}
				break
			}
		}
	}
	return
}

func findBigZhaDan(cards []*Card, count [18]int, info CardTypeInfo) (ret []*Card) {
	var zhaDan int
	for i := info.MinValue + 1; i <= 17; i++ {
		if count[i] == 4 {
			zhaDan = i
			break
		}
	}

	if zhaDan == 0 {
		return
	}

	for i := 0; i < len(cards); i++ {
		if int(cards[i].Num) == zhaDan {
			for k := 0; k < 4; k++ {
				ret = append(ret, cards[i+k])
			}
			break
		}
	}
	return
}

func findBigHuoJian(cards []*Card, count [18]int) (ret []*Card) {
	return findLianZha(cards, count)
}

func findBigLianZha(cards []*Card, count [18]int, info CardTypeInfo) (ret []*Card) {
	valueRange := info.MaxValue - info.MinValue + 1

	var zhaDanList []int
	for i := info.MinValue + 1; i <= 14; i++ {
		exist := true
		for j := i; j < i+valueRange; j++ {
			if count[j] < 4 {
				exist = false
				break
			}
		}

		if exist {
			for j := i; j < i+valueRange; j++ {
				zhaDanList = append(zhaDanList, j)
			}
			break
		}
	}

	if len(zhaDanList) == 0 {
		return
	}

	for _, zhaDan := range zhaDanList {
		for i := 0; i < len(cards); i++ {
			if int(cards[i].Num) == zhaDan {
				for k := 0; k < 4; k++ {
					ret = append(ret, cards[i+k])
				}
				break
			}
		}
	}
	return
}
