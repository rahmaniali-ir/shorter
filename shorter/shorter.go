package shorter

import (
	"errors"
	"math/rand"
	"time"
)

var Records = make(map[string]string)
var shortURLCharset = "0123456789_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var shortURLLength = 6

func getByLong(long string) (string, bool) {
	shortURL, exists := Records[long]

	return shortURL, exists
}

func getByShort(short string) (string, bool) {
	for long := range Records {
		if Records[long] == short {
			return long, true
		}
	}

	return "", false
}

func generateRandomURL() string {
	short := ""

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < shortURLLength; i++ {
		j := rand.Intn(len(shortURLCharset))
		short += string(shortURLCharset[j])
	}

	return short
}

func generateShortURL() string {
	var ok bool

	for !ok {
		url := generateRandomURL()
		_, exists := getByShort(url)
		
		if !exists {
			return url
		}
	}

	return ""
}

func MakeShorter(long string) string {
	_, exists := getByLong(long)
	
	if exists {
		return Records[long]
	}
	
	key := generateShortURL()
	Records[long] = key

	return key
}

func MakeLonger(short string) (string, error) {
	long, exists := getByShort(short)

	if exists {
		return long, nil
	}

	return "", errors.New("There's no such link!")
}
