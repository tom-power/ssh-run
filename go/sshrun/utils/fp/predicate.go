package fp

func ReplaceIf[V any](value *V, newValue V, predicate func(V) bool) {
	if predicate(newValue) {
		*value = newValue
	}
}
