package main

import (
	"fmt"
	"net/url"
	)
func main() {
	fmt.Print("Enter Your URL : ")
	var link string
	fmt.Scanln(&link)
	u, err := url.Parse(link)
	if err != nil {
		fmt.Println("[-] ERROR")
	} else {
		parserURL(u)
//		fmt.Println(u)
//		fmt.Println(u.Path)
	}
}


