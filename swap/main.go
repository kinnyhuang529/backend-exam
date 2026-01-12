package main

import (
	"fmt"
)

func swap[T any](a, b T) {
	switch va := any(a).(type) {
	case *int:
		vb := any(b).(*int)
		*va, *vb = *vb, *va
	default:
		panic("不支援的類型交換")
	}
}

func main() {
	a := 10
	b := 20

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)

	swap(&a, &b)

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)
}
