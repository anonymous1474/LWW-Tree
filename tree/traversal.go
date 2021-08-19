package tree

func (node *ReplicaNode) isCycle(a, b string) bool {
	for b != "ROOT" {
		if b == a {
			return true
		}
		b = node.store.db[b]
	}
	return false
}

func (node *ReplicaNode) findLast(a, b string) string {
	var latest, tiebreaker int32
	var answer string
	latest = 0
	tiebreaker = 0

	for b != a {
		if (node.store.lww[b] > latest) || (node.store.lww[b] == latest && node.store.tieID[b] > tiebreaker) {
			latest = node.store.lww[b]
			answer = b
		}
		b = node.store.db[b]
	}

	return answer
}
