package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"log"
)

var siteName = "M-TEAM"
var siteUrl = "https://tp.m-team.cc/torrents.php"
var siteCookie = ""

// It always would be the first part of your site, like: urlHalf = "https://tp.m-team.cc/"
var urlHalf = "https://tp.m-team.cc/"

// If your site is a Gazelle Site, please change this varible to True, like: isGazelle = True
var isGazelle = false

// If the torrents' download url in your site were encrypted like HDC, please change this varible to True, like: isEncrypted = True
var isEncrypted = false

// Check the download url, especially when you are using a https(SSL) url.
// Some torrents' download pages url could be "https://tp.m-team.cc/download.php?id=xxxxx&https=1", in this case, you need to assign the variable of "url_last". Examples:
// url_last = "&https=1"

// If you couldn't downlaod the torrents to your directory where the *.py script is, you could just define the variables below. Make sure the format of your path because of the difference between Windows and Linux.
var monitorPath = "/opt/freetorrent/torrent/"

// Other informations for safer sites. Complete it if you cannot download torrents.
var userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"

// You don't need to define the variables shows below unless you couldn't download the torrents after defined the above one
var referer = "https://tp.m-team.cc/login.php"
var host = "tp.m-team.cc"

// You don't need to define the variables shows below unless you couldn't download the torrents after defined the above two
var upgradeInsecureRequests = ""
var dnt = ""
var acceptLanguage = ""
var acceptEncoding = ""
var accept = ""
var cacheControl = ""
var contentLength = ""
var contentType = ""
var origin = ""

// Only if you just want to check the first 10 torrents is free in the page and download the free torrents in this small amount, please change it to 10
// We always grab all the torrents in the whole page, but you can define the amount of grabing torrents by defining the variable below
var torrentsAmount = 0

// You don't need to change this variables unless you cannot download from your GAZELLE site
// check this value from the page source code
var colspan = "3"

// You don't need to change this variables unless you cannot find free torrents correctly
var freeTag = "pro_free"
var freeTag2 = "pro_free2up"
var DicFreeTag = "torrent_label tooltip tl_free"

//PTP_free_tag = 'torrent-info__download-modifier--free'

var torrentsClassName = ".torrentname"
var HdcTorrentsClassName = ".t_name"
var DicTorrentsClassName = ".td_info"

//PTP_torrents_class_name = '.basic-movie-list__torrent-row--user-seeding'

var downloadClassName = ".rowfollow"
var HdcDownloadClassName = ".torrentdown_button"

var headers = map[string]string{
	"User-Agent":                userAgent,
	"Referer":                   referer,
	"Host":                      host,
	"accept":                    accept,
	"accept-language":           acceptLanguage,
	"accept-encoding":           acceptEncoding,
	"origin":                    origin,
	"dnt":                       dnt,
	"upgrade-insecure-requests": upgradeInsecureRequests,
	"cache-control":             cacheControl,
	"content-length":            contentLength,
	"content-type":              contentType,
}

func main() {

	fetch()
}

func fetch() soup.Root {
	fmt.Println("Fetch Url", siteUrl)
	soup.Headers = setHeader()
	soup.Cookie("tp", siteCookie)

	source, err := soup.Get(siteUrl)
	if err != nil {
		log.Fatal(err)
	}
	doc := soup.HTMLParse(source)
	trs := doc.Find("table", "class", "torrents").FindAll("td", "class", "torrenttr")
	var res []string
	for _, tr := range trs {

		img := tr.Find("td", "class", "embedded").Find("img", "class", "pro_free")

		if img.Error == nil {
			link := tr.Find("td", "class", "embedded").Find("a")
			fmt.Println(link.Attrs()["href"])
			res = append(res, link.Attrs()["href"])
		}

	}
	return doc
}

func setHeader() map[string]string {
	var res = map[string]string{}
	for s := range headers {
		if headers[s] != "" {
			res[s] = headers[s]
		}
	}
	return res
}
