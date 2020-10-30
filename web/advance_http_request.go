package main

import (
	"fmt"
	"github.com/asmcos/requests"
	"strings"
	)


func main() {
	fmt.Print("Enter Your URL : ")
	var url string
	fmt.Scanln(&url)
	fmt.Print("Enter The Method : ")
	var method string
	fmt.Scanln(&method)
	the_url := strings.ReplaceAll(url,"*","<img src=x onerror=alert(1)>")
	req := requests.Requests()
	req.Debug = 1
	req.SetTimeout(20)
	req.Header.Set("User-agent","ScanT3r(github.com/knassar702/scant3r)V0.1")
	if method == "POST" {
		resp1,err1 := req.Post(the_url)
		if err1 != nil {
			return
		} else {
			fmt.Println(resp1.Text())
		}
	}
	resp,err := req.Get(the_url)
	if err != nil {
		return
	} else {
		fmt.Println(resp.Text())
	}
}
