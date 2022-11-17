package main

import (
	"flag"
	"fmt"
)

func main() {
	isQuietF := flag.Bool("q", false, "isQuiet")
	flag.Parse()
	isQuiet := *isQuietF
	ips := GetIps()
	for _, v := range ips {
		if isQuiet {
			fmt.Printf("%s github.com\n", v.address)
			continue
		}
		fmt.Printf("%s github.com %dms\n", v.address, v.t)
	}
}
