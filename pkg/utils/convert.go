package utils

import (
	"strconv"
	"strings"
)

func StrToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}

func StrToUint64(str string) uint64 {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func StrToFloat64(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return num
}

func StrToBool(str string) bool {
	num, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return num
}

func IntToStr(num int) string {
	return strconv.Itoa(num)
}

func Uint64ToStr(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func NumbersJoin[T int | int64 | uint64 | float64](nums []T, separator string) string {
	var strs = make([]string, 0, len(nums))
	for _, num := range nums {
		switch v := any(num).(type) {
		case int:
			strs = append(strs, strconv.Itoa(v))
		case int64:
			strs = append(strs, strconv.FormatInt(v, 10))
		case uint64:
			strs = append(strs, strconv.FormatUint(v, 10))
		case float64:
			strs = append(strs, strconv.FormatFloat(v, 'f', -1, 64))
		}
	}
	return strings.Join(strs, separator)
}
