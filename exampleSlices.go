package main

import "fmt"

// Function was defined in godoc...
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func PrintSlice(slice []byte) {
	fmt.Printf("slice[:]   %d\n", slice[:])   // slice[:]  [1 2 3 4 5 6]
	fmt.Printf("slice[0:]  %d\n", slice[0:])  // slice[0:] [1 2 3 4 5 6]
	fmt.Printf("slice[:0]  %d\n", slice[:0])  // slice[:0] []
	fmt.Printf("slice[0:0] %d\n", slice[0:0]) // slice[0:0] []
	fmt.Printf("slice[1:]  %d\n", slice[1:])  // slice[1:] [2 3 4 5 6]
	fmt.Printf("slice[:1]  %d\n", slice[:1])  // slice[:1] [1]
	fmt.Printf("slice[:2]  %d\n", slice[:2])  // slice[:1] [1 2]
	fmt.Printf("slice[3:5] %d\n", slice[3:5]) // slice[3:5] [4 5]
}

func RemoveByte(slice []byte, index int) (int, []byte) {
	var rc int
	slice_len := len(slice)
	if index > slice_len {
		fmt.Println("Index out of a range.")
		rc = -1
	} else {
		slice = append(slice[:index], slice[index+1:]...)
		rc = 0
	}
	return rc, slice
}

func main() {
	fmt.Println("Below is presented some information about slices in Goland.")

	var slice_1, slice_2, slice_3 []byte
	slice_4 := []byte{1, 2, 3, 4, 5, 6}
	slice_1 = make([]byte, 3, 5)
	slice_1 = append(slice_1, 6, 6, 6) // append() automatically extend a slice
	slice_2 = []byte{1, 2, 3, 4}
	fmt.Println(slice_1)
	fmt.Println(slice_2)
	slice_1 = AppendByte(slice_1, 9, 9, 9)
	slice_3 = AppendByte(slice_1, slice_2...)
	fmt.Println(slice_3)
	PrintSlice(slice_4)
	fmt.Println(slice_4)
	_, slice_4 = RemoveByte(slice_4, 8)
	fmt.Println(slice_4)
}
