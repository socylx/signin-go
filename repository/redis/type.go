package redis

type RenewTargeValue map[string]uint64

type RedisType string

const (
	List    RedisType = "List"
	Set     RedisType = "Set"
	SortSet           = "SortSet"
)
