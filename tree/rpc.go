package tree

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/anonymous1474/lww-tree/protos"
)

// RPC Call
func (node *ReplicaNode) Propogate(ctx context.Context, in *protos.UpdateValue) (*protos.Void, error) {
	crdt.Lock()
	start := time.Now()
	totalops++
	//fmt.Printf(Yellow+"Replica %v : %s -> %s --- Time = %v\n", in.ID, in.Key, in.Value, in.Clock)

	val := node.store.lww[in.Key]
	prevID := node.store.tieID[in.Key]

	// Check if timestamp is lower
	if (val > in.Clock) || (val == in.Clock && prevID > in.ID) {
		elapsed := time.Since(start)
		measure2 += elapsed
		crdt.Unlock()
		return &protos.Void{}, nil
	}

	// Check if it will form cycle
	cycle := node.isCycle(in.Key, in.Value)

	if cycle {
		// find last moved node in cycle
		resolves += 1
		tempLast := node.findLast(in.Key, in.Value)
		var toMove string

		if (node.store.lww[tempLast] > in.Clock) || (node.store.lww[tempLast] == in.Clock && node.store.tieID[tempLast] > in.ID) {
			toMove = tempLast
		} else {
			toMove = in.Key
		}

		// Update Lamport Timestamp
		if in.Clock > node.lamport {
			node.lamport = in.Clock
		}
		node.lamport += 1
		sendClock := node.lamport

		// Do Local Update, make sure to send with this replicaID
		node.store.db[toMove] = "CONFLICT"
		node.store.lww[toMove] = sendClock
		node.store.tieID[toMove] = node.replicaID

		msg := &protos.UpdateValue{
			Key:   toMove,
			Value: "CONFLICT",
			Clock: sendClock,
			ID:    node.replicaID,
		}
		// send to other peers
		go node.SendOps(msg)

		if toMove == in.Key { // Since a later operation has been applied on it
			elapsed := time.Since(start)
			measure2 += elapsed
			crdt.Unlock()
			//fmt.Println(Reset + "Released Lock")
			return &protos.Void{}, nil
		}
	}

	// Change on all 3 maps
	node.store.db[in.Key] = in.Value
	node.store.lww[in.Key] = in.Clock
	node.store.tieID[in.Key] = in.ID

	// Update Lamport Timestamp
	if in.Clock > node.lamport {
		node.lamport = in.Clock
	}

	// Release Lock
	elapsed := time.Since(start)
	measure2 += elapsed
	crdt.Unlock()
	return &protos.Void{}, nil
}

func (node *ReplicaNode) CheckAnswer(ctx context.Context, in *protos.Verify) (*protos.Void, error) {
	crdt.Lock()
	count := 0
	for i := 1; i <= int(in.ID); i++ {
		key := strconv.Itoa(i)
		if node.store.db[key] != in.Vertices[i-1] {
			count++
		}
	}
	fmt.Printf(Red+"Matching with replica %v\n", count)
	crdt.Unlock()
	return &protos.Void{}, nil
}
