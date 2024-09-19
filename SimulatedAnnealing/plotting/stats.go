package plotting

import (
	"bufio"
	"gonum.org/v1/plot/plotter"
	"log"
	"os"
	"strconv"
	"strings"
)

type stats struct {
	iteration           int
	temperature         float64
	bestEnergy          float64
	nBadChoicesAccepted int
}

func getStats() []*stats {
	file, err := os.Open("stats.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(file)
	result := make([]*stats, 0)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		tokens := strings.Split(line, " ")

		iter, _ := strconv.Atoi(tokens[0])
		temperature, _ := strconv.ParseFloat(tokens[1], 64)
		bestEnergy, _ := strconv.ParseFloat(tokens[2], 64)
		badChoices, _ := strconv.Atoi(tokens[3])
		result = append(result, &stats{iter, temperature, bestEnergy, badChoices})
	}
	return result
}

func getPoints(stats []*stats) (*plotter.XYs, *plotter.XYs, *plotter.XYs) {
	tempPlt := make(plotter.XYs, len(stats))
	energyPlt := make(plotter.XYs, len(stats))
	choicesPlt := make(plotter.XYs, len(stats))

	for _, stat := range stats {
		tempPlt[stat.iteration].X = float64(stat.iteration)
		tempPlt[stat.iteration].Y = stat.temperature

		energyPlt[stat.iteration].X = float64(stat.iteration)
		energyPlt[stat.iteration].Y = stat.bestEnergy

		choicesPlt[stat.iteration].X = float64(stat.iteration)
		choicesPlt[stat.iteration].Y = float64(stat.nBadChoicesAccepted)
	}

	return &tempPlt, &energyPlt, &choicesPlt
}
