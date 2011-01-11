package main

import "google/urlshortener"
import "bytes"
import "flag"
import "fmt"
import "os"
import "template"

var expand = flag.Bool("e", false, "expand URL")
var info = flag.Bool("i", false, "show info about URL")

func analytics(a urlshortener.AnalyticsInfo) {
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
{.repeated section Referrers}    {Id} ({Count})
{.end}
  Countries:
{.repeated section Countries}    {Id} ({Count})
{.end}
  Browsers:
{.repeated section Browsers}    {Id} ({Count})
{.end}
  Platforms:
{.repeated section Platforms}    {Id} ({Count})
{.end}
{.end}
`, nil)
	tmpl.Execute(a, &b)
	println(b.String())
}

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "usage: urlshortener [-e|-i] URL\n")
		os.Exit(2)
	}
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
	}

	if *expand {
		u, e := urlshortener.ExpandURL(args[0])
		if e != nil {
			panic(e.String())
		} else {
			println(u)
		}
	} else if *info {
		a, e := urlshortener.AnalyticsURL(args[0])
		if e != nil {
			panic(e.String())
		} else {
			analytics(a)
		}
	} else {
		u, e := urlshortener.ShortenURL(args[0])
		if e != nil {
			panic(e.String())
		} else {
			println(u)
		}
	}
}
