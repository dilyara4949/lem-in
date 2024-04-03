package internal

func Dfs(graph map[string][]string, start string, end string, path []string, allPaths *[][]string) {
	path = append(path, start)
	if start == end {
		*allPaths = append(*allPaths, append([]string{}, path...))
		return
	}
	for _, node := range graph[start] {
		if !contains(path, node) {
			Dfs(graph, node, end, path, allPaths)
		}
	}
}

func contains(path []string, node string) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}

func AllPossiblePaths(colony Colony) [][]string {
	startPoint := colony.Start
	endPoint := colony.End

	var allPaths [][]string
	Dfs(colony.Ants, startPoint, endPoint, []string{}, &allPaths)

	return (allPaths)
}

func AllCombinations(allpaths [][]string) [][][]string {
	res := [][][]string{}

	allpaths = Sort(allpaths)

	for i := 0; i < len(allpaths); i++ {
		curr := [][]string{}
		curr = append(curr, allpaths[i])
		for j := i + 1; j < len(allpaths); j++ {
			if !IsIntersects(allpaths[i], allpaths[j]) && !IsIntersectsCurr(allpaths[j], curr) {
				res = append(res, curr)
				curr = append(curr, allpaths[j])
			}
		}
		res = append(res, curr)
	}

	return res
}

func IsIntersectsCurr(p []string, curr [][]string) bool {
	for _, path := range curr {
		for j, r := range path {
			for i := 0; i < len(p)-1; i++ {
				if (i > 0 && j > 0) && r == p[i] {
					return true
				}
			}
		}
	}
	return false
}

func IsIntersects(p1 []string, p2 []string) bool {
	for i := 1; i < len(p1)-1; i++ {
		for j := 1; j < len(p2)-1; j++ {
			if p1[i] == p2[j] {
				return true
			}
		}
	}
	return false
}

func DeleteDublicates(paths *[][][]string) {
	for i := 0; i < len(*paths); i++ {
		for j := i + 1; j < len(*paths); j++ {
			if len((*paths)[i]) == len((*paths)[j]) {
				x := 0
				for k, path := range (*paths)[i] {
					if len(path) == len((*paths)[j][k]) {
						x++
					}
				}
				if x == len((*paths)[i]) {
					(*paths) = append((*paths)[:j], (*paths)[j+1:]...)
					j--
				}

			}
		}
	}
}

func FinalPath(allcombo [][][]string, cntAnt int) ([][]string, []int) {
	minTime := 9223372036854775807
	resAnts := []int{}
	res := [][]string{}
	for _, paths := range allcombo {
		if cntAnt == 0 {
			break
		}

		maxHeight := len(paths[len(paths)-1])
		underHeightTotal := 0
		for i := 0; i < len(paths)-1; i++ {
			underHeightTotal += maxHeight - len(paths[i])
		}
		time := 0
		var ants []int
		if cntAnt <= underHeightTotal { // dopolnit'
			ants, time = CoverUnder(paths, cntAnt)
		} else { // dopolnit' i sverhu raspredelit
			ants, time = CoverUp(paths, cntAnt, underHeightTotal)
		}
		if time < minTime {
			minTime = time
			res = paths
			resAnts = ants
		}
	}
	return res, resAnts
}

// 3 ants

//    #
// **##
// *###
// ####

// 6 ants

// ***#
// **##
// *###
// ####
func CoverUnder(paths [][]string, cntAnt int) ([]int, int) {

	i := 0
	ants := make([]int, len(paths))
	maxAnt := 0
	if len(paths) < 2 {
		ants[0] = cntAnt
		maxAnt = cntAnt

		return ants, maxAnt
	}
	for cntAnt != 0 {
		if ants[i]+len(paths[i]) < ants[i+1]+len(paths[i+1]) {
			d := len(paths[i+1]) + ants[i+1] - (len(paths[i]) + ants[i])
			if d <= cntAnt {
				ants[i] += d
				cntAnt -= d
			} else {
				ants[i] += cntAnt
				cntAnt -= cntAnt
			}
			maxAnt = max(maxAnt, ants[i])
		}

		i++
		if i == len(paths)-1 {
			if ants[i-1]+len(paths[i-1]) == len(paths[i]) {
				break
			}
			i = 0
		}
	}
	return ants, maxAnt
}

// 8 ants

// **
// ***#
// **##
// *###
// ####
func CoverUp(paths [][]string, cntAnt, underHeightTotal int) ([]int, int) {

	ants := make([]int, len(paths))

	if underHeightTotal > 0 {
		ants, _ = CoverUnder(paths, cntAnt)

	}

	cntAnt -= underHeightTotal
	maxAnt := 0
	m := cntAnt / len(paths)
	rem := cntAnt % len(paths)

	if len(paths) < 2 {
		ants[0] = cntAnt
		maxAnt = cntAnt

		return ants, maxAnt

	}
	for i := range paths {
		ants[i] += m
		if rem > 0 {
			ants[i]++
			rem--
		}
		maxAnt = max(maxAnt, ants[i])
	}

	return ants, maxAnt
}
