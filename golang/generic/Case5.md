# Case4

## MissionChain[T]

* Old
```go

type Request struct {
	SomeDataA string
	SomeDataB int
	SomeDataC bool
}

func ProcessA(ctx context.Context, s *Request) error {
	fmt.Println("ProcessA")
	if !s.SomeDataC {
		s.SomeDataA = "test"
	}
	time.Sleep(2 * time.Second)
	return nil
}
func ProcessB(ctx context.Context, s *Request) error {
	fmt.Println("ProcessB")
	if s.SomeDataA == "test" {
		s.SomeDataB = 1
	}
	return nil
}

func main() {
	var request = &Request{}
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	err := ProcessA(ctx, request)
	if err != nil {
		panic(err)
	}
	select {
	case <-ctx.Done():
	default:
		err = ProcessB(ctx, request)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("result: ", request)

}
```

* New
```go
package main

import (
	"context"
	"fmt"
	"time"
)

type Request struct {
	SomeDataA string
	SomeDataB int
	SomeDataC bool
}

func ProcessA(ctx context.Context, s *Request) error {
	fmt.Println("ProcessA")
	if !s.SomeDataC {
		s.SomeDataA = "test"
	}
	time.Sleep(2 * time.Second)
	return nil
}
func ProcessB(ctx context.Context, s *Request) error {
	fmt.Println("ProcessB")
	if s.SomeDataA == "test" {
		s.SomeDataB = 1
	}
	return nil
}

type layer[T any] func(ctx context.Context, t T) error

func (l layer[T]) and(nl layer[T]) layer[T] {
	if l == nil {
		return nl
	}
	return func(ctx context.Context, t T) error {
		err := (l)(ctx, t)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			err = (nl)(ctx, t)
			if err != nil {
				return err
			}
			return nil
		}
	}
}

type MissionChain[T any] struct {
	l   *layer[T]
	ctx context.Context
	val T
}

func (i *MissionChain[T]) Run() error {
	return (*i.l)(i.ctx, i.val)
}

func (i *MissionChain[T]) Do(nl layer[T]) *MissionChain[T] {
	if i.l == nil {
		i.l = &nl
		return i
	}
	*i.l = (*i.l).and(nl)
	return i
}

func NewMissionChain[T any](ctx context.Context, val T) MissionChain[T] {
	return MissionChain[T]{
		ctx: ctx,
		val: val,
	}
}

func main() {
	var request = &Request{}
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	inj := NewMissionChain(ctx, request)

	inj.Do(ProcessA).Do(ProcessB)

	err := inj.Run()
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	fmt.Println("result: ", request)
}

```