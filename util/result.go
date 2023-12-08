package util

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"time"
)

type Result struct {
	Year      string
	Day       string
	Part1     any
	Part2     any
	StartTime time.Time
	ParseTime time.Time
	EndTime   time.Time
}

func PrintResults(results ...*Result) {
	data := make([][]string, 0, len(results))
	for _, result := range results {
		data = append(data, []string{
			result.Year,
			result.Day,
			fmt.Sprintf("%v", result.Part1),
			fmt.Sprintf("%v", result.Part2),
			result.ParseTime.Sub(result.StartTime).String(),
			result.EndTime.Sub(result.ParseTime).String(),
			result.EndTime.Sub(result.StartTime).String(),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Year", "Day", "Part 1", "Part 2", "Parse", "Execute", "Total"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(data)
	table.Render()
}
