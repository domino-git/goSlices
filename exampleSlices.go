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

func main() {
	fmt.Println("Below is presented some information about slices in Goland.")

	var slice_1, slice_2, slice_3 []byte
	slice_1 = make([]byte, 3, 5)
	slice_1 = append(slice_1, 6, 6, 6) // append() automatically extend a slice
	slice_2 = []byte{1, 2, 3, 4}
	fmt.Println(slice_1)
	fmt.Println(slice_2)
	slice_1 = AppendByte(slice_1, 9, 9, 9)
	slice_3 = AppendByte(slice_1, slice_2...)
	fmt.Println(slice_3)
}
