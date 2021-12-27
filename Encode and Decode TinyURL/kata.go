package kata

import (
	"math/rand"
	"time"
)

type Codec struct {
	urlContainer map[string]string
}

func Constructor() Codec {
	return Codec{
		urlContainer: make(map[string]string),
	}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longURL string) string {
	v := genID(5)
	url := "http://tinyurl.com/" + v
	c.urlContainer[url] = longURL
	return url
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	return c.urlContainer[shortUrl]
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func genID(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
