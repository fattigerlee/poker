package ddz

import "strconv"

// 牌值
type NumType int

func (n NumType) String() string {
	switch n {
	case 3, 4, 5, 6, 7, 8, 9, 10:
		return strconv.Itoa(int(n))
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	case 14:
		return "A"
	case 15:
		return "2"
	case 16:
		return "小王"
	case 17:
		return "大王"
	default:
		return "牌值错误"
	}
}

// 花色
type SuitType int

const (
	SuitTypeJoker   SuitType = iota
	SuitTypeHeart            // 红桃
	SuitTypeSpade            // 黑桃
	SuitTypeDiamond          // 方块
	SuitTypeClub             // 草花
)

func (s SuitType) String() string {
	switch s {
	case SuitTypeJoker:
		return "鬼牌"
	case SuitTypeHeart:
		return "红桃"
	case SuitTypeSpade:
		return "黑桃"
	case SuitTypeDiamond:
		return "方块"
	case SuitTypeClub:
		return "草花"
	default:
		return "花色错误"
	}
}

// 牌型
type CardType int

const (
	CardTypeNone       CardType = iota
	CardTypeDan                 // 单
	CardTypeDui                 // 对
	CardTypeLianDui             // 连对
	CardTypeSanBuDai            // 三不带
	CardTypeSanDaiYi            // 三带一
	CardTypeSanDaiEr            // 三带二
	CardTypeSiDaiDan            // 四带单(炸弹带两张单)
	CardTypeSiDaiDui            // 四带对(炸弹带两对)
	CardTypeShunZi              // 顺子
	CardTypeFeiJiBuDai          // 飞机不带
	CardTypeFeiJiDaiYi          // 飞机带一
	CardTypeFeiJiDaiEr          // 飞机带二
	CardTypeZhaDan              // 炸弹
	CardTypeHuoJian             // 火箭
	CardTypeLianZha             // 连炸
)

func (c CardType) String() string {
	switch c {
	case CardTypeDan:
		return "单"
	case CardTypeDui:
		return "对"
	case CardTypeLianDui:
		return "连对"
	case CardTypeSanBuDai:
		return "三不带"
	case CardTypeSanDaiYi:
		return "三带一"
	case CardTypeSanDaiEr:
		return "三带二"
	case CardTypeSiDaiDan:
		return "四带单(炸弹带两张单)"
	case CardTypeSiDaiDui:
		return "四带对(炸弹带两对)"
	case CardTypeShunZi:
		return "顺子"
	case CardTypeFeiJiBuDai:
		return "飞机不带"
	case CardTypeFeiJiDaiYi:
		return "飞机带一"
	case CardTypeFeiJiDaiEr:
		return "飞机带二"
	case CardTypeZhaDan:
		return "炸弹"
	case CardTypeHuoJian:
		return "火箭"
	case CardTypeLianZha:
		return "连炸"
	default:
		return "牌型错误"
	}
}
