package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tech-yush/bittorent-client/bencode"
	"github.com/tech-yush/bittorent-client/peers"
)

func main() {
	path := os.Args[1]
	torrentfile, err := bencode.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	resp, _ := bencode.ParseResp(torrentfile)
	peers, _ := peers.SeparatePeers([]byte(resp.Peers))
	fmt.Println(peers)
}
