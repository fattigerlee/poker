package ddz

// 获取牌型
func GetType(cards []*Card) (cardType Type) {
	switch len(cards) {
	case 1:
		// 单
		return isDan(cards)
	case 2:
		// 对
		if cardType = isDui(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 火箭
		if cardType = isHuoJian(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 3:
		// 三不带
		if cardType = isSanBuDai(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 4:
		// 三带一
		if cardType = isSanDaiYi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 炸弹
		if cardType = isZhaDan(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 5:
		// 顺子
		if cardType = isShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 三带二
		if cardType = isSanDaiEr(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 6:
		// 顺子
		if cardType = isShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 连对
		if cardType = isLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机不带
		if cardType = isFeiJiBuDai(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 四带单
		if cardType = isSiDaiDan(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 7:
		// 顺子
		if cardType = isShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 8:
		// 连对
		if cardType = isLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机带一
		if cardType = isFeiJiDaiYi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 连炸
		if cardType = isLianZha(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 顺子
		if cardType = isShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 四带对
		if cardType = isSiDaiDui(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 9:
		// 顺子
		if cardType = isShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机不带
		if cardType = isFeiJiBuDai(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 10:
		// 飞机带二
		if cardType = isFeiJiDaiEr(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 连对
		if cardType = isLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 顺子
		if cardType = isShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 11:
		// 顺子
		if cardType = isShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 12:
		// 连对
		if cardType = isLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 连炸
		if cardType = isLianZha(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机不带
		if cardType = isFeiJiBuDai(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机带一
		if cardType = isFeiJiDaiYi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 顺子
		if cardType = isShunZi(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 14:
		// 连对
		if cardType = isLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 15:
		// 飞机不带
		if cardType = isFeiJiBuDai(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机带二
		if cardType = isFeiJiDaiEr(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 16:
		// 连对
		if cardType = isLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 连炸
		if cardType = isLianZha(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机带一
		if cardType = isFeiJiDaiYi(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 18:
		// 连对
		if cardType = isLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}
	case 20:
		// 连对
		if cardType = isLianDui(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 连炸
		if cardType = isLianZha(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机带一
		if cardType = isFeiJiDaiYi(cards); cardType.CardType != CardTypeNone {
			return
		}

		// 飞机带二
		if cardType = isFeiJiDaiEr(cards); cardType.CardType != CardTypeNone {
			return
		}
	}
	return
}
