package modules

import (
	"time"
)

func CheckTime_now(variable string) string {
	now := time.Now()
	return now.Format("2006-01-02")
}
