package main

import "fmt"

// BubbleSort 冒泡排序（升序）
func BubbleSort(arr []int) []int {
	n := len(arr)
	sortedArr := make([]int, n)
	copy(sortedArr, arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sortedArr[j] > sortedArr[j+1] {
				sortedArr[j], sortedArr[j+1] = sortedArr[j+1], sortedArr[j]
			}
		}
	}
	return sortedArr
}

// BubbleSortOptimized 优化版（提前终止）
func BubbleSortOptimized(arr []int) []int {
	n := len(arr)
	sortedArr := make([]int, n)
	copy(sortedArr, arr)

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if sortedArr[j] > sortedArr[j+1] {
				sortedArr[j], sortedArr[j+1] = sortedArr[j+1], sortedArr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return sortedArr
}

// BubbleSortDesc 冒泡排序（降序）
func BubbleSortDesc(arr []int) []int {
	n := len(arr)
	sortedArr := make([]int, n)
	copy(sortedArr, arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sortedArr[j] < sortedArr[j+1] {
				sortedArr[j], sortedArr[j+1] = sortedArr[j+1], sortedArr[j]
			}
		}
	}
	return sortedArr
}

func maopao() {
	fmt.Println("====== 冒泡排序测试 ======")

	arr1 := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("原始数组: %v\n", arr1)
	fmt.Printf("升序排序: %v\n", BubbleSort(arr1))
	fmt.Printf("降序排序: %v\n", BubbleSortDesc(arr1))
	fmt.Println()

	arr2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("已排序数组: %v\n", arr2)
	fmt.Printf("优化排序: %v\n", BubbleSortOptimized(arr2))
	fmt.Println()

	arr3 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Printf("逆序数组: %v\n", arr3)
	fmt.Printf("升序排序: %v\n", BubbleSort(arr3))
	fmt.Println()

	arr4 := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	fmt.Printf("包含重复元素: %v\n", arr4)
	fmt.Printf("升序排序: %v\n", BubbleSort(arr4))
	fmt.Println()

	arr5 := []int{42}
	fmt.Printf("单个元素: %v\n", arr5)
	fmt.Printf("升序排序: %v\n", BubbleSort(arr5))
}

func main() {
	maopao()
}


func 