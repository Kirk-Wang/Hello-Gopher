package conf

import "time"

const SysTimeform = "2006-01-02 15:04:05"
const SysTimeformShort = "2006-01-02"

var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

var SignSecret = []byte("0123456789abcdef")

var CookieSecret = "hellolottery"