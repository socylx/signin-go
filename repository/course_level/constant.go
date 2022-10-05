package course_level

type CourseLevelID = uint32

const (
	ChangGui CourseLevelID = 1 //常规课
	XiaoBan  CourseLevelID = 2 //小班课
	JiXun                  = 3 //集训课
	Master                 = 4 //大师课
	SiJiao                 = 5 //私教
	TeXun                  = 6 //特训
	TiYan                  = 7 //体验
	JiChu                  = 8 //基础课
)
