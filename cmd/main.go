package main

import (
	"fmt"
	"lem-in/internal"
	"log"
)

func main() {

	colony, input := internal.ReadArg()

	allpaths := internal.AllPossiblePaths(colony)
	allcombs := internal.AllCombinations(allpaths)
	internal.DeleteDublicates(&allcombs)

	if len(allpaths) == 0 {
		log.Fatalln("Path from start till end does not exist")
	}
	finalpath, cntAnts := internal.FinalPath(allcombs, colony.Cnt)

	fmt.Println(input)
	fmt.Println()
	internal.DistributeAnts(finalpath, cntAnts, colony.Cnt, colony.End)

}
