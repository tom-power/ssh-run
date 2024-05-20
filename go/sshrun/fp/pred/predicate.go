package pred

func IsEqualFn[V comparable](value V) func(other V) bool {
	return func(other V) bool { return value == other }
}

func IsNotEmpty(value string) bool {
	return len(value) > 0
}
