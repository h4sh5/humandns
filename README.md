# humandns

Liberating DNS for humans. No more registrars; generate, register and act as authoritative server for human readable servers.

This works by hashing the IP address the request comes from using [humanhash](https://github.com/wolfeidau/humanhash), which gives each IP a unique hash. 

The dns name is in the form of `one-two-three-four.ip4` for ipv4 and `one-two-three-four-five.ip6` for ipv6 addresses.

To register, simply send a HTTP request (any HTTP request) to the endpoint. A DNS name will be allocated to the IP address it came from and it will stay the same for each IP. IPv4 and 6 are both supported for registration.

A reference implementation of the DNS server that queries the database is [humandns53](https://github.com/h4sh5/humandns53)

