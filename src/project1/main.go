package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	initFlg()
	flag.Parse()

	var (
		startPosX int
		startPosY int
		goalPosX  int
		goalPosY  int
		width     int
		height    int
		i         int
		j         int
	)

	fmap, err := os.Open("p1Map.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fmap.Close()

	scanner := bufio.NewScanner(fmap)
	lCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		mInfo := strings.Split(line, " ")
		if len(mInfo) > 1 {
			i, err = strconv.Atoi(mInfo[0])
			if err != nil {
				log.Fatal(err)
				return
			}
			j, err = strconv.Atoi(mInfo[1])
			if err != nil {
				log.Fatal(err)
				return
			}
			// means actual map hasn't started yet
			switch lCount {
			case 0: // define width and height
				width = i
				height = j
			case 1: // define start coords
				startPosX = i
				startPosY = j
			case 2: // define goal coords
				goalPosX = i
				goalPosY = j
			default: // shouldn't get here
				log.Fatal("incorrect map definition")
			}
		} else {
			// defining the actual map
			for _, space := range line {
				s := rune(space)
				switch s {
				case 'R':
				case 'f':
				case 'F':
				case 'h':
				case 'r':
				case 'M':
				case 'W':
				}
			}
		}
		lCount++
	}
	// mapSpec := structs.MapSpec{
	// 	startPosX: startPosX,
	// 	startPosY: startPosY,
	// 	goalPosX:  goalPosX,
	// 	goalPosY:  goalPosY,
	// }
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
