
# Case1

#### Old:
```go
package main

import "fmt"

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}

func main() {
    // Initialize a map for the integer values
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))
}
```

* New
```go
package main

import "fmt"
// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

func main() {
    // Initialize a map for the integer values
    ints := map[string]int64{
    "first":  34,
    "second": 12,
    }
    
    // Initialize a map for the float values
    floats := map[string]float64{
    "first":  35.98,
    "second": 26.99,
    }
    
	//fmt.Printf("Generic Sums: %v and %v\n",
    //SumIntsOrFloats[string, int64](ints),
    //SumIntsOrFloats[string, float64](floats))
	
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
}
```

* With type constraint
```go
package main

import "fmt"

type Number interface {
	int64 | float64
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

func main() {
    // Initialize a map for the integer values
    ints := map[string]int64{
    "first":  34,
    "second": 12,
    }
    
    // Initialize a map for the float values
    floats := map[string]float64{
    "first":  35.98,
    "second": 26.99,
    }
	
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
}
```


```go
package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](a T, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	x := Min(1, 2)
	fmt.Println(x)
	y := Min(1.342653343222, 1.342653343223)
	fmt.Println(y)
	z := Min("dfsde", "Afdskldfjls")
	fmt.Println(z)
}

```


```go
package main

import (
	"fmt"
)

// IFF if yes return a else b
func IFF[T any](yes bool, a, b T) T {
	if yes {
		return a
	}
	return b
}

func main() {
	a := -1
	fmt.Println(IFF(a > 0, a, 0), 0)
	fmt.Println(IFF(a > 0, func() string { return "a" }, func() string { return "B" })())
}

```

```go
package main
import "strconv"

type Number interface {
	int | int32 | int64 | uint32 | uint64 | float64
}

func Str2Number[N Number](strNumber string) (N, error) {
	var num N
	switch (interface{})(num).(type) {
	case int:
		cn, err := strconv.Atoi(strNumber)
		return N(cn), err
	case int32:
		cn, err := strconv.ParseInt(strNumber, 10, 32)
		return N(cn), err
	case int64:
		cn, err := strconv.ParseInt(strNumber, 10, 64)
		return N(cn), err
	case uint32:
		cn, err := strconv.ParseUint(strNumber, 10, 32)
		return N(cn), err
	case uint64:
		cn, err := strconv.ParseUint(strNumber, 10, 64)
		return N(cn), err
	case float64:
		cn, err := strconv.ParseFloat(strNumber, 64)
		return N(cn), err
	}
	return 0, nil
}

```