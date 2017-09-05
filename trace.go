package go_commons

import (
	"fmt"
	"time"
)

func TraceTime() func() {
	now := time.Now()
	return func() {
		fmt.Println(time.Since(now))
	}
}
