package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int
}

func main() {
	measurements, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer measurements.Close()

	data := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		rawData := scanner.Text()
		semicolonIndex := strings.Index(rawData, ";")
		location := rawData[:semicolonIndex]
		temperature := rawData[semicolonIndex+1:]
		temperatureFloat, _ := strconv.ParseFloat(temperature, 64)

		measurement, ok := data[location]
		if !ok {
			measurement = Measurement{Min: temperatureFloat, Max: temperatureFloat, Sum: temperatureFloat, Count: 1}
		} else {
			measurement.Sum += temperatureFloat
			measurement.Count++
			measurement.Min = math.Min(measurement.Min, temperatureFloat)
			measurement.Max = math.Max(measurement.Max, temperatureFloat)
		}
		data[location] = measurement
	}

	locations := make([]string, 0, len(data))
	for name := range data {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	for _, name := range locations {
		measurement := data[name]
		fmt.Printf("%s=%.1f/%.1f/%.1f, ",
			name,
			measurement.Min,
			measurement.Max,
			measurement.Sum/float64(measurement.Count),
		)
	}

}
