package main

import (
	"fmt"
	"sort"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	lanMap := getLANMap(input)

	fmt.Println("answer to part 1: ", findT(find3Connected(lanMap)))
	fmt.Println("answer to part 2: ", findMaxCliques(lanMap))
}

func getLANMap(input []string) map[string]map[string]struct{} {
	lanMap := make(map[string]map[string]struct{})

	for _, line := range input {
		comp1, comp2 := line[:2], line[3:]

		if _, ok := lanMap[comp1]; !ok {
			lanMap[comp1] = make(map[string]struct{})
		}

		if _, ok := lanMap[comp2]; !ok {
			lanMap[comp2] = make(map[string]struct{})
		}

		lanMap[comp1][comp2] = struct{}{}
		lanMap[comp2][comp1] = struct{}{}
	}
	return lanMap
}

type lan struct {
	a, b, c string
}

func find3Connected(lanMap map[string]map[string]struct{}) map[lan]struct{} {
	result := make(map[lan]struct{})
	for key, value := range lanMap {
		if len(value) <= 1 {
			continue
		}

		for key2, _ := range value {
			for key3, _ := range value {
				if key2 == key3 {
					continue
				}

				if _, ok := lanMap[key2][key3]; ok {
					r := []string{key, key2, key3}
					sort.Strings(r)
					result[lan{r[0], r[1], r[2]}] = struct{}{}
				}
			}
		}
	}
	return result
}

func findT(input map[lan]struct{}) int {
	count := 0
	for key, _ := range input {
		if string(key.a[0]) == "t" || string(key.b[0]) == "t" || string(key.c[0]) == "t" {
			count++
		}
	}
	return count
}

func BronKerbosch(currentClique []string, yetToConsider []string, alreadyConsidered []string, lanMap map[string]map[string]struct{}, cliques [][]string) [][]string {
	if len(yetToConsider) == 0 && len(alreadyConsidered) == 0 {
		cliques = append(cliques, append([]string{}, currentClique...))
		return cliques
	}

	for index := 0; index < len(yetToConsider); {
		node := yetToConsider[index]
		newYetToConsider := []string{}
		newAlreadyConsidered := []string{}

		for _, n := range yetToConsider {
			if _, ok := lanMap[node][n]; ok {
				newYetToConsider = append(newYetToConsider, n)
			}
		}

		for _, n := range alreadyConsidered {
			if _, ok := lanMap[node][n]; ok {
				newAlreadyConsidered = append(newAlreadyConsidered, n)
			}
		}

		cliques = BronKerbosch(append(currentClique, node), newYetToConsider, newAlreadyConsidered, lanMap, cliques)

		yetToConsider = append(yetToConsider[:index], yetToConsider[index+1:]...)
		alreadyConsidered = append(alreadyConsidered, node)
	}
	return cliques
}

func findMaxCliques(lanMap map[string]map[string]struct{}) string {
	maxClique := []string{}
	allComputers := []string{}
	for key, _ := range lanMap {
		allComputers = append(allComputers, key)
	}
	cliques := BronKerbosch([]string{}, allComputers, []string{}, lanMap, [][]string{})
	for _, c := range cliques {
		if len(c) > len(maxClique) {
			maxClique = c
		}
	}
	sort.Strings(maxClique)
	return strings.Join(maxClique, ",")
}
