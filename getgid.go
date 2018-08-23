package main

import (
	"net/http"
	"io/ioutil"
	"regexp"
)

func Getgid(url string)string{

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authority", "www.javbus.in")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("dnt", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("referer", "https://www.javbus.in/")
	req.Header.Add("accept-encoding", "deflate")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Add("cookie", "__cfduid=dfa619587ab865b77e201ff4d19b143c41535036165; PHPSESSID=u46kverte713bnl6d8rbceqd84; HstCfa3159616=1535036168467; HstCmu3159616=1535036168467; HstCnv3159616=1; HstCns3159616=3; HstCla3159616=1535041470519; HstPn3159616=8; HstPt3159616=8")
	req.Header.Add("postman-token", "be7db0de-7fdf-e3f4-ff0a-08f1c7824aae")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	content:=string(body)

	reg := regexp.MustCompile(`var gid = (\d+);`)
	allMagent:=reg.FindStringSubmatch(content)
	gid:=allMagent[1]
	//fmt.Print(gid)

	return gid

}