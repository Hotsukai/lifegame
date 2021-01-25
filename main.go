package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Hotsukai/lifegame/components"
)

func parseCommandLine() (int, int, float64, int, bool, bool) {
	args := os.Args
	const usage = `
positional arguments:
	height      Field height
	width       Field width
	init_rate   Percentage of surviving cells of the first generation
	interval    Time to evolve to the next generatio
optional arguments:
	-r, --routine use Go Routine
	-p, --print show field
	-h, --help  show this help message and exit`
	isPrint := true
	useRoutine := false
	for i, arg := range args {
		if arg == "-h" || arg == "--help" {
			fmt.Println(usage)
			os.Exit(0)
		}
		if arg == "-r" || arg == "--routine" {
			useRoutine = true
		}
		if arg == "-p" || arg == "--print" {
			if len(args) > i+1 {
				isPrint, _ = strconv.ParseBool(args[i+1])
			} else {
				fmt.Println(usage)
				os.Exit(0)
			}
		}
	}
	if len(args) < 5 {
		fmt.Println("Error : Arguments in Valid")
		fmt.Println(usage)
		os.Exit(1)
	}

	height, _ := strconv.Atoi(args[1])
	width, _ := strconv.Atoi(args[2])
	initRate, _ := strconv.ParseFloat(args[3], 64)
	interval, _ := strconv.Atoi(args[4])
	return height, width, initRate, interval, useRoutine, isPrint
}

func main() {
	height, width, initRate, interval, useRoutine, isPrint := parseCommandLine()
	lifegame := components.NewLifeGame(height, width, initRate, interval, useRoutine, isPrint)
	lifegame.MainLoop()
}
