package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AlexCrane/tvslogparser/analysis/pickup"
	"github.com/AlexCrane/tvslogparser/parser"
)

func main() {
	logPtr := flag.String("f", "", "path to TVS full log file")
	pickupAnaysisPtr := flag.Bool("pickups", false, "print pickup analysis")
	flag.Parse()

	if len(*logPtr) == 0 {
		log.Fatalf("You must specify the -f flag, giving a path to a TVS full log file")
	}

	logFile, err := os.Open(*logPtr)
	if err != nil {
		log.Fatalf("Failed to read TVS log file: %s", err)
	}

	actions, err := parser.ParseTVSLog(logFile)
	if err != nil {
		log.Fatalf("Failed to parse TVS log file: %s", err)
	}

	if *pickupAnaysisPtr {
		analysis, err := pickup.NewAnalysis(actions)
		if err != nil {
			log.Fatalf("Pickup analysis failed: %s", err)
		}
		fmt.Println(analysis.FormatAsString())
	}
}
