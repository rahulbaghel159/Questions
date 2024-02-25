package maxsquare

/*
*
*

Input matrix = [
	["1","0","1","0","0"],
	["1","0","1","1","1"],
	["1","1","1","1","1"],
	["1","0","0","1","0"]
]

Output = 2

res = 2
1. i=0, j=0, tmp=1
	x, y, flag = 1, 1, true
		k = 0...1
		flag = false
2. i=0, j=1
3. i=0, j=2, tmp=1
	x, y, flag = 1, 3, true
	k = 2..3
	k = 0..1
4. i=0, j=3
5. i=0, j=4
6. i=1, j=0, tmp=1
	x, y, flag = 2, 1, true
	k = 0..1
	k = 1..2
7. i=1, j=1
8. i=1, j=2, tmp=1
	x, y, flag = 2, 3, true
	k = 2..3
	k = 1..2
	x, y, tmp = 3, 4, 2

*/

func maxSquareSize(matrix [][]int) int {
	if len(matrix) == 0 {
		return -1
	}

	res := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			tmp := 0
			if matrix[i][j] != 1 {
				continue
			}
			tmp++
			x, y, flag := i+1, j+1, true
			for x < len(matrix) && y < len(matrix[0]) && matrix[x][y] == 1 {
				for k := j; k <= y; k++ {
					if matrix[x][k] == 0 {
						flag = false
						break
					}
				}
				if !flag {
					break
				}
				for k := i; k <= x; k++ {
					if matrix[k][y] == 0 {
						flag = false
						break
					}
				}
				if flag {
					tmp++
					x, y = x+1, y+1
				} else {
					break
				}
			}
			res = max(res, tmp)
		}
	}

	return res
}
