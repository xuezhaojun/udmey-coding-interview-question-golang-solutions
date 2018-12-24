package _3_sort

// BubbleSort  n^2
// 冒泡排序的大错觉： 是j 和 j+1 的对比，而不是i和j的比较
func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	return nums
}

// SelectSort 选择排序 n^2
// SelectSort 是 min 和 j 比， min 和 i 换
// 假设i是最小的那个数，然后向后比较，一旦发现比i小的数字，则用min记录，最终将min和i交换
// j 只是用于记录sorted 和 unsorted 的界限
func SelectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		var min int = i
		for j := i; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}

// Merge Sort n*log(n)
// MergeSort = Merge将两个array合并 + MergeSort
func MergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	left := nums[:len(nums)/2]
	right := nums[len(nums)/2:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	var result []int
	var l, r = 0, 0
	for {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}

		if l == len(left) {
			result = append(result, right[r:]...)
			return result
		}
		if r == len(right) {
			result = append(result, left[l:]...)
			return result
		}
	}
}
