package main

import (
	"net/http"
	"io/ioutil"
)

func Get_url_content(gid string)string {

	url := "https://www.javbus.in/ajax/uncledatoolsbyajax.php?gid="+gid+"&lang=zh&img=https%3A%2F%2Fpics.javcdn.pw%2Fcover%2F6ooo_b.jpg&uc=0&floor=403"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cookie", "__cfduid=dfa619587ab865b77e201ff4d19b143c41535036165; PHPSESSID=u46kverte713bnl6d8rbceqd84; HstCfa3159616=1535036168467; HstCmu3159616=1535036168467; HstCnv3159616=1; HstCla3159616=1535037887419; HstPn3159616=4; HstPt3159616=4; HstCns3159616=2")
	req.Header.Add("dnt", "1")
	req.Header.Add("accept-encoding", "deflate")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	req.Header.Add("accept", "*/*")
	req.Header.Add("referer", "https://www.javbus.in/HUNTA-484")
	req.Header.Add("authority", "www.javbus.in")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "974d1672-6289-7d1b-6728-4aea38a29cfd")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))
	return string(body)

}