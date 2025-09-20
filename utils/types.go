package utils

func Contains[T comparable](s []T, v T) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}
	return false
}
