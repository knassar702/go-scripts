
package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main(){
	resp,err := http.Get("http://httpbin.org/")
	if err != nil {
		fmt.Println("HTTP ERROR ..!")
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}
}
