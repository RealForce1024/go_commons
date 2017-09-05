package go_commons

import (
	"testing"
	"time"
)

func Test_TraceTime(t *testing.T) {
	test()
}

func test() {
	time.Sleep(time.Second)
	defer TraceTime()
}
