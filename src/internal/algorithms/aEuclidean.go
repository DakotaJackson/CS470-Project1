package algorithms

import (
	"errors"

	"github.com/DakotaJackson/CS470-Project1/src/internal/dataStructs"
	"github.com/DakotaJackson/CS470-Project1/src/internal/structs"
)

// Following is the algorithm used for an a star search algorithm
// 	with a Euclidean distance heuristic.
// 	Some help used here: http://theory.stanford.edu/~amitp/GameProgramming/Heuristics.html

// AESHelper contains all the info needed to conduct the a star search.
type AESHelper struct {
	graph     dataStructs.Graph
	output    structs.OutputSpec
	startPos  int
	startN    *AESNode
	targetPos int
	targetN   *AESNode
	vVisited  []bool
	mapGrid   [][]int
}

// AESNode is a helper for the algorithm.
type AESNode struct {
	vertNum int
	vertX   int
	vertY   int
	parent  *AESNode
	estDist int
	pCost   int
}

// InitAES creates all the necessary info from the graph&mapInfo.
func InitAES(mapInfo structs.MapSpec, graph dataStructs.Graph) *AESHelper {
	helper := new(AESHelper)
	helper.graph = graph
	helper.startPos = mapInfo.StartVert
	helper.targetPos = mapInfo.GoalVert
	helper.mapGrid = mapInfo.OrigMap
	helper.vVisited = make([]bool, graph.GetNumVerticies())

	helper.startN = &AESNode{
		vertNum: mapInfo.StartVert,
		vertX:   mapInfo.StartPosX,
		vertY:   mapInfo.StartPosY,
		parent:  nil,
		estDist: 0,
		pCost:   0,
	}

	helper.targetN = &AESNode{
		vertNum: mapInfo.GoalVert,
		vertX:   mapInfo.GoalPosX,
		vertY:   mapInfo.GoalPosY,
		estDist: 0,
		pCost:   0,
	}

	return helper
}

// FindPathAES returns the output struct with a path found from the aes search alg.
func (aes *AESHelper) FindPathAES() (structs.OutputSpec, error) {
	output := structs.OutputSpec{}
	var finalPath, openS, closeS []*AESNode

	openS = append(openS, aes.startN)

	for len(openS) != 0 {
		currN := aes.getLowestEuclidean(openS)

		if currN.parent != nil {
			vCst, err := aes.graph.GetWeight(currN.vertNum)
			if err != nil {
				return output, errors.New("aes error getting cost")
			}
			currN.pCost = currN.parent.pCost + vCst.(int)
		}

		if currN.vertNum == aes.targetN.vertNum {
			// made it to end node, get final path and break out of loop
			finalPath = aes.getFinalPath(currN)
			break
		}

		if len(finalPath) > 1 {
			break
		}

		openS = aes.removeN(openS, currN)
		closeS = append(closeS, currN)

		// get all adjacent verticies
		edges := aes.graph.GetEdgesForVerticy(currN.vertNum)

		for val := range edges.Data {
			if !aes.hasNode(closeS, val.(int)) {
				aes.vVisited[val.(int)] = true

				// create/define adjacent node
				nextN := aes.getNodeFromVertNum(val.(int))
				estCst := aes.getEuclidean(nextN)
				if estCst == -1 {
					return output, errors.New("aes error calculating Euclidean")
				}
				nextN.estDist = currN.estDist + estCst

				// append it to the open nodes (to reference later)
				if !aes.hasNode(openS, nextN.vertNum) {
					openS = append(openS, nextN)
				}

				nextN.parent = currN
				if val.(int) == aes.targetN.vertNum {
					finalPath = aes.getFinalPath(nextN)
					break
				}
			}
		}
	}

	if finalPath == nil {
		return output, errors.New("aes error final path doesn't exist")
	}

	output.AlgType = "A* Search With Euclidean Heuristic"
	output.Ppath = aes.getVertPathFromNodes(finalPath)
	output.Pmoves = len(output.Ppath)
	output.Pvisited = aes.getVisitedVerticies()
	output.Pcost = aes.getPathCost(output.Ppath)

	if output.Pcost == -1 {
		return output, errors.New("bfs can't find cost of path")
	}

	return output, nil
}

// getEuclidean returns the heuristic for a specific verticy.
func (aes *AESHelper) getEuclidean(vert *AESNode) int {
	// will add vertical and horizontal path to the target position
	// 	regardless of water, but water will add 5 to path instead of 0
	vNum := 0
	estCostR := 0
	estCostC := 0
	estCost := 0
	for rowNum, row := range aes.mapGrid {
		for colNum := range row {
			if rowNum == vert.vertX && vert.vertNum < vNum {
				tmp, err := aes.graph.GetWeight(vNum)
				if err != nil {
					return -1
				}
				if tmp.(int) == 0 {
					estCostR += 5
				}
				estCostR += tmp.(int)
			}
			if colNum == vert.vertY && vert.vertNum < vNum {
				tmp, err := aes.graph.GetWeight(vNum)
				if err != nil {
					return -1
				}
				if tmp.(int) == 0 {
					estCostC += 5
				}
				estCostC += tmp.(int)
			}
			vNum++
		}
	}
	estCost = (estCostR*estCostR + estCostC*estCostC)
	return estCost
}

// returns the verticy with the lowest calculated Euclidean distance.
func (aes *AESHelper) getLowestEuclidean(verts []*AESNode) *AESNode {
	if len(verts) == 0 {
		return nil
	}
	lowest := verts[0]
	bestEuclidean := lowest.estDist

	for _, v := range verts {
		if v.estDist < bestEuclidean {
			bestEuclidean = v.estDist
			lowest = v
		}
	}

	return lowest
}

func (aes *AESHelper) getNodeFromVertNum(vert int) *AESNode {
	vNum := 0
	tmp := new(AESNode)
	for rowNum, row := range aes.mapGrid {
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

func (aes *AESHelper) getFinalPath(endVert *AESNode) []*AESNode {
	var fPath []*AESNode
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

func (aes *AESHelper) hasNode(verts []*AESNode, sVert int) bool {
	for _, v := range verts {
		if v.vertNum == sVert {
			return true
		}
	}
	return false
}

func (aes *AESHelper) removeN(verts []*AESNode, vert *AESNode) []*AESNode {
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

func (aes *AESHelper) getVertPathFromNodes(finalPath []*AESNode) []int {
	finalVertPath := make([]int, 0)
	for _, v := range finalPath {
		finalVertPath = append(finalVertPath, v.vertNum)
	}
	return finalVertPath
}

func (aes *AESHelper) getPathCost(path []int) int {
	totalCost := 0
	for _, vert := range path {
		singleCost, err := aes.graph.GetWeight(vert)
		if err != nil {
			return -1
		}
		totalCost += singleCost.(int)
	}
	return totalCost
}

func (aes *AESHelper) getVisitedVerticies() []int {
	visited := make([]int, 0)
	for i := 0; i < aes.graph.GetNumVerticies(); i++ {
		if aes.vVisited[i] {
			visited = append(visited, i)
		}
	}
	return visited
}
