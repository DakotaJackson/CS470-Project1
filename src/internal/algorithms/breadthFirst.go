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
	queue *dataStructs.Queue
	stack *dataStructs.Stack
	graph dataStructs.Graph
	// start and target verticies (need to be found)
	startPos  int
	targetPos int
	// v[variablename] are used in the bfs
	vVisited []bool
	vOrder   []int
	vParents []int
	vPaths   map[int][]int
}

// InitBFS creates all the necessary info from the graph&mapInfo then
// 	executes the breadth first search
func InitBFS(mapInfo structs.MapSpec, graph dataStructs.Graph) *BFSHelper {
	// create a new instance
	helper := new(BFSHelper)

	// gather needed data structures
	helper.queue = dataStructs.InitQueue()
	helper.stack = dataStructs.InitStack()
	helper.graph = graph
	// define the start and target verticies
	helper.startPos, helper.targetPos = getGoals(mapInfo)

	// init the needed helper variables
	helper.vVisited = make([]bool, graph.GetNumVerticies())
	helper.vOrder = make([]int, 1, graph.GetNumVerticies())
	helper.vParents = make([]int, graph.GetNumVerticies())
	helper.vPaths = make(map[int][]int)

	// populate the helper variables with their initial values
	helper.vVisited[helper.startPos] = true
	helper.vOrder[0] = helper.startPos
	helper.queue.Enqueue(helper.startPos)

	// return the newly made bfs helper object
	return helper
}

// FindPathBFS returns an integer slice of the verticies traveled from start to target.
func (bfs *BFSHelper) FindPathBFS() ([]int, error) {
	// can't reach target (eg. surrounded by water)
	if !bfs.canReachVerticy(bfs.targetPos) {
		return nil, errors.New("bfs target verticy is unreachable")
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
				return nil, errors.New("bfs error in popping from stack searching for path")
			}
		}
		bfs.vPaths[bfs.targetPos] = path
		return path, nil
	} else if cached != nil {
		return cached, nil
	}
	return nil, errors.New("bfs can't find path")
}

// Order returns the order of verticies from start to target position.
func (bfs *BFSHelper) Order() []int {
	if !bfs.queue.IsEmpty() {
		bfs.search(-1)
	}
	return bfs.vOrder
}

// canReachVerticy checks if the path from start to target is possible.
func (bfs *BFSHelper) canReachVerticy(target int) bool {
	if !bfs.vVisited[target] && !bfs.queue.IsEmpty() {
		bfs.search(target)
	}
	return bfs.vVisited[target]
}

func (bfs *BFSHelper) search(target int) {
	for !bfs.queue.IsEmpty() {
		src, exists := bfs.queue.Dequeue()
		if exists {
			bfs.graph.DoVerticy(src, func(peer interface{}) {
				if p := peer.(int); !bfs.vVisited[p] {
					bfs.vVisited[p] = true
					bfs.vOrder = append(bfs.vOrder, p)
					bfs.vParents[p] = src
					bfs.queue.Enqueue(p)
				}
			})
		}

		if target == src {
			return
		}
	}
}

// getGoals returns the number of the verticies for both the start and goal
// 	positions based on the x/y coordinates in mapSpec.go.
func getGoals(mapInfo structs.MapSpec) (int, int) {
	// TODO: add this functionality
	vertNum := 0
	startVertNum := 0
	targetVertNum := 0
	for i := 0; i < mapInfo.Height; i++ {
		for j := 0; j < mapInfo.Width; j++ {
			if i == mapInfo.StartPosY && j == mapInfo.StartPosX {
				startVertNum = vertNum
			}
			if i == mapInfo.GoalPosY && j == mapInfo.GoalPosX {
				targetVertNum = vertNum
			}
			vertNum++
		}
	}
	return startVertNum, targetVertNum
}
