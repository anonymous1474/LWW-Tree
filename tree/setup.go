package tree

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/anonymous1474/lww-tree/protos"
	"google.golang.org/grpc"
)

type Store struct {
	db    map[string]string
	lww   map[string]int32
	tieID map[string]int32
}

type ReplicaNode struct {
	replicaID   int32
	peerReplica []protos.ChatServiceClient
	//parentLog   [][]string
	address  string
	myServer *grpc.Server
	peers    int
	lamport  int32
	store    *Store
}

var crdt sync.Mutex
var measure1, measure2 time.Duration
var resolves, totalops int

func SetupReplica(num, id, rounds, rate, siz int) {
	grpc_address := ":500" + strconv.Itoa(id)
	lis, err := net.Listen("tcp", grpc_address)
	CheckFatalError(err)
	measure1 = 0
	measure2 = 0
	resolves = 0
	totalops = 0

	grpcServer := grpc.NewServer()
	node := &ReplicaNode{
		replicaID:   int32(id),
		peerReplica: make([]protos.ChatServiceClient, num),
		//parentLog:   make([][]string, 11), // number of dummy keys
		address:  grpc_address,
		myServer: grpcServer,
		peers:    int(num),
		lamport:  0,
		store: &Store{
			db:    make(map[string]string),
			lww:   make(map[string]int32),
			tieID: make(map[string]int32),
		},
	}
	protos.RegisterChatServiceServer(grpcServer, node)

	rep_addrs := make([]string, num)
	// to run on Cloud VMs"
	/*Uncomment and update from here
	rep_addrs[0] = "10.1.0.5:5000"
	rep_addrs[1] = "10.2.0.5:5001"
	rep_addrs[2] = "10.0.0.7:5002"
	*/ // Till here

	// Comment out to stop local run from here
	for i := 0; i < num; i++ {
		rep_addrs[i] = ":500" + strconv.Itoa(i)
	}
	// till here
	node.connectRest(rep_addrs)
	go node.experiment(rounds, rate, siz)

	node.store.db["ROOT"] = "ROOT"
	node.store.lww["ROOT"] = 0
	node.store.tieID["ROOT"] = 0
	node.store.db["TRASH"] = "ROOT"
	node.store.lww["TRASH"] = 0
	node.store.tieID["TRASH"] = 0
	node.store.db["CONFLICT"] = "ROOT"
	node.store.lww["CONFLICT"] = 0
	node.store.tieID["CONFLICT"] = 0

	for i := 1; i <= siz; i++ {
		//node.parentLog[i] = make([]string, 0)
		key := strconv.Itoa(i)

		node.store.db[key] = "ROOT"
		node.store.lww[key] = 0
		node.store.tieID[key] = 0
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// Connect to other Peers
func (node *ReplicaNode) connectRest(rep_addrs []string) {
	client_obj := make([]protos.ChatServiceClient, node.peers)

	for i := 0; i < node.peers; i++ {
		if int32(i) == node.replicaID {
			continue
		}

		connxn, err := grpc.Dial(rep_addrs[i], grpc.WithInsecure())
		CheckFatalError(err)
		fmt.Println("Connected to replica ", i)
		cli := protos.NewChatServiceClient(connxn)
		client_obj[i] = cli
	}
	node.peerReplica = client_obj
}
