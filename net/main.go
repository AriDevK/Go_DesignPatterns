package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "Site to scann ports")

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	for port := 0; port < 65325; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}

			_ = conn.Close()
			fmt.Println("Port", port, "is open")
		}(port)
	}
	wg.Wait()
}
