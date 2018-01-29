package structs

// MapSpec contains all info for defining the map.
type MapSpec struct {
	OrigMap   [][]int
	TrSpaces  [][]int
	CurrTrMap [][]int
	StartPosX int
	StartPosY int
	GoalPosX  int
	GoalPosY  int
	Width     int
	Height    int
}
