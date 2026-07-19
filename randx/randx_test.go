package randx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	n := Int(100)
	assert.True(t, n >= 0 && n < 100)

	assert.Equal(t, 0, Int(0))
	assert.Equal(t, 0, Int(-1))
}

func TestIntRange(t *testing.T) {
	n := IntRange(10, 20)
	assert.True(t, n >= 10 && n < 20)

	assert.Equal(t, 5, IntRange(5, 5))
	assert.Equal(t, 5, IntRange(5, 3))
}

func TestFloat64(t *testing.T) {
	f := Float64()
	assert.True(t, f >= 0.0 && f < 1.0)
}

func TestBool(t *testing.T) {
	// 只是测试不 panic，概率性结果不好断言
	_ = Bool()
}

func TestChoice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	v := Choice(s)
	assert.Contains(t, s, v)

	// 空切片
	var zero int
	assert.Equal(t, zero, Choice([]int{}))
}

func TestShuffle(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	original := make([]int, len(s))
	copy(original, s)

	Shuffle(s)

	// 长度不变
	assert.Equal(t, 5, len(s))

	// 元素还在，只是顺序可能变了
	sumOriginal := 0
	for _, v := range original {
		sumOriginal += v
	}
	sumShuffled := 0
	for _, v := range s {
		sumShuffled += v
	}
	assert.Equal(t, sumOriginal, sumShuffled)
}

func TestString(t *testing.T) {
	s := String(10)
	assert.Equal(t, 10, len(s))

	s = String(0)
	assert.Equal(t, "", s)
}

func TestStringWithCharset(t *testing.T) {
	s := StringWithCharset(8, "abc")
	assert.Equal(t, 8, len(s))

	// 所有字符都来自 charset
	for _, c := range s {
		assert.Contains(t, "abc", string(c))
	}

	assert.Equal(t, "", StringWithCharset(5, ""))
}

func TestHex(t *testing.T) {
	h := Hex(4)
	assert.Equal(t, 8, len(h))

	h = Hex(0)
	assert.Equal(t, "", h)
}
