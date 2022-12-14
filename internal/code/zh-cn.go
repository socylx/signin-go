package code

var text = map[int]string{
	ServerError:        "内部服务器错误",
	TooManyRequests:    "请求过多",
	ParamBindError:     "参数信息错误",
	AuthorizationError: "签名信息错误",
	UrlSignError:       "参数签名错误",
	CacheSetError:      "设置缓存失败",
	CacheGetError:      "获取缓存失败",
	CacheDelError:      "删除缓存失败",
	CacheNotExist:      "缓存不存在",
	ResubmitError:      "请勿重复提交",
	HashIdsEncodeError: "HashID 加密失败",
	HashIdsDecodeError: "HashID 解密失败",
	PermissionError:    "无权限",
	RedisConnectError:  "Redis 连接失败",
	MySQLConnectError:  "MySQL 连接失败",
	WriteConfigError:   "写入配置文件失败",
	SendEmailError:     "发送邮件失败",
	MySQLExecError:     "SQL 执行失败",
	GoVersionError:     "Go 版本不满足要求",
	SocketConnectError: "Socket 未连接",
	SocketSendError:    "Socket 消息发送失败",

	UsersCreateError:             "创建用户失败",
	UsersListError:               "获取用户列表失败",
	UsersDeleteError:             "删除用户失败",
	UsersUpdateError:             "更新用户失败",
	UsersModifyPersonalInfoError: "修改个人信息失败",
	UsersDetailError:             "获取个人信息失败",
	UsersDataError:               "获取学员数据失败",

	StrategyCreateError: "创建策略失败",
	StrategyQueryError:  "查询策略失败",
	StrategyDeleteError: "删除策略失败",
	StrategyUpdateError: "修改失败",

	StrategyIndicatorQueryError: "获取策略指标失败",

	StudioQueryError:             "查询门店失败",
	StudioStrategyMapCreateError: "应用门店失败",

	UserBeforeMemberQueryError: "查询线索失败",
}
