package utils

import (
	"math/rand/v2"
	"strings"
	"time"
)

type CharSet string

const (
	Digits     CharSet = "0123456789"
	AlphaUpper CharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphaLower CharSet = "abcdefghijklmnopqrstuvwxyz"
	AlphaAll   CharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

var (
	rs = rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 1))
)

func RandomCode(length int, charset ...CharSet) string {
	var chars string
	for _, c := range charset {
		chars += string(c)
	}
	code := make([]byte, length)
	for i := range code {
		code[i] = chars[rs.IntN(len(chars))]
	}
	return string(code)
}

func RandomUserCode() string {
	code := RandomCode(16, AlphaUpper, Digits)
	var result strings.Builder
	for i := 0; i < len(code); i += 4 {
		if i > 0 {
			result.WriteString("-")
		}
		result.WriteString(code[i : i+4])
	}
	return result.String()
}
