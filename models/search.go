package models

//BinarySearch adalah algoritma pencarian dasar
func BinarySearch(arr []int, l int, r int, x int) int {

	if r >= l {
		mid := l + (r-l)/2

		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			return BinarySearch(arr, l, mid-1, x)
		}

		return BinarySearch(arr, mid+1, r, x)
	}
	return -1
}
