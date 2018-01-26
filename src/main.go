package main

import "flag"

func main() {
	initFlg()
	flag.Parse()
}

func initFlg() {
	flag.Bool("all", false, "run and get output for all algorithms")
	flag.Bool("breadth", false, "run for breadth first algorithm")
	flag.Bool("lowCost", false, "run for lowest cost algorithm")
	flag.Bool("gBestFirst", false, "run for greedy best first algorithm")
	flag.Bool("aEuclidean", false, "run for A* w/Euclidean heuristic algorithm")
	flag.Bool("aOctile", false, "run for A* w/Octile heuristic algorithm")
	flag.String("map", "", "specify the map file to build from")
}
