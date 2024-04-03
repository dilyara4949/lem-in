package internal

import "sort"


func Sort(arr [][]string) [][]string {
    comparator := func(i, j int) bool {
        return len(arr[i]) < len(arr[j])
    }
    sort.Slice(arr, comparator)
    return arr
}