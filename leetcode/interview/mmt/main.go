package mmt

/*
_H_H_H_H_H_

_HRH_HRH_HR => 3

l := 0
r := 0
T = {}
H = {1, 1, 1, }

T - 3, 5 i.e. 2

_, => 0
_H => 1
_H_ => 1
_H_H => 1
_H_H_ => 1
_H_H_H => 2
_H_H_H_ => 2
_H_H_H_H => 3
_H_H_H_H_ => 3
_H_H_H_H_H_ => 3

{
H => -1
_H => 1
H_ => 1
H_H => 1
HH => -1
_HH => -1
HH_ => -1
H_H_H => 2
}

_H_H_HH_ (tricky)

_HRHRHHR => 3

_H_H => 1
_H_H_H => 2
_H_H_HH => -1

_HH_ => 2

RHHR => 2



H - houses
_ - tank
house water from either side
min water tanks for each house

__HH

T -1

__HH_H

T - 2
*/

// func countTank(str string) int {
// 	if str == "" {
// 		return -1
// 	}

// 	arr := []byte(str)

// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == '_' {
// 			arr[i] = 'T'
// 		}
// 	}

// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == 'H' {
// 			tankCount := 0
// 			if i-1 >= 0 && (arr[i-1] == 'T' || arr[i-1] == 'R' || arr[i-1] == '_') {
// 				tankCount++
// 			}
// 			if i+1 < len(arr) && (arr[i+1] == 'T' || arr[i+1] == 'R' || arr[i+1] == '_') {
// 				tankCount++
// 			}
// 			if tankCount == 0 {
// 				return -1
// 			} else if tankCount == 1 {
// 				continue
// 			} else {
// 				if arr[i-1] == 'R' && arr[i+1] == 'R' {
// 					continue
// 				}
// 				if arr[i-1] != 'R' {
// 					arr[i-1] = '_'
// 					arr[i+1] = 'R'
// 				} else {
// 					arr[i+1] = '_'
// 					arr[i-1] = 'R'
// 				}
// 			}
// 		}
// 	}

// 	count := 0
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == 'R' {
// 			count++
// 		}
// 	}

// 	return count
// }

func countTank(str string) int {
	if str == "" {
		return -1
	}

	tmp := []byte(str)
	watered := make([]int, len(tmp))
	for i := 0; i < len(tmp); i++ {
		if tmp[i] == 'H' {
			watered[i] = 0
		} else {
			watered[i] = 2
		}
	}

	arr := make([]byte, len(tmp)+2)
	for i := 1; i < len(tmp); i++ {
		arr[i] = tmp[i-1]
	}
	arr[len(arr)-2] = tmp[len(tmp)-1]

	for i := 1; i < len(tmp); i++ {
		if arr[i] == 'H' {
			if arr[i-1] == 'H' && arr[i+1] == 'H' {
				return -1
			}

			if arr[i-1] == 'T' {
				watered[i-1] = 1
				continue
			}
			if arr[i-1] == '_' {
				if arr[i+1] == '_' {
					arr[i+1] = 'T'
					watered[i-1] = 1
				} else {
					arr[i-1] = 'T'
					watered[i-1] = 1
				}
			} else if arr[i+1] == '_' {
				arr[i+1] = 'T'
				watered[i-1] = 1
			}
		}
	}

	for i := 0; i < len(watered); i++ {
		if watered[i] == 0 {
			return -1
		}
	}

	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 'T' {
			count++
		}
	}

	return count
}
