package fp

func GetOr[V any](values []V, at int, defaultValue V) V {
	if len(values) > at {
		return values[at]
	}
	return defaultValue
}
