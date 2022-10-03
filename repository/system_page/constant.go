package system_page

type SystemPageData struct {
	ID      uint32
	PageKey string
}

var (
	Master,
	Analysis,
	Studio,
	Sale,
	Consultant,
	Service,
	Photographer,
	TeacherManager,
	SaleManager,
	HR,
	Designer,
	Market,
	Quality,
	NewMedia,
	Finance,
	JianZhi,
	Shop,
	FullTeacher,
	Product,
	MediaPartTime *SystemPageData
)

func init() {
	Master = &SystemPageData{ID: 1, PageKey: "master"}                  // 超级管理员
	Analysis = &SystemPageData{ID: 2, PageKey: "analysis"}              // 数据分析
	Studio = &SystemPageData{ID: 3, PageKey: "studio"}                  // 店长
	Sale = &SystemPageData{ID: 4, PageKey: "sale"}                      // 销售
	Consultant = &SystemPageData{ID: 5, PageKey: "consultant"}          // 课程顾问
	Service = &SystemPageData{ID: 6, PageKey: "service"}                // 客服
	Photographer = &SystemPageData{ID: 7, PageKey: "photographer"}      // 摄像师
	TeacherManager = &SystemPageData{ID: 8, PageKey: "teacher-manager"} // 师资管理
	SaleManager = &SystemPageData{ID: 9, PageKey: "sale-manager"}       // 销售运营
	HR = &SystemPageData{ID: 10, PageKey: "hr"}                         // 人力资源
	Market = &SystemPageData{ID: 11, PageKey: "market"}                 // 市场
	Quality = &SystemPageData{ID: 12, PageKey: "quality"}               // 质检
	NewMedia = &SystemPageData{ID: 13, PageKey: "new-media"}            // 新媒体运营
	Finance = &SystemPageData{ID: 14, PageKey: "finance"}               // 财务
	Designer = &SystemPageData{ID: 15, PageKey: "designer"}             // 设计师
	JianZhi = &SystemPageData{ID: 16, PageKey: "jianzhi"}               // 兼职
	Shop = &SystemPageData{ID: 17, PageKey: "shop"}                     // 商城管理
	FullTeacher = &SystemPageData{ID: 18, PageKey: "full-teacher"}      // 全职老师
	Product = &SystemPageData{ID: 19, PageKey: "product"}               // 产品
	MediaPartTime = &SystemPageData{ID: 20, PageKey: "media_part_time"} // 新媒体兼职
}
