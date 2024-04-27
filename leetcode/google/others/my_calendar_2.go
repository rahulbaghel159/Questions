package others

type MyCalendarTwo struct {
	bookings        [][2]int
	double_bookings [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		bookings:        make([][2]int, 0),
		double_bookings: make([][2]int, 0),
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, booking := range this.double_bookings {
		if booking[0] < end && booking[1] > start {
			return false
		}
	}

	for _, booking := range this.bookings {
		if booking[0] < end && booking[1] > start {
			this.double_bookings = append(this.double_bookings, [2]int{max(start, booking[0]), min(end, booking[1])})
		}
	}

	this.bookings = append(this.bookings, [2]int{start, end})

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
