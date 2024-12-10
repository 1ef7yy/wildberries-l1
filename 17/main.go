package main

import (
	"fmt"
)

// бинарный поиск работает только на отсортированном массиве
// 1. Разбиваем массив пополам
// 2. Если искомый элемент больше или меньше среднего элемента, то ищем в другой половине
// 3. Если искомый элемент равен среднему элементу, то возвращаем индекс

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 6
	idx := binarySearch(numbers, target)
	if idx == -1 {
		fmt.Printf("%d is not present in array\n", target)
		return
	}
	fmt.Printf("%d is present at index %d\n", target, idx)
}
