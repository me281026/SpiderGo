package etime

import (
	"github.com/jinzhu/now"
	"time"
)

func MustDateToUnix(date string) int64 {
	if date == `` || date == `0000-00-00` {
		return 0
	}
	location, _ := time.LoadLocation("Asia/Beijing")
	inLocation, err := now.ParseInLocation(location, date)
	if err != nil {
		return 0
	}
	return inLocation.Unix()

}
