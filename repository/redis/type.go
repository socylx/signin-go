package redis

/*
门店/顾问 的 续卡目标值
*/
type RenewTargeValue map[string]uint64

/*
redis 中的数据类型
*/
type RedisType string

const (
	List    RedisType = "List"
	Set     RedisType = "Set"
	SortSet           = "SortSet"
)
