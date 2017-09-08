package constant

import (
	"time"

	"github.com/shimokp/takizawa-garbage-bot/model"
)

const GARBAGE_NAME_NORMAL = "普通ゴミ"
const GARBAGE_NAME_METAL = "金属"
const GARBAGE_NAME_GLASS = "ガラス"
const GARBAGE_NAME_PET = "ペットボトル"
const GARBAGE_NAME_NEWS = "新聞紙・衣類"
const GARBAGE_NAME_MEDIUM = "中型ごみ"
const GARBAGE_NAME_HOLIDAY = "休み"
const GARBAGE_NAME_UNKNOWN = "不明"

const MESSAGE_PREFFIX_TODAY = "今日は"
const MESSAGE_PREFFIX_TOMORROW = "明日は"
const MESSAGE_SUFFIX = "の日です。"

var BiweeklyTuesdayStartGarbageForA = model.Medium
var BiweeklyTuesdaySecondGarbageForA = model.News
var BiweeklyFridayStartGarbageForA = model.PET
var BiweeklyFridaySecondGarbageForA = model.Glass
var BiweeklyTuesdayStartDateForA = time.Date(2017, 4, 4, 0, 0, 0, 0, &time.Location{})
var BiweeklyFridayStartDateForA = time.Date(2017, 4, 7, 0, 0, 0, 0, &time.Location{})

var BiweeklyWednesdayStartGarbageForB = model.Glass
var BiweeklyWednesdaySecondGarbageForB = model.PET
var BiweeklyThursdayStartGarbageForB = model.News
var BiweeklyThursdaySecondGarbageForB = model.Medium
var BiweeklyWednesdayStartDateForB = time.Date(2017, 4, 5, 0, 0, 0, 0, &time.Location{})
var BiweeklyThursdayStartDateForB = time.Date(2017, 4, 6, 0, 0, 0, 0, &time.Location{})

var HolidaysForA = []model.Hoge{
	{5, 3},
	{5, 5},
	{8, 11},
	{11, 3},
	{1, 1},
	{1, 2},
	{1, 3},
	{3, 21},
}

var HolidaysForB = []model.Hoge{
	{5, 3},
	{7, 17},
	{9, 18},
	{10, 9},
	{1, 1},
	{1, 2},
	{1, 3},
	{2, 12},
	{2, 21},
}
