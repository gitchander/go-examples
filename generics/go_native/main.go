package main

// https://blog.logrocket.com/understanding-generics-go-1-18/

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Signed integers
type Signed interface {
	int | int8 | int16 | int32 | int64
}

// Unsigned integers
type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Floats interface {
	float32 | float64
}

type Integers interface {
	Signed | Unsigned
}

type Ordered interface {
	Integers | Floats
}

//------------------------------------------------------------------------------
// comparable
func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func sum[T Ordered](as ...T) T {
	var sum T
	for _, a := range as {
		sum += a
	}
	return sum
}

func Keys[K comparable, V any](m map[K]V) []K {
	ks := make([]K, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

type Stringer interface {
	String() string
}

func toString[T Stringer](x T) string {
	return x.String()
}

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) User {
	return User{
		Name: name,
		Age:  age,
	}
}

func (u User) String() string {
	return fmt.Sprintf("(%s %d)", u.Name, u.Age)
}

func Contains[T comparable](a T, bs []T) bool {
	for _, b := range bs {
		if a == b {
			return true
		}
	}
	return false
}

func CloneSlice[T any](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	return b
}

func QuoRem[T Integers](a, b T) (quo, rem T) {
	quo = a / b
	rem = a % b
	return
}

func Reverse[T any](a []T) []T {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i, j = i+1, j-1
	}
	return a
}

func testMore() {
	fmt.Println(min(56.9, 2.78))
	fmt.Println(sum[int]())
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	ks := Keys(m)
	fmt.Println(ks)

	u := NewUser("Isaac", 33)
	fmt.Println(toString(u))

	fmt.Println(Contains(72, []int{1, 2, 3, 71, 72, 20, 31}))

	fmt.Println(CloneSlice([]int{1, 2, 3}))

	fmt.Println(QuoRem(7, 3))
	fmt.Println(Reverse([]int{1, 2, 3, 4}))

}

func main() {
	testMin()
}
