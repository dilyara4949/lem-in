package internal

import (
	"fmt"
)

func DistributeAnts(path [][]string, antsPathCnt []int, totalAnts int, end string) {
	res := make([][]string, 1000009)
	Ants := MakeAnts(totalAnts, path[0][0])
	anoth := make([]int, len(antsPathCnt))
	curAnt := 1

	copy(anoth[:], antsPathCnt[:])
	for curAnt <= totalAnts {

		for i := 0; i < len(path); i++ {
			if antsPathCnt[i] > 0 {
				for j := 1; j < len(path[i]); j++ {
					ans := fmt.Sprint("L", curAnt, "-", path[i][j], " ")
					Ants[curAnt-1].PrevRooms++

					res[Ants[curAnt-1].PrevRooms+anoth[i]-antsPathCnt[i]] = append(res[Ants[curAnt-1].PrevRooms+anoth[i]-antsPathCnt[i]], ans)

					if path[i][j] == end {
						antsPathCnt[i]--
						curAnt++
					}
				}
			} else {
				break
			}

		}
	}
	for _, s := range res {
		if len(s) > 0 {
			for _, c := range s {
				fmt.Print(c, " ")
			}
			fmt.Println()
		}
	}
}

func MakeAnts(totalAnts int, start string) []Ant {
	res := []Ant{}
	for i := 1; i <= totalAnts; i++ {
		res = append(res, Ant{ID: i, Room: start, PrevRooms: 0})
	}
	return res
}
