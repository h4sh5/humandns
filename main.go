package main


import (
    "log"
    "net"
    "github.com/wolfeidau/humanhash"
)

func main() {
	ip4str := "192.168.1.2"
	ip6str := "2606:4700:4700::1111"
	ip4 := net.ParseIP(ip4str) //[]byte{192,168,1,1}
	ip6 := net.ParseIP(ip6str)
	log.Printf("ipv4 parsed: %#v", ip4)
	log.Printf("ipv6 parsed: %#v", ip6)

	result4, _ := humanhash.Humanize(ip4, 4)
	result6_round1, _ := humanhash.Humanize(ip6, 8)
	// double hash it so it has enough entropy
	result6, _ := humanhash.Humanize([]byte(result6_round1), 5) // make this one 5 words so that it's distinguishable from ipv4
	log.Printf("ipv4 name for address %s = %s.ip4", ip4str, result4)
	log.Printf("ipv4 name for address %s = %s.ip6", ip6str, result6)

}