package ports

import (
	"strconv"
	"net"
)

const BEGIN_RANGE = 49152
const END_RANGE   = 65535

func NextAvailPort() int {
	for port := BEGIN_RANGE; port <= END_RANGE; port++ {
		if(IsOpen(port)) {
			return port
		}
	}
	return -1;
}

func ListUsedPorts() []int {
	used := make([]int, 0)

	for port := BEGIN_RANGE; port <= END_RANGE; port++ {
		if(!IsOpen(port)) {
			used = append(used, port)
		}
	}

	return used
}

func IsOpen(port int) bool {
	conn, _ := net.Dial("tcp", net.JoinHostPort("", strconv.Itoa(port)))
	if conn != nil {
		conn.Close();
		return false;
	}
	return true;
}
