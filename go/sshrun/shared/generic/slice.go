package generic

import (
	"errors"
)

func Map[V any, W any](slice []V, fn func(V) W) []W {
	tmp := make([]W, len(slice))
	for i, s := range slice {
		tmp[i] = fn(s)
	}
	return tmp
}

func Filter[V any](slice []V, predicate func(V) bool) []V {
	var filtered []V
	for _, s := range slice {
		if predicate(s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func Single[V any](values []V, predicate func(V) bool) (*V, error) {
	for _, value := range values {
		if predicate(value) {
			return &value, nil
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
	for _, item := range values {
		if item == value {
			return true
		}
	}
	return false
}

func All[V comparable](values []V, value V) bool {
	for _, item := range values {
		if item != value {
			return false
		}
	}
	return true
}

func Intersect[V comparable](values []V, otherValues []V) bool {
	for _, item := range values {
		for _, otherItem := range otherValues {
			if item == otherItem {
				return true
			}
		}
	}
	return false
}

func GetOr[V any](slice []V, at int, defaultValue V) V {
	if len(slice) > at {
		return slice[at]
	}
	return defaultValue
}
