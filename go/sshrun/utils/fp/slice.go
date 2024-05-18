package fp

import (
	"errors"
)

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

func GetOr[V any](values []V, at int, defaultValue V) V {
	if len(values) > at {
		return values[at]
	}
	return defaultValue
}
