package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AlexCrane/tvslogparser/analysis/attack"
	"github.com/AlexCrane/tvslogparser/analysis/pickup"
	"github.com/AlexCrane/tvslogparser/analysis/repair"
	"github.com/AlexCrane/tvslogparser/parser"
)

func main() {
	logPtr := flag.String("f", "", "path to TVS full log file")
	pickupAnalysisPtr := flag.Bool("pickups", false, "print pickup analysis")
	attackAnalysisPtr := flag.Bool("attacks", false, "print attack analysis")
	repairsAnalysisPtr := flag.Bool("repairs", false, "print repair analysis")
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

	if *pickupAnalysisPtr {
		analysis, err := pickup.NewAnalysis(actions)
		if err != nil {
			log.Fatalf("Pickup analysis failed: %s", err)
		}
		fmt.Println(analysis.FormatAsString())
	}
	if *attackAnalysisPtr {
		analysis, err := attack.NewAnalysis(actions)
		if err != nil {
			log.Fatalf("Attack analysis failed: %s", err)
		}
		fmt.Println(analysis.FormatAsString())
	}
	if *repairsAnalysisPtr {
		analysis, err := repair.NewAnalysis(actions)
		if err != nil {
			log.Fatalf("Repair analysis failed: %s", err)
		}
		fmt.Println(analysis.FormatAsString())
	}
}
