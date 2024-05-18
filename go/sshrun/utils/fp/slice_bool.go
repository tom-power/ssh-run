package fp

func Any[V comparable](values []V, value V) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func All[V comparable](values []V, value V) bool {
	for _, v := range values {
		if v != value {
			return false
		}
	}
	return true
}

func Intersect[V comparable](values []V, otherValues []V) bool {
	for _, v := range values {
		for _, otherItem := range otherValues {
			if v == otherItem {
				return true
			}
		}
	}
	return false
}
