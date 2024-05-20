package fp

func IsEqualFn[V comparable](value V) func(other V) bool {
	return func(other V) bool { return value == other }
}

func IsNotEqualFn[V comparable](value V) func(other V) bool {
	return func(other V) bool { return value != other }
}
