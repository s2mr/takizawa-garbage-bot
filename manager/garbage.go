package manager

import "time"

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

func GetGarbageName(date time.Time) string {

	return "ゴミです"
}
