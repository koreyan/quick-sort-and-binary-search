package main

import (
	"fmt"
)

//var count int = 1

func swap(a, b *int) {

	temp := *a
	*a = *b
	*b = temp
}

func partition(dataArr []int, left int, right int) int {

	//fmt.Printf("partition %d번째 호출\n", count)
	//count++
	//fmt.Printf("시작 : left->%d right->%d\n", left, right)
	pivotLoc := left
	pivot := dataArr[pivotLoc]
	left++

	for left <= right {
		for dataArr[left] <= pivot && left < right { //left는 pivot보다 큰 수를 찾음
			left++
			//fmt.Printf("left증가 현재 left %d\n", left)
		}
		//fmt.Println("현재 left가 가리키고 있는것", dataArr[left])

		for dataArr[right] > pivot && left <= right {
			right--
			//fmt.Printf("right증가 현재 right %d\n", right)
		}
		//fmt.Println("현재 right가 가리키고 있는것", dataArr[right])
		if left < right {
			//fmt.Printf("swap : %d, %d\n", dataArr[left], dataArr[right])
			swap(&dataArr[left], &dataArr[right])
			//fmt.Println("현재 상태 :", dataArr)
		} else { // left와 right가 같고 가리키는 값이 pivot보다 작으면 무한 반복된다.
			break
		}
	}

	swap(&dataArr[pivotLoc], &dataArr[right])

	//fmt.Println("현재 상태 :", dataArr)
	return right
}

func quick_sort(dataArr []int, left, right int) {

	if left < right {
		mid := partition(dataArr, left, right)
		quick_sort(dataArr, left, mid-1)
		quick_sort(dataArr, mid+1, right)
	}

}

func binary_search(dataArr []int, target int) int {
	first := 0
	end := len(dataArr) - 1
	mid := (first + end) / 2
	for first <= end {
		if target < dataArr[mid] {
			end = mid - 1
		} else if target > dataArr[mid] {
			first = mid + 1
		} else {
			return mid
		}
		mid = (first + end) / 2
	}
	return -1
}

func main() {
	s := []int{5, 1, 3, 4, 8, 6}
	fmt.Println(s)
	fmt.Println("after quick sort..")
	quick_sort(s, 0, len(s)-1)
	fmt.Println(s)

	d := binary_search(s, 8)
	if d == -1 {
		fmt.Println("찾는 데이터가 없습니다.")
	} else {
		fmt.Println("target index", d)
	}

	d = binary_search(s, 16)
	if d == -1 {
		fmt.Println("찾는 데이터가 없습니다.")
	} else {
		fmt.Println("target index", d)
	}
}
