package peer

import (
	"encoding/binary"
	"fmt"
	"net"
)

type Peer struct {
	IP   net.IP
	Port uint16
}

func SeparatePeers(peersResp []byte) ([]Peer, error) {
	const peerSize = 6 // 4 for IP, 2 for port
	numPeers := len(peersResp) / peerSize
	if len(peersResp)%peerSize != 0 {
		err := fmt.Errorf("peers len err")
		return nil, err
	}
	peers := make([]Peer, numPeers)
	for i := 0; i < numPeers; i++ {
		offset := i * peerSize
		peers[i].IP = net.IP(peersResp[offset : offset+4])
		peers[i].Port = binary.BigEndian.Uint16(peersResp[offset+4 : offset+6])
	}
	return peers, nil
}
