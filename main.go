package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tech-yush/bittorent-client/bencode"
)

func main() {
	path := os.Args[1]
	torrentfile, err := bencode.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(torrentfile.Announce)
}
