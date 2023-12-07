package main

import (
	"advent/2022/day01"
	"advent/2022/day02"
	"advent/2022/day03"
	"advent/2022/day04"
	"advent/2022/day05"
	"advent/2022/day06"
	"advent/2022/day07"
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
	{year: "2022", day: "1", runner: day01.Run},
	{year: "2022", day: "2", runner: day02.Run},
	{year: "2022", day: "3", runner: day03.Run},
	{year: "2022", day: "4", runner: day04.Run},
	{year: "2022", day: "5", runner: day05.Run},
	{year: "2022", day: "6", runner: day06.Run},
	{year: "2022", day: "7", runner: day07.Run},
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
