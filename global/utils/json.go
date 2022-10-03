package utils

import jsoniter "github.com/json-iterator/go"

var Json jsoniter.API

func init() {
	Json = jsoniter.ConfigCompatibleWithStandardLibrary
}
