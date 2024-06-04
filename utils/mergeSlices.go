package utils

import "KarimovbejanHW/hw_tasks/required"

func MergeSlices(first, second []int) []int {
	mergeSlice := append(first, second...)
	var resultSlice []int

	duplicates := required.GetDuplicates(mergeSlice)

	if duplicates == nil {
		return mergeSlice
	} else {
		for i := 0; i < len(mergeSlice); i++ {
			isDuplicate := false
			for j := 0; j < len(duplicates); j++ {
				if mergeSlice[i] == duplicates[j] {
					isDuplicate = true
					break
				}
			}

			if !isDuplicate {
				resultSlice = append(resultSlice, mergeSlice[i])
			}
		}

		return resultSlice
	}
}
