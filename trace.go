package go_commons

import (
	"fmt"
	"time"
)

func traceTime() func() {
	now := time.Now()
	return func() {
		fmt.Println(time.Since(now))
	}
}
