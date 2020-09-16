package main

import "testing"

func Test_QSORT(t *testing.T) {
	a := []int{-1, 0, 1, 2, -1, -4}
	qsort(a)
	t.Log(a)
}

func qsort(array []int) {
	if len(array) <= 1 {
		return
	}
	var (
		left  = 0
		right = len(array) - 1
	)
	mod := array[0]
	for left < right {
		if mod > array[left+1] {
			array[left+1], array[left] = array[left], array[left+1]
			left++
		} else {
			array[left+1], array[right] = array[right], array[left+1]
			right--
		}
	}
	array[left] = mod
	qsort(array[:left])
	qsort(array[left+1:])
}
