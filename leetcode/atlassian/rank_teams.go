package atlassian

import "sort"

// https://leetcode.com/problems/rank-teams-by-votes/
type VoteRecord struct {
	team  string
	ranks []int
}

func rankTeams(votes []string) string {
	if len(votes) == 0 {
		return ""
	}

	if len(votes) == 1 {
		return votes[0]
	}

	votes_map := make(map[string]VoteRecord)
	for _, vote := range votes {
		for i := 0; i < len(vote); i++ {
			team := string(vote[i])
			if _, ok := votes_map[team]; !ok {
				votes_map[team] = VoteRecord{
					team:  team,
					ranks: make([]int, len(vote)),
				}
			}
			votes_map[team].ranks[i]++
		}
	}

	votesList, k := make([]VoteRecord, len(votes_map)), 0
	for _, v := range votes_map {
		votesList[k] = v
		k++
	}

	sort.Slice(votesList, func(i, j int) bool {
		idx := 0
		for idx < len(votesList[i].ranks) && idx < len(votesList[j].ranks) {
			if votesList[i].ranks[idx] == votesList[j].ranks[idx] {
				idx++
			} else {
				return votesList[j].ranks[idx] < votesList[i].ranks[idx]
			}
		}
		return votesList[i].team < votesList[j].team
	})

	res := ""
	for _, record := range votesList {
		res += record.team
	}

	return res
}
