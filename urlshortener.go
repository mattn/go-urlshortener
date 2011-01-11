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
	err = dec.Decode(&out)
	if err != nil {
		return
	}
	expandedUrl = out["longUrl"].(string)
	return
}

type AnalyticsCount struct {
    Count string
    Id string
}

type AnalyticsItem struct {
	ShortUrlClicks string
	LongUrlClicks string
	Referrers []AnalyticsCount
	Countries []AnalyticsCount
	Browsers []AnalyticsCount
	Platforms []AnalyticsCount
}

type AnalyticsInfo struct {
 Kind string
 Id string
 LongUrl string
 Status string
 Created string
 Analytics struct {
  AllTime AnalyticsItem
  Month AnalyticsItem
  Week AnalyticsItem
  Day AnalyticsItem
  TwoHours AnalyticsItem
 }
}

func AnalyticsURL(shortUrl string) (info AnalyticsInfo, err os.Error) {
	param := http.EncodeQuery(map[string][]string{"shortUrl": {shortUrl}, "projection": {"FULL"}})
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
	err = json.Unmarshal(b, &info)
	return
}
