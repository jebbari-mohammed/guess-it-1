package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	content := bufio.NewScanner(os.Stdin)

	Num := []int{}

	for content.Scan() {

		nb, err := strconv.Atoi(content.Text())
		if err != nil {
			continue
		}

		Num = append(Num, nb)

		if len(Num) > 1 {
			min, max := Guess(Num)
			fmt.Printf("%d %d\n", int(math.Round(min)), int(math.Round(max)))
		}

	}

	if err := content.Err(); err != nil {
		return
	}
}

func Average(nbs []int) float64 {
	if len(nbs) == 0 {
		return 0
	}

	avg := 0.0
	for i := 0; i < len(nbs); i++ {
		avg += float64(nbs[i])
	}

	return avg / float64(len(nbs))
}

func Median(nbs []int) float64 {
	if len(nbs) == 0 {
		return 0
	}

	sort.Ints(nbs)
	if len(nbs)%2 == 0 {
		return (float64(nbs[len(nbs)/2-1]) + float64(nbs[len(nbs)/2])) / 2
	} else {
		return float64(nbs[len(nbs)/2])
	}
}

func Variance(nbs []int) float64 {
	if len(nbs) == 0 {
		return 0
	}

	variance := 0.0
	for i := 0; i < len(nbs); i++ {
		pow := float64(nbs[i]) - Average(nbs)
		variance += pow * pow
	}

	return variance / float64(len(nbs))
}

func StandardDeviation(nbs []int) float64 {
	return math.Sqrt(Variance(nbs))
}

func Guess(numbers []int) (float64, float64) {
	start := len(numbers) - 4
	if start < 0 {
		start = 0
	}
	nb := numbers[start:]
	std := StandardDeviation(nb)
	Average := Average(nb)
	min := Average - (std * 2)
	max := Average + (std * 2)
	return min, max
}
