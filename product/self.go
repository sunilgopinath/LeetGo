package product

// ExceptSelf returns all the products without the curr index
func ExceptSelf(nums []int) []int {
	pp := make([]int, len(nums))
	psf := 1
	for i := 0; i < len(nums); i++ {
		pp[i] = psf
		psf *= nums[i]
	}
	psf = 1
	for i := len(nums) - 1; i >= 0; i-- {
		pp[i] *= psf
		psf *= nums[i]
	}
	return pp
}
