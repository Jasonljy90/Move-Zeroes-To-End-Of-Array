package main

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	j := 0 // Keeps track where the 0 is at
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i] // Swap nearest non 0 with nearest 0
			}
			j++ // Keeps track where the 0 is at
		}
	}
}
