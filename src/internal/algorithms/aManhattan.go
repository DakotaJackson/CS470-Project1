package algorithms

import (
	"errors"

	"github.com/DakotaJackson/CS470-Project1/src/internal/dataStructs"
	"github.com/DakotaJackson/CS470-Project1/src/internal/structs"
)

// Following is the algorithm used for an a star search algorithm
// 	with a manhattan distance heuristic.
// 	Some help used here: http://theory.stanford.edu/~amitp/GameProgramming/Heuristics.html

// AMSHelper contains all the info needed to conduct the a star search.
type AMSHelper struct {
	graph     dataStructs.Graph
	output    structs.OutputSpec
	startPos  int
	startN    *AMSNode
	targetPos int
	targetN   *AMSNode
	vVisited  []bool
	mapGrid   [][]int
}

// AMSNode is a helper for the algorithm.
type AMSNode struct {
	vertNum int
	vertX   int
	vertY   int
	parent  *AMSNode
	estDist int
	pCost   int
}

// InitAMS creates all the necessary info from the graph&mapInfo.
func InitAMS(mapInfo structs.MapSpec, graph dataStructs.Graph) *AMSHelper {
	helper := new(AMSHelper)
	helper.graph = graph
	helper.startPos = mapInfo.StartVert
	helper.targetPos = mapInfo.GoalVert
	helper.mapGrid = mapInfo.OrigMap
	helper.vVisited = make([]bool, graph.GetNumVerticies())

	helper.startN = &AMSNode{
		vertNum: mapInfo.StartVert,
		vertX:   mapInfo.StartPosX,
		vertY:   mapInfo.StartPosY,
		parent:  nil,
		estDist: 0,
		pCost:   0,
	}

	helper.targetN = &AMSNode{
		vertNum: mapInfo.GoalVert,
		vertX:   mapInfo.GoalPosX,
		vertY:   mapInfo.GoalPosY,
		estDist: 0,
		pCost:   0,
	}

	return helper
}

// FindPathAMS returns the output struct with a path found from the ams search alg.
func (ams *AMSHelper) FindPathAMS() (structs.OutputSpec, error) {
	output := structs.OutputSpec{}
	var finalPath, openS, closeS []*AMSNode

	openS = append(openS, ams.startN)

	for len(openS) != 0 {
		currN := ams.getLowestManhattan(openS)

		if currN.parent != nil {
			vCst, err := ams.graph.GetWeight(currN.vertNum)
			if err != nil {
				return output, errors.New("ams error getting cost")
			}
			currN.pCost = currN.parent.pCost + vCst.(int)
		}

		if currN.vertNum == ams.targetN.vertNum {
			// made it to end node, get final path and break out of loop
			finalPath = ams.getFinalPath(currN)
			break
		}

		if len(finalPath) > 1 {
			break
		}

		openS = ams.removeN(openS, currN)
		closeS = append(closeS, currN)

		// get all adjacent verticies
		edges := ams.graph.GetEdgesForVerticy(currN.vertNum)

		for val := range edges.Data {
			if !ams.hasNode(closeS, val.(int)) {
				ams.vVisited[val.(int)] = true

				// create/define adjacent node
				nextN := ams.getNodeFromVertNum(val.(int))
				estCst := ams.getManhattan(nextN)
				if estCst == -1 {
					return output, errors.New("ams error calculating manhattan")
				}
				nextN.estDist = currN.estDist + estCst

				// append it to the open nodes (to reference later)
				if !ams.hasNode(openS, nextN.vertNum) {
					openS = append(openS, nextN)
				}

				nextN.parent = currN
				if val.(int) == ams.targetN.vertNum {
					finalPath = ams.getFinalPath(nextN)
					break
				}
			}
		}
	}

	if finalPath == nil {
		return output, errors.New("ams error final path doesn't exist")
	}

	output.AlgType = "A* Search With Manhattan Heuristic"
	output.Ppath = ams.getVertPathFromNodes(finalPath)
	output.Pmoves = len(output.Ppath)
	output.Pvisited = ams.getVisitedVerticies()
	output.Pcost = ams.getPathCost(output.Ppath)

	if output.Pcost == -1 {
		return output, errors.New("bfs can't find cost of path")
	}

	return output, nil
}

