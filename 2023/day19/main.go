package main

import (
	"fmt"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	instSet, parts := getStepsAndParts(input)
	fmt.Println("answer for part 1: ", ans1(instSet, parts))
	fmt.Println("answer for part 2: ", ans2(instSet))
}

func ans1(instSet map[string][]inst, parts []part) int {
	accepted := evaluateParts(parts, instSet)

	sum := 0
	for _, i := range accepted {
		sum += i.s + i.a + i.x + i.m
	}
	return sum
}

func ans2(instSet map[string][]inst) int {
	accepted := evaluateParts2(instSet)
	sum := 0
	for _, i := range accepted {
		sum += (i.x.end - i.x.start + 1) * (i.m.end - i.m.start + 1) * (i.a.end - i.a.start + 1) * (i.s.end - i.s.start + 1)
	}
	return sum
}

type part struct {
	x int
	m int
	a int
	s int
}

type inst struct {
	label  string
	symbol rune
	num    int
	dest   string
}

func getStepsAndParts(input []string) (map[string][]inst, []part) {
	isParts := false
	parts := []part{}
	instSet := make(map[string][]inst)
	for _, line := range input {
		if len(line) == 0 {
			isParts = true
			continue
		}

		if !isParts {
			name := line[:strings.Index(line, "{")]
			instSet[name] = []inst{}
			insts := strings.Split(line[strings.Index(line, "{")+1:len(line)-1], ",")
			for _, current := range insts {
				i := inst{}
				if len(current) < 2 {
					i.dest = current
					instSet[name] = append(instSet[name], i)
					continue
				}
				if current[1] == '>' || current[1] == '<' {
					i.label = string(current[0])
					i.symbol = rune(current[1])
					semicolonIndex := strings.Index(current, ":")
					i.num = aoc.FetchNumFromStringIgnoringNonNumeric(current[2:semicolonIndex])
					i.dest = current[semicolonIndex+1:]
				} else {
					i.dest = current
				}
				instSet[name] = append(instSet[name], i)
			}
		} else {
			line = line[1 : len(line)-1]
			componenets := strings.Split(line, ",")
			p := part{}
			for _, c := range componenets {
				num := aoc.FetchNumFromStringIgnoringNonNumeric(c[2:])
				switch c[0] {
				case 'x':
					p.x = num
				case 'm':
					p.m = num
				case 'a':
					p.a = num
				case 's':
					p.s = num
				}
			}
			parts = append(parts, p)
		}
	}
	return instSet, parts
}

func evaluateParts(parts []part, instSet map[string][]inst) []part {
	accepted := []part{}

	for _, p := range parts {
		stop := false
		label := "in"
		for !stop {
			if label == "A" {
				accepted = append(accepted, p)
				stop = true
			}

			if label == "R" {
				stop = true
			}

			label = findNextLabel(instSet, p, label)
		}
	}
	return accepted
}

type part2 struct {
	x ranges
	m ranges
	a ranges
	s ranges
}

type ranges struct {
	start int
	end   int
}

func evaluateParts2(instSet map[string][]inst) []part2 {
	parts := []part2{
		{
			x: ranges{1, 4000},
			m: ranges{1, 4000},
			a: ranges{1, 4000},
			s: ranges{1, 4000},
		},
	}

	accepted := []part2{}

	for len(parts) != 0 {
		stop := false
		label := "in"
		p := parts[0]
		for !stop {
			if label == "A" {
				accepted = append(accepted, p)
				stop = true
			}

			if label == "R" {
				stop = true
			}

			newlabel, new := findNextLabel2(instSet, p, label)
			if len(new) > 1 {
				parts = append(parts, new[1])
			}
			p = new[0]
			label = newlabel
		}
		parts = parts[1:]
	}
	return accepted
}

