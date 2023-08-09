package optimalaccountbalancing

//https://leetcode.com/problems/optimal-account-balancing/

func minTransfers(transactions [][]int) int {
	balance := make([]int, 12)

	for i := 0; i < len(transactions); i++ {
		balance[transactions[i][0]] -= transactions[i][2]
		balance[transactions[i][1]] += transactions[i][2]
	}

	// count, posIndex := 0, 0
	// for i := 0; i < len(balance); i++ {
	// 	if balance[i] > 0 {
	// 		posIndex = i
	// 		break
	// 	}
	// }

	// fmt.Println("balance", balance)
	// fmt.Println("posIndex", posIndex)

	// for i := 0; i < len(balance); {
	// 	if balance[i] < 0 {
	// 		settlement := min(abs(balance[i]), abs(balance[posIndex]))
	// 		balance[posIndex] -= settlement
	// 		balance[i] += settlement
	// 		count++

	// 		if balance[posIndex] == 0 {
	// 			for j := posIndex + 1; j < len(balance); j++ {
	// 				if balance[j] > 0 {
	// 					posIndex = j
	// 					break
	// 				}
	// 			}
	// 		}
	// 	}

	// 	if balance[i] >= 0 {
	// 		i++
	// 	}
	// }

	// return count

	posBalance := 0
	for _, v := range balance {
		if v > 0 {
			posBalance += v
		}
	}

	count := 0
	for _, v := range balance {
		if v < 0 {
			posBalance -= v
			count++
		}
		if posBalance < 0 {
			return -1
		}
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
