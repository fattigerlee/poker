package ddz

// 获取牌型
func GetType(cards []*Card) (cardType Type) {
	switch len(cards) {
	case 1:
		// 单
		return IsDan(cards)
	case 2:
		// 对
		if cardType = IsDui(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 火箭
		if cardType = IsHuoJian(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 3:
		// 三不带
		if cardType = IsSanBuDai(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 4:
		// 三带一
		if cardType = IsSanDaiYi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 炸弹
		if cardType = IsZhaDan(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 5:
		// 顺子
		if cardType = IsShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 三带二
		if cardType = IsSanDaiEr(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 6:
		// 顺子
		if cardType = IsShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 连对
		if cardType = IsLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机不带
		if cardType = IsFeiJiBuDai(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 四带单
		if cardType = IsSiDaiDan(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 7:
		// 顺子
		if cardType = IsShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 8:
		// 连对
		if cardType = IsLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机带一
		if cardType = IsFeiJiDaiYi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 连炸
		if cardType = IsLianZha(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 顺子
		if cardType = IsShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 四带对
		if cardType = IsSiDaiDui(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 9:
		// 顺子
		if cardType = IsShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机不带
		if cardType = IsFeiJiBuDai(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 10:
		// 飞机带二
		if cardType = IsFeiJiDaiEr(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 连对
		if cardType = IsLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 顺子
		if cardType = IsShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 11:
		// 顺子
		if cardType = IsShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 12:
		// 连对
		if cardType = IsLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 连炸
		if cardType = IsLianZha(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机不带
		if cardType = IsFeiJiBuDai(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机带一
		if cardType = IsFeiJiDaiYi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 顺子
		if cardType = IsShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 14:
		// 连对
		if cardType = IsLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 15:
		// 飞机不带
		if cardType = IsFeiJiBuDai(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机带二
		if cardType = IsFeiJiDaiEr(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 16:
		// 连对
		if cardType = IsLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 连炸
		if cardType = IsLianZha(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机带一
		if cardType = IsFeiJiDaiYi(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 18:
		// 连对
		if cardType = IsLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 20:
		// 连对
		if cardType = IsLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 连炸
		if cardType = IsLianZha(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机带一
		if cardType = IsFeiJiDaiYi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机带二
		if cardType = IsFeiJiDaiEr(cards); cardType.CardType != CardTypeNone {
			return
		}
	}
	return
}
