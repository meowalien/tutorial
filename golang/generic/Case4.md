# Case4

## Queue[T]

```go
package main

func NewQueue[T any](i int) Queue[T] {
	return make(chan T, i)
}

type Queue[T any] chan T

func (q Queue[T]) Push(val T) {
	q <- val
}
func (q Queue[T]) Pop() T {
	return <-q
}

func (q Queue[T]) TryPop() (T, bool) {
	select {
	case val := <-q:
		return val, true
	default:
		var zero T
		return zero, false
	}
}

```
