package main

import "fmt"

type Int interface {
	~int | ~uint
}

func IsSigned[T Int](n T) {
	switch (interface{})(n).(type) {
	case int:
		fmt.Println("signed")
	default:
		fmt.Println("unsigned")
	}
}

func main() {
	type MyInt int
	IsSigned(1)
	IsSigned(MyInt(1))
}

// Output:
// signed
// unsigned
