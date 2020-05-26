//Go Algorithm for fetching a local IPv4 address
//Implementation 1 for Network Dynamics
//Loopback method for finding the "external" IP address for the machine.
//

package main

import (
	"errors"
	"fmt"
	"net"
)

func main() {
	ip, err := localIP()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ip)
	}
}

func localIP() (string, error) {

	netv1, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, netv1 := range netv1 {
		if netv1.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if netv1.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := netv1.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // for a non-IPv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("error: check network connection")

}
