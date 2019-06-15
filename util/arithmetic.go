package util

import "strings"

func GetIndexByDichotomy(arr []int, target int) int {
	return getIndexByDichotomy(arr, target)
}

// 二分法查找某个数 -- 排序数组
func getIndexByDichotomy(arr []int, aim int) (aimIndex int) {
	// start和end最大最小的角标
	start := 0
	end := len(arr) - 1
	for start <= end {
		aimIndex = (start + end) / 2
		if aim == arr[aimIndex] {
			return
		} else if aim > arr[aimIndex] {
			start = aimIndex + 1
		} else {
			end = aimIndex - 1
		}
	}
	return -1
}

// 选择问题:
func SelectionProblem(k int, arr []int) []int {
	return selectionProblem(k, arr)
}

// 找出数组arr中k个最大值,其中,k < len(arr);
func selectionProblem(k int, arr []int) []int {
	if k >= len(arr) {
		return arr
	}
	//arr = bubbleSort(arr,"desc")
	arr = bubbleSort(arr, "asc")
	return arr[:k]
}

// 冒泡排序
func BubbleSort(arr []int, order string) []int {
	return bubbleSort(arr, order)
}

// 冒泡排序
// asc从小到大,DESC从大到小
func bubbleSort(arr []int, order string) []int {
	if strings.ToUpper(order) != "DESC" {
		order = "ASC"
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if order == "ASC" {
				if arr[i] > arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			} else {
				if arr[i] < arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}

		}
	}
	return arr
}
