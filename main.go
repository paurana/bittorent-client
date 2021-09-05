package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tech-yush/bittorent-client/bencode"
)

func main() {
	path := os.Args[1]
	torrentfile, err := bencode.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	var peerID [20]byte
	_, _ = rand.Read(peerID[:])
	fmt.Println(peerID)
	url, _ := torrentfile.BuildTrackerURL(peerID, 6969)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
