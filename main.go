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

type Host struct {
	address string
	t       int64
}

type Hosts []Host

func (h Hosts) Len() int {
	return len(h)
}

func (h Hosts) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hosts) Less(i, j int) bool {
	return h[i].t < h[j].t
}

func Ping(address string) (int64, error) {
	now := time.Now()
	u := fmt.Sprintf("%s:22", address)
	dialer := net.Dialer{
		Timeout: time.Second * time.Duration(5),
	}
	conn, err := dialer.Dial("tcp", u)
	if err != nil {
		return 0, err
	}
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)
	t := time.Now().Sub(now).Milliseconds()
	return t, nil
}

func main() {
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

	var hosts Hosts
	var lock sync.Mutex
	var wg = sync.WaitGroup{}
	for _, v := range meta.Git {
		wg.Add(1)
		go func(i string) {
			defer wg.Done()
			if strings.Contains(i, ":") {
				return
			}
			ips, _ := iputils.CidrToIpsClean(i)
			address := ips[0]
			t, err := Ping(address)
			if err != nil {
				return
			}
			lock.Lock()
			hosts = append(hosts, Host{
				address: address,
				t:       t,
			})
			lock.Unlock()
		}(v)
	}
	wg.Wait()

	sort.Sort(hosts)

	for _, v := range hosts {
		fmt.Printf("%s github.com\n", v.address)
	}
}
