package algorithm

import "github.com/phial3/container"

func IsSorted[T any](arr []T, cmp func(a, b T) container.Ordering) bool {
	var n = len(arr)
	if n == 0 {
		return true
	}

	var flag1 = false
	var flag2 = false
	for i := 1; i < n; i++ {
		var result = cmp(arr[i], arr[i-1])
		if result == container.Less {
			flag1 = true
		} else if result == container.Greater {
			flag2 = true
		}
	}
	return !(flag1 && flag2)
}

func Sort[T any](arr []T, cmp func(a, b T) container.Ordering) {
	QuickSort(arr, 0, len(arr)-1, cmp)
}

func getMedium[T any](arr []T, begin int, end int, cmp func(a, b T) container.Ordering) int {
	var mid = (begin + end) / 2
	if cmp(arr[end], arr[begin])+cmp(arr[begin], arr[mid]) != 0 {
		return begin
	} else if cmp(arr[end], arr[mid])+cmp(arr[mid], arr[begin]) != 0 {
		return mid
	} else {
		return end
	}
}

func insertionSort[T any](arr []T, a, b int, cmp func(a, b T) container.Ordering) {
	for i := a + 1; i <= b; i++ {
		for j := i; j > a && cmp(arr[j], arr[j-1]) == container.Less; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

func QuickSort[T any](arr []T, begin int, end int, cmp func(a, b T) container.Ordering) {
	if begin >= end {
		return
	}
	if end-begin <= 15 {
		insertionSort(arr, begin, end, cmp)
		return
	}

	var index = begin
	var mid = getMedium(arr, begin, end, cmp)
	arr[mid], arr[begin] = arr[begin], arr[mid]
	for i := begin + 1; i <= end; i++ {
		var flag = cmp(arr[i], arr[begin])
		if flag == container.Less || (flag == container.Equal && i%2 == 0) {
			index++
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index], arr[begin] = arr[begin], arr[index]

	QuickSort(arr, begin, index-1, cmp)
	QuickSort(arr, index+1, end, cmp)
}

// not exist return -1
func BinarySearch[T any](arr []T, target T, cmp func(a, b T) container.Ordering) int {
	var n = len(arr)
	if n == 0 {
		return -1
	}

	var left = 0
	var right = n - 1
	for right-left > 1 {
		var mid = (left + right) / 2
		var flag = cmp(arr[mid], target)
		if flag == container.Equal {
			return mid
		} else if flag == container.Greater {
			right = mid
		} else {
			left = mid
		}
	}

	if cmp(arr[left], target) == container.Equal {
		return left
	} else if cmp(arr[right], target) == container.Equal {
		return right
	} else {
		return -1
	}
}
