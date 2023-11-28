package main

import (
	"encoding/json"
	"fmt"
	"github.com/lizongying/go-ip-utils/iputils"
	"io"
	"log"
	"net"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

type Meta struct {
	VerifiablePasswordAuthentication bool `json:"verifiable_password_authentication"`
	SSHKeyFingerprints               struct {
		SHA256RSA     string `json:"SHA256_RSA"`
		SHA256ECDSA   string `json:"SHA256_ECDSA"`
		SHA256ED25519 string `json:"SHA256_ED25519"`
	} `json:"ssh_key_fingerprints"`
	SSHKeys    []string `json:"ssh_keys"`
	Hooks      []string `json:"hooks"`
	Web        []string `json:"web"`
	API        []string `json:"api"`
	Git        []string `json:"git"`
	Packages   []string `json:"packages"`
	Pages      []string `json:"pages"`
	Importer   []string `json:"importer"`
	Actions    []string `json:"actions"`
	Dependabot []string `json:"dependabot"`
}

type Ip struct {
	address string
	time    int64
}

type Ips []Ip

func (i Ips) Len() int {
	return len(i)
}

func (i Ips) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func (i Ips) Less(a, b int) bool {
	return i[a].time < i[b].time
}

type Github struct {
	port   int
	mode   string
	dialer *net.Dialer
}

func (g *Github) Ping(address string) (int64, error) {
	now := time.Now()
	u := fmt.Sprintf("%s:%d", address, g.port)
	conn, err := g.dialer.Dial("tcp", u)
	if err != nil {
		return 0, err
	}
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)
	t := time.Now().Sub(now).Milliseconds()
	return t, nil
}

func (g *Github) GetIps() (ips Ips) {
	u := "https://api.github.com/meta"
	r, err := http.Get(u)
	if err != nil {
		log.Println(err)
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	var meta Meta
	err = json.Unmarshal(b, &meta)
	if err != nil {
		log.Println(err)
		return
	}

	var lock sync.Mutex
	var wg = sync.WaitGroup{}
	var ipList []string
	switch g.mode {
	case "web":
		ipList = meta.Web
	case "git":
		ipList = meta.Git
	case "api":
		ipList = meta.API
	default:
	}
	for _, v := range ipList {
		wg.Add(1)
		go func(i string) {
			defer wg.Done()
			if strings.Contains(i, ":") {
				return
			}
			ip, _ := iputils.CidrToIpsClean(i)
			if len(ip) == 0 {
				return
			}
			address := ip[0]
			t, e := g.Ping(address)
			if e != nil {
				return
			}
			lock.Lock()
			ips = append(ips, Ip{
				address: address,
				time:    t,
			})
			lock.Unlock()
		}(v)
	}
	wg.Wait()

	sort.Sort(ips)
	return ips
}

func NewGithub(port int, mode string) (github *Github) {
	github = &Github{
		port: port,
		mode: mode,
		dialer: &net.Dialer{
			Timeout: time.Second * time.Duration(5),
		},
	}
	return
}
