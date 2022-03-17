# slices

Go 1.18 utility functions when working with slice.

## Equal

Equal reports whether two slices are equal: the same length and all elements equal.

```go
package main

import (
	"fmt"
	"github.com/creekorful/slices"
)

func main() {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"a", "b"}

	fmt.Println(slices.Equal(s1, s2)) // False
}
```

## EqualFunc

EqualFunc reports whether two slices are equal using a comparison function on each pair of elements.

```go
package main

import (
	"fmt"
	"github.com/creekorful/slices"
)

func main() {
	type user struct {
		name string
	}

	s1 := []user{{name: "Aloïs"}, {name: "Creekorful"}}
	s2 := []user{{name: "Aloïs"}, {name: "Creekorful"}}

	f := func(left user, right user) bool {
		return left.name == right.name
	}

	fmt.Println(slices.EqualFunc(s1, s2, f)) // True
}
```

## Index

Index returns the index of the first occurrence of v in s, or -1 if not present.

```go
package main

import (
	"fmt"
	"github.com/creekorful/slices"
)

func main() {
	s := []string{"a", "b", "c"}

	fmt.Println(slices.Index(s, "b")) // 1
	fmt.Println(slices.Index(s, "z")) // -1
}
```

## IndexFunc

IndexFunc returns the index into s of the first element satisfying f(c), or -1 if none do.

```go
package main

import (
	"fmt"
	"github.com/creekorful/slices"
)

func main() {
	type user struct {
		name string
	}

	s1 := []user{{name: "Aloïs"}, {name: "Creekorful"}}
	s2 := []user{{name: "Creekorful"}}

	f := func(e user) bool {
		return e.name == "Aloïs"
	}

	fmt.Println(slices.IndexFunc(s1, f)) // 0
	fmt.Println(slices.IndexFunc(s2, f)) // -1
}
```

## Contains

Contains reports whether v is present in s.

```go
package main

import (
	"fmt"
	"github.com/creekorful/slices"
)

func main() {
	s := []string{"a", "b", "c"}

	fmt.Println(slices.Contains(s, "b")) // True
	fmt.Println(slices.Contains(s, "z")) // False
}
```

## ContainsFunc

ContainsFunc reports whether v is present in s using given f() as predicate.

```go
package main

import (
	"fmt"
	"github.com/creekorful/slices"
)

func main() {
	type user struct {
		name string
	}

	s1 := []user{{name: "Aloïs"}, {name: "Creekorful"}}
	s2 := []user{{name: "Creekorful"}}

	f := func(e user) bool {
		return e.name == "Aloïs"
	}

	fmt.Println(slices.ContainsFunc(s1, f)) // True
	fmt.Println(slices.ContainsFunc(s2, f)) // False
}
```

## Compact

Compact replaces consecutive runs of equal elements with a single copy.

```go
package main

import (
	"fmt"
	"github.com/creekorful/slices"
)

func main() {
	s := []string{"b", "c", "c", "d", "a", "b", "a"}

	fmt.Println(slices.Compact(s)) // b, c, d, a
}
```

## CompactFunc

CompactFunc is like Compact, but uses a comparison function.

```go
package main

import (
	"fmt"
	"github.com/creekorful/slices"
)

func main() {
	type user struct {
		name string
	}

	s := []user{{name: "Creekorful"}, {name: "Aloïs"}, {name: "Aloïs"}, {name: "Creekorful"}, {name: "Aloïs"}}

	f := func(left user, right user) bool {
		return left.name == right.name
	}

	fmt.Println(slices.CompactFunc(s, f)) // {name: "Creekorful"}, {name: "Aloïs"}
}
```

## Map

Map creates a new slice with contains every element of s mapped using f()

```go
package main

import (
	"fmt"
	"github.com/creekorful/slices"
)

func main() {
	s := []int{2, 4, 8, 16}

	f := func(e int) int {
		return e * 2
	}

	fmt.Println(slices.Map(s, f)) // 4, 8, 16, 32
}
```

## Filter

Filter creates a new slice with element of s that satisfy the predicate f()

```go
package main

import (
	"fmt"
	"github.com/creekorful/slices"
)

func main() {
	s := []int{25, 12, 6, 100, 65, 44}

	f := func(e int) bool {
		return e > 50
	}

	fmt.Println(slices.Filter(s, f)) // 100, 65
}
```