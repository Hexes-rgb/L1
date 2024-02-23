package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 150, 156}

	groupedTemps := group(temperatures)

	for key, _ := range groupedTemps {
		fmt.Printf("Group %d: %v\n", key, groupedTemps[key])
	}
}

func group(temperatures []float64) map[int][]float64 {
	groupedTemps := make(map[int][]float64) //make map for temp groups

	for _, temp := range temperatures {
		group := int(temp) / 10 * 10                            //get temp group
		groupedTemps[group] = append(groupedTemps[group], temp) //append temp to group,
	}
	return groupedTemps
}
