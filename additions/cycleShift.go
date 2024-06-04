package additions

import (
	"KarimovbejanHW/hw_tasks/required"
	"KarimovbejanHW/hw_tasks/utils"
)

func CycleShift(slice []int, cycleShifts int) []int {
	var sliceCycleShiftObjects []int
	var resulOfCycleShift []int

	if cycleShifts > len(slice) {
		for len(slice) < cycleShifts {
			cycleShifts -= len(slice)
		}
	}

	sliceCycleShiftObjects = slice[:len(slice)-cycleShifts]
	sliceCycleShiftObjects = slice[len(slice)-cycleShifts:]

	slice = append(slice, sliceCycleShiftObjects...)

	duplicates := required.GetDuplicates(slice)

	if duplicates == nil {
		return slice
	} else {
		resulOfCycleShift = utils.MergeSlices(slice, sliceCycleShiftObjects)
		sliceCycleShiftObjects = append(sliceCycleShiftObjects, resulOfCycleShift...)
		slice = sliceCycleShiftObjects
	}

	return slice
}
