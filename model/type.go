package model

import (
	"time"
)

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

type MessageText struct {
	Events []Event `json:"events"`
}

type Event struct {
	ReplyToken string `json:"replyToken"`
	Type       string `json:"type"`
	Timestamp  int64  `json:"timestamp"`
	Source     struct {
		Type   string `json:"type"`
		UserID string `json:"userId"`
	} `json:"source"`
	Message struct {
		ID   string `json:"id"`
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"message"`
}

type Profile struct {
	DisplayName   string `json:"displayName"`
	UserID        string `json:"userId"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

type User struct {
	ID      int64
	UserID  string
	Region  Region
	Created time.Time
}

func ConvertStringToRegion(region string) Region {
	switch region {
	case "A":
		return A
	default:
		return B
	}
}

func ConvertRegionToString(region Region) string {
	switch region {
	case A:
		return "A"
	default:
		return "B"
	}
}

func ConvertStringToDateType(dateType string) DateType {
	switch dateType {
	case "today":
		return Today
	default:
		return Tomorrow
	}
}
