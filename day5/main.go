package main

import (
	"fmt"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	fmt.Println("answer for part1: ", createMap(input))
	fmt.Println("answer for part2: ", part2(input))
}

type zoo struct {
	end        int
	startValue int
}

func part2(input []string) int {
	seeds := make(map[int]int)
	seedSoil := make(map[int]zoo)
	soilFert := make(map[int]zoo)
	fertWater := make(map[int]zoo)
	waterLight := make(map[int]zoo)
	lightTemp := make(map[int]zoo)
	tempHum := make(map[int]zoo)
	humLoc := make(map[int]zoo)

	var mapPointer map[int]zoo
	for _, line := range input {
		if strings.Index(line, ":") != -1 && line[:strings.Index(line, ":")] == "seeds" {
			inputSeeds := aoc.FetchSliceOfIntsInString(line[strings.Index(line, ":"):])
			for i := 0; i < len(inputSeeds); i++ {
				seeds[inputSeeds[i]] = inputSeeds[i+1] + inputSeeds[i] - 1
				i++
			}
			continue
		}

		if strings.Index(line, "seed-to-soil") != -1 {
			mapPointer = seedSoil
			continue
		}

		if strings.Index(line, "soil-to-fertilizer") != -1 {
			mapPointer = soilFert
			continue
		}

		if strings.Index(line, "fertilizer-to-water") != -1 {
			mapPointer = fertWater
			continue
		}

		if strings.Index(line, "water-to-light") != -1 {
			mapPointer = waterLight
			continue
		}

		if strings.Index(line, "light-to-temperature") != -1 {
			mapPointer = lightTemp
			continue
		}

		if strings.Index(line, "temperature-to-humidity") != -1 {
			mapPointer = tempHum
			continue
		}

		if strings.Index(line, "humidity-to-location") != -1 {
			mapPointer = humLoc
			continue
		}

		nums := aoc.FetchSliceOfIntsInString(line)
		if len(nums) == 3 {
			valueStart := nums[0]
			keyStart := nums[1]
			count := nums[2] - 1

			mapPointer[valueStart] = zoo{
				end:        valueStart + count,
				startValue: keyStart,
			}
		}

	}
	i := 0
	for ; ; i++ {

		hum := findValue(humLoc, &i)
		temp := findValue(tempHum, hum)
		light := findValue(lightTemp, temp)
		water := findValue(waterLight, light)
		fert := findValue(fertWater, water)
		soil := findValue(soilFert, fert)
		seed := findValue(seedSoil, soil)

		for k, v := range seeds {
			if *seed >= k && *seed <= v {
				return i
			}
		}
	}

}

func createMap(input []string) int {
	inputSeeds := []int{}
	seedSoil := make(map[int]zoo)
	soilFert := make(map[int]zoo)
	fertWater := make(map[int]zoo)
	waterLight := make(map[int]zoo)
	lightTemp := make(map[int]zoo)
	tempHum := make(map[int]zoo)
	humLoc := make(map[int]zoo)

	var mapPointer map[int]zoo
	for _, line := range input {
		if strings.Index(line, ":") != -1 && line[:strings.Index(line, ":")] == "seeds" {
			inputSeeds = aoc.FetchSliceOfIntsInString(line[strings.Index(line, ":"):])
			continue
		}

		if strings.Index(line, "seed-to-soil") != -1 {
			mapPointer = seedSoil
			continue
		}

		if strings.Index(line, "soil-to-fertilizer") != -1 {
			mapPointer = soilFert
			continue
		}

		if strings.Index(line, "fertilizer-to-water") != -1 {
			mapPointer = fertWater
			continue
		}

		if strings.Index(line, "water-to-light") != -1 {
			mapPointer = waterLight
			continue
		}

		if strings.Index(line, "light-to-temperature") != -1 {
			mapPointer = lightTemp
			continue
		}

		if strings.Index(line, "temperature-to-humidity") != -1 {
			mapPointer = tempHum
			continue
		}

		if strings.Index(line, "humidity-to-location") != -1 {
			mapPointer = humLoc
			continue
		}

		nums := aoc.FetchSliceOfIntsInString(line)
		if len(nums) == 3 {
			valueStart := nums[0]
			keyStart := nums[1]
			count := nums[2] - 1

			mapPointer[keyStart] = zoo{
				end:        keyStart + count,
				startValue: valueStart,
			}
		}

	}

	var low *int

	for _, value := range inputSeeds {
		soil := findValue(seedSoil, &value)
		fert := findValue(soilFert, soil)
		water := findValue(fertWater, fert)
		light := findValue(waterLight, water)
		temp := findValue(lightTemp, light)
		hum := findValue(tempHum, temp)
		loc := findValue(humLoc, hum)

		locValue := *loc
		if low == nil {
			low = &locValue
		} else {
			if locValue < *low {
				low = &locValue
			}
		}
	}

	return *low

}

func findValue(lookIn map[int]zoo, key *int) *int {
	var val *int
	for k, v := range lookIn {
		if *key <= v.end && *key >= k {
			l := v.startValue + (*key - k)
			val = &l
		}
	}
	if val == nil {
		val = key
	}
	return val
}
