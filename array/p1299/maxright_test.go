package p1299

import "fmt"

func Example_one() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(arr))
	// Output: [18 6 6 6 1 -1]
}

func Example_two() {
	arr := []int{400}
	fmt.Println(replaceElements(arr))
	// Output: [-1]
}
