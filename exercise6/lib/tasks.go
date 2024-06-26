package lib

func SearchMaxSumSubSlice(slice []uint) uint {
	switch len(slice) {
	case 0:
		return 0
	case 1:
		return slice[0]
	case 2:
		return slice[0] + slice[1]
	}

	var (
		one             = slice[0]
		two             = one
		sum        uint = 0
		currentSum uint = 0
	)

	for i := range slice {
		if two != slice[i] {
			two = slice[i]
			break
		}
	}

	for i := range slice {
		if slice[i] != one && slice[i] != two {
			if currentSum > sum {
				sum = currentSum
			}
			currentSum = 0
			one = slice[i]
		} else {
			currentSum += slice[i]
		}
	}

	return sum
}
