package model

//指定日時
type DateType int

//DateType
const (
	Today DateType = iota + 1
	Tomorrow
)

// ゴミ出しの地区
type Region int

// Region
const (
	A Region = iota + 1
	B
)

// ゴミの種類
type GarbageType int

// GarbageType
const (
	Normal GarbageType = iota + 1
	Metal
	Glass
	PET
	News
	Medium
	Holiday
	Unknown
)

//TODO: rename
type Hoge struct {
	Month int
	Day   int
}
