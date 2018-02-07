package util

import (
	"fmt"
	"strconv"

	"github.com/DakotaJackson/CS470-Project1/src/internal/structs"
	"github.com/fatih/color"
)

// PrintOutput is the temporary way to print output to terminal.
// 	To achieve the output with color (see key for meanings), the package
//  https://github.com/fatih/color    was used.
func PrintOutput(output structs.OutputSpec) {

	// set up colors to print initial information
	info := color.New(color.FgHiYellow, color.Bold, color.BgBlue)
	name := color.New(color.FgHiYellow)

	name.Printf("CURRENT ALGORITHM: ")
	info.Printf(output.AlgType)

	fmt.Printf("\n")
	name.Printf("TOTAL NUMBER OF MOVES: ")
	info.Printf(strconv.Itoa(output.Pmoves))

	fmt.Printf("\n")
	name.Printf("TOTAL COST OF MOVES: ")
	info.Printf(strconv.Itoa(output.Pcost))

	fmt.Printf("\n\n")
	name.Printf("...PRINTING OUT MAP...")
	fmt.Printf("\n\n")

	// prints out the map with a key
	printKey()
	printColor(output)

}

func printPath(output structs.OutputSpec) {
	vert := 0
	var line string
	var div string
	for x, z := range output.OrigMap {
		line = ""
		div = ""
		for j, i := range z {
			if j == 0 {
				line = "|"
				div = div + " "
			}
			if isInSlice(vert, output.Ppath) {
				if i == 10 {
					line = line + " " + strconv.Itoa(i) + "|"
				} else {
					line = line + " " + strconv.Itoa(i) + " |"
				}
			} else {
				line = line + "   |"
			}
			div = div + "--- "
			vert++
		}
		if x == 0 {
			fmt.Println(div)
		}
		fmt.Println(line)
		fmt.Println(div)
	}
}

func printMap(output structs.OutputSpec) {
	var line string
	var div string
	for x, z := range output.OrigMap {
		line = ""
		div = ""
		for j, i := range z {
			if j == 0 {
				line = "|"
				div = div + " "
			}
			if i == 10 {
				line = line + " " + strconv.Itoa(i) + "|"
			} else {
				line = line + " " + strconv.Itoa(i) + " |"
			}
			div = div + "--- "
		}
		if x == 0 {
			fmt.Println(div)
		}
		fmt.Println(line)
		fmt.Println(div)
	}
}

func printVisited(output structs.OutputSpec) {
	vert := 0
	var line string
	var div string
	for x, z := range output.OrigMap {
		line = ""
		div = ""
		for j, i := range z {
			if j == 0 {
				line = "|"
				div = div + " "
			}
			if isInSlice(vert, output.Pvisited) {
				if i == 10 {
					line = line + " " + strconv.Itoa(i) + "|"
				} else {
					line = line + " " + strconv.Itoa(i) + " |"
				}
			} else {
				line = line + "   |"
			}
			div = div + "--- "
			vert++
		}
		if x == 0 {
			fmt.Println(div)
		}
		fmt.Println(line)
		fmt.Println(div)
	}
}

func printColor(output structs.OutputSpec) {
	pathV := color.New(color.FgHiBlue, color.Underline)
	pathC := color.New(color.FgHiGreen, color.Bold, color.Underline)
	vert := 0
	var line string
	var div string
	for x, z := range output.OrigMap {
		line = ""
		div = ""
		if x == 0 {
			for p := range z {
				if p == 0 {
					div = div + " "
				}
				div = div + "--- "
			}
			fmt.Println(div)
			div = ""
		}
		for j, i := range z {
			if j == 0 {
				//line = "|"
				fmt.Printf("|")
				div = div + " "
			}
			if isInSlice(vert, output.Ppath) {
				if i == 10 {
					//line = " " + strconv.Itoa(i) + "|"
					fmt.Printf(" ")
					pathC.Printf(strconv.Itoa(i))
					fmt.Printf("|")
				} else {
					//line = " " + strconv.Itoa(i) + " |"
					fmt.Printf(" ")
					pathC.Printf(strconv.Itoa(i))
					fmt.Printf(" |")
				}
			} else if isInSlice(vert, output.Pvisited) {
				if i == 10 {
					//line = " " + strconv.Itoa(i) + "|"
					fmt.Printf(" ")
					pathV.Printf(strconv.Itoa(i))
					fmt.Printf("|")
				} else {
					//line = " " + strconv.Itoa(i) + " |"
					fmt.Printf(" ")
					pathV.Printf(strconv.Itoa(i))
					fmt.Printf(" |")
				}
			} else if i == 10 {
				line = " " + strconv.Itoa(i) + "|"
				fmt.Printf(line)
			} else {
				line = " " + strconv.Itoa(i) + " |"
				fmt.Printf(line)
			}

			div = div + "--- "
			vert++
		}
		fmt.Printf("\n")

		fmt.Println(div)

	}
}

func printKey() {

	fmt.Println(" ------------------------------------- ")
	fmt.Printf("| ")
	key := color.New(color.BgHiWhite, color.FgBlack)
	key.Printf("KEY FOR MAP:")
	pathC := color.New(color.FgHiGreen, color.Bold, color.Underline)
	fmt.Printf("\n| ")
	pathC.Println("Path Traveled = Bold Green")

	pathV := color.New(color.FgHiBlue, color.Underline)
	fmt.Printf("| ")
	pathV.Println("Visited Verticies = Underlined Blue")
	fmt.Printf("| ")
	fmt.Println("Unvisited Verticies = White")
	fmt.Println(" ------------------------------------- ")
}

func isInSlice(s int, sl []int) bool {
	for _, a := range sl {
		if s == a {
			return true
		}
	}
	return false
}
