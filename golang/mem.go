package main

import (
	"fmt"
	"unsafe"
)

type Employee struct {
   IsActive  bool
   Age       int64
   IsMarried bool
   Name      string
   Photo     float32
}

func main() {
	var employee Employee
	fmt.Printf("Size of %T struct; %d bytes", employee, unsafe.Sizeof(employee))
}
