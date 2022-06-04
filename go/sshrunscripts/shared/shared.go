package shared

func Map[T any, S any](arr []T, f func(T) S) []S {
	tmp := make([]S, len(arr))
	for i, v := range arr {
		tmp[i] = f(v)
	}
	return tmp
}

func Filter[V any](vs []V, predicate func(V) bool) []V {
	var filtered []V
	for _, v := range vs {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
