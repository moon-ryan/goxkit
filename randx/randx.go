package randx

import (
	"math/rand"
	"time"
)

func init() {
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))
}

// Int returns randx int in [0, max).
// 返回 [0, max) 范围内的随机整数。
func Int(max int) int {
	if max <= 0 {
		return 0
	}
	return rand.Intn(max)
}

// IntRange returns randx int in [0, max).
// 返回 [min, max) 范围内的随机整数。
func IntRange(min, max int) int {
	if min >= max {
		return min
	}
	return min + rand.Intn(max-min)
}

// Float64 returns randx float64 in [0.0, 1.0).
// 返回 [0.0, 1.0) 范围内的随机整数。
func Float64() float64 {
	return rand.Float64()
}

// Bool returns randx boolean.
// 返回随机的 boolean。
func Bool() bool {
	return rand.Intn(2) == 0
}

// Choice returns randx element from slice.
// 从切片中返回随机元素。
func Choice[T any](s []T) T {
	var zero T
	if len(s) == 0 {
		return zero
	}
	return s[rand.Intn(len(s))]
}

// Shuffle Reorder the slices randomly and shuffle the elements.
// 打乱切片顺序，把元素随机重新排列。
func Shuffle[T any](s []T) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

// String generates randx string of given length.
// 生成指定长度的随机字符串。
func String(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// StringWithCharset generates randx string from custom charset.
// 从自定义字符集中生成随机字符串。
func StringWithCharset(n int, charset string) string {
	if n <= 0 || len(charset) == 0 {
		return ""
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Hex generates randx hex string of given byte length.
// 生成指定字节长度的随机十六进制字符串。
func Hex(n int) string {
	const hexChars = "0123456789abcdef"
	if n <= 0 {
		return ""
	}
	b := make([]byte, n*2)
	for i := range b {
		b[i] = hexChars[rand.Intn(len(hexChars))]
	}
	return string(b)
}
