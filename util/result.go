package util

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Problem struct {
	Year   string
	Day    string
	Runner func([]byte) (Result, error)
	Input  []byte
}

type Result struct {
	Year      string
	Day       string
	Part1     any
	Part2     any
	StartTime time.Time
	ParseTime time.Time
	EndTime   time.Time
}

func PrintResults(results ...Result) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Year", "Day", "Part 1", "Part 2", "Parse", "Execute", "Total"})
	for _, result := range results {
		err := table.Append(
			result.Year,
			result.Day,
			fmt.Sprintf("%v", result.Part1),
			fmt.Sprintf("%v", result.Part2),
			result.ParseTime.Sub(result.StartTime).String(),
			result.EndTime.Sub(result.ParseTime).String(),
			result.EndTime.Sub(result.StartTime).String(),
		)
		if err != nil {
			log.Panic(err)
		}
	}
	err := table.Render()
	if err != nil {
		log.Panic(err)
	}
}
