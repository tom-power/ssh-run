package fp

func IsEqual[V comparable](value V) func(other V) bool {
	return func(other V) bool { return value == other }
}

func IsNotEqual[V comparable](value V) func(other V) bool {
	return func(other V) bool { return value != other }
}

func IsEqualWithIndex[V comparable](value V, index int) func(other V) bool {
	return IsEqual(value)
}
