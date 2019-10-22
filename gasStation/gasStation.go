package gasStation

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		if checkThisRoad(i, gas, cost) {
			return i
		}
	}
	return -1
}

func checkThisRoad(ind int, gas []int, cost []int) bool {
	rank := 0
	for i, val := range gas[ind:] {
		rank += val - cost[i + ind]
		if rank >= 0{
			continue
		} else {
			return false
		}
	}

	for i, val := range gas[:ind] {
		rank += val - cost[i]
		if rank >= 0{
			continue
		} else {
			return false
		}
	}
	return true
}
