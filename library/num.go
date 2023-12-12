package aoc

import (
	"fmt"
	"strings"
)

func ConvertIntSliceToString(intSlice []int, sep string) string {
	var build strings.Builder
	for i, num := range intSlice {
		build.WriteString(fmt.Sprintf("%d", num))
		if i != len(intSlice) {
			build.WriteString(sep)
		}
	}
	return build.String()
}
