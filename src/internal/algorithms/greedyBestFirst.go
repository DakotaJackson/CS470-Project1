package algorithms

import (
	"errors"

	"github.com/DakotaJackson/CS470-Project1/src/internal/dataStructs"
	"github.com/DakotaJackson/CS470-Project1/src/internal/structs"
)

// Following is the algorithm used for a greedy best first search algorithm
// 	with a manhattan distance heuristic.
// 	Some help used here: http://theory.stanford.edu/~amitp/GameProgramming/AStarComparison.html

// GBFHelper contains all the info needed to conduct the a star search.
type GBFHelper struct {
	graph     dataStructs.Graph
	output    structs.OutputSpec
	startPos  int
	startN    *GBFNode
	targetPos int
	targetN   *GBFNode
	vVisited  []bool
	mapGrid   [][]int
	vPath     []int
}

// GBFNode is a helper for the algorithm.
type GBFNode struct {
	vertNum int
	vertX   int
	vertY   int
	parent  *GBFNode
	estDist int
	pCost   int
}

// InitGBF creates all the necessary info from the graph&mapInfo.
func InitGBF(mapInfo structs.MapSpec, graph dataStructs.Graph) *GBFHelper {
	helper := new(GBFHelper)
	helper.graph = graph
	helper.startPos = mapInfo.StartVert
	helper.targetPos = mapInfo.GoalVert
	helper.mapGrid = mapInfo.OrigMap
	helper.vPath = make([]int, 0)
	helper.vVisited = make([]bool, graph.GetNumVerticies())

	helper.startN = &GBFNode{
		vertNum: mapInfo.StartVert,
		vertX:   mapInfo.StartPosX,
		vertY:   mapInfo.StartPosY,
		parent:  nil,
		estDist: 0,
		pCost:   0,
	}

	helper.targetN = &GBFNode{
		vertNum: mapInfo.GoalVert,
		vertX:   mapInfo.GoalPosX,
		vertY:   mapInfo.GoalPosY,
		estDist: 0,
		pCost:   0,
	}

	return helper
}

// FindPathGBF returns the output struct with a path found from the gbf search alg.
func (gbf *GBFHelper) FindPathGBF() (structs.OutputSpec, error) {
	output := structs.OutputSpec{}
	var finalPath, openS, closeS []*GBFNode

	openS = append(openS, gbf.startN)
	gbf.vPath = append(gbf.vPath, gbf.startN.vertNum)

	for len(openS) != 0 {
		currN := gbf.getLowestManhattan(openS)
		gbf.vPath = append(gbf.vPath, currN.vertNum)
		if currN.parent != nil {
			vCst, err := gbf.graph.GetWeight(currN.vertNum)
			if err != nil {
				return output, errors.New("gbf error getting cost")
			}
			currN.pCost = currN.parent.pCost + vCst.(int)
		}

		if currN.vertNum == gbf.targetN.vertNum {
			// made it to end node, get final path and break out of loop
			finalPath = gbf.getFinalPath(currN)
			break
		}

		if len(finalPath) > 1 {
			break
		}

		openS = gbf.removeN(openS, currN)
		closeS = append(closeS, currN)

		// get all adjacent verticies
		edges := gbf.graph.GetEdgesForVerticy(currN.vertNum)

		for val := range edges.Data {
			if !gbf.hasNode(closeS, val.(int)) {
				gbf.vVisited[val.(int)] = true

				// create/define adjacent node
				nextN := gbf.getNodeFromVertNum(val.(int))
				estCst := gbf.getManhattan(nextN)
				if estCst == -1 {
					return output, errors.New("gbf error calculating manhattan")
				}
				nextN.estDist = currN.estDist + estCst

				// append it to the open nodes (to reference later)
				if !gbf.hasNode(openS, nextN.vertNum) {
					openS = append(openS, nextN)
				}

				nextN.parent = currN
				if val.(int) == gbf.targetN.vertNum {
					finalPath = gbf.getFinalPath(nextN)
					break
				}
			}
		}
	}

	if finalPath == nil {
		return output, errors.New("gbf error final path doesn't exist")
	}

	output.AlgType = "Greedy Best First Search"
	output.Ppath = gbf.getVertPathFromNodes(finalPath)
	//output.Ppath = gbf.vPath
	output.Pmoves = len(output.Ppath)
	output.Pvisited = gbf.getVisitedVerticies()
	output.Pcost = gbf.getPathCost(output.Ppath)

	if output.Pcost == -1 {
		return output, errors.New("bfs can't find cost of path")
	}

	return output, nil
}

// getManhattan returns the heuristic for a specific verticy.
func (gbf *GBFHelper) getManhattan(vert *GBFNode) int {
	// will add vertical and horizontal path to the target position
	// 	regardless of water, but water will add 5 to path instead of 0
	vNum := 0
	estCost := 0
	for rowNum, row := range gbf.mapGrid {
		for colNum := range row {
			if rowNum == vert.vertX && vert.vertNum < vNum {
				tmp, err := gbf.graph.GetWeight(vNum)
				if err != nil {
					return -1
				}
				if tmp.(int) == 0 {
					estCost += -3
				}
				estCost += tmp.(int)
			}
			if colNum == vert.vertY && vert.vertNum < vNum {
				tmp, err := gbf.graph.GetWeight(vNum)
				if err != nil {
					return -1
				}
				if tmp.(int) == 0 {
					estCost += -3
				}
				estCost += tmp.(int)
			}
			vNum++
		}
	}
	return estCost
}

// returns the verticy with the lowest calculated manhattan distance.
func (gbf *GBFHelper) getLowestManhattan(verts []*GBFNode) *GBFNode {
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

func (gbf *GBFHelper) getNodeFromVertNum(vert int) *GBFNode {
	vNum := 0
	tmp := new(GBFNode)
	for rowNum, row := range gbf.mapGrid {
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

func (gbf *GBFHelper) getFinalPath(endVert *GBFNode) []*GBFNode {
	var fPath []*GBFNode
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

func (gbf *GBFHelper) hasNode(verts []*GBFNode, sVert int) bool {
	for _, v := range verts {
		if v.vertNum == sVert {
			return true
		}
	}
	return false
}

func (gbf *GBFHelper) removeN(verts []*GBFNode, vert *GBFNode) []*GBFNode {
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

func (gbf *GBFHelper) getVertPathFromNodes(finalPath []*GBFNode) []int {
	finalVertPath := make([]int, 0)
	for _, v := range finalPath {
		finalVertPath = append(finalVertPath, v.vertNum)
	}
	return finalVertPath
}

func (gbf *GBFHelper) getPathCost(path []int) int {
	totalCost := 0
	for _, vert := range path {
		singleCost, err := gbf.graph.GetWeight(vert)
		if err != nil {
			return -1
		}
		totalCost += singleCost.(int)
	}
	return totalCost
}

func (gbf *GBFHelper) getVisitedVerticies() []int {
	visited := make([]int, 0)
	for i := 0; i < gbf.graph.GetNumVerticies(); i++ {
		if gbf.vVisited[i] {
			visited = append(visited, i)
		}
	}
	return visited
}
