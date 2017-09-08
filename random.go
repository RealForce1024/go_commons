package go_commons

import (
	"math/rand"
	"time"
)

func RandInt(floor, ceiling int) int {
	rand.Seed(time.Now().UnixNano())
	if floor == ceiling {
		return floor
	} else if floor > ceiling {
		floor, ceiling = ceiling, floor
	}
	return rand.Intn(ceiling-floor) + floor
}
