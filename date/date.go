package date

import (
	"strings"
	"time"
)

func GetCurrentDate() string {
	current := time.Now().String()
	indexOfSpace := strings.Index(current, " ")
	dateFromCurrent := current[:indexOfSpace]
	return dateFromCurrent
}
