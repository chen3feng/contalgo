package contalgo

import (
	"testing"
)

func Test_IsSorted(t *testing.T) {
	expectTrue(t, IsSorted([]int{}))
	expectTrue(t, IsSorted([]int{1, 2, 3, 4}))
	expectTrue(t, IsSorted([]int{1, 2, 2, 3, 4}))
	expectFalse(t, IsSorted([]int{1, 2, 2, 3, 2}))
}

func Test_IsDescSorted(t *testing.T) {
	expectTrue(t, IsDescSorted([]int{}))
	expectFalse(t, IsDescSorted([]int{1, 2, 3, 4}))
	expectFalse(t, IsDescSorted([]int{1, 2, 2, 3, 4}))
	expectTrue(t, IsDescSorted([]int{5, 4, 3, 3, 2}))
}

func Test_Sort(t *testing.T) {
	a := []int{5, 4, 3, 2, 1}
	Sort(a)
	expectTrue(t, IsSorted(a))
}

func Test_DescSort(t *testing.T) {
	a := []int{1, 2, 3, 4}
	DescSort(a)
	expectTrue(t, IsDescSorted(a))
}

func Test_StableSort(t *testing.T) {
	a := []int{5, 4, 3, 2, 1}
	StableSort(a)
	expectTrue(t, IsSorted(a))
}

func Test_DescStableSort(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	DescStableSort(a)
	expectTrue(t, IsDescSorted(a))
}

func greater(x, y int) bool { return x > y }

func Test_SortFunc(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	SortFunc(a, greater)
	expectTrue(t, IsDescSorted(a))
}

func Test_StableSortFunc(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	StableSortFunc(a, greater)
	expectTrue(t, IsDescSorted(a))
}