func findNextLabel2(instSet map[string][]inst, p part2, label string) (string, []part2) {
	newLabel := label
	newPart := p
	for _, i := range instSet[label] {
		if len(i.label) != 0 {
			switch i.label {
			case "x":
				switch i.symbol {
				case '>':
					if p.x.start > i.num {
						newLabel = i.dest
					} else if p.x.start < i.num && p.x.end > i.num {
						newPart.x.start = p.x.start
						newPart.x.end = i.num
						p.x.start = i.num + 1
						newLabel = i.dest
					}
				case '<':
					if p.x.end < i.num {
						newLabel = i.dest
					} else if p.x.start < i.num && p.x.end > i.num {
						newPart.x.start = i.num
						newPart.x.end = p.x.end
						p.x.end = i.num - 1
						newLabel = i.dest
					}
				}
			case "m":
				switch i.symbol {
				case '>':
					if p.m.start > i.num {
						newLabel = i.dest
					} else if p.m.start < i.num && p.m.end > i.num {
						newPart.m.start = p.m.start
						newPart.m.end = i.num
						p.m.start = i.num + 1
						newLabel = i.dest
					}
				case '<':
					if p.m.end < i.num {
						newLabel = i.dest
					} else if p.m.start < i.num && p.m.end > i.num {
						newPart.m.start = i.num
						newPart.m.end = p.m.end
						p.m.end = i.num - 1
						newLabel = i.dest
					}
				}
			case "a":
				switch i.symbol {
				case '>':
					if p.a.start > i.num {
						newLabel = i.dest
					} else if p.a.start < i.num && p.a.end > i.num {
						newPart.a.start = p.a.start
						newPart.a.end = i.num
						p.a.start = i.num + 1
						newLabel = i.dest
					}
				case '<':
					if p.a.end < i.num {
						newLabel = i.dest
					} else if p.a.start < i.num && p.a.end > i.num {
						newPart.a.start = i.num
						newPart.a.end = p.a.end
						p.a.end = i.num - 1
						newLabel = i.dest
					}
				}
			case "s":
				switch i.symbol {
				case '>':
					if p.s.start > i.num {
						newLabel = i.dest
					} else if p.s.start < i.num && p.s.end > i.num {
						newPart.s.start = p.s.start
						newPart.s.end = i.num
						p.s.start = i.num + 1
						newLabel = i.dest
					}
				case '<':
					if p.s.end < i.num {
						newLabel = i.dest
					} else if p.s.start < i.num && p.s.end > i.num {
						newPart.s.start = i.num
						newPart.s.end = p.s.end
						p.s.end = i.num - 1
						newLabel = i.dest
					}
				}
			default:
				panic("none matched")
			}
		} else {
			newLabel = i.dest
			break
		}
		if newLabel != label {
			break
		}
	}

	final := []part2{p}

	if p != newPart {
		final = append(final, newPart)
	}
	return newLabel, final
}

func findNextLabel(instSet map[string][]inst, p part, label string) string {
	newLabel := label
	for _, i := range instSet[label] {
		if len(i.label) != 0 {
			switch i.label {
			case "x":
				switch i.symbol {
				case '>':
					if p.x > i.num {
						newLabel = i.dest
					}
				case '<':
					if p.x < i.num {
						newLabel = i.dest
					}
				}
			case "m":
				switch i.symbol {
				case '>':
					if p.m > i.num {
						newLabel = i.dest
					}
				case '<':
					if p.m < i.num {
						newLabel = i.dest
					}
				}
			case "a":
				switch i.symbol {
				case '>':
					if p.a > i.num {
						newLabel = i.dest
					}
				case '<':
					if p.a < i.num {
						newLabel = i.dest
					}
				}
			case "s":
				switch i.symbol {
				case '>':
					if p.s > i.num {
						newLabel = i.dest
					}
				case '<':
					if p.s < i.num {
						newLabel = i.dest
					}
				}
			default:
				panic("none matched")
			}
		} else {
			newLabel = i.dest
			break
		}
		if newLabel != label {
			break
		}
	}

	return newLabel
}
