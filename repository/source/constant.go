package source

type SourceID = uint32

const (
	DazhongDianpingMeituan   SourceID = 1
	Meituan                  SourceID = 2
	Fission                           = 3
	JingDianZiXun                     = 4
	DiTuiHuoDong                      = 5
	WeiXingMiniProgram                = 10
	SOHODiTui                         = 11
	FuLiMiniProgram                   = 12
	DouYinTuiGuang                    = 13
	Other                             = 14
	WeiXinPengYouQuan                 = 15
	GongZhongHaoWenZhang              = 16
	H5HuoDong                         = 17
	MiniProgramLiveGood               = 18
	BaiDuSouSuo                       = 19
	SanFangGongZhongHao               = 20
	XiaoNiuHeZuo                      = 21
	LiNingHeZuo                       = 22
	FengZhongXianXiaGuangGao          = 23
	WeiXinSouSuoMiniProgram           = 24
	KuaiShouQiYeHeZuo                 = 25
	SinaWeiBo                         = 26
	WuBangHeZuo                       = 27
	XiaoHongShu                       = 28
	ZheJieLianHe                      = 29
)

const (
	TotalType   string = "total"
	FissionType string = "fission"
	OtherType   string = "other"
)

var SourceType = map[SourceID]string{
	DazhongDianpingMeituan:   TotalType,
	Meituan:                  TotalType,
	Fission:                  FissionType,
	JingDianZiXun:            OtherType,
	DiTuiHuoDong:             TotalType,
	WeiXingMiniProgram:       TotalType,
	SOHODiTui:                OtherType,
	FuLiMiniProgram:          OtherType,
	DouYinTuiGuang:           TotalType,
	Other:                    OtherType,
	WeiXinPengYouQuan:        TotalType,
	GongZhongHaoWenZhang:     TotalType,
	H5HuoDong:                TotalType,
	MiniProgramLiveGood:      TotalType,
	BaiDuSouSuo:              TotalType,
	SanFangGongZhongHao:      TotalType,
	XiaoNiuHeZuo:             TotalType,
	LiNingHeZuo:              TotalType,
	FengZhongXianXiaGuangGao: OtherType,
	WeiXinSouSuoMiniProgram:  OtherType,
	KuaiShouQiYeHeZuo:        OtherType,
	SinaWeiBo:                TotalType,
	WuBangHeZuo:              OtherType,
	XiaoHongShu:              TotalType,
	ZheJieLianHe:             TotalType,
}
