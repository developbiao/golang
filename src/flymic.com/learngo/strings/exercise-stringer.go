package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (addr IPAddr) String() string {
	s := make([]string, len(addr))
	for i, val := range addr {
		s[i] = strconv.Itoa(int(val))
	}
	return strings.Join(s, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

}
