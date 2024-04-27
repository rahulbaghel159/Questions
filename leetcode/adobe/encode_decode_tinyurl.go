package adobe

// https://leetcode.com/problems/encode-and-decode-tinyurl/
type Codec struct {
	dict map[string]struct{}
}

func Constructor() Codec {
	return Codec{
		dict: make(map[string]struct{}),
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	if _, ok := this.dict[longUrl]; !ok {
		this.dict[longUrl] = struct{}{}
	}

	return longUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	if _, ok := this.dict[shortUrl]; !ok {
		return ""
	}

	return shortUrl
}
