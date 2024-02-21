package main

import "fmt"

/*
In Go, arrays and slices can indeed have similar syntax, especially when initializing them.

  - Arrays when Fixed Size:
    Arrays in Go have a fixed size that is determined at the time of declaration.
    Once the size is set, it cannot be changed.
    Append also does not work on it unless it converted to slice
    It passed as Value type not refernce so any change into func will be reflected into copy of arr inside func only, unless it gets returned and updated to orignal

  - Slices:
    Dynamic Size:
    Slices are dynamic and can change in size during runtime.
    Reference Type:
    When you pass a slice to a function or assign it to another variable, you're working with a reference to the same underlying array.
    NOTE: Append won't reflect as after append a copy of slice gets created so that belongs only to that func scope only and not referencing to actual slice

    Slices can be created by slicing an existing array or another slice.
    arr := [5]int{1, 2, 3, 4, 5} // Array
    convert arr to Slice from index 1 to 3 (4 not included)
    slice1 := arr[1:4]

    append works on slice
    slice := []int{1, 2, 3}
    slice = append(slice, 4, 5) // copy of slice is returned here

    Slices are more flexible and are often preferred over arrays when working with collections of data whose size may change.
*/
func main() {
	// array
	arr := [4]int{1, 2, 3, 4}
	fmt.Println("Arr: ", arr)
	// arr = append(arr, 5) -> error
	modifyArr(arr, 0)
	fmt.Println("modifyArr: ", arr)
	// array is passed as copy not reference so it should be returned from func
	arr = modifyArrAndReturn(arr, 0)
	fmt.Println("modifyArrAndReturn: ", arr)

	//slice
	slice := []int{1, 2, 3, 4}
	slice = append(slice, 5, 6)
	fmt.Println("slice: ", slice)

	// convert array to slice
	slice2 := arr[0:]
	// now elems can be appended to this slice
	slice2 = append(slice2, 5, 6)
	fmt.Println("slice2: ", slice2)

	// modify slice elem directly as it is passed as reference type
	modifySlice(slice, 0) // change 0th elem to -1
	fmt.Println("modifySlice: ", slice)
	// append needs to returned from func as it create copy of slice and return it so it would be treated as local copy inside func
	slice = addElem(slice, 10)
	fmt.Println("addElem: ", slice)

	addElemByRef(&slice, 20)
	fmt.Println("addElemByRef: ", slice)

	// len - represents the no of elements present in slice
	// cap - represents max holding capacity of slice
	// after elem is appended then len gets increment by +1 but cap gets doubled of prev cap in new copy
	// ex. here slice has len = 6 and cap = 8
	slice = []int{1, 2}
	fmt.Println("len1: ", len(slice))
	fmt.Println("cap1: ", cap(slice))
	slice = append(slice, 3, 4, 5)
	fmt.Println("len2: ", len(slice))
	fmt.Println("cap2: ", cap(slice))
	slice = append(slice, 6, 7, 8)
	fmt.Println("len3: ", len(slice))
	fmt.Println("cap3: ", cap(slice))
	slice = append(slice, 9, 10, 11, 12, 13)
	fmt.Println("len4: ", len(slice))
	fmt.Println("cap4: ", cap(slice))
}

func modifyArr(arr [4]int, i int) {
	arr[i] = -1
}

// change value to -1 at given index
func modifyArrAndReturn(arr [4]int, index int) [4]int {
	arr[index] = -1
	return arr
}

// slices work as reference type when passed to func
func modifySlice(slice []int, i int) {
	slice[i] = -1
}

func addElem(slice []int, elem int) []int {
	slice = append(slice, elem)
	return slice
}

// append elem to slice by reference -> No return is required here
func addElemByRef(slice *[]int, elem int) {
	*slice = append(*slice, elem)
}
