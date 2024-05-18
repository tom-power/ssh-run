package fp

func IsEqual[V comparable](value V) func(other V) bool {
	return func(other V) bool { return value == other }
}
