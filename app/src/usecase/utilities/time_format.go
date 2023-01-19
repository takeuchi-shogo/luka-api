package utilities

import "time"

func SetCurrentUnixTime() int64 {
	return time.Now().Unix()
}

func FormatStr(unixtime int64) string {
	layout := "2006年01月02日　15:04:05"
	return time.Unix(unixtime, 0).Format(layout)
}
