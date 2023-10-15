package main

import (
	. "com.vorto.vehiclerouter/pkg/models"
	"com.vorto.vehiclerouter/pkg/reader"
	"com.vorto.vehiclerouter/pkg/routing/loadproximity"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// main function reads loads from given file path, processes them
// and uses a greedy algorithm to deliver the loads with minimal cost
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter the path to an input file")
		return
	}
	filePath := os.Args[1]

	loads := getLoads(filePath)
	Process(loads)
}

func Process(loads []Load) {
	routingAlgorithm := &loadproximity.ClosestLoadGreedy{}
	scheduledLoads := routingAlgorithm.ScheduleLoads(loads)
	printSolution(scheduledLoads)
}

func getLoads(filePath string) []Load {
	loadReader := reader.NewLoadsFileReader()
	loads, err := loadReader.Read(filePath)
	if err != nil {
		panic(err)
	}
	return loads
}

// printSolution prints the loads assigned to each driver.
// Each driver's assignments are displayed on a separate line enclosed in brackets.
//
// An example output could be:
// [1]
// [4,2]
// [3]
//
// This means one driver does load 1; another driver does load 4 followed by load 2;
// and a final driver does load 3.
func printSolution(assignments []DriverDeliveryAssignment) {
	for _, assignment := range assignments {
		fmt.Printf("[%s]\n", strings.Join(extractLoadIdAndConvertToStrArr(assignment.Loads), ","))
	}
}

func extractLoadIdAndConvertToStrArr(loads []Load) []string {
	strArr := make([]string, len(loads))
	for i, load := range loads {
		strArr[i] = strconv.Itoa(load.LoadId)
	}
	return strArr
}
