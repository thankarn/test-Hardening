package utils

func IIf[T any](condition bool, x, y T) T {
	if condition {
		return x
	} else {
		return y
	}
}
