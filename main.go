package main

import (
	"strconv"
	"net"
	"fmt"
)

const BEGIN_RANGE = 49152
const END_RANGE   = 65535

func main() {
	used := make([]int, 0)

	for port := BEGIN_RANGE; port <= END_RANGE; port++ {
		portString := ":" + strconv.Itoa(port)
		conn, err := net.Listen("tcp", portString)

		if err != nil {
			used = append(used, port)
		} else {
			conn.Close()
		}
	}

	fmt.Println(used)
}
