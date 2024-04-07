package kata

func longestMonotonicSubarray(nums []int) int {
	var lsi int
	var lsd int
	var currentLsi = 1
	var currentLsd = 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			currentLsi++
			if lsd < currentLsd {
				lsd = currentLsd
			}
			currentLsd = 1
		} else if nums[i-1] > nums[i] {
			currentLsd++
			if lsi < currentLsi {
				lsi = currentLsi
			}
			currentLsi = 1
		} else {
			if lsi < currentLsi {
				lsi = currentLsi
			}
			if lsd < currentLsd {
				lsd = currentLsd
			}
			currentLsd = 1
			currentLsi = 1
		}
	}

	if lsd < currentLsd {
		lsd = currentLsd
	}
	if lsi < currentLsi {
		lsi = currentLsi
	}

	if lsi < lsd {
		return lsd
	}

	return lsi
}
