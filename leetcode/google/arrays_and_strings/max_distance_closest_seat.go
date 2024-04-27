package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3058/
func maxDistToClosest(seats []int) int {
	if len(seats) == 0 {
		return 0
	}

	seats_index := make([]int, 0)
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			seats_index = append(seats_index, i)
		}
	}

	if len(seats_index) == 1 {
		if seats[0] == 1 || seats[len(seats)-1] == 1 {
			return len(seats) - 1
		} else {
			if len(seats)-1-seats_index[0] > seats_index[0] {
				return len(seats) - 1 - seats_index[0]
			} else {
				return seats_index[0]
			}
		}
	}

	maxDist := 0
	for i := 0; i < len(seats_index)-1; i++ {
		maxDist = max(maxDist, seats_index[i+1]-seats_index[i])
	}
	maxDist /= 2

	if seats[0] == 0 && len(seats_index) > 0 {
		maxDist = max(maxDist, seats_index[0])
	}

	if seats[len(seats)-1] == 0 && len(seats_index) > 0 {
		maxDist = max(maxDist, len(seats)-1-seats_index[len(seats_index)-1])
	}

	return maxDist
}
