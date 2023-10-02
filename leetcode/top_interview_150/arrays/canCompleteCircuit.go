package arrays

//https://leetcode.com/problems/gas-station

func canCompleteCircuit(gas []int, cost []int) int {
	currGain, totalGain, answer := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		totalGain += gas[i] - cost[i]
		currGain += gas[i] - cost[i]

		if currGain < 0 {
			answer = i + 1
			currGain = 0
		}
	}

	if totalGain >= 0 {
		return answer
	}

	return -1
}
