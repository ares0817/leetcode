package problem_8

import (
	"strings"
)

func myAtoi(str string) int {

	intMax := int32(^uint32(0) >> 1)
	intMin := ^intMax

	var num int64 = 0
	flag := true

	str = strings.TrimLeft(str, " ")
	if str == "" {
		return 0
	}
	if str[0] == 43 {
		str = str[1:]
	} else if str[0] == 45 {
		flag = false
		str = str[1:]
	}
	str = strings.TrimLeft(str, "0")
	for idx, c := range str {
		if idx == 12 {
			break
		}
		if c < 48 || c > 57 {
			break
		}
		num = num * 10 + int64(c) - 48
	}

	if !flag {
		num = 0 - num
	}

	if num > int64(intMax) {
		return int(intMax)
	}

	if num < int64(intMin) {
		return int(intMin)
	}

	return int(num)
}