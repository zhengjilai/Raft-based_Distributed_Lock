package node

import (
	"errors"
)

var ConfigEmptyError = errors.New("dlock_raft.init_node: init peer node list fails for " +
	"node config is nil")
var ConfigPeerListError = errors.New("dlock_raft.init_node: init peer node list fails for " +
	"peer node list is empty, or id and address mismatch")

type PeerNode struct {

	// peer id
	PeerId uint32
	// the peer address name, format ip:port
	AddressName string
	// the state of the peer
	PeerState int
	// nextIndex is used for leader to find the next index to append for follower
	// first set equal to the last commit index, then decrement continuously to find the last common entry
	NextIndex uint64
	// matchIndex indicates the last match index in the peer's local LogMemory
	// often used for leader's committing process, judging whether majority of followers have logged an Entry
	MatchIndex uint64

	// the gprc client instance, for network transport
	GrpcClient *GrpcClientImpl
	// reference to node object
	NodeRef *Node

}

func NewPeerNode(peerId uint32, addressName string, peerState int, nextIndex uint64, matchIndex uint64,
	grpcClient *GrpcClientImpl, node *Node) *PeerNode {
	return &PeerNode{
		PeerId: peerId,
		AddressName: addressName,
		PeerState: peerState,
		NextIndex: nextIndex,
		MatchIndex: matchIndex,
		GrpcClient: grpcClient,
		NodeRef: node,
	}
}

func NewPeerNodeListFromConfig(node *Node) ([]*PeerNode, error) {

	// config errors test, e.g. empty peer list
	if node.NodeConfigInstance == nil {
		return nil, ConfigEmptyError
	}
	idList := node.NodeConfigInstance.Id.PeerId
	addressList := node.NodeConfigInstance.Network.PeerAddress
	// error happens if
	if len(addressList) == 0 || len(idList) == 0 || len(addressList) != len(idList){
		return nil, ConfigPeerListError
	}

	// make the feedback object
	peerList := make([]*PeerNode, len(addressList))
	for i := 0 ; i < len(addressList) ; i++ {
		grpcClient := NewGrpcClient(addressList[i], node)
		peerList[i] = NewPeerNode(idList[i], addressList[i], Unknown, 0, 0, grpcClient, node)
	}
	return peerList, nil
}

