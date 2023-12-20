package main

import (
	"fmt"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	fmt.Println("answer for part 1: ", ans1(input))
	fmt.Println("answer for part 2: ", ans2(input))
}

type flipFlopModule struct {
	destination []string
	state       string
}

type conjunctionModule struct {
	source      map[string]pulse
	destination []string
}

type broadcast struct {
	destination []string
}

func makeMap(input []string) (broadcast, map[string]conjunctionModule, map[string]flipFlopModule) {
	b := broadcast{}
	c := make(map[string]conjunctionModule)
	f := make(map[string]flipFlopModule)

	for _, line := range input {
		dest := strings.Split(line[strings.Index(line, ">")+2:], ", ")
		if strings.Contains(line, "broadcaster") {
			b = broadcast{
				destination: dest,
			}
			continue
		}

		name := line[1:strings.Index(line, " ")]
		switch line[0] {
		case '&':
			c[name] = conjunctionModule{
				destination: dest,
			}
		case '%':
			f[name] = flipFlopModule{
				destination: dest,
				state:       "off",
			}
		}
	}

	for name, m := range f {
		for _, i := range m.destination {
			if _, ok := c[i]; ok {
				if len(c[i].source) == 0 {
					s := c[i]
					s.source = make(map[string]pulse)
					c[i] = s
				}
				c[i].source[name] = low
			}
		}
	}

	for name, m := range c {
		for _, i := range m.destination {
			if _, ok := c[i]; ok {
				if len(c[i].source) == 0 {
					s := c[i]
					s.source = make(map[string]pulse)
					c[i] = s
				}
				c[i].source[name] = low
			}
		}
	}

	return b, c, f
}

type pulse string

const (
	high pulse = "high"
	low  pulse = "low"
)

func handleFlipFlop(typeOfPulse pulse, source string, m flipFlopModule, c map[string]conjunctionModule) (flipFlopModule, pulse, []string) {
	switch typeOfPulse {
	case high:
		return m, "", nil
	case low:
		switch m.state {
		case "off":
			m.state = "on"
			updateConjunctionSources(high, source, m.destination, c)
			return m, high, m.destination
		case "on":
			m.state = "off"
			updateConjunctionSources(low, source, m.destination, c)
			return m, low, m.destination
		}
	}
	panic("invalid flip flop encountered")
}

func updateConjunctionSources(typeOfPulse pulse, source string, destinations []string, c map[string]conjunctionModule) {
	for _, dest := range destinations {
		if _, ok := c[dest]; ok {
			c[dest].source[source] = typeOfPulse
		}
	}
}

func handleConjunction(typeOfPulse pulse, c conjunctionModule, source string, cMap map[string]conjunctionModule) (pulse, []string) {
	countHigh := 0
	for _, s := range c.source {
		if s == high {
			countHigh++
		}
	}

	// if all sources sent a high pulse, send low
	if countHigh == len(c.source) {
		updateConjunctionSources(low, source, c.destination, cMap)
		return low, c.destination
	}
	updateConjunctionSources(high, source, c.destination, cMap)
	return high, c.destination
}

type pulseIn struct {
	p    pulse
	name string
	from string
}

func ans1(input []string) int {
	broad, conMap, ffMap := makeMap(input)
	toEvaluateInitial := []pulseIn{}
	// add all initial pulses
	for _, b := range broad.destination {
		toEvaluateInitial = append(toEvaluateInitial, pulseIn{
			p:    low,
			name: b,
		})
	}

	countLow := 0
	countHigh := 0

	for i := 0; i < 1000; i++ {
		toEvaluate := toEvaluateInitial
		countLow++

		for len(toEvaluate) != 0 {
			start := toEvaluate[0]

			if start.p == low {
				countLow++
			}

			if start.p == high {
				countHigh++
			}

			// if node is a flip flop
			if ff, ok := ffMap[start.name]; ok {
				ff, newPulse, dest := handleFlipFlop(start.p, start.name, ff, conMap)
				ffMap[start.name] = ff
				toEvaluate = append(toEvaluate, fetchNewPulseIn(newPulse, dest, start.name)...)
			}

			// if the node is a conjucation
			if con, ok := conMap[start.name]; ok {
				newPulse, dest := handleConjunction(start.p, con, start.name, conMap)
				toEvaluate = append(toEvaluate, fetchNewPulseIn(newPulse, dest, start.name)...)
			}

			toEvaluate = toEvaluate[1:]
		}

	}
	return countHigh * countLow
}

type loop struct {
	start  int
	lenght int
}

func ans2(input []string) int {
	broad, conMap, ffMap := makeMap(input)
	toEvaluateInitial := []pulseIn{}
	// add all initial pulses
	for _, b := range broad.destination {
		toEvaluateInitial = append(toEvaluateInitial, pulseIn{
			p:    low,
			name: b,
		})
	}

	toReadStates := findStatesLeadingToRX(conMap)
	loopRanges := make(map[pulseIn]loop)

	for _, state := range toReadStates {
		count := 0
		i := 0
		for count != 2 {
			i++
			toEvaluate := toEvaluateInitial
			for len(toEvaluate) != 0 {
				start := toEvaluate[0]

				if start == state && count == 0 {
					loopRanges[state] = loop{
						start: i,
					}
					count++
				} else if start == state && count == 1 {
					s := loopRanges[state]
					s.lenght = i - s.start
					loopRanges[state] = s
					count++
					break
				}

				// if node is a flip flop
				if ff, ok := ffMap[start.name]; ok {
					ff, newPulse, dest := handleFlipFlop(start.p, start.name, ff, conMap)
					ffMap[start.name] = ff
					toEvaluate = append(toEvaluate, fetchNewPulseIn(newPulse, dest, start.name)...)
				}

				// if the node is a conjucation
				if con, ok := conMap[start.name]; ok {
					newPulse, dest := handleConjunction(start.p, con, start.name, conMap)
					toEvaluate = append(toEvaluate, fetchNewPulseIn(newPulse, dest, start.name)...)
				}

				toEvaluate = toEvaluate[1:]
			}
		}
	}

	findLCM := []int{}
	for _, l := range loopRanges {
		findLCM = append(findLCM, l.lenght)
	}
	return aoc.LCM(findLCM)

}

func findStatesLeadingToRX(conMap map[string]conjunctionModule) []pulseIn {
	output := []pulseIn{}
	for name, c := range conMap {
		for _, d := range c.destination {
			if d == "rx" {
				for n, _ := range c.source {
					output = append(output, pulseIn{
						p:    high,
						name: name,
						from: n,
					})
				}
			}
		}
	}
	return output
}

func fetchNewPulseIn(p pulse, dest []string, source string) []pulseIn {
	output := []pulseIn{}
	if p == "" {
		return output
	}
	for _, d := range dest {
		o := pulseIn{
			p:    p,
			name: d,
			from: source,
		}
		output = append(output, o)
	}
	return output
}
