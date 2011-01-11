package main

import "google/urlshortener"
import "bytes"
import "template"

func main() {
	a, e := urlshorter.AnalyticsURL("http://goo.gl/M1xh")
	if e != nil {
		println(e.String())
	} else {
		var b bytes.Buffer
		tmpl, _ := template.Parse(`
Kind: {Kind}
Id: {Id}
LongUrl: {LongUrl}
Status: {Status}
Created: {Created}
AllTime:
{.section Analytics.AllTime}
  ShortUrlClicks: {ShortUrlClicks}
  LongUrlClicks: {LongUrlClicks}
  Referrers:
    {.repeated section Referrers}{@}{.end}
  Countries:
    {.repeated section Countries}{@}{.end}
  Browsers:
    {.repeated section Browsers}{@}{.end}
  Platforms:
    {.repeated section Platforms}{@}{.end}
{.end}
`,
			nil)
		tmpl.Execute(a, &b)
		println("", b.String())
	}
}
