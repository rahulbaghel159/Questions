package atlassian

// https://leetcode.com/problems/logger-rate-limiter/
type Logger struct {
	index map[string]int
}

func Constructor() Logger {
	return Logger{
		index: make(map[string]int),
	}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if time, exists := this.index[message]; exists {
		if timestamp >= time {
			this.index[message] = timestamp + 10
			return true
		} else {
			return false
		}
	} else {
		this.index[message] = timestamp + 10
		return true
	}
}
