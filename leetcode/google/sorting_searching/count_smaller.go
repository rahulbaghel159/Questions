package sortingsearching

// https://leetcode.com/explore/interview/card/google/63/sorting-and-searching-4/3083/
func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	counts := make([]int, len(nums))
	minimum, maximum := 1000000, -1000000

	for _, num := range nums {
		minimum = min(minimum, num)
		maximum = max(maximum, num)
	}

	st := buildSegmentTree(minimum, maximum)

	for i := len(nums) - 1; i >= 0; i-- {
		st.update(nums[i])
		counts[i] = st.query(minimum, nums[i]-1)
	}

	return counts
}

type SegmentTreeNode struct {
	left  *SegmentTreeNode
	right *SegmentTreeNode

	start, end, sum int
}

func buildSegmentTree(start, end int) *SegmentTreeNode {
	if start > end {
		return nil
	}

	node := &SegmentTreeNode{
		start: start,
		end:   end,
		sum:   0,
	}

	if start == end {
		return node
	}

	mid := start + (end-start)/2
	node.left = buildSegmentTree(start, mid)
	node.right = buildSegmentTree(mid+1, end)

	return node
}

func (s *SegmentTreeNode) update(index int) {
	if s == nil {
		return
	}

	if s.start == index && s.end == index {
		s.sum += 1
		return
	}

	mid := s.start + (s.end-s.start)/2
	if index <= mid {
		s.left.update(index)
	} else {
		s.right.update(index)
	}

	s.sum = s.left.sum + s.right.sum
}

func (s *SegmentTreeNode) query(start, end int) int {
	if s == nil || start > end {
		return 0
	}

	if s.start == start && s.end == end {
		return s.sum
	}

	mid := s.start + (s.end-s.start)/2
	if end <= mid {
		return s.left.query(start, end)
	} else if start > mid {
		return s.right.query(start, end)
	}

	return s.left.query(start, mid) + s.right.query(mid+1, end)
}
