package main

import (
	"fmt"

	luannh "../luannh"
)

func main() {
	fmt.Printf("Hello, world.\n")
	var a = true
	fmt.Println(a)
	b := "klgt"
	fmt.Println(b)
	animals := []luannh.Animal{luannh.Dog{}, luannh.Cat{}, luannh.Llama{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	s := luannh.NewStack()
	s.Push("100")
	s.Push("luannh")
	s1 := s.Pop()
	s2 := s.Pop()
	fmt.Println(s1)
	fmt.Println(s2)

	var v = 102
	value, ok := v.(int)
	fmt.Println(value)
	fmt.Print(ok)

	x := 1.5
	square(&x)
	fmt.Println(x)
}

func printValue(v interface{}) {
	value, ok := v.(int)

	if ok {
		fmt.Printf("The value of v is: %v", v.(int))
		fmt.Println(value)
	} else {
		fmt.Println("loi roi")
	}

}
func square(x *float64) {
	*x = *x * *x
}
