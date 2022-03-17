package slices

// Equal reports whether two slices are equal: the same length and all elements equal.
func Equal[T comparable](s1, s2 []T) bool {
	return EqualFunc(s1, s2, func(left T, right T) bool {
		return left == right
	})
}

// EqualFunc reports whether two slices are equal using a comparison function on each pair of elements.
func EqualFunc[T1, T2 any](s1 []T1, s2 []T2, f func(T1, T2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, e := range s1 {
		if !f(e, s2[i]) {
			return false
		}
	}

	return true
}

// Index returns the index of the first occurrence of v in s, or -1 if not present.
func Index[T comparable](s []T, v T) int {
	return IndexFunc(s, func(e T) bool {
		return e == v
	})
}

// IndexFunc returns the index into s of the first element satisfying f(c), or -1 if none do.
func IndexFunc[T any](s []T, f func(T) bool) int {
	for i, e := range s {
		if f(e) {
			return i
		}
	}

	return -1
}

// Contains reports whether v is present in s.
func Contains[T comparable](s []T, v T) bool {
	return Index(s, v) != -1
}

// ContainsFunc reports whether v is present in s using given f() as predicate.
func ContainsFunc[T any](s []T, f func(T) bool) bool {
	return IndexFunc(s, f) != -1
}

// Compact replaces consecutive runs of equal elements with a single copy.
func Compact[T comparable](s []T) []T {
	return CompactFunc(s, func(left T, right T) bool {
		return left == right
	})
}

// CompactFunc is like Compact, but uses a comparison function.
func CompactFunc[T any](s []T, f func(T, T) bool) []T {
	var r []T

	for _, e1 := range s {
		found := false

		for _, e2 := range r {
			if f(e1, e2) {
				found = true
				break
			}
		}

		if !found {
			r = append(r, e1)
		}
	}

	return r
}

// Map creates a new slice with contains every element of s mapped using f()
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))

	for i, e := range s {
		r[i] = f(e)
	}

	return r
}

// Filter creates a new slice with element of s that satisfy the predicate f()
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T

	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}

	return r
}
