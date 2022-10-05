package utils

import "math"

type Location struct {
	Longitude float64 //经度
	Latitude  float64 //纬度
}

var TianAnMen = &Location{Longitude: 116.38, Latitude: 39.90}

const (
	BeiJindRadius float64 = 50       //千米
	EarthRadius   float64 = 6378.137 //地球半径
)

func Distance(lng1, lat1, lng2, lat2 float64) float64 {
	rad := math.Pi / 180.0

	lng1, lat1 = lng1*rad, lat1*rad
	lng2, lat2 = lng2*rad, lat2*rad
	theta := lng2 - lng1

	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * BeiJindRadius
}
