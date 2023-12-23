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
	day2214 "advent/2022/day14"
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
	day2318 "advent/2023/day18"
	day2319 "advent/2023/day19"
	day2320 "advent/2023/day20"
	day2321 "advent/2023/day21"
	day2322 "advent/2023/day22"
	day2323 "advent/2023/day23"
	"advent/util"
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var problems = []util.Problem{
	// 2022
	day2201.Problem,
	day2202.Problem,
	day2203.Problem,
	day2204.Problem,
	day2205.Problem,
	day2206.Problem,
	day2207.Problem,
	day2208.Problem,
	day2209.Problem,
	day2210.Problem,
	day2211.Problem,
	day2212.Problem,
	day2213.Problem,
	day2214.Problem,
	// 2023
	day2301.Problem,
	day2302.Problem,
	day2303.Problem,
	day2304.Problem,
	day2305.Problem,
	day2306.Problem,
	day2307.Problem,
	day2308.Problem,
	day2309.Problem,
	day2310.Problem,
	day2311.Problem,
	day2312.Problem,
	day2313.Problem,
	day2314.Problem,
	day2315.Problem,
	day2316.Problem,
	day2317.Problem,
	day2318.Problem,
	day2319.Problem,
	day2320.Problem,
	day2321.Problem,
	day2322.Problem,
	day2323.Problem,
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
