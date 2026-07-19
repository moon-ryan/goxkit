package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	assert.Equal(t, "", ToString(nil))
	assert.Equal(t, "hello", ToString("hello"))
	assert.Equal(t, "42", ToString(42))
	assert.Equal(t, "3.14", ToString(3.14))
	assert.Equal(t, "true", ToString(true))
	assert.Equal(t, "123", ToString(int8(123)))
	assert.Equal(t, "456", ToString(uint16(456)))
}

func TestToInt(t *testing.T) {
	n, err := ToInt("42")
	assert.NoError(t, err)
	assert.Equal(t, 42, n)

	n, err = ToInt(3.14)
	assert.NoError(t, err)
	assert.Equal(t, 3, n)

	_, err = ToInt("abc")
	assert.Error(t, err)
}

func TestToInt64(t *testing.T) {
	n, err := ToInt64("100")
	assert.NoError(t, err)
	assert.Equal(t, int64(100), n)
}

func TestToFloat64(t *testing.T) {
	f, err := ToFloat64("3.14")
	assert.NoError(t, err)
	assert.Equal(t, 3.14, f)
}

func TestToBool(t *testing.T) {
	b, err := ToBool("true")
	assert.NoError(t, err)
	assert.True(t, b)

	b, err = ToBool(1)
	assert.NoError(t, err)
	assert.True(t, b)

	b, err = ToBool(0)
	assert.NoError(t, err)
	assert.False(t, b)
}

func TestMustToInt(t *testing.T) {
	assert.Equal(t, 42, MustToInt("42"))
	assert.Panics(t, func() { MustToInt("abc") })
}

func TestMustToInt64(t *testing.T) {
	assert.Equal(t, int64(100), MustToInt64("100"))
	assert.Panics(t, func() { MustToInt64("abc") })
}

func TestMustToFloat64(t *testing.T) {
	assert.Equal(t, 3.14, MustToFloat64("3.14"))
	assert.Panics(t, func() { MustToFloat64("abc") })
}

func TestMustToBool(t *testing.T) {
	assert.True(t, MustToBool("true"))
	assert.Panics(t, func() { MustToBool("abc") })
}
