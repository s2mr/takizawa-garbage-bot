package model

// ゴミ出しの地区
type Region int

// Region
const (
	A Region = iota
	B
)

// ゴミの種類
type GarbageType int

// Region
const (
	Normal GarbageType = iota
	Metal
	Glass
	PET
	News
	Medium
)
