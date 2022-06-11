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

func Single[V any](values []V, predicate func(V) bool) (*V, error) {
	for _, value := range values {
		if predicate(value) {
			return &value, nil
		}
	}
	return nil, errors.New("couldn't find value")
}

func Unique(strSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func UpdateIf[V any](value *V, newValue V, predicate func(V) bool) {
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
