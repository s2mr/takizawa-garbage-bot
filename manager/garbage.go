package manager

import (
	"time"

	"github.com/shimokp/takizawa-garbage-bot/constant"
	"github.com/shimokp/takizawa-garbage-bot/model"
)

type GarbageManager struct {
	// Some fields
}

var sharedInstance *GarbageManager = newGarbageManager()

func newGarbageManager() *GarbageManager {
	// 何かしらの初期化処理
	return &GarbageManager{ /* 初期化 */ }
}

func GetInstance() *GarbageManager {
	return sharedInstance
}

func GetGarbageName(date time.Time, region model.Region) string {
	switch region {
	case model.A:
		return garbageTypeToString(getGarbageForA(date)) + constant.GARBAGE_NAME_SUFFIX
	case model.B:
		return garbageTypeToString(getGarbageForB(date)) + constant.GARBAGE_NAME_SUFFIX
	default:
		return constant.GARBAGE_NAME_UNKNOWN
	}
}

func garbageTypeToString(garbageType model.GarbageType) string {
	switch garbageType {
	case model.Normal:
		return constant.GARBAGE_NAME_NORMAL
	case model.Metal:
		return constant.GARBAGE_NAME_METAL
	case model.Glass:
		return constant.GARBAGE_NAME_GLASS
	case model.PET:
		return constant.GARBAGE_NAME_PET
	case model.News:
		return constant.GARBAGE_NAME_NEWS
	case model.Medium:
		return constant.GARBAGE_NAME_MEDIUM
	case model.Holiday:
		return constant.GARBAGE_NAME_HOLIDAY
	default:
		return constant.GARBAGE_NAME_UNKNOWN
	}
}

func getGarbageForA(date time.Time) model.GarbageType {
	if date.Weekday() == time.Sunday || date.Weekday() == time.Saturday {
		return model.Holiday
	}

	//休みの日かどうかを判定
	if isHoliday(date, model.A) {
		return model.Holiday
	}

	//隔週じゃないやつは返す
	switch date.Weekday() {
	case time.Monday:
		return model.Normal
	case time.Thursday:
		return model.Normal
	case time.Wednesday:
		return model.Metal
	}

	//隔週のやつがどっちかを判定
	return getGarbageBiweeklyForA(date)
}

func getGarbageForB(date time.Time) model.GarbageType {
	return model.Unknown
}

func isHoliday(date time.Time, region model.Region) bool {
	switch region {
	case model.A:
		for i := 0; i < len(constant.HolidaysForA); i++ {
			if constant.HolidaysForA[i].Month == int(date.Month()) && constant.HolidaysForA[i].Day == date.Day() {
				return true
			}
		}
	case model.B:
		for i := 0; i < len(constant.HolidaysForB); i++ {
			if constant.HolidaysForB[i].Month == int(date.Month()) && constant.HolidaysForB[i].Day == date.Day() {
				return true
			}
		}
	}
	return false
}

func getGarbageBiweeklyForA(date time.Time) model.GarbageType {
	switch date.Weekday() {
	case time.Tuesday:
		duration := date.Sub(constant.BiweeklyTuesdayStartDateForA)
		days := int(duration.Hours()) / 24
		if days%2 == 0 {
			return constant.BiweeklyTuesdayStartGarbageForA
		} else {
			return constant.BiweeklyTuesdaySecondGarbageForA
		}
	case time.Friday:
		duration := date.Sub(constant.BiweeklyFridayStartDateForA)
		days := int(duration.Hours()) / 24
		if days%2 == 0 {
			return constant.BiweeklyFridayStartGarbageForA
		} else {
			return constant.BiweeklyFridaySecondGarbageForA
		}
	}
	return model.Unknown
}
