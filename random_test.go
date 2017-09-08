package go_commons

import (
	"testing"
	"extra/testify/assert"
	"fmt"
)

func Test_RandomInt(t *testing.T) {
	randInt := RandInt(10, 20)
	fmt.Println(randInt)
	assert.Equal(t, true, randInt > 10 && randInt < 20, "错误")
}
