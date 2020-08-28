/****************************************

* File Name : stringer.go

* Creation Date : 27-08-2020

* Last Modified : Thursday 27 August 2020 10:48:06 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import "fmt"

type IPAddr [4]byte

// Add a "String() string" method to IPAddr. Stringer is exported by fmt and is an interface that describes itself as string
func (i IPAddr) String() string {
    return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
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

