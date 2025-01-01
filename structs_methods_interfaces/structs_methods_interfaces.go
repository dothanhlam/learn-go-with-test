package structs_methods_interfaces


func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

func Area(nums ...float64) float64 {

	if len(nums) == 1 {
		return 3.141592653589793 * nums[0] * nums[0]
	}
	return nums[0] * nums[1]
}