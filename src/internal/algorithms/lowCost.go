package algorithms

import (
	"errors"
	"fmt"

	"github.com/DakotaJackson/CS470-Project1/src/internal/dataStructs"
	"github.com/DakotaJackson/CS470-Project1/src/internal/structs"
)

// Following is the algorithm used for a lowest cost search.
// 	The algorithm being used is dijkstra's algorithm.
// 	Basic functionality and help gathered from here:
// 	https://github.com/flemeur/go-shortestpath

// LCSHelper contains all the information needed to conduct the lcs.
type LCSHelper struct {
	queue     *dataStructs.Queue
	graph     dataStructs.Graph
	startPos  int
	targetPos int
	vVisited  []bool
	vWeights  map[int]int
	vPrev     map[int]int
	vPath     []int
}

// InitLCS creates all the necessary info from the graph&mapInfo.
func InitLCS(mapInfo structs.MapSpec, graph dataStructs.Graph) *LCSHelper {
	// init a new helper object for the algorithm
	helper := new(LCSHelper)

	// define and create the necessary initial values
	helper.queue = dataStructs.InitQueue()
	helper.graph = graph
	helper.vWeights = make(map[int]int)
	helper.vVisited = make([]bool, graph.GetNumVerticies())
	helper.vPrev = make(map[int]int)
	helper.vPath = make([]int, 0)
	helper.startPos = mapInfo.StartVert
	helper.targetPos = mapInfo.GoalVert

	helper.vWeights[helper.startPos] = 0
	helper.queue.Enqueue(helper.startPos)

	return helper
}

// FindPathLCS returns the path found from the lcs search algorithm.
func (lcs *LCSHelper) FindPathLCS() ([]int, error) {
	// as long as there is an item in the queue, execute the loop
	for lcs.queue.GetLenQ() > 0 {
		vert, exists := lcs.queue.Dequeue()
		fmt.Println("verticy being analyzed:", vert)
		if !exists {
			return nil, errors.New("lcs current vert dequeued doesn't exist")
		}

		if lcs.vVisited[lcs.targetPos] {
			// found the target verticy! Done with loop
			break
		}

		edges := lcs.graph.GetEdgesForVerticy(vert)
		eWeight, err := lcs.graph.GetWeight(vert)
		if err != nil {
			return nil, errors.New("lcs unable to get weight for verticy")
		}

		// iterate over and find least cost of each edge from a verticy
		for val := range edges.Data {
			dest := val.(int)
			cost, err := lcs.graph.GetWeight(val.(int))
			if err != nil {
				return nil, errors.New("lcs unable to get weight in edges loop")
			}
			// total cost of moving to tile
			cost = cost.(int) + eWeight.(int)
			if tentDist, ok := lcs.vWeights[dest]; !ok || cost.(int) < tentDist {
				// prevent two nodes mapping to each other
				if lcs.vPrev[vert] != dest {
					lcs.vWeights[dest] = cost.(int)
					lcs.vPrev[dest] = vert
					lcs.queue.Enqueue(dest)
				}
			}
		}

		// set the current verticy being analyzed as true
		lcs.vVisited[vert] = true
	}

	if !lcs.vVisited[lcs.targetPos] {
		return nil, errors.New("lcs unable to reach destination")
	}

	// gathers the lowest cost path based on the vPrev map
	lcs.vPath = append(lcs.vPath, lcs.targetPos)
	for n, ok := lcs.vPrev[lcs.targetPos]; ok; n = lcs.vPrev[n] {
		fmt.Println("CURRENT PATH VERT: ", n)
		lcs.vPath = append(lcs.vPath, n)
		if n == lcs.startPos {
			break
		}
	}

	for i, j := 0, len(lcs.vPath)-1; i < j; i, j = i+1, j-1 {
		lcs.vPath[i], lcs.vPath[j] = lcs.vPath[j], lcs.vPath[i]
	}

	return lcs.vPath, nil
}
