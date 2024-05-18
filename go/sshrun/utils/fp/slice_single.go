package fp

import (
	"errors"
)

func Single[V any](values []V, predicate func(V) bool) (*V, error) {
	for _, v := range values {
		if predicate(v) {
			return &v, nil
		}
	}
	return nil, errors.New("couldn't find value")
}

func LastOr[V any](values []V, defaultV V) V {
	if len(values) == 0 {
		return defaultV
	}
	return values[len(values)-1]
}

func GetOr[V any](values []V, at int, defaultValue V) V {
	if len(values) > at {
		return values[at]
	}
	return defaultValue
}
