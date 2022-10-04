package strategy

type StrategyType = uint32

const (
	Xuka            StrategyType = 100
	LaXinCoupon     StrategyType = 201
	LaXinSubscribe  StrategyType = 202
	LaXinTry        StrategyType = 203
	LaXinMembership StrategyType = 204
)

var StrategyTypeCheck = map[uint32]bool{
	Xuka:            true,
	LaXinCoupon:     true,
	LaXinSubscribe:  true,
	LaXinTry:        true,
	LaXinMembership: true,
}

var StrategyTypes = []map[string]interface{}{
	{"value": Xuka, "name": "续卡"},
	{"value": LaXinCoupon, "name": "拉新需促买券"},
	{"value": LaXinSubscribe, "name": "拉新需促约课"},
	{"value": LaXinTry, "name": "拉新需促试课"},
	{"value": LaXinMembership, "name": "拉新需促成交"},
}
