package design

// https://leetcode.com/explore/interview/card/google/65/design-4/3093/
type Logger struct {
	dict map[string]int
}

func Constructor3() Logger {
	return Logger{
		dict: make(map[string]int, 0),
	}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if _, ok := this.dict[message]; !ok {
		this.dict[message] = timestamp
		return true
	}

	last_time := this.dict[message]
	if timestamp >= last_time+10 {
		this.dict[message] = timestamp
		return true
	}

	return false
}
