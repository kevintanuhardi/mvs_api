package time

import "time"

func GetCurrentTime() (timestamp time.Time) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	return now
}

func GetCurrentTimeAdd15Min() (timestamp time.Time) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc).Add(time.Minute * 15)
	return now
}
