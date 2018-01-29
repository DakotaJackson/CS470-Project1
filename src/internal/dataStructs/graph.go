package dataStructs

import "errors"

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
	data map[interface{}]int
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

	// TODO: finish this functionality
	vertNum := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
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
func (g *Graph) MakeWeightsFromMap(mapDef [][]int) {
	// TODO: finish this functionality
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

// DoVerticy executes the Do function on a verticy, which
func (g *Graph) DoVerticy(verticy int, f func(interface{})) error {
	if verticy <= g.verticies {
		g.edges[verticy].DoEdgeForVerticy(f)
		return nil
	}
	return errors.New("can't execute do, verticy does not exist")
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
	e.data[value]++
	e.size++
}

// DoEdgeForVerticy executes a function on each connection for a verticy.
func (e *EdgeHelper) DoEdgeForVerticy(f func(interface{})) {
	for val, count := range e.data {
		for ; count > 0; count-- {
			f(val)
		}
	}
}
