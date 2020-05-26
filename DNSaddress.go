//Go Algorithm developed for fetching DNS address of given URLs
//Implementation 2 for Network Dynamics

package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("Enter URL:")
	var IP string
	fmt.Scanf("%s", &IP)
	fmt.Println()
	fmt.Println("URL resolves to:")
	fmt.Println(dnsLookup(IP))

}

func dnsLookup(ipStr string) (string, error) {
	ipAddr, err := net.ResolveIPAddr("ip", ipStr)
	return ipAddr.String(), err
}
