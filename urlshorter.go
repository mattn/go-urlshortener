package urlshorter

import (
	"bytes"
	"json"
	"http"
	"io/ioutil"
	"os"
	"strings"
)

func ShortenURL(longUrl string) (shortenUrl string, err os.Error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	err = enc.Encode(map[string]string{"longUrl": longUrl})
	if err != nil {
		return
	}
	res, err := http.Post("https://www.googleapis.com/urlshortener/v1/url", "application/json", strings.NewReader(buf.String()))
	if err != nil {
		return
	}
	if res.StatusCode != 200 {
		err = os.NewError("failed to post")
		return
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	shortenUrl = string(b)
	return
}

func ExpandURL(shortUrl string) (expandedUrl string, err os.Error) {
	param := http.EncodeQuery(map[string][]string{"shortUrl": {shortUrl}})
	res, _, err := http.Get("https://www.googleapis.com/urlshortener/v1/url?" + param)
	if err != nil {
		return
	}
	if res.StatusCode != 200 {
		err = os.NewError("failed to post")
		return
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	var buf bytes.Buffer
	buf.Write(b)
	dec := json.NewDecoder(&buf)
	var out map[string]interface{}
	dec.Decode(&out)
	expandedUrl = out["longUrl"].(string)
	return
}
