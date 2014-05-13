package euler

import (
	"fmt"
	"time"
)

func P019() string {
	start, _ := time.Parse("2 Jan 2006", "1 Jan 1901")
	end, _ := time.Parse("2 Jan 2006", "31 Dec 2000")
	result := 0

	for t := start; t.Unix() <= end.Unix(); t = t.AddDate(0, 1, 0) {
		if t.Weekday() == time.Sunday {
			result++
		}
	}
	return fmt.Sprintf("%v", result)
}
