package main

import "fmt"

func mergeSort(arr []int) []int {
	return reMergeSort(arr)
}
func reMergeSort(arr []int) []int {
	fmt.Println("i", arr)
	var half1, half2 []int
	var out, rem []int
	if len(arr) > 1 {
		half1 = reMergeSort(arr[:int(len(arr)/2)])
		half2 = reMergeSort(arr[int(len(arr)/2):])
	} else {
		return arr
	}
	for len(half1) != 0 && len(half2) != 0 {
		out = append(out, min(half1[0], half2[0]))
		if half1[0] < half2[0] {
			if len(half1) == 0 {
				half1 = []int{}
			} else {
				half1 = half1[1:]
			}
		} else {
			if len(half2) == 0 {
				half2 = []int{}
			} else {
				half2 = half2[1:]
			}
		}
	}
	if len(half1) != 0 {
		rem = half1
	} else {
		rem = half2
	}
	out = append(out, rem...)

	fmt.Println("0", out)

	return out
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
