package slices

// FilterAll фильтрует слайс, пропуская только значения, проходящие все фильтры.
// Работает как оператор &&.
func FilterAll[T any](s []T, match ...func(i int, v T) bool) []T {
	out := make([]T, 0, len(s))
	for i, v := range s {
		if !matchAll(i, v, match) {
			continue
		}
		out = append(out, v)
	}
	return out
}

// matchAll работает как оператор &&, то есть возвращает первую ложную проверку.
func matchAll[T any](i int, v T, match []func(i int, v T) bool) bool {
	for _, m := range match {
		if !m(i, v) {
			return false
		}
	}
	return true
}

// FilterAny фильтрует слайс, пропуская только значения, проходящие все фильтры.
// Работает как оператор ||.
func FilterAny[T any](s []T, match ...func(i int, v T) bool) []T {
	out := make([]T, 0, len(s))
	for i, v := range s {
		if !matchAny(i, v, match) {
			continue
		}
		out = append(out, v)
	}
	return out
}

// matchAll  работает  как  оператор  ||,  то есть  возвращает  первую  истинную
// проверку.
func matchAny[T any](i int, v T, match []func(i int, v T) bool) bool {
	for _, m := range match {
		if m(i, v) {
			return true
		}
	}
	return false
}
