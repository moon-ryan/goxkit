package convert

// MustToInt 转换为 int，错误时会崩溃结束。
func MustToInt(v any) int {
	i, err := ToInt(v)
	if err != nil {
		panic(err)
	}
	return i
}

// MustToInt64 转换为 int64，错误时会崩溃结束。
func MustToInt64(v any) int64 {
	i, err := ToInt64(v)
	if err != nil {
		panic(err)
	}
	return i
}

// MustToFloat64 转换为 float64，错误时会崩溃结束。
func MustToFloat64(v any) float64 {
	f, err := ToFloat64(v)
	if err != nil {
		panic(err)
	}
	return f
}

// MustToBool 转换为 bool，错误时会崩溃结束。
func MustToBool(v any) bool {
	b, err := ToBool(v)
	if err != nil {
		panic(err)
	}
	return b
}
