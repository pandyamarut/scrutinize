package main 

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func portScanner(protocol, hname string, port int) string {
	addr := hname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, addr, 60*time.Second)
	if err != nil {
		return "YES, secured port"
	}

	defer conn.Close()
	return "NO, unsecured port"
}

func main() {
	fmt.Println("Checking for the secured port...")
	isopen := portScanner("tcp", "localhost", 1331) // portScanner("UDP", "192.122.12.34", portnumber)
	fmt.Println("is port secured?", isopen)
}