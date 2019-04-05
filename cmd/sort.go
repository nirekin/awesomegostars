package cmd

import (
	"fmt"
	"sort"
)

func sortResponses(sortingKey string, toSort []Response) ([]Response, error) {
	r := make([]Response, 0)
	temp := make(map[int]map[string]Response)

	switch sortingKey {
	case keyStart, keyFork, keyWatch, keyIssues:
		for _, v := range toSort {
			i := v.sortingValue(sortingKey)
			if val, ok := temp[i]; ok {
				val[v.name] = v
				temp[i] = val
			} else {
				iN := make(map[string]Response)
				iN[v.name] = v
				temp[i] = iN
			}
		}

		// Sorting base on the sorting key
		var keys []int
		for k := range temp {
			keys = append(keys, k)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))

		for _, k := range keys {
			byKey := temp[k]
			// Sorting base on the project name
			var names []string
			for k := range byKey {
				names = append(names, k)
			}
			sort.Strings(names)
			for _, k := range names {
				r = append(r, byKey[k])
			}
		}
	default:
		return r, fmt.Errorf("Invalid sorting key \"%s\"", sortingKey)
	}
	return r, nil
}
