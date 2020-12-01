package parser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/AlexCrane/tvslogparser/action"
	"github.com/AlexCrane/tvslogparser/common"
)

type parserState struct {
	myPlayer *common.Player
}

func (ps *parserState) NewActionFromCSVRecord(record []string) (action.Action, error) {
	if len(record) != 6 {
		return nil, fmt.Errorf("unexpected record len %d", len(record))
	}

	actionRecord := record[3]

	// Start off wth logs we'll ignore
	if strings.HasPrefix(actionRecord, "Set autopilot to") {
		return nil, nil
	}
	if strings.HasPrefix(actionRecord, "Unset autopilot") {
		return nil, nil
	}
	if strings.HasPrefix(actionRecord, "Entered carrier") {
		return nil, nil
	}
	for _, s := range []string{
		"Taken sector damage",
		"Earned domination bonus",
		"Hit by atmospheric anomaly",
		"Exited carrier",
		"Hit by base explosion",
	} {

		if actionRecord == s {
			return nil, nil
		}
	}

	// Now logs we're interested in
	if strings.HasPrefix(actionRecord, "Pickup") {
		return action.PickupFromCSVRecord(record)
	}
	if actionRecord == "Added to the game" {
		return action.AddedFromCSVRecord(record)
	}
	if strings.HasPrefix(actionRecord, "Selected to pilot a") {
		return action.SelectShipFromCSVRecord(record)
	}
	if strings.HasPrefix(actionRecord, "Hypered to ") {
		return action.HyperFromCSVRecord(record)
	}
	if strings.Contains(actionRecord, "attacked") {
		return action.AttackFromCSVRecord(record, ps.myPlayer)
	}
	if strings.HasPrefix(actionRecord, "Achieved level") || strings.HasPrefix(actionRecord, "Awarded cruiser") {
		return action.LevelUpFromCSVRecord(record)
	}
	if strings.HasSuffix(actionRecord, "scrap") {
		return action.ScrapFromCSVRecord(record)
	}
	if strings.Contains(actionRecord, "repaired") {
		if actionRecord == "Self repaired" {
			return action.SelfRepFromCSVRecord(record)
		}
		return action.RepairFromCSVRecord(record, ps.myPlayer)
	}

	return nil, errors.New("not implemented")
}

func ParseTVSLog(r io.Reader) ([]action.Action, error) {
	actions := make([]action.Action, 0)

	logReader := csv.NewReader(r)
	parserState := &parserState{}

	// First line is headings, just discard
	_, err := logReader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to parse TVS log file: %w", err)
	}

	for {
		csvSlice, err := logReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to parse TVS log file: %w", err)
		}

		turnAction, err := parserState.NewActionFromCSVRecord(csvSlice)
		if err != nil {
			log.Printf("Error: failed to parse CSV record %s; %s", strings.Join(csvSlice, ","), err)
			continue
		}
		if turnAction == nil {
			// considered inconsequential (i.e. sector damage)
			continue
		}

		if turnAction.ActionType() == action.ActionTypeAdded {
			asAdded := turnAction.(*action.Added)
			parserState.myPlayer = common.NewPlayer(asAdded.PilotName)
		} else if turnAction.ActionType() == action.ActionTypeSelectShip {
			asSS := turnAction.(*action.SelectShip)
			parserState.myPlayer.SelectShip(asSS.ShipType)
		} else if turnAction.ActionType() == action.ActionTypeLevelUp {
			asLevelUp := turnAction.(*action.LevelUp)
			parserState.myPlayer.GetShip().LevelUp(asLevelUp.Level)
		}

		actions = append(actions, turnAction)
	}

	return actions, nil
}
