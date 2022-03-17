package slices

import "testing"

func TestEqual(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{3, 2, 1}
	s3 := []int{1, 2, 3, 4}

	if !Equal(s1, s1) {
		t.Fail()
	}
	if Equal(s1, s2) {
		t.Fail()
	}
	if Equal(s1, s3) {
		t.Fail()
	}
}

func TestEqualFunc(t *testing.T) {
	type user struct {
		name string
	}

	s1 := []user{{name: "Aloïs"}}
	s2 := []user{{name: "Aloïs"}, {name: "Creekorful"}}
	s3 := []user{{name: "Creekorful"}, {name: "Aloïs"}}

	f := func(left user, right user) bool {
		return left.name == right.name
	}

	if !EqualFunc(s1, s1, f) {
		t.Fail()
	}
	if EqualFunc(s1, s2, f) {
		t.Fail()
	}
	if EqualFunc(s1, s3, f) {
		t.Fail()
	}
	if EqualFunc(s2, s3, f) {
		t.Fail()
	}
}

func TestIndex(t *testing.T) {
	s := []int{1, 2, 3}

	if Index(s, 2) != 1 {
		t.Fail()
	}

	if Index(s, 12) != -1 {
		t.Fail()
	}
}

func TestIndexFunc(t *testing.T) {
	type user struct {
		name string
	}

	s := []user{{name: "Aloïs"}}

	f := func(value string) func(u user) bool {
		return func(u user) bool {
			return u.name == value
		}
	}

	if IndexFunc(s, f("Aloïs")) != 0 {
		t.Fail()
	}

	if IndexFunc(s, f("other")) != -1 {
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	s := []int{1, 2, 3}

	if !Contains(s, 2) {
		t.Fail()
	}

	if Contains(s, 12) {
		t.Fail()
	}
}

func TestContainsFunc(t *testing.T) {
	type user struct {
		name string
	}

	s := []user{{name: "Aloïs"}}

	f := func(value string) func(u user) bool {
		return func(u user) bool {
			return u.name == value
		}
	}

	if !ContainsFunc(s, f("Aloïs")) {
		t.Fail()
	}

	if ContainsFunc(s, f("other")) {
		t.Fail()
	}
}

func TestCompact(t *testing.T) {
	s := []int{2, 3, 4, 2, 1, 5, 1}

	actual := Compact(s)
	expected := []int{2, 3, 4, 1, 5}

	if !Equal(actual, expected) {
		t.Fail()
	}
}

func TestCompactFunc(t *testing.T) {
	type user struct {
		name string
	}

	s := []user{{name: "Aloïs"}, {name: "Creekorful"}, {name: "Aloïs"}}

	f := func(left user, right user) bool {
		return left.name == right.name
	}

	actual := CompactFunc(s, f)
	expected := []user{{name: "Aloïs"}, {name: "Creekorful"}}

	if !Equal(actual, expected) {
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	type user struct {
		name string
	}

	s := []user{{name: "Aloïs"}, {name: "Creekorful"}}

	actual := Map(s, func(u user) string { return u.name })
	expected := []string{"Aloïs", "Creekorful"}

	if !Equal(actual, expected) {
		t.Fail()
	}
}

func TestFilter(t *testing.T) {
	type user struct {
		name string
	}

	s := []user{{name: "Aloïs"}, {name: "Creekorful"}}

	actual := Filter(s, func(u user) bool { return u.name == "Aloïs" })
	expected := []user{{name: "Aloïs"}}

	if !Equal(actual, expected) {
		t.Fail()
	}
}
