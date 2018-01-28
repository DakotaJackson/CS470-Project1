package internal

// Following is the algorithm used for a breadth first shirt.
// 	it is called from main and will ultimately return the data needed
// 	for the outputSpec struct

// BFSHelper contains all the information needed to conduct the bfs.
type BFSHelper struct {
	mapGraph *graph.Graph
	queue    *queue.Queue

	vVisited  []bool
	startPos  int
	targetPos int
}

func initBFS() {
	helper := new(BGSHelper)
}
