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
	default:
		return constant.GARBAGE_NAME_UNKNOWN
	}
}

func getGarbageForA(date time.Time) model.GarbageType {
	return model.Glass
}

func getGarbageForB(date time.Time) model.GarbageType {
	return model.Glass
}
