package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Conf struct {
	SiteCookie  string  `yaml:"siteCookie"`
	PassKey     string  `yaml:"passKey"`
	TorrentPath string  `yaml:"torrentPath"`
	FreeDays    int     `yaml:"freeDays"`
	FreeSize    float64 `yaml:"freeSize"`
}

var siteUrl = "https://tp.m-team.cc/torrents.php"
var userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"
var referer = "https://tp.m-team.cc/login.php"
var downloadURL = "https://kp.m-team.cc/download.php"
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

var c Conf

func main() {

	c.getConf()
	fetch()
}

func fetch() soup.Root {
	fmt.Println("Fetch Url", siteUrl)
	soup.Headers = setHeader()
	soup.Cookie("tp", c.SiteCookie)

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
			span := tr.Find("td", "class", "embedded").Find("span")
			if span.Pointer != nil {
				date := strings.Split(span.Pointer.FirstChild.Data, "：")
				dateValue := date[1]
				if strings.Contains(dateValue, "日") {
					num, err := strconv.Atoi(dateValue[0:1])
					if err != nil || num < c.FreeDays {
						continue
					}
				} else {
					continue
				}
			}

			sizeStr := tr.FindNextSibling().FindNextSibling().FindNextSibling().Pointer.FirstChild.Data
			size, err := strconv.ParseFloat(sizeStr, 32)
			if err != nil || size > c.FreeSize {
				continue
			}

			link := tr.Find("td", "class", "embedded").Find("a")
			fmt.Println(link.Attrs()["href"])
			res = append(res, link.Attrs()["href"])
		}
	}

	for _, detail := range res {
		tmp := strings.Split(detail, "?")
		tmp1 := strings.Split(tmp[1], "&")
		id := strings.Split(tmp1[0], "=")

		downloadURL := downloadURL + "?" + tmp1[0] + "&passkey=" + c.PassKey + "&https=1"
		err := DownloadFile(c.TorrentPath+"[M-TEAM]"+id[1]+".torrent", downloadURL)
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded: " + downloadURL)
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

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func (c *Conf) getConf() *Conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
