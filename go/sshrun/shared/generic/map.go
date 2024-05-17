package generic

func KeyOr[V comparable, K comparable](val V, theMap map[K]V, defaultKey K) K {
	for k, v := range theMap {
		if v == val {
			return k
		}
	}
	return defaultKey
}
