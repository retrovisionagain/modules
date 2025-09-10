package modules

import (
	"time"
)

func Check_now(variable string) string {
	now := time.Now()
	return now.Format("2006-01-02")
}
