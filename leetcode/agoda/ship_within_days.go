package agoda

// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/
func shipWithinDays(weights []int, days int) int {
	if len(weights) == 0 {
		return 0
	}

	if days == 1 {
		sum := 0
		for _, w := range weights {
			sum += w
		}
		return sum
	}

	max_weight := weights[0]
	for _, w := range weights {
		max_weight = max(max_weight, w)
	}

	sum := 0
	for _, w := range weights {
		sum += w
	}

	left, right := max_weight, sum
	for left < right {
		mid := left + (right-left)/2
		if feasible(weights, mid, days) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func feasible(weights []int, capacity, days int) bool {
	days_needed, current_load := 1, 0
	for _, w := range weights {
		current_load += w
		if current_load > capacity {
			days_needed++
			current_load = w
		}
	}

	return days_needed <= days
}
