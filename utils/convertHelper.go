package utils

func Ptr[T any](v T) *T {
	return &v
}

func PtrToString(val *string) string {
	if val == nil {
		return ""
	}

	return *val
}
