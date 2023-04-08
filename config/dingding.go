package config

var DingDingSignCheck = false

var DingDingAppSecret = ""

var Timestamp int64
var Sign string

func StartDingDingSignCheck(AppSecret string) {
	DingDingSignCheck = true
	DingDingAppSecret = AppSecret
}
