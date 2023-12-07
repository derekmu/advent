package main

import (
	day2201 "advent/2022/day01"
	day2202 "advent/2022/day02"
	day2203 "advent/2022/day03"
	day2204 "advent/2022/day04"
	day2205 "advent/2022/day05"
	day2206 "advent/2022/day06"
	day2207 "advent/2022/day07"
	day2208 "advent/2022/day08"
	day2209 "advent/2022/day09"
	day2210 "advent/2022/day10"
	day2211 "advent/2022/day11"
	day2212 "advent/2022/day12"
	day2213 "advent/2022/day13"
	day2301 "advent/2023/day01"
	day2302 "advent/2023/day02"
	day2303 "advent/2023/day03"
	day2304 "advent/2023/day04"
	day2305 "advent/2023/day05"
	day2306 "advent/2023/day06"
	day2307 "advent/2023/day07"
	"advent/util"
	"context"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

type problem struct {
	year   string
	day    string
	runner func([]byte) error
}

var problems = []problem{
	// 2022
	{year: "2022", day: "1", runner: day2201.Run},
	{year: "2022", day: "2", runner: day2202.Run},
	{year: "2022", day: "3", runner: day2203.Run},
	{year: "2022", day: "4", runner: day2204.Run},
	{year: "2022", day: "5", runner: day2205.Run},
	{year: "2022", day: "6", runner: day2206.Run},
	{year: "2022", day: "7", runner: day2207.Run},
	{year: "2022", day: "8", runner: day2208.Run},
	{year: "2022", day: "9", runner: day2209.Run},
	{year: "2022", day: "10", runner: day2210.Run},
	{year: "2022", day: "11", runner: day2211.Run},
	{year: "2022", day: "12", runner: day2212.Run},
	{year: "2022", day: "13", runner: day2213.Run},
	// 2023
	{year: "2023", day: "1", runner: day2301.Run},
	{year: "2023", day: "2", runner: day2302.Run},
	{year: "2023", day: "3", runner: day2303.Run},
	{year: "2023", day: "4", runner: day2304.Run},
	{year: "2023", day: "5", runner: day2305.Run},
	{year: "2023", day: "6", runner: day2306.Run},
	{year: "2023", day: "7", runner: day2307.Run},
}

func main() {
	var yearCommands []*cli.Command
	yearMap := make(map[string]*cli.Command, 2)
	for _, pLoop := range problems {
		p := pLoop
		yearCommand, ok := yearMap[p.year]
		if !ok {
			yearCommand = &cli.Command{Name: p.year}
			yearMap[p.year] = yearCommand
			yearCommands = append(yearCommands, yearCommand)
		}
		yearCommand.Commands = append(yearCommand.Commands, &cli.Command{
			Name: p.day,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				var day string
				if len(p.day) == 1 {
					day = "day0" + p.day
				} else {
					day = "day" + p.day
				}
				input := util.ReadDayInput(p.year, day)
				return p.runner(input)
			},
		})
	}

	cmd := &cli.Command{
		Name:            "advent",
		Usage:           "Run Advent of Code problems",
		HideHelpCommand: true,
		Commands:        yearCommands,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
