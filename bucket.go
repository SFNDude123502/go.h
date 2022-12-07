package main

import "fmt"

// MUST ALSO IMPORT MERGE SORT

func bucketSort(arr *[]int) {
	var buckets = make([][]int, 25)
	var lrg int
	var arrRep []int
	for _, ival := range *arr {
		if lrg < ival {
			lrg = ival
		}
	}
	lrg++
	for _, ival := range *arr {
		curBucket := int(ival * 25 / lrg)
		buckets[curBucket] = append(buckets[curBucket], ival)
	}
	for i, bucket := range buckets {
		if len(bucket) == 0 {
			continue
		}
		fmt.Println("b", bucket)
		buckets[i] = mergeSort(bucket)
		fmt.Println("a", bucket)
	}
	for _, bucket := range buckets {
		arrRep = append(arrRep, bucket...)
	}
	*arr = arrRep
}
