package page_access

type PageAccessType string

const (
	MiniProgram     PageAccessType = "mini"
	Admin           PageAccessType = "admin"
	FuliMiniProgram                = "fuli"
)

const (
	CourseDetailIndex string = "pages/course/course-detail/index"
	BuyCardIndex      string = "sub-2/shop/buy-card/index"
)
