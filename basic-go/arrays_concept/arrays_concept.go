package arrays_concept

import (
	"fmt"
)

func declare_arr() {
	var arr = [3]int{1, 2, 3} // array of 3 integers
	fmt.Println(arr)          // [1 2 3]
}

// Exec /* Exec - Here capital "E" meaning this function is exported and can be use by other modules */
func Exec() {
	declare_arr()
}
