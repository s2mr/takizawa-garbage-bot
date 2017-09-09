package manager

import (
	"time"

	"github.com/shimokp/takizawa-garbage-bot/constant"
	"github.com/shimokp/takizawa-garbage-bot/model"
)

type GarbageManager struct {
	// Some fields
}

var sharedGarbageManagerInstance *GarbageManager = newGarbageManager()

func newGarbageManager() *GarbageManager {
	// 何かしらの初期化処理
	return &GarbageManager{ /* 初期化 */ }
}

func GetGarbageManagerInstance() *GarbageManager {
	return sharedGarbageManagerInstance
}

func GetMessage(dateType model.DateType, region model.Region) string {
	var message string

	switch dateType {
	case model.Today:
		message = constant.MESSAGE_PREFFIX_TODAY + getGarbageName(time.Now(), region)
	case model.Tomorrow:
		tomorrow := time.Now().AddDate(0, 0, 1)
		message = constant.MESSAGE_PREFFIX_TOMORROW + getGarbageName(tomorrow, region)
	}

	return message + constant.MESSAGE_SUFFIX
}

func getGarbageName(date time.Time, region model.Region) string {
	switch region {
	case model.A:
		return garbageTypeToString(getGarbageForA(date))
	case model.B:
		return garbageTypeToString(getGarbageForB(date))
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
	case time.Wednesday:
		return model.Metal
	case time.Thursday:
		return model.Normal
	}

	//隔週のやつがどっちかを判定
	return getGarbageBiweeklyForA(date)
}

func getGarbageForB(date time.Time) model.GarbageType {
	if date.Weekday() == time.Sunday || date.Weekday() == time.Saturday {
		return model.Holiday
	}

	//休みの日かどうかを判定
	if isHoliday(date, model.B) {
		return model.Holiday
	}

	//隔週じゃないやつは返す
	switch date.Weekday() {
	case time.Monday:
		return model.Metal
	case time.Tuesday:
		return model.Normal
	case time.Friday:
		return model.Normal
	}

	//隔週のやつがどっちかを判定
	return getGarbageBiweeklyForB(date)
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

func getGarbageBiweeklyForB(date time.Time) model.GarbageType {
	switch date.Weekday() {
	case time.Wednesday:
		duration := date.Sub(constant.BiweeklyWednesdayStartDateForB)
		days := int(duration.Hours()) / 24
		if days%2 == 0 {
			//FIXME: 曜日に依存しない変数名
			return constant.BiweeklyWednesdayStartGarbageForB
		} else {
			return constant.BiweeklyWednesdaySecondGarbageForB
		}
	case time.Thursday:
		duration := date.Sub(constant.BiweeklyThursdayStartDateForB)
		days := int(duration.Hours()) / 24
		if days%2 == 0 {
			return constant.BiweeklyThursdayStartGarbageForB
		} else {
			return constant.BiweeklyThursdaySecondGarbageForB
		}
	}
	return model.Unknown
}
