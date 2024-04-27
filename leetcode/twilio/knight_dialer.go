package twilio

// https://leetcode.com/problems/knight-dialer/

/*
1 -> 6, 8
2 -> 7, 9
3 -> 4, 8
4 -> 0, 3, 9
5 ->
6 -> 0, 1, 7
7 -> 2, 6
8 -> 1, 3
9 -> 2, 4
0 -> 4, 6

1, 3, 7, 9 -> 1*4
2, 8 -> 2*2
4,6 -> 4*2
0
5
*/

var (
	dict map[int][]int
	MOD  int
)

func knightDialer(n int) int {
	if n == 1 {
		return 10
	}

	dict = map[int][]int{
		1: {6, 8},
		2: {7, 9},
		3: {4, 8},
		4: {0, 3, 9},
		5: {},
		6: {0, 1, 7},
		7: {2, 6},
		8: {1, 3},
		9: {2, 4},
		0: {4, 6},
	}
	MOD = 1e9 + 7

	memo := make(map[int]map[int]int)

	n1 := backtrackKnight(1, 1, n, memo) * 4
	n2 := backtrackKnight(2, 1, n, memo) * 2
	n4 := backtrackKnight(4, 1, n, memo) * 2
	n0 := backtrackKnight(0, 1, n, memo)
	n5 := backtrackKnight(5, 1, n, memo)

	return (n1 + n2 + n4 + n0 + n5) % MOD
}

func backtrackKnight(curr_num, len_so_far, n int, memo map[int]map[int]int) int {
	if len_so_far == n {
		return 1
	}

	if _, ok := memo[curr_num][n-len_so_far]; ok {
		return memo[curr_num][n-len_so_far]
	}

	next := dict[curr_num]
	count := 0
	for _, dial := range next {
		count += backtrackKnight(dial, len_so_far+1, n, memo)
	}

	if _, ok := memo[curr_num]; !ok {
		memo[curr_num] = make(map[int]int)
	}

	memo[curr_num][n-len_so_far] = count % MOD
	return memo[curr_num][n-len_so_far]
}
