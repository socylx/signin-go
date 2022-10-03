package system_page

type PageKey = string

const (
	Master         PageKey = "master"          // 超级管理员
	Analysis       PageKey = "analysis"        // 数据分析
	Studio                 = "studio"          // 店长
	Sale                   = "sale"            // 销售
	Consultant             = "consultant"      // 课程顾问
	Service                = "service"         // 客服
	Photographer           = "photographer"    // 摄像师
	TeacherManager         = "teacher-manager" // 师资管理
	SaleManager            = "sale-manager"    // 销售运营
	HR                     = "hr"              // 人力资源
	Designer               = "designer"        // 设计师
	Market                 = "market"          // 市场
	Quality                = "quality"         // 质检
	NewMedia               = "new-media"       // 新媒体运营
	Finance                = "finance"         // 财务
	JianZhi                = "jianzhi"         // jianzhi
	Shop                   = "shop"            // 商城管理
	FullTeacher            = "full-teacher"    // 全职老师
	Product                = "product"         // 产品
	MediaPartTime          = "media_part_time" // 新媒体兼职
)
