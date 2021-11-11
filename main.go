package main

import (
	"flag"
	"github.com/anaskhan96/soup"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Conf struct {
	SiteCookie  string   `yaml:"siteCookie"`
	PassKey     string   `yaml:"passKey"`
	UserAgent   string   `yaml:"userAgent"`
	TorrentPath string   `yaml:"torrentPath"`
	FreeDays    int      `yaml:"freeDays"`
	FreeSize    float64  `yaml:"freeSize"`
	BlockList   []string `yaml:"blockList"`
}

type Torrent struct {
	ID   string
	Name string
	Size float64
}

var host = "kp.m-team.cc"
var baseUrl = "https://" + host
var siteUrl = baseUrl + "/torrents.php"
var referer = baseUrl + "/login.php"
var downloadURL = baseUrl + "/download.php"

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
var configFlag string

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	flag.StringVar(&configFlag, "c", "conf.yaml", "config file path")
	file, err := os.OpenFile("freeTorrent.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	flag.Parse()
	c.getConf()
	deleteTorrents()
	fetchTorrents()
}

func deleteTorrents() {
	files, err := filepath.Glob(c.TorrentPath + "*")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
}

func fetchTorrents() soup.Root {
	Info.Println(time.Now())
	soup.Headers = setHeader(c)
	soup.Cookie("tp", c.SiteCookie)

	source, err := soup.Get(siteUrl)
	if err != nil {
		Error.Fatal(err)
	}
	doc := soup.HTMLParse(source)
	trs := doc.Find("table", "class", "torrents").FindAll("td", "class", "torrenttr")
	var res []*Torrent
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

			sizeUnit := tr.FindNextSibling().FindNextSibling().FindNextSibling().Pointer.LastChild.Data
			sizeStr := tr.FindNextSibling().FindNextSibling().FindNextSibling().Pointer.FirstChild.Data
			size, err := strconv.ParseFloat(sizeStr, 32)

			switch sizeUnit {
			case "MB":
				size = size / 1024
			case "TB":
				size = size * 1024
			}

			if err != nil{
				continue
			}

			//torrent id
			link := tr.Find("td", "class", "embedded").Find("a")
			//href will be like "details.php?id=523177&hit=1"
			href := link.Attrs()["href"]
			tmp := strings.Split(href, "=")
			tmp1 := strings.Split(tmp[1], "&")
			id := tmp1[0]

			//torrent name
			title := link.Attrs()["title"]

			Info.Println(id)
			Info.Println(title)
			t := NewTorrent(id, title, size)

			res = append(res, t)
		}
	}

	for _, t := range res {
		if t.Size > c.FreeSize {
			continue
		}

		matched := false
		for _, b := range c.BlockList {
			matched, _ = regexp.MatchString(b, t.Name)
			if matched {
				Info.Println("Matched: " + b)
				break
			}
		}
		if matched {
			continue
		}

		id := t.ID
		torrentURL := downloadURL + "?id=" + id + "&passkey=" + c.PassKey + "&https=1"

		err := DownloadFile(c.TorrentPath+"[M-TEAM]"+t.Name+".torrent", torrentURL)
		if err != nil {
			panic(err)
		}
		Info.Println("Downloaded: " + torrentURL)
	}
	return doc
}

func setHeader(c Conf) map[string]string {
	var res = map[string]string{}
	for s := range headers {
		if headers[s] != "" {
			res[s] = headers[s]
		}
	}
	res["User-Agent"] = c.UserAgent
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
	yamlFile, err := ioutil.ReadFile(configFlag)
	if err != nil {
		Error.Printf("yamlFile.Get err #%v", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		Error.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func NewTorrent(id string, name string, size float64) *Torrent {
	t := new(Torrent)
	t.ID = id
	t.Name = name
	t.Size = size
	return t
}