package event

type EventQueue []Event

func (h EventQueue) Len() int { return len(h) }

func (h EventQueue) Less(i, j int) bool {
	return h[i].Date < h[j].Date
}

func (h EventQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *EventQueue) Push(x interface{}) {
	*h = append(*h, x.(Event))
}

func (h *EventQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