// getManhattan returns the heuristic for a specific verticy.
func (ams *AMSHelper) getManhattan(vert *AMSNode) int {
	// will add vertical and horizontal path to the target position
	// 	regardless of water, but water will add 5 to path instead of 0
	vNum := 0
	estCost := 0
	for rowNum, row := range ams.mapGrid {
		for colNum := range row {
			if rowNum == vert.vertX && vert.vertNum < vNum {
				tmp, err := ams.graph.GetWeight(vNum)
				if err != nil {
					return -1
				}
				if tmp.(int) == 0 {
					estCost += 5
				}
				estCost += tmp.(int)
			}
			if colNum == vert.vertY && vert.vertNum < vNum {
				tmp, err := ams.graph.GetWeight(vNum)
				if err != nil {
					return -1
				}
				if tmp.(int) == 0 {
					estCost += 5
				}
				estCost += tmp.(int)
			}
			vNum++
		}
	}
	return estCost
}

// returns the verticy with the lowest calculated manhattan distance.
func (ams *AMSHelper) getLowestManhattan(verts []*AMSNode) *AMSNode {
	if len(verts) == 0 {
		return nil
	}
	lowest := verts[0]
	bestManhattan := lowest.estDist

	for _, v := range verts {
		if v.estDist < bestManhattan {
			bestManhattan = v.estDist
			lowest = v
		}
	}

	return lowest
}

func (ams *AMSHelper) getNodeFromVertNum(vert int) *AMSNode {
	vNum := 0
	tmp := new(AMSNode)
	for rowNum, row := range ams.mapGrid {
		for colNum := range row {
			if vNum == vert {
				tmp.vertX = rowNum
				tmp.vertY = colNum
				tmp.vertNum = vert
			}
			vNum++
		}
	}

	return tmp
}

func (ams *AMSHelper) getFinalPath(endVert *AMSNode) []*AMSNode {
	var fPath []*AMSNode
	fPath = append(fPath, endVert)

	for endVert.parent != nil {
		fPath = append(fPath, endVert.parent)
		endVert = endVert.parent
	}

	for i, j := 0, len(fPath)-1; i < j; i, j = i+1, j-1 {
		fPath[i], fPath[j] = fPath[j], fPath[i]
	}
	return fPath
}

func (ams *AMSHelper) hasNode(verts []*AMSNode, sVert int) bool {
	for _, v := range verts {
		if v.vertNum == sVert {
			return true
		}
	}
	return false
}

func (ams *AMSHelper) removeN(verts []*AMSNode, vert *AMSNode) []*AMSNode {
	indx := -1

	for ind, v := range verts {
		if v == vert {
			indx = ind
			break
		}
	}

	if indx != -1 {
		copy(verts[indx:], verts[indx+1:])
		verts = verts[:len(verts)-1]
	}

	return verts
}

func (ams *AMSHelper) getVertPathFromNodes(finalPath []*AMSNode) []int {
	finalVertPath := make([]int, 0)
	for _, v := range finalPath {
		finalVertPath = append(finalVertPath, v.vertNum)
	}
	return finalVertPath
}

func (ams *AMSHelper) getPathCost(path []int) int {
	totalCost := 0
	for _, vert := range path {
		singleCost, err := ams.graph.GetWeight(vert)
		if err != nil {
			return -1
		}
		totalCost += singleCost.(int)
	}
	return totalCost
}

func (ams *AMSHelper) getVisitedVerticies() []int {
	visited := make([]int, 0)
	for i := 0; i < ams.graph.GetNumVerticies(); i++ {
		if ams.vVisited[i] {
			visited = append(visited, i)
		}
	}
	return visited
}
