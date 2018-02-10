package algorithms

import (
	"errors"

	"github.com/DakotaJackson/CS470-Project1/src/internal/dataStructs"
	"github.com/DakotaJackson/CS470-Project1/src/internal/structs"
)

// Following is the algorithm used for a breadth first search.
// 	it is called from main and will ultimately return the data needed
// 	for the outputSpec struct
// Some functionality and help was derived from here:
//  https://github.com/karalabe/cookiejar/blob/v2/graph/bfs/bfs.go

// BFSHelper contains all the information needed to conduct the bfs.
type BFSHelper struct {
	// needed data structures (graph passed in)
	queue  *dataStructs.Queue
	stack  *dataStructs.Stack
	graph  dataStructs.Graph
	output structs.OutputSpec
	// start and target verticies (need to be found)
	startPos  int
	targetPos int
	// v[variablename] are used in the bfs
	vVisited []bool
	vParents []int
	vPaths   map[int][]int
	success  bool
}

// InitBFS creates all the necessary info from the graph&mapInfo.
func InitBFS(mapInfo structs.MapSpec, graph dataStructs.Graph) *BFSHelper {
	// create a new instance
	helper := new(BFSHelper)

	// gather needed data structures
	helper.queue = dataStructs.InitQueue()
	helper.stack = dataStructs.InitStack()
	helper.graph = graph
	// define the start and target verticies
	helper.startPos = mapInfo.StartVert
	helper.targetPos = mapInfo.GoalVert

	// init the needed helper variables
	helper.vVisited = make([]bool, graph.GetNumVerticies())
	helper.vParents = make([]int, graph.GetNumVerticies())
	helper.vPaths = make(map[int][]int)
	helper.success = false

	// populate the helper variables with their initial values
	helper.vVisited[helper.startPos] = true
	helper.queue.Enqueue(helper.startPos)

	// return the newly made bfs helper object
	return helper
}

// FindPathBFS returns the output struct with path from start to target.
func (bfs *BFSHelper) FindPathBFS() (structs.OutputSpec, error) {
	output := structs.OutputSpec{}
	// can't reach target (eg. surrounded by water)
	if !bfs.canReachVerticy(bfs.targetPos) {
		return output, errors.New("bfs target verticy is unreachable")
	}

	if cached, ok := bfs.vPaths[bfs.targetPos]; !ok {
		for cur := bfs.targetPos; cur != bfs.startPos; {
			bfs.stack.Push(cur)
			cur = bfs.vParents[cur]
		}
		bfs.stack.Push(bfs.startPos)

		path := make([]int, bfs.stack.GenLenS())
		var exists bool
		for i := 0; i < len(path); i++ {
			path[i], exists = bfs.stack.Pop()
			if !exists {
				return output, errors.New("bfs error in popping from stack searching for path")
			}
		}
		bfs.vPaths[bfs.targetPos] = path
		output.Ppath = path
		bfs.success = true
	} else if cached != nil {
		output.Ppath = cached
		bfs.success = true
	}
	if !bfs.success {
		return output, errors.New("bfs can't find path")
	}
	output.Pvisited = bfs.getVisitedVerticies()
	output.AlgType = "Breadth First Search"
	output.Pmoves = len(output.Ppath)
	output.Pcost = bfs.getPathCost(output.Ppath)
	if output.Pcost == -1 {
		return output, errors.New("bfs can't find cost of path")
	}
	return output, nil
}

// canReachVerticy checks if the path from start to target is possible.
func (bfs *BFSHelper) canReachVerticy(target int) bool {
	if !bfs.vVisited[target] && !bfs.queue.IsEmpty() {
		bfs.search(target)
	}
	return bfs.vVisited[target]
}

// search executes the actual breadth first search.
func (bfs *BFSHelper) search(target int) {
	// execute loop while there is a value in the queue.
	for !bfs.queue.IsEmpty() {
		vert, exists := bfs.queue.Dequeue()
		if exists {
			bfs.graph.DoVerticy(vert, func(peer interface{}) {
				if p := peer.(int); !bfs.vVisited[p] {
					bfs.vVisited[p] = true
					bfs.vParents[p] = vert
					bfs.queue.Enqueue(p)
				}
			})
		}

		// if node is found, return
		if target == vert {
			return
		}
	}
}

// getVisitedVerticies returns all verticies visited for output.
func (bfs *BFSHelper) getVisitedVerticies() []int {
	visited := make([]int, 0)
	for i := 0; i < bfs.graph.GetNumVerticies(); i++ {
		if bfs.vVisited[i] {
			visited = append(visited, i)
		}
	}
	return visited
}

// getPathCost returns the total cost of movement for output.
func (bfs *BFSHelper) getPathCost(path []int) int {
	totalCost := 0
	for _, vert := range path {
		singleCost, err := bfs.graph.GetWeight(vert)
		if err != nil {
			return -1
		}
		totalCost += singleCost.(int)
	}
	return totalCost
}
