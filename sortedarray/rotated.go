package sortedarray

// Search returns the index of the target in the input slice
func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// find pivot element
	pvt := FindPivot(nums)
	if target == pvt {
		return pvt
	}
	if pvt == -1 {
		return binarySearch(nums, target, 0, len(nums)-1)
	}
	if nums[target] > nums[pvt] {
		// perform binary search with hi as pvt
	}
}

func search(nums []int, target, lo, hi int) int {
	return 0
}

// FindPivot returns the pivot element index
func FindPivot(nums []int) int {
	return findPivot(nums, 0, len(nums)-1)
}

func findPivot(nums []int, lo, hi int) int {
	// base cases
	if hi < lo {
		return -1
	}
	if hi == lo {
		return lo
	}

	/* low + (high - low)/2; */
	mid := (lo + hi) / 2
	if mid < hi && nums[mid] > nums[mid+1] {
		return mid
	}
	if mid > lo && nums[mid] < nums[mid-1] {
		return mid - 1
	}
	if nums[lo] >= nums[mid] {
		return findPivot(nums, lo, mid-1)
	}
	return findPivot(nums, mid+1, hi)
}
