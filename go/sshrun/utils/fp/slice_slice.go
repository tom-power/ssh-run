package fp

func Map[V any, W any](values []V, fn func(V) W) []W {
	tmp := make([]W, len(values))
	for i, v := range values {
		tmp[i] = fn(v)
	}
	return tmp
}

func Filter[V any](values []V, predicate func(V) bool) []V {
	var filtered []V
	for _, v := range values {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
