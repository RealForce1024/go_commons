package go_commons

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_compare(t *testing.T) {
	//result := compare([]int{1, 2, 3}, []int{2, 3, 1})
	result := compare([]int{1, 2, 3}, []int{1,2,3})
	assert.Equal(t, result, true, "对比错误")
}
