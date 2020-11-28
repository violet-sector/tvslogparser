package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AlexCrane/tvslogparser/action"
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

	for _, a := range actions {
		fmt.Println(a)
	}
	fmt.Println()

	if *pickupAnaysisPtr {
		pickupsByType := make(map[action.PickupType]int)
		totalPickups := 0

		for _, a := range actions {
			if a.ActionType() == action.ActionTypePickup {
				pickup := a.(*action.Pickup)

				val, ok := pickupsByType[pickup.Type]
				if !ok {
					pickupsByType[pickup.Type] = 1
				} else {
					pickupsByType[pickup.Type] = val + 1
				}

				totalPickups++
			}
		}

		fmt.Println("Total pickups:", totalPickups)
		for _, k := range []action.PickupType{
			action.PTNone,
			action.PTFreeAttack,
			action.PTTripleAttack,
			action.PTExplosive50,
			action.PTExplosive250,
			action.PTExplosive450,
			action.PTRepair150,
			action.PTRepair200,
			action.PTRepair300,
			action.PTPoints100,
			action.PTPoints150,
			action.PTPoints250,
			action.PTPoints300,
			action.PTInvulnerability,
			action.PTHalfMaxRepair,
			action.PTGateIntel,
			action.PTDecoy,
			action.PTRandomHyperGate,
		} {
			v := pickupsByType[k]
			fmt.Printf("%-24s : %-4d (%03f%%)\n", k, v, float64(v)/float64(totalPickups)*100)
		}
	}
}
