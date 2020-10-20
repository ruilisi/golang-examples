## flag

### [cidr.go](cidr.go)

#### introduction

This program implements a CIDR subnet calculator.
It takes a CIDR address prefix an calculates ip-ranges,
total hosts, wildcard mask, etc.

#### the function

Given IPv4 block 192.168.100.14/24
The followings uses IPNet to get:
- The routing address for the subnet (i.e. 192.168.100.0)
- one-bits of the network mask (24 out of 32 total)
- The subnetmask (i.e. 255.255.255.0)
- Total hosts in the network (2 ^(host identifer bits) or 2^8)
- Wildcard the inverse of subnet mask (i.e. 0.0.0.255)
- The maximum address of the subnet (i.e. 192.168.100.255)

#### how to use

`flag.StringVar(&cidr, "c", "", "the CIDR address")`

run code `go run cidr.go -c ip/cidr`


### [ipinfo.go](ipinfo.go) 

#### intruduction

Easiest way to create net.IP value is to use
net.ParseIP which parses a string value representation
of a IPv4 dot-separated or IPv6 colon-separated address.
This example uses net.ParseIP to parse an IP address provided
from the command line and prints information about the address.

#### how to use

run code `go run ipinfo.go ip`

