package structs

// MapSpec contains all info for defining the map.
type MapSpec struct {
	origMap   [][]int
	trSpaces  [][]int
	currTrMap [][]int
	startPosX int
	startPosY int
	goalPosX  int
	goalPosY  int
}
