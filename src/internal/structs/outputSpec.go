package structs

// OutputSpec contains all the info necessary to display the output.
type OutputSpec struct {
	AlgType  string
	OrigMap  [][]int
	Ppath    []int
	Pvisited []int
	Pmoves   int
	Pcost    int
}
