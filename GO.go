package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
    "sync"
    )
// Usage :
//        $ cat hosts | ./Go
func main() {
    sc := bufio.NewScanner(os.Stdin)
    var wg sync.WaitGroup
    for sc.Scan(){
        domain := sc.Text()
        wg.Add(1)
        go func(){
            defer wg.Done()
            _,err := net.LookupHost(domain)
            if err != nil {
                return
            }
            fmt.Println(domain)
        }()
    }
    wg.Wait()
}
