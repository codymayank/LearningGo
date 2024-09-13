package loops

import "fmt"

func simple_loop() {
	fmt.Println("Simple loop")
	for i := 0; i < 5; i++ {
		println(i)
	}
}

func range_loop() {
	fmt.Println("Range loop")
	s := []int{1, 2, 3}
	for i, v := range s {
		println(i, v)
	}
}

func Exec() {
	simple_loop()
	range_loop()
}
