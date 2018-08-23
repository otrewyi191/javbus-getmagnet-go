package main

import (
	"regexp"
	"strings"
)

func GetMagnet(content string)string{
	reg := regexp.MustCompile(`'(magnet:.*?)'`)
	allMagent:=reg.FindAllString(content, -1)
	result:=strings.Join(allMagent,"\n")
	result=strings.Replace(result,"'","",-1)
	//fmt.Print(allMagent)
	return result
}
