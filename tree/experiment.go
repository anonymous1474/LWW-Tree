package tree

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strconv"
	"time"

	pb "github.com/anonymous1474/lww-tree/protos"
)

// Make random updates
func (node *ReplicaNode) experiment(rounds, rate, siz int) {
	time.Sleep((5 - time.Duration(node.replicaID)) * time.Second)

	done := make(chan bool, 10)

	var peerStream []pb.StreamService_FetchResponseClient
	peerStream = make([]pb.StreamService_FetchResponseClient, 3)
	//vc := make([]int, 3)

	for i := 0; i < 3; i++ {
		if int32(i) == node.replicaID {
			continue
		}
		in := &pb.Request{Id: node.replicaID}
		var err error
		peerStream[i], err = node.peerReplica[i].FetchResponse(context.Background(), in)
		CheckFatalError(err)

		id := i
		go func(id int) {
			for {
				in, err := peerStream[id].Recv()
				if err == io.EOF {
					done <- true //means stream is finished
					return
				}
				if err != nil {
					log.Fatalf("cannot receive %v", err)
				}

				totalops2++
				crdt.Lock()
				start := time.Now()
				//crdt.Lock()
				val := node.store.lww[in.Key]
				prevID := node.store.tieID[in.Key]
				//vc[in.ID] += 1
				//fmt.Println(Reset, vc)

				// Check if timestamp is lower
				if (val > in.Clock) || (val == in.Clock && prevID > in.ID) {
					elapsed := time.Since(start)
					measure2 += elapsed
					crdt.Unlock()
					continue
				}

				// Check if it will form cycle
				//fmt.Println(Purple + "Cycle check")
				cycle := node.isCycle(in.Key, in.Value)

				if cycle {
					// find last moved node in cycle
					//fmt.Println(Cyan + "Cycle Found")
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

					// Find Previous Parent Location
					keyMove, _ := strconv.Atoi(toMove)
					var nextLocation string

					// Till we get safe location
					unsafe := true
					for unsafe == true {
						prevLen := len(node.parentLog[keyMove])
						if prevLen == 0 {
							nextLocation = "CONFLICT"
							unsafe = false
						} else {
							nextLocation = node.parentLog[keyMove][prevLen-1]
							node.parentLog[keyMove] = node.parentLog[keyMove][0:(prevLen - 1)]
							// not in subtree of parent
							unsafe = node.isCycle(in.Key, nextLocation)
						}
					}

					// Do Local Update, make sure to send with this replicaID
					node.store.db[toMove] = nextLocation
					node.store.lww[toMove] = sendClock
					node.store.tieID[toMove] = node.replicaID

					msg := move{toMove, nextLocation, sendClock, node.replicaID}
					// send to other peers
					ch1 <- msg
					ch2 <- msg
					ch3 <- msg

					if toMove == in.Key { // Since a later operation has been applied on it
						elapsed := time.Since(start)
						measure2 += elapsed
						crdt.Unlock()
						continue
					}
				}

				// Change on all 3 maps
				prevParent := node.store.db[in.Key]
				node.store.db[in.Key] = in.Value
				node.store.lww[in.Key] = in.Clock
				node.store.tieID[in.Key] = in.ID

				// Add previous parent
				genKey, _ := strconv.Atoi(in.Key)
				node.parentLog[genKey] = append(node.parentLog[genKey], prevParent)
				if len(node.parentLog[genKey]) == 5 {
					node.parentLog[genKey] = node.parentLog[genKey][1:5]
				}

				// Update Lamport Timestamp
				if in.Clock > node.lamport {
					node.lamport = in.Clock
				}

				elapsed := time.Since(start)
				measure2 += elapsed
				crdt.Unlock()
			}
		}(id)
	}

	//Run experiment mentioned number of times

	for i := 0; i < rounds; i++ {

		// Generate random operations & Sleep for interval time
		time.Sleep(time.Duration(rate) * time.Microsecond)
		genKey := (rand.Intn(10000) % siz) + 1
		genVal := (rand.Intn(10000) % siz) + 1
		sendKey := strconv.Itoa(genKey)
		sendVal := strconv.Itoa(genVal)

		crdt.Lock()
		start := time.Now()
		//crdt.Lock()
		node.lamport += 1
		sendClock := node.lamport
		//vc[node.replicaID] += 1
		//fmt.Println(Reset, vc)
		totalops++

		if node.isCycle(sendKey, sendVal) {
			elapsed := time.Since(start)
			measure1 += elapsed
			crdt.Unlock()
			continue
		}

		// Do Local Update
		prevParent := node.store.db[sendKey]
		node.store.db[sendKey] = sendVal
		node.store.lww[sendKey] = sendClock
		node.store.tieID[sendKey] = node.replicaID

		// Add previous parent
		node.parentLog[genKey] = append(node.parentLog[genKey], prevParent)
		if len(node.parentLog[genKey]) == 5 {
			node.parentLog[genKey] = node.parentLog[genKey][1:5]
		}

		elapsed := time.Since(start)
		measure1 += elapsed
		crdt.Unlock()

		ch1 <- move{sendKey, sendVal, sendClock, node.replicaID}
		ch2 <- move{sendKey, sendVal, sendClock, node.replicaID}
		ch3 <- move{sendKey, sendVal, sendClock, node.replicaID}
	}
	time.Sleep(5 * time.Second)

	ch1 <- move{"zero", "zero", -1, node.replicaID}
	ch2 <- move{"zero", "zero", -1, node.replicaID}
	ch3 <- move{"zero", "zero", -1, node.replicaID}

	<-done
	<-done
	fmt.Printf(Green+"Local time = %s, Remote Time = %s, totalops = %v, Conflicts = %v\n", measure1/time.Duration(totalops), measure2/time.Duration(totalops2), totalops, resolves)
	vertices := make([]string, 0, siz)

	for i := 1; i <= siz; i++ {
		key := strconv.Itoa(i)
		vertices = append(vertices, node.store.db[key])
	}

	terminate := &pb.Verify{
		Vertices: vertices,
		ID:       int32(siz),
	}

	node.SendOps(terminate)
}

func (node *ReplicaNode) SendOps(msg *pb.Verify) {
	for i := 0; i < node.peers; i++ {
		if int32(i) == node.replicaID {
			continue
		}
		clientObj := node.final[i]

		go func(node *ReplicaNode, clientObj pb.ChatServiceClient, msg *pb.Verify) {

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
