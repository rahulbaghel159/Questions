package google

//https://leetcode.com/explore/interview/card/google/59/array-and-strings/436/
/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */
var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	prevBuf, prevSize, prevIndex := make([]byte, 4), 0, 0
	return func(buf []byte, n int) int {
		counter := 0
		for counter < n {
			if prevIndex < prevSize {
				buf[counter] = prevBuf[prevIndex]
				counter++
				prevIndex++
			} else {
				prevSize = read4(prevBuf)
				prevIndex = 0
				if prevSize == 0 {
					break
				}
			}
		}

		return counter
	}
}
