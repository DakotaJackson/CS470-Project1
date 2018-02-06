package dataStructs

import (
	"errors"

	"github.com/DakotaJackson/CS470-Project1/src/internal/structs"
)

// Following is a graph which is used for the traversal algorithms.
// 	It should be noted that basic functionality and ideas were referenced
// 	from here: https://github.com/karalabe/cookiejar/blob/v2/graph/graph.go

// Graph is the data structure that is traversed.
type Graph struct {
	verticies int
	edges     []*EdgeHelper
	gInfo     map[int]interface{}
}

// EdgeHelper is a helper struct for info about edges
type EdgeHelper struct {
	size int
	Data map[interface{}]int
}

// InitGraph creates a blank graph, with specified number of verticies.
func InitGraph(numVerticies int) *Graph {
	g := &Graph{
		verticies: numVerticies,
		edges:     make([]*EdgeHelper, numVerticies),
		gInfo:     make(map[int]interface{}),
	}

	for i := 0; i < numVerticies; i++ {
		g.edges[i] = InitEdgeHelper()
	}
	return g
}

// MakeGraphFromMap takes in the width and height from the specified map
// 	and creates all the necessary connections.
func (g *Graph) MakeGraphFromMap(mapDef [][]int, width, height int) {
	// will connect the NSEW verticies to all parts of the 2d map array
	// 	(it will ignore the water squares and not connect them)

	vertNum := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if mapDef[i][j] == 0 {
				// current vert is water (still increment counter)
				vertNum++
				continue
			} else if i == 0 {
				// top row
				if j != width-1 && mapDef[i][j+1] != 0 {
					g.ConnectVerticies(vertNum, vertNum+1) // right
				}
				if j != 0 && mapDef[i][j-1] != 0 {
					g.ConnectVerticies(vertNum, vertNum-1) // left
				}
				if mapDef[i+1][j] != 0 {
					g.ConnectVerticies(vertNum, vertNum+width) // below
				}
			} else if i == height-1 {
				// bottom row
				if j != width-1 && mapDef[i][j+1] != 0 {
					g.ConnectVerticies(vertNum, vertNum+1) // right
				}
				if j != 0 && mapDef[i][j-1] != 0 {
					g.ConnectVerticies(vertNum, vertNum-1) // left
				}
				if mapDef[i-1][j] != 0 {
					g.ConnectVerticies(vertNum, vertNum-width) // above
				}
			} else if j == 0 {
				// left col
				if mapDef[i][j+1] != 0 {
					g.ConnectVerticies(vertNum, vertNum+1) // right
				}
				if i != 0 && mapDef[i-1][j] != 0 {
					g.ConnectVerticies(vertNum, vertNum-width) // above
				}
				if i != height-1 && mapDef[i+1][j] != 0 {
					g.ConnectVerticies(vertNum, vertNum+width) // below
				}
			} else if j == width-1 {
				// right col
				if mapDef[i][j-1] != 0 {
					g.ConnectVerticies(vertNum, vertNum-1) // left
				}
				if i != 0 && mapDef[i-1][j] != 0 {
					g.ConnectVerticies(vertNum, vertNum-width) // above
				}
				if i != height-1 && mapDef[i+1][j] != 0 {
					g.ConnectVerticies(vertNum, vertNum+width) // below
				}
			} else {
				// all 4 directions can be connected
				if mapDef[i][j+1] != 0 {
					g.ConnectVerticies(vertNum, vertNum+1) // right
				}
				if mapDef[i][j-1] != 0 {
					g.ConnectVerticies(vertNum, vertNum-1) // left
				}
				if mapDef[i-1][j] != 0 {
					g.ConnectVerticies(vertNum, vertNum-width) // above
				}
				if mapDef[i+1][j] != 0 {
					g.ConnectVerticies(vertNum, vertNum+width) // below
				}
			}
			vertNum++
		}
	}
}

// ConnectVerticies takes two verticies of a map and creates a connection.
func (g *Graph) ConnectVerticies(v1, v2 int) {
	g.edges[v1].MakeConnection(v2)
	if v1 != v2 {
		g.edges[v2].MakeConnection(v1)
	}
}

// MakeWeightsFromMap takes the 2d array of weighted movement and adds them
// 	to the graph data structure.
func (g *Graph) MakeWeightsFromMap(mapDef [][]int, width, height int) {
	vertNum := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			g.AssignWeight(vertNum, mapDef[i][j])
			vertNum++
		}
	}
}

// AssignWeight sets the movement penalty for a specific verticy.
func (g *Graph) AssignWeight(verticy int, data interface{}) error {
	// only give a weight to the
	if verticy <= g.verticies {
		g.gInfo[verticy] = data
		return nil
	}
	return errors.New("can't assign weight, verticy does not exist")
}

// GetWeight finds the movement penalty for a specific verticy.
func (g *Graph) GetWeight(verticy int) (interface{}, error) {
	if verticy <= g.verticies {
		return g.gInfo[verticy], nil
	}
	return nil, errors.New("can't get weight, verticy does not exist")
}

// DoVerticy executes the Do function on a verticy.
func (g *Graph) DoVerticy(verticy int, f func(interface{})) error {
	if verticy <= g.verticies {
		g.edges[verticy].DoEdgeForVerticy(f)
		return nil
	}
	return errors.New("can't execute do, verticy does not exist")
}

// GetNumVerticies simply returns the number of verticies in the graph.
func (g *Graph) GetNumVerticies() int {
	return g.verticies
}

// GetStartEndVerticies returns the start and end verticy number.
func (g *Graph) GetStartEndVerticies(mapInfo structs.MapSpec) (int, int) {
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

// GetEdgesForVerticy returns the edges for a specified verticy.
func (g *Graph) GetEdgesForVerticy(vertNum int) *EdgeHelper {
	return g.edges[vertNum].GetEdges()
}

/*
 * ------------------------------
 * START OF EDGE HELPER FUNCTIONS
 * ------------------------------
 */

// InitEdgeHelper creates a blank edge helper for use on each edge.
func InitEdgeHelper() *EdgeHelper {
	return &EdgeHelper{
		0,
		make(map[interface{}]int),
	}
}

// MakeConnection adds information to the edge needed for proper functionality.
func (e *EdgeHelper) MakeConnection(value interface{}) {
	e.Data[value]++
	e.size++
}

// DoEdgeForVerticy executes a function on each connection for a verticy.
func (e *EdgeHelper) DoEdgeForVerticy(f func(interface{})) {
	for val, count := range e.Data {
		for ; count > 0; count-- {
			f(val)
		}
	}
}

// GetEdges returns all the edges of the node.
func (e *EdgeHelper) GetEdges() *EdgeHelper {
	return e
}
