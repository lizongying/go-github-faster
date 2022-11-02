package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
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

func Ping(address string) (int64, error) {
	u := fmt.Sprintf("%s:22", address)
	now := time.Now()
	dialer := net.Dialer{
		Timeout: time.Second * time.Duration(5),
	}
	_, err := dialer.Dial("tcp", u)
	if err != nil {
		return 0, err
	}
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

	var hosts []Host
	for _, v := range meta.Git {
		if strings.Contains(v, ":") {
			continue
		}
		idx := strings.Index(v, "/")
		last := v[idx:]
		if last != "/32" {
			continue
		}
		address := v[:strings.Index(v, "/")]
		t, err := Ping(address)
		if err != nil {
			continue
		}
		if len(hosts) == 0 {
			hosts = []Host{{
				address: address,
				t:       t,
			}}
			continue
		}
		ok := false
		for kk, vv := range hosts {
			if t < vv.t {
				ok = true
				hosts = append(hosts[:kk], append([]Host{{
					address: address,
					t:       t,
				}}, hosts[kk:]...)...)
				break
			}
		}
		if !ok {
			hosts = append(hosts, Host{
				address: address,
				t:       t,
			})
		}
	}
	for _, v := range hosts {
		fmt.Printf("%s github.com\n", v.address)
	}
}
