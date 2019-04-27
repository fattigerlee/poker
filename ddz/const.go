package ddz

import "strconv"

// 牌值
type NumType int

const (
	NumTypeNone       NumType = iota
	NumTypeThree              = 3
	NumTypeAce                = 14
	NumTypeTwo                = 15
	NumTypeSmallJoker         = 16
	NumTypeBigJoker           = 17
)

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
	case NumTypeAce:
		return "A"
	case NumTypeTwo:
		return "2"
	case NumTypeSmallJoker:
		return "小王"
	case NumTypeBigJoker:
		return "大王"
	default:
		return "牌值错误"
	}
}

// 花色
type SuitType int

const (
	SuitTypeNone    SuitType = iota
	SuitTypeJoker            // 鬼牌
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
	CardTypeNone            CardType = iota
	CardTypeDan                      // 单
	CardTypeDui                      // 对
	CardTypeSanBuDai                 // 三不带
	CardTypeSanDaiYi                 // 三带一
	CardTypeSanDaiEr                 // 三带二
	CardTypeSiDaiDan                 // 四带单(炸弹带两张单)
	CardTypeSiDaiDui                 // 四带对(炸弹带两对)
	CardTypeShunZi                   // 顺子
	CardTypeLianDui                  // 连对
	CardTypeFeiJiBuDai               // 飞机不带
	CardTypeFeiJiDaiYi               // 飞机带一
	CardTypeFeiJiDaiEr               // 飞机带二
	CardTypeRuanZhaDan4              // 四软炸
	CardTypeZhaDan                   // 硬炸弹
	CardTypeLaiZiZhaDan4             // 四癞子炸
	CardTypeChunLaiZiZhaDan          // 四纯癞子炸
	CardTypeRuanZhaDan5              // 五软炸
	CardTypeLaiZiZhaDan5             // 五癞子炸
	CardTypeRuanZhaDan6              // 六软炸
	CardTypeLaiZiZhaDan6             // 六癞子炸
	CardTypeRuanZhaDan7              // 七软炸
	CardTypeLaiZiZhaDan7             // 七癞子炸
	CardTypeRuanZhaDan8              // 八软炸
	CardTypeLaiZiZhaDan8             // 八癞子炸
	CardTypeRuanZhaDan9              // 九软炸
	CardTypeRuanZhaDan10             // 十软炸
	CardTypeRuanZhaDan11             // 十一软炸
	CardTypeRuanZhaDan12             // 十二软炸
	CardTypeHuoJian                  // 火箭
	CardTypeRuanLianZha              // 软连炸(癞子只有大小王)
	CardTypeLianZha                  // 硬连炸
)

func (c CardType) String() string {
	switch c {
	case CardTypeDan:
		return "单"
	case CardTypeDui:
		return "对"
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
	case CardTypeLianDui:
		return "连对"
	case CardTypeFeiJiBuDai:
		return "飞机不带"
	case CardTypeFeiJiDaiYi:
		return "飞机带一"
	case CardTypeFeiJiDaiEr:
		return "飞机带二"
	case CardTypeRuanZhaDan4:
		return "四软炸"
	case CardTypeZhaDan:
		return "硬炸弹"
	case CardTypeLaiZiZhaDan4:
		return "四癞子炸弹"
	case CardTypeChunLaiZiZhaDan:
		return "四纯癞子炸弹"
	case CardTypeRuanZhaDan5:
		return "五软炸"
	case CardTypeLaiZiZhaDan5:
		return "五癞子炸弹"
	case CardTypeRuanZhaDan6:
		return "六软炸"
	case CardTypeLaiZiZhaDan6:
		return "六癞子炸弹"
	case CardTypeRuanZhaDan7:
		return "七软炸"
	case CardTypeLaiZiZhaDan7:
		return "七癞子炸弹"
	case CardTypeRuanZhaDan8:
		return "八软炸"
	case CardTypeLaiZiZhaDan8:
		return "八癞子炸弹"
	case CardTypeRuanZhaDan9:
		return "九软炸"
	case CardTypeRuanZhaDan10:
		return "十软炸"
	case CardTypeRuanZhaDan11:
		return "十一软炸"
	case CardTypeRuanZhaDan12:
		return "十二软炸"
	case CardTypeHuoJian:
		return "火箭"
	case CardTypeRuanLianZha:
		return "软连炸"
	case CardTypeLianZha:
		return "硬连炸"
	default:
		return "牌型错误"
	}
}

// 底牌牌型
type BottomCardType int

const (
	BottomCardTypeNone        BottomCardType = iota
	BottomCardTypeWangZha                    // 王炸
	BottomCardTypeDanWang                    // 单王
	BottomCardTypeTongHuaShun                // 同花顺
	BottomCardTypeTongHua                    // 同花
	BottomCardTypeShunZi                     // 顺子
	BottomCardTypeSanBuDai                   // 三不带
)

func (c BottomCardType) String() string {
	switch c {
	case BottomCardTypeWangZha:
		return "王炸"
	case BottomCardTypeDanWang:
		return "单王"
	case BottomCardTypeTongHuaShun:
		return "同花顺"
	case BottomCardTypeTongHua:
		return "同花"
	case BottomCardTypeShunZi:
		return "顺子"
	case BottomCardTypeSanBuDai:
		return "三不带"
	default:
		return "底牌牌型错误"
	}
}

// 结果类型
type ResultType int

const (
	ResultTypeNone ResultType = iota
	ResultTypeMin             // 最小值
	ResultTypeMax             // 最大值
)
