package queue

var q []string

// GetQueue: get current queue
func GetQueue() []string {
	return q
}

// Push: push element in queue
func Push(str string) {
	q = append(q, str)
}

// Pop: pop element from queue
func Pop() string {
	lenQueue := len(q)

	if lenQueue == 0 {
		return "queue is empty"
	}

	popElem := q[0]
	q = q[1:]

	return popElem
}