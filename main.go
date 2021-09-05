package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"

	"github.com/tech-yush/bittorent-client/bencode"
	"github.com/tech-yush/bittorent-client/peers"
	"github.com/tech-yush/bittorent-client/tcp"
)

func main() {
	path := os.Args[1]
	torrentfile, err := bencode.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	resp, _ := bencode.ParseResp(torrentfile)
	peers, _ := peers.SeparatePeers([]byte(resp.Peers))
	// fmt.Println(peers)
	var peerID [20]byte
	_, _ = rand.Read(peerID[:])
	conn, err := tcp.New(peers[0], peerID, torrentfile.InfoHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn)
}
