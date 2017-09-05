package go_commons

import (
	"testing"
	"time"
)

func Test_traceTime(t *testing.T) {
	test()
}

func test() {
	time.Sleep(time.Second)
	defer traceTime()
}
