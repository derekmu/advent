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
	"context"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

type problem struct {
	year   string
	day    string
	runner func([]byte) error
	input  []byte
}

var problems = []problem{
	// 2022
	{year: "2022", day: "1", runner: day2201.Run, input: day2201.Input},
	{year: "2022", day: "2", runner: day2202.Run, input: day2202.Input},
	{year: "2022", day: "3", runner: day2203.Run, input: day2203.Input},
	{year: "2022", day: "4", runner: day2204.Run, input: day2204.Input},
	{year: "2022", day: "5", runner: day2205.Run, input: day2205.Input},
	{year: "2022", day: "6", runner: day2206.Run, input: day2206.Input},
	{year: "2022", day: "7", runner: day2207.Run, input: day2207.Input},
	{year: "2022", day: "8", runner: day2208.Run, input: day2208.Input},
	{year: "2022", day: "9", runner: day2209.Run, input: day2209.Input},
	{year: "2022", day: "10", runner: day2210.Run, input: day2210.Input},
	{year: "2022", day: "11", runner: day2211.Run, input: day2211.Input},
	{year: "2022", day: "12", runner: day2212.Run, input: day2212.Input},
	{year: "2022", day: "13", runner: day2213.Run, input: day2213.Input},
	// 2023
	{year: "2023", day: "1", runner: day2301.Run, input: day2301.Input},
	{year: "2023", day: "2", runner: day2302.Run, input: day2302.Input},
	{year: "2023", day: "3", runner: day2303.Run, input: day2303.Input},
	{year: "2023", day: "4", runner: day2304.Run, input: day2304.Input},
	{year: "2023", day: "5", runner: day2305.Run, input: day2305.Input},
	{year: "2023", day: "6", runner: day2306.Run, input: day2306.Input},
	{year: "2023", day: "7", runner: day2307.Run, input: day2307.Input},
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
				return p.runner(p.input)
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
