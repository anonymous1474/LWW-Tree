package tree

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/anonymous1474/lww-tree/protos"
)

// Make random updates
func (node *ReplicaNode) experiment(rounds, rate, siz int) {
	time.Sleep((25 - time.Duration(node.replicaID)) * time.Second)
	//Run experiment mentioned number of times
	for i := 0; i < rounds; i++ {
		// Generate Operation Key and Value
		genKey := (rand.Intn(10000) % siz) + 1
		genVal := (rand.Intn(10000) % siz) + 1
		sendKey := strconv.Itoa(genKey)
		sendVal := strconv.Itoa(genVal)

		// Sleep for random time
		time.Sleep(time.Duration(rate) * time.Microsecond)

		crdt.Lock()
		start := time.Now()
		totalops++

		// check if key value pair is valid
		localCycle := node.isCycle(sendKey, sendVal)
		if genKey == genVal || localCycle == true {
			//fmt.Printf(Red + "Invalid key value pair\n")
			elapsed := time.Since(start)
			measure1 += elapsed
			crdt.Unlock()
			continue
		}

		node.lamport += 1
		sendClock := node.lamport
		sendID := node.replicaID
		//fmt.Printf(Yellow+"Replica %v : %s -> %s --- Time = %v\n", sendID, sendKey, sendVal, sendClock)
		// Do Local Update
		node.store.db[sendKey] = sendVal
		node.store.lww[sendKey] = sendClock
		node.store.tieID[sendKey] = sendID

		elapsed := time.Since(start)
		measure1 += elapsed
		crdt.Unlock()

		msg := &protos.UpdateValue{
			Key:   sendKey,
			Value: sendVal,
			Clock: sendClock,
			ID:    sendID,
		}

		go node.SendOps(msg)
	}

	time.Sleep(5 * time.Second)
	fmt.Printf(Cyan+"Finished Replica %v\n", node.replicaID)
	var done int
	fmt.Scanf("%v", &done)

	crdt.Lock()

	vertices := make([]string, 0, siz)

	//vertices = append(vertices, "ROOT")
	//vertices = append(vertices, "TRASH")
	//vertices = append(vertices, "CONFLICT")

	for i := 1; i <= siz; i++ {
		key := strconv.Itoa(i)
		vertices = append(vertices, node.store.db[key])
	}

	terminate := &protos.Verify{
		Vertices: vertices,
		ID:       int32(siz),
	}

	node.SendOps2(terminate)

	crdt.Unlock()
	fmt.Printf(Green+"Local time = %s, Remote Time = %s, totalops = %v, Conflicts = %v\n", measure1, measure2, totalops, resolves)

}

func (node *ReplicaNode) SendOps(msg *protos.UpdateValue) {
	for i := 0; i < node.peers; i++ {
		if int32(i) == node.replicaID {
			continue
		}
		clientObj := node.peerReplica[i]

		go func(node *ReplicaNode, clientObj protos.ChatServiceClient, msg *protos.UpdateValue) {

			ctx, cancel := context.WithCancel(context.Background())
			var ack error
			_, ack = clientObj.Propogate(ctx, msg)
			cancel()

			for ack != nil {
				_, ack = clientObj.Propogate(ctx, msg)
			}
		}(node, clientObj, msg)
	}
}

func (node *ReplicaNode) SendOps2(msg *protos.Verify) {
	for i := 0; i < node.peers; i++ {
		if int32(i) == node.replicaID {
			continue
		}
		clientObj := node.peerReplica[i]

		go func(node *ReplicaNode, clientObj protos.ChatServiceClient, msg *protos.Verify) {

			ctx, cancel := context.WithCancel(context.Background())
			var ack error
			_, ack = clientObj.CheckAnswer(ctx, msg)
			cancel()

			for ack != nil {
				_, ack = clientObj.CheckAnswer(ctx, msg)
			}
		}(node, clientObj, msg)
	}
}
