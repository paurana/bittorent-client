package bencode

import (
	"os"

	"github.com/jackpal/bencode-go"
)

//I will be using github.com/jackpal/bencode-go to parse the .torrent file

type bencodeInfo struct {
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"` //binary blob of hashes of each piece
}

type parsedBencode struct {
	Announce     string      `bencode:"announce"`
	Info         bencodeInfo `bencode:"info"`
	Comment      string      `bencode:"comment"`
	CreationDate int         `bencode:"creation date"`
}

func Open(path string) (*parsedBencode, error) {
	file, err := os.Open(path)
	if err != nil {
		return &parsedBencode{}, err
	}
	defer file.Close()

	torrent := parsedBencode{}
	err = bencode.Unmarshal(file, &torrent)
	if err != nil {
		return &parsedBencode{}, err
	}
	return &torrent, nil
}
