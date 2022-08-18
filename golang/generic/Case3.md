# Case3

## Walk[T]

#### Old:

```go
package main

import "fmt"

type Root struct {
	L1    *L1    `json:"l1"`
	Data1 string `json:"data1"`
}
type L1 struct {
	L2    *L2    `json:"l2"`
	Data2 string `json:"data2"`
}
type L2 struct {
	L3    *L3    `json:"l3"`
	Data3 string `json:"data3"`
}
type L3 struct {
	Data4 string `json:"data4"`
}

var root = &Root{
	L1: &L1{
		L2: &L2{
			L3: &L3{
				Data4: "data4",
			},
			Data3: "data3",
		},
		Data2: "data2",
	},
	Data1: "data1",
}

func Walk(x any, f func(any) bool) {
	if f(x) {
		switch x := x.(type) {
		case *Root:
			Walk(x.L1, f)
		case *L1:
			Walk(x.L2, f)
		case *L2:
			Walk(x.L3, f)
		case *L3:
			// nothing to do
		}
	}
}

func main() {
	var l2 *L2

	Walk(root, func(x any) bool {
		switch y := x.(type) {
		case *L2:
			l2 = y
			return false
		}
		return true
	})
	fmt.Println("l2", l2.Data3)
	return
}
```

* New
```go
package main

import "fmt"

type Root struct {
	L1    *L1    `json:"l1"`
	Data1 string `json:"data1"`
}
type L1 struct {
	L2    *L2    `json:"l2"`
	Data2 string `json:"data2"`
}
type L2 struct {
	L3    *L3    `json:"l3"`
	Data3 string `json:"data3"`
}
type L3 struct {
	Data4 string `json:"data4"`
}

var root = &Root{
	L1: &L1{
		L2: &L2{
			L3: &L3{
				Data4: "data4",
			},
			Data3: "data3",
		},
		Data2: "data2",
	},
	Data1: "data1",
}

func main() {
	var l2 *L2
	Walk(root, func(x *L2) bool {
		l2 = x
		return false
	})
	fmt.Println("l2", l2.Data3)
}

func Walk[T any](r any, f func(x T) bool) {
	if a, ok := r.(T); ok {
		if !f(a) {
			return
		}
	}
	switch x := r.(type) {
	case *Root:
		Walk(x.L1, f)
	case *L1:
		Walk(x.L2, f)
	case *L2:
		Walk(x.L3, f)
	case *L3:
		// nothing to do
	}
}
```

* Advanced
```go
package main

import "fmt"

type Root struct {
	L1    *L1    `json:"l1"`
	Data1 string `json:"data1"`
}
type L1 struct {
	L2    *L2    `json:"l2"`
	Data2 string `json:"data2"`
}
type L2 struct {
	L3    *L3    `json:"l3"`
	Data3 string `json:"data3"`
}
type L3 struct {
	L3_L2 *L2    `json:"l3>l2"`
	Data4 string `json:"data4"`
}

var root = &Root{
	L1: &L1{
		L2: &L2{
			L3: &L3{
				Data4: "data4",
				L3_L2: &L2{
					Data3: "L3_L2",
				},
			},
			Data3: "data3",
		},
		Data2: "data2",
	},
	Data1: "data1",
}


func FindAll[T any](r any) (ans []T) {
	WalkNew(r, func(x T) bool {
		ans = append(ans, x)
		return true
	})
	return
}

func WalkNew[T any](r any, f func(x T) bool) {
	if a, ok := r.(T); ok {
		if !f(a) {
			return
		}
	}

	//fmt.Println("x: ", r == nil)
	//fmt.Println("x: ", r)
	switch x := r.(type) {
	case *Root:
		if x != nil {
			WalkNew(x.L1, f)
		}
	case *L1:
		if x != nil {
			WalkNew(x.L2, f)
		}
	case *L2:
		if x != nil {
			WalkNew(x.L3, f)
		}
	case *L3:
		if x != nil {
			WalkNew(x.L3_L2, f)
		}
	}
}

func main() {
	l2 := FindAll[*L2](root)
	for i, l := range l2 {
		fmt.Println(i, l.Data3)
	}
}
```
