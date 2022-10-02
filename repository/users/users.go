package users

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
	"sync"
	"time"

	"gorm.io/gorm"
)

// Users 用户信息表
type Users struct {
	ID              int       `gorm:"primaryKey;column:id" json:"id"`
	Nickname        string    `gorm:"column:nickname" json:"nickname"` // 微信昵称
	Sex             uint32    `gorm:"column:sex" json:"sex"`           // 性别
	Country         string    `gorm:"column:country" json:"country"`
	Province        string    `gorm:"column:province" json:"province"`
	City            string    `gorm:"column:city" json:"city"`
	UnionID         string    `gorm:"column:union_id" json:"union_id"` // 微博账号id
	OpenID          string    `gorm:"column:open_id" json:"open_id"`
	OpenIDFuli      string    `gorm:"column:open_id_fuli" json:"open_id_fuli"` // 福利小程序的open_id
	Headimgurl      string    `gorm:"column:headimgurl" json:"headimgurl"`     // 头像（后台获取设置）
	Department      string    `gorm:"column:department" json:"department"`     // 部门
	Name            string    `gorm:"column:name" json:"name"`                 // 姓名
	Job             string    `gorm:"column:job" json:"job"`                   // 职务
	Phone           string    `gorm:"column:phone" json:"phone"`               // 电话
	Email           string    `gorm:"column:email" json:"email"`               // 邮箱
	Password        string    `gorm:"column:password" json:"password"`         // 密码（暂时不用）
	Wechat          string    `gorm:"column:wechat" json:"wechat"`             // 微信号
	JoinTime        time.Time `gorm:"column:join_time" json:"join_time"`       // 加入时间（以前的后台设置，新来的自动获取）
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`
	Birthday        time.Time `gorm:"column:birthday" json:"birthday"`                   // 会员生日
	Address         string    `gorm:"column:address" json:"address"`                     // 详细地址
	Remark          string    `gorm:"column:remark" json:"remark"`                       // 备注
	Source          int       `gorm:"column:source" json:"source"`                       // 来源，例如熟人推荐，大众点评等, 10 小程序
	ReferrerID      int       `gorm:"column:referrer_id" json:"referrer_id"`             // 推荐人user id
	BelongsStudioID uint32    `gorm:"column:belongs_studio_id" json:"belongs_studio_id"` // 用户所属门店，用户最新办卡的门店
	FuliCreateTime  time.Time `gorm:"column:fuli_create_time" json:"fuli_create_time"`   // 第一次注册或者访问小猫咪的时间
	Type            uint32    `gorm:"column:type" json:"type"`                           // 1-正常用户, 2-虚拟用户
	ManagerUserID   uint32    `gorm:"column:manager_user_id" json:"manager_user_id"`     // 会员负责人，顾问
}

type UsersRepo interface {
	i()
	TableName() string

	// dao.go
	Detail(ctx core.Context, userID uint32) (user *Users, err error)
	List(ctx core.Context, filter *Filter) (users []*Users, err error)
	Update(ctx core.Context, filter *Filter, data map[string]interface{}) (err error)
}

type users struct {
	db *gorm.DB
	// redis *redis.Client
}

func New() *users {
	once.Do(func() {
		u = &users{
			db: mysql.DB,
		}
	})
	return u
}

var u *users
var once sync.Once
var _ UsersRepo = (*users)(nil)

func (u *users) i() {}

// TableName get sql table name.获取数据库表名
func (u *users) TableName() string {
	return "users"
}
