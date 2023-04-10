package main

import (
	"flag"
	"fmt"
)

func main() {
	quietPtr := flag.Bool("q", false, "quiet")
	portPtr := flag.Int("p", 22, "port")
	modePtr := flag.String("m", "git", "mode web/api/git")
	flag.Parse()

	github := NewGithub(*portPtr, *modePtr)
	ips := github.GetIps()
	for _, v := range ips {
		if *quietPtr {
			fmt.Printf("%s github.com\n", v.address)
			continue
		}
		fmt.Printf("%s github.com %dms\n", v.address, v.time)
	}
}
