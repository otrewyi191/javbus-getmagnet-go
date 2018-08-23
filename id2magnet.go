package main

func Id2Magnet(id string)string {
	url:="https://www.javbus.in/"+id
	gid:=Getgid(url)
	//fmt.Print(gid)

	content:=Get_url_content(gid)
	result:=GetMagnet(content)
	//fmt.Print(content)
	//fmt.Print(result)
	return result
}