package main

import (
	"flag"
	"fmt"
)

func main() {
	quietPtr := flag.Bool("q", false, "quiet")
	protocolPtr := flag.String("p", "tcp", "protocol tcp/ssh")
	modePtr := flag.String("m", "web", "mode web/api/git")
	flag.Parse()

	port := 443
	switch *protocolPtr {
	case "ssh":
		port = 22
	}

	github := NewGithub(port, *modePtr)
	ips := github.GetIps()
	for _, v := range ips {
		if *quietPtr {
			fmt.Printf("%s github.com\n", v.address)
			continue
		}
		fmt.Printf("%s github.com %dms\n", v.address, v.time)
	}
}
