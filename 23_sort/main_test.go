package _3_sort

import "testing"

func TestSort(t *testing.T) {
	raw := []int{2, 31, 3, 4, 5, 12, 3, 4, 3, 8, 54, 1, 531241, 676}
	t.Log(BubbleSort(raw))
	t.Log(SelectSort(raw))
	t.Log(MergeSort(raw))
}
