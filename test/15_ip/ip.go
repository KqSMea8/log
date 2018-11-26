package main

import (
	"fmt"
	"net"
	"strings"
)

func getLocalIp() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	addrMap := make(map[string]string, 0)

	for _, inter := range interfaces {
		addrs, _ := inter.Addrs()

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					addrMap[inter.Name] = strings.Replace(ipnet.IP.String(),".","_",-1)
				}
			}
		}
	}
	if v, ok := addrMap["bond0"]; ok {

		return v, nil
	} else if v, ok := addrMap["eth1"]; ok {
		return v, nil
	} else if v, ok := addrMap["eth2"]; ok {
		return v, nil
	} else {
		return "", fmt.Errorf("get fail")
	}

}

func main() {
	addr, err := getLocalIp()
	fmt.Println(addr, err)

}

////以太网网卡名称为eth0
//inter, err := net.InterfaceByName("eth0")
//if err != nil {
//	log.Fatalln(err)
//}
////mac地址
//fmt.Println("str:", inter.HardwareAddr.String())
//addrs, err := inter.Addrs()
//if err != nil {
//	log.Fatalln(err)
//}
////ip地址一个ip4一个ip6
//for _, addr := range addrs {
//	fmt.Println(addr.String())
//}
//}
