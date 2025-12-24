package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"mime/multipart"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	APIKey      string   `yaml:"apiKey"`
	TorrentPath string   `yaml:"torrentPath"`
	UserAgent   string   `yaml:"userAgent"`
	FreeDays    int      `yaml:"freeDays"`
	FreeSize    float64  `yaml:"freeSize"`
	BlockList   []string `yaml:"blockList"`
}

type Torrent struct {
	ID   string
	Name string
	Size float64
}

type DlTokenResponse struct {
	Message string `json:"message"`
	Data    string `json:"data"`
	Code    string `json:"code"`
}

type TorrentSearchRequest struct {
	PageNumber int    `json:"pageNumber"`
	PageSize   int    `json:"pageSize"`
	Mode       string `json:"mode"`
	Categories []int  `json:"categories"`
	Visible    int    `json:"visible"`
}

type TorrentSearchResponse struct {
	Message string          `json:"message"`
	Data    TorrentListData `json:"data"`
	Code    string          `json:"code"`
}

type TorrentListData struct {
	PageNumber string        `json:"pageNumber"`
	PageSize   string        `json:"pageSize"`
	Total      string        `json:"total"`
	TotalPages string        `json:"totalPages"`
	Data       []TorrentInfo `json:"data"`
}

type TorrentInfo struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Size   string `json:"size"`
	Status struct {
		Discount        string `json:"discount"`
		DiscountEndTime string `json:"discountEndTime"`
	} `json:"status"`
}

var host = "api.m-team.cc"
var baseUrl = "https://" + host
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

// 在 fetchTorrents 函数中修改大小转换部分
func fetchTorrents() {
	Info.Println(time.Now())

	searchReq := TorrentSearchRequest{
		PageNumber: 1,
		PageSize:   100,
		Mode:       "normal",
		Categories: []int{},
		Visible:    1,
	}

	jsonData, err := json.Marshal(searchReq)
	if err != nil {
		Error.Fatal(err)
	}

	req, err := http.NewRequest("POST", baseUrl+"/api/torrent/search", bytes.NewBuffer(jsonData))
	if err != nil {
		Error.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.APIKey)
	req.Header.Set("User-Agent", c.UserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		Error.Fatal(err)
	}
	defer resp.Body.Close()

	var searchResp TorrentSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		Error.Fatal(err)
	}

	var res []*Torrent
	for _, t := range searchResp.Data.Data {
		if t.Status.Discount == "FREE" {
			// 如果discountEndTime为空则代表永久免费
			var expireTime time.Time
			if t.Status.DiscountEndTime == "" {
				expireTime = time.Now().AddDate(100, 0, 0) // 设置一个很久以后的时间作为到期时间
			} else {
				expireTime, err = time.Parse("2006-01-02 15:04:05", t.Status.DiscountEndTime)
			}
			if err != nil {
				continue
			}

			daysLeft := int(time.Until(expireTime).Hours() / 24)
			if daysLeft < c.FreeDays {
				continue
			}

			size, _ := strconv.ParseFloat(t.Size, 64)
			sizeGB := size / (1024 * 1024 * 1024)

			torrent := NewTorrent(t.ID, t.Name, sizeGB)
			res = append(res, torrent)
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
		err := DownloadFile(c.TorrentPath+"[M-TEAM]"+t.Name+".torrent", id)
		if err != nil {
			panic(err)
		}
		Info.Println("Downloaded torrent: " + t.Name)
	}
}

func DownloadFile(filepath string, torrentId string) error {
	// 构造获取下载链接的请求
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("id", torrentId)
	writer.Close()

	req, err := http.NewRequest("POST", baseUrl+"/api/torrent/genDlToken", body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("x-api-key", c.APIKey)
	req.Header.Set("User-Agent", c.UserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var tokenResp DlTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return err
	}

	if tokenResp.Code != "0" {
		return fmt.Errorf("get download token failed: %s", tokenResp.Message)
	}

	// 使用获取到的下载链接下载种子文件
	dlReq, err := http.NewRequest("GET", tokenResp.Data, nil)
	if err != nil {
		return err
	}

	// 设置请求头
	dlReq.Header.Set("User-Agent", c.UserAgent)

	// 发送请求
	dlResp, err := client.Do(dlReq)
	if err != nil {
		return err
	}
	defer dlResp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, dlResp.Body)
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
