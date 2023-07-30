package maxruntime

//https://leetcode.com/problems/maximum-running-time-of-n-computers

/*
//sorting approach
func maxRunTime(n int, batteries []int) int64 {
	sort.Slice(batteries, func(i, j int) bool {
		return batteries[i] < batteries[j]
	})

	live := batteries[len(batteries)-n:]

	extra := int64(0)
	for i := 0; i < len(batteries)-n; i++ {
		extra += int64(batteries[i])
	}

	for i := 0; i < n-1; i++ {
		if extra < int64(i+1)*int64((live[i+1]-live[i])) {
			return int64(live[i]) + extra/int64((i+1))
		}

		extra -= int64(i+1) * int64((live[i+1] - live[i]))
	}

	return int64(live[n-1]) + extra/int64(n)
}
*/

// binary search approach
func maxRunTime(n int, batteries []int) int64 {
	sumPower := int64(0)

	for _, v := range batteries {
		sumPower += int64(v)
	}

	left, right := int64(1), sumPower/int64(n)

	for left < right {
		target, extra := right-(right-left)/2, int64(0)

		for _, v := range batteries {
			extra += min(int64(v), target)
		}

		if extra >= int64(n)*target {
			left = target
		} else {
			right = target - 1
		}
	}

	return left
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
