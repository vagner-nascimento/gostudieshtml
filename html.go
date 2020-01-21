package gostudieshtml

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// GetTitles Get titles of the websites url
func GetTitles(urls ...string) <-chan string {
	ch := make(chan string)
	for _, url := range urls {
		go func(urlParam string) { // Calling an anonymous function in a go routine
			resp, _ := http.Get(urlParam)
			html, _ := ioutil.ReadAll(resp.Body)
			regX, _ := regexp.Compile("<title.*?>(.*?)<\\/title>")
			matches := regX.FindStringSubmatch(string(html))

			ch <- matches[1]
		}(url) // Passing each URL in URLs array
	}

	return ch
}
