package shared

import (
	"errors"
)

func Map[V any, W any](values []V, fn func(V) W) []W {
	tmp := make([]W, len(values))
	for i, value := range values {
		tmp[i] = fn(value)
	}
	return tmp
}

func Filter[V any](values []V, predicate func(V) bool) []V {
	var filtered []V
	for _, value := range values {
		if predicate(value) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func LastOr[V any](values []V, defaultV V) V {
	if len(values) == 0 {
		return defaultV
	}
	return values[len(values)-1]
}

func Single[V any](values []V, predicate func(V) bool) (*V, error) {
	for _, value := range values {
		if predicate(value) {
			return &value, nil
		}
	}
	return nil, errors.New("couldn't find value")
}

func ReplaceIf[V any](value *V, newValue V, predicate func(V) bool) {
	if predicate(newValue) {
		*value = newValue
	}
}

func SafeSlice[V any](slice []V, at int, defaultValue V) V {
	if len(slice) > at {
		return slice[at]
	}
	return defaultValue
}

func KeyFor[V comparable, K comparable](val V, theMap map[K]V, defaultKey K) K {
	for k, v := range theMap {
		if v == val {
			return k
		}
	}
	return defaultKey
}

func ValueFor[V comparable, K comparable](key K, theMap map[K]V) V {
	return theMap[key]
}
