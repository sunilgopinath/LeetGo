package merge

//SortedArray returns the merge of two sorted arrays
func SortedArray(nums1 []int, m int, nums2 []int, n int) []int {
	k := m + n - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			// move to last position
			nums1[k] = nums1[m]
			m--
		} else {
			nums1[k] = nums2[n]
			n--
		}
		k--
	}
	for n >= 0 {
		nums1[k] = nums2[n]
		k--
		n--
	}
	return nums1
}
