package main


import (
    "log"
    "net"
    "fmt"
    "net/http"
    "strings"
    // "context"
    "github.com/wolfeidau/humanhash"
    "github.com/go-redis/redis" // use redis to store things
)


// setup redis
var redisClient *redis.Client 

/**
 * takes an IPv4 or IPv6 address and convert it to humandns name
 * return "error" if errored
 * 
 * make sure the port part e.g. the :1234 in 127.0.0.1:1234 is NOT included in the ip
 */
func IPtoHumanDNS(ipstr string) string {

	parsed := net.ParseIP(ipstr)
	if parsed == nil {
		return "error"
	}

	if strings.Contains(ipstr, ":") { // identified ipv6
		log.Printf("[IPtoHash] ipv6 parsed: %#v", parsed)

		result6_round1, err1 := humanhash.Humanize(parsed, 8)
		if err1 != nil {
			log.Printf("%s", err1)
			return "error"
		}
		// double hash it so it has enough entropy
		result6, err2 := humanhash.Humanize([]byte(result6_round1), 5) // make this one 5 words so that it's distinguishable from ipv4
		if err2 != nil {
			log.Printf("%s", err2)
			return "error"
		}

		return result6 + ".ip6"
	} else {
		log.Printf("[IPtoHash] ipv4 parsed: %#v", parsed)
		result4, err := humanhash.Humanize(parsed, 4)
		if err != nil {
			log.Printf("%s", err)
			return "error"
		}
		return result4

	}
}

func storeMapping(dns string, ip net.IP) {
	// ctx := context.TODO()
	// result, err := client.Append()
	dnsRes := redisClient.Get(dns)
	// log.Printf("[storeMapping] val %#v", dnsRes.Val())
	if dnsRes.Val() == "" { // not in redis yet
		// log.Printf("[storeMapping] from is nil", dnsRes)
		ipString := ip.String()
		// key, value, expiration time in nanoseconds (0 means no expiration)
		setRes := redisClient.Set(dns, ipString, 0)

		log.Printf("[storeMapping] adding name %s as %s (result:%v) ", dns, ipString, setRes)
	}
	log.Printf("[storeMapping] from db: %v", dnsRes)

}

func homePage(w http.ResponseWriter, r *http.Request){
    
    // if r.Method == "GET" {

	// log.Printf("r.RemoteAddr: %s", r.RemoteAddr)
	remoteAddrParts := strings.Split(r.RemoteAddr, ":")
	remoteIP := strings.Join(remoteAddrParts[:len(remoteAddrParts)-1], ":") // must handle both ipv4 and 6

	if strings.Contains(remoteIP, "[") { // ipv6, take the brackets away before parsing
		remoteIP = strings.Split(remoteIP, "[")[1]
		remoteIP = strings.Split(remoteIP, "]")[0]

		// log.Printf("ipv6 string processed: %#v", ipstr)
	}

	// log.Printf("remoteIP: %s", remoteIP)
	resultDNS := IPtoHumanDNS(remoteIP)
	if resultDNS != "error" {
		go storeMapping(resultDNS, net.ParseIP(remoteIP))
	}
	
	log.Printf("%s -> %s", remoteIP, resultDNS)
	fmt.Fprintf(w, "%s", resultDNS)

    
}


func handleRequests() {
	log.Printf("handling requests now..")
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":80", nil))
}
 

func main() {
	ip4str := "192.168.1.2"
	ip6str := "2606:4700:4700::1111"
	ip4 := net.ParseIP(ip4str) //[]byte{192,168,1,1}
	ip6 := net.ParseIP(ip6str)
	log.Printf("(example) ipv4 parsed: %#v", ip4)
	log.Printf("(example) ipv6 parsed: %#v", ip6)

	result4, _ := humanhash.Humanize(ip4, 4)
	result6_round1, _ := humanhash.Humanize(ip6, 8)
	// double hash it so it has enough entropy
	result6, _ := humanhash.Humanize([]byte(result6_round1), 5) // make this one 5 words so that it's distinguishable from ipv4
	log.Printf("(example) ipv4 name for address %s = %s.ip4", ip4str, result4)
	log.Printf("(example) ipv4 name for address %s = %s.ip6", ip6str, result6)

	redisClient = redis.NewClient(&redis.Options{
	    Addr: "localhost:6379",
	    Password: "",
	    DB: 0,
	})

	handleRequests()

}