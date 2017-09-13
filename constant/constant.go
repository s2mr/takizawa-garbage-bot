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

const MESSAGE_FIRST_RESPONSE = `友達追加ありがとうございます！
まず初めにあなたの地区を登録してください。
[A]
小岩井、大釜上、大釜南、 篠木、大沢、姥屋敷、元村南、元村中央、元村東、元村西、元村北、法誓寺、国分、あすみ野、室小路、柳沢、南一本木、北一本木、いずみ巣子ニュータウン
[B]
鵜飼中央、滝沢パークタウン、上の山、上鵜飼、鵜飼南、鵜飼温泉、 滝沢ニュータウン、巣子、南巣子、長根、川前

このメッセージにAまたはBと返信してください。`
const MESSAGE_COMMAND_INSTRUCTION = `以下のコマンドが利用できます。
[今日]=>今日のゴミ収集情報をお伝えします。
[明日]=>明日のゴミ収集情報をお伝えします。
[A]=>地区をAに変更します。
[B]=>地区をBに変更します。
コマンドの利用は[]内の文字のみをご入力ください

また、１日２回（朝は7:00~8:00, 夜は19:00~20:00）今日または明日のゴミ収集情報を配信します。
それでは良いごみLIFEを。`
const MESSAGE_COMMAND_NOTFOUND = `そのコマンドは存在しません。打ち間違えがないかお確かめください。
尚、英字は半角のみの入力ができます。`

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
