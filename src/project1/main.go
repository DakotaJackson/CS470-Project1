package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/DakotaJackson/CS470-Project1/src/internal/dataStructs"
	"github.com/DakotaJackson/CS470-Project1/src/internal/structs"
)

var (
	allFlg        bool
	breadthFlg    bool
	lowCostFlg    bool
	gBestFirstFlg bool
	aEuclideanFlg bool
	aOctileFlg    bool
	mapFile       string
)

func main() {
	initFlg()
	flag.Parse()

	mapInfo, err := gatherMapInfo(mapFile)

	if err != nil {
		log.Fatal("can't gather initial map info: ", err)
		return
	}
	fmt.Println(mapInfo)

	graph, err := createGraph(mapInfo.OrigMap, mapInfo.Width, mapInfo.Height)

	if err != nil {
		log.Fatal("can't create graph from map 2d array", err)
		return
	}
	fmt.Println(graph)

	// run the proper command for each flag/algorithm needed
	if allFlg {

	} else if breadthFlg {

	} else if lowCostFlg {

	} else if gBestFirstFlg {

	} else if aEuclideanFlg {

	} else if aOctileFlg {

	}
}

// gatherMapInfo takes in a text file of the form specified by the assignment
// 	then creates a struct with all the necessary information
func gatherMapInfo(mapFile string) (structs.MapSpec, error) {
	var (
		mapArray  [][]int
		startPosX int
		startPosY int
		goalPosX  int
		goalPosY  int
		width     int
		height    int
		i         int
		j         int
	)
	fmap, err := os.Open(mapFile)
	if err != nil {
		return structs.MapSpec{}, err
	}
	defer fmap.Close()

	scanner := bufio.NewScanner(fmap)
	lCount := 0
	rowNum := 0
	colNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		mInfo := strings.Split(line, " ")
		if len(mInfo) > 1 {
			i, err = strconv.Atoi(mInfo[0])
			if err != nil {
				return structs.MapSpec{}, err
			}
			j, err = strconv.Atoi(mInfo[1])
			if err != nil {
				return structs.MapSpec{}, err
			}
			// means actual map hasn't started yet
			switch lCount {
			case 0: // define width and height
				width = i
				height = j
				mapArray = make([][]int, height)
				for w := range mapArray {
					mapArray[w] = make([]int, width)
				}
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
			colNum = 0
			for _, space := range line {
				s := rune(space)
				switch s {
				// defining movement costs
				case 'R':
					mapArray[rowNum][colNum] = 1
				case 'f':
					mapArray[rowNum][colNum] = 2
				case 'F':
					mapArray[rowNum][colNum] = 4
				case 'h':
					mapArray[rowNum][colNum] = 5
				case 'r':
					mapArray[rowNum][colNum] = 7
				case 'M':
					mapArray[rowNum][colNum] = 10
				case 'W':
					// 0 means can't travel on
					mapArray[rowNum][colNum] = 0
				default:
					mapArray[rowNum][colNum] = -1
					log.Fatal("error populating map")
				}
				colNum++
			}
			rowNum++
		}
		lCount++
	}
	// defines everything as a default, with values gathered from the map def.
	return structs.MapSpec{
		OrigMap:   mapArray,
		TrSpaces:  mapArray,
		CurrTrMap: mapArray,
		StartPosX: startPosX,
		StartPosY: startPosY,
		GoalPosX:  goalPosX,
		GoalPosY:  goalPosY,
		Width:     width,
		Height:    height,
	}, nil
}

func createGraph(mapDef [][]int, width, height int) (dataStructs.Graph, error) {
	numVerticies := width * height
	g := dataStructs.InitGraph(numVerticies)
	g.MakeGraphFromMap(mapDef, width, height)
	g.MakeWeightsFromMap(mapDef)
	return *g, nil
}

func initFlg() {
	flag.BoolVar(&allFlg, "all", false, "run and get output for all algorithms")
	flag.BoolVar(&breadthFlg, "breadth", false, "run for breadth first algorithm")
	flag.BoolVar(&lowCostFlg, "lowCost", false, "run for lowest cost algorithm")
	flag.BoolVar(&gBestFirstFlg, "gBestFirst", false, "run for greedy best first algorithm")
	flag.BoolVar(&aEuclideanFlg, "aEuclidean", false, "run for A* w/Euclidean heuristic algorithm")
	flag.BoolVar(&aOctileFlg, "aOctile", false, "run for A* w/Octile heuristic algorithm")
	flag.StringVar(&mapFile, "map", "", "specify the map file to build from")
}
