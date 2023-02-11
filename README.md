# humandns

Liberating DNS for humans. No more registrars; generate, register and act as authoritative server for human readable servers.

Imagine a world where every single person and business have a `/64` IPv6 block (that's the mininum assignment!). That's 18 Billion Billion addresses! Even if you give them a `/96`, that's 4 Billion addresses, the size of the IPv4 internet in its entirety.

With that many number of addresses, there's enough for everyone; but IPv6 addresses are very long, and impossible to send to someone via a human readable format unless you buy a domain. Forget about domains; those should be free!

This is where humandns comes in - it works by hashing the IP address the request comes from using [humanhash](https://github.com/wolfeidau/humanhash), which gives each IP a unique hash. You just get what you're given, the hashing algorithm ensures (mostly) that the same IP gets the same name. Just proof you own the IP address by sending a HTTP request from it.

The dns name is in the form of `one-two-three-four.ip4` for ipv4 and `one-two-three-four-five.ip6` for ipv6 addresses.

To register, simply send a HTTP request (any HTTP request) to `/rego`. A DNS name will be allocated to the IP address it came from and it will stay the same for each IP. IPv4 and 6 are both supported for registration.

To visit/resolve a humandns name in the absense of DNS (which should be available on the same host via port 53), use `/visit?d=your-domain-name`. This will test if your domain name resolves properly.

No more DNS sync issues! To update humandns servers, simply send a `/rego` request to all of them. By default, the TTL is 30 minutes so unused IPs are deleted from the Redis cache to save space.

A reference implementation of the DNS server that queries the database is [humandns53](https://github.com/h4sh5/humandns53)

