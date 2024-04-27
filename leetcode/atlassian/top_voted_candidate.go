package atlassian

// https://leetcode.com/problems/online-election/
type TopVotedCandidate struct {
	winner_map map[int]int
	times      []int
}

func Constructor5(persons []int, times []int) TopVotedCandidate {
	helping_map, winner_map, winner := make(map[int]int), make(map[int]int), []int{persons[0], 1}

	for i := 0; i < len(times); i++ {
		helping_map[persons[i]]++
		if helping_map[persons[i]] > winner[1] || helping_map[persons[i]] == winner[1] {
			winner = []int{persons[i], helping_map[persons[i]]}
		}
		winner_map[times[i]] = winner[0]
	}

	return TopVotedCandidate{
		winner_map: winner_map,
		times:      times,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	left, right, mid := 0, len(this.times)-1, 0
	for left <= right {
		mid = left + (right-left)/2
		if this.times[mid] == t {
			return this.winner_map[this.times[mid]]
		} else if this.times[mid] < t {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return this.winner_map[this.times[right]]
}
