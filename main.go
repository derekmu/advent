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
	day2308 "advent/2023/day08"
	day2309 "advent/2023/day09"
	day2310 "advent/2023/day10"
	day2311 "advent/2023/day11"
	day2312 "advent/2023/day12"
	day2313 "advent/2023/day13"
	day2314 "advent/2023/day14"
	day2315 "advent/2023/day15"
	day2316 "advent/2023/day16"
	day2317 "advent/2023/day17"
	"advent/util"
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var problems = []util.Problem{
	// 2022
	day2201.Problem,
	{Year: "2022", Day: "02", Runner: day2202.Run, Input: day2202.Input},
	{Year: "2022", Day: "03", Runner: day2203.Run, Input: day2203.Input},
	{Year: "2022", Day: "04", Runner: day2204.Run, Input: day2204.Input},
	{Year: "2022", Day: "05", Runner: day2205.Run, Input: day2205.Input},
	{Year: "2022", Day: "06", Runner: day2206.Run, Input: day2206.Input},
	{Year: "2022", Day: "07", Runner: day2207.Run, Input: day2207.Input},
	{Year: "2022", Day: "08", Runner: day2208.Run, Input: day2208.Input},
	{Year: "2022", Day: "09", Runner: day2209.Run, Input: day2209.Input},
	{Year: "2022", Day: "10", Runner: day2210.Run, Input: day2210.Input},
	{Year: "2022", Day: "11", Runner: day2211.Run, Input: day2211.Input},
	{Year: "2022", Day: "12", Runner: day2212.Run, Input: day2212.Input},
	{Year: "2022", Day: "13", Runner: day2213.Run, Input: day2213.Input},
	// 2023
	{Year: "2023", Day: "01", Runner: day2301.Run, Input: day2301.Input},
	{Year: "2023", Day: "02", Runner: day2302.Run, Input: day2302.Input},
	{Year: "2023", Day: "03", Runner: day2303.Run, Input: day2303.Input},
	{Year: "2023", Day: "04", Runner: day2304.Run, Input: day2304.Input},
	{Year: "2023", Day: "05", Runner: day2305.Run, Input: day2305.Input},
	{Year: "2023", Day: "06", Runner: day2306.Run, Input: day2306.Input},
	{Year: "2023", Day: "07", Runner: day2307.Run, Input: day2307.Input},
	{Year: "2023", Day: "08", Runner: day2308.Run, Input: day2308.Input},
	{Year: "2023", Day: "09", Runner: day2309.Run, Input: day2309.Input},
	{Year: "2023", Day: "10", Runner: day2310.Run, Input: day2310.Input},
	{Year: "2023", Day: "11", Runner: day2311.Run, Input: day2311.Input},
	{Year: "2023", Day: "12", Runner: day2312.Run, Input: day2312.Input},
	{Year: "2023", Day: "13", Runner: day2313.Run, Input: day2313.Input},
	{Year: "2023", Day: "14", Runner: day2314.Run, Input: day2314.Input},
	{Year: "2023", Day: "15", Runner: day2315.Run, Input: day2315.Input},
	{Year: "2023", Day: "16", Runner: day2316.Run, Input: day2316.Input},
	{Year: "2023", Day: "17", Runner: day2317.Run, Input: day2317.Input},
}

func main() {
	var yearCommands []*cli.Command
	yearMap := make(map[string]*cli.Command, 2)
	for _, p := range problems {
		yearCommand, ok := yearMap[p.Year]
		if !ok {
			yearCommand = makeYearCommand(p.Year)
			yearMap[p.Year] = yearCommand
			yearCommands = append(yearCommands, yearCommand)
		}
		yearCommand.Commands = append(yearCommand.Commands, makeProblemCommand(p))
	}

	cmd := &cli.Command{
		Name:            "advent",
		Usage:           "Run Advent of Code problems",
		HideHelpCommand: true,
		Commands:        yearCommands,
		Action: func(_ context.Context, _ *cli.Command) error {
			var results []*util.Result
			for _, p := range problems {
				result, err := p.Runner(p.Input)
				if err != nil {
					return err
				}
				result.Year = p.Year
				result.Day = p.Day
				results = append(results, result)
			}
			util.PrintResults(results...)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func makeYearCommand(year string) *cli.Command {
	return &cli.Command{
		Name: year,
		Action: func(_ context.Context, _ *cli.Command) error {
			var results []*util.Result
			for _, p := range problems {
				if p.Year == year {
					result, err := p.Runner(p.Input)
					if err != nil {
						return err
					}
					result.Year = p.Year
					result.Day = p.Day
					results = append(results, result)
				}
			}
			util.PrintResults(results...)
			return nil
		},
	}
}

func makeProblemCommand(p util.Problem) *cli.Command {
	return &cli.Command{
		Name: p.Day,
		Action: func(_ context.Context, _ *cli.Command) error {
			result, err := p.Runner(p.Input)
			if err != nil {
				return err
			}
			result.Year = p.Year
			result.Day = p.Day
			util.PrintResults(result)
			return nil
		},
	}
}
