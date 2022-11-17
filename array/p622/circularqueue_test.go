package p622

import "fmt"

func Example_one() {
	q := NewMyCircularQueue(3)
	fmt.Printf("%t,%t,%t,%t,%d,%t,%t,%t,%d",
		q.EnQueue(1),
		q.EnQueue(2),
		q.EnQueue(3),
		q.EnQueue(4),
		q.Rear(),
		q.IsFull(),
		q.DeQueue(),
		q.EnQueue(4),
		q.Rear())
	// Output: true,true,true,false,3,true,true,true,4
}
