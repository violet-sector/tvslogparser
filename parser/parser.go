package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/AlexCrane/tvslogparser/action"
)

func ParseTVSLog(r io.Reader) ([]action.Action, error) {
	actions := make([]action.Action, 0)

	logReader := csv.NewReader(r)

	// First line is headings, just discard
	_, err := logReader.Read()
	if err != nil {
		return nil, fmt.Errorf("Failed to parse TVS log file: %w", err)
	}

	for {
		csvSlice, err := logReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Failed to parse TVS log file: %w", err)
		}

		turnAction, err := action.NewActionFromCSVRecord(csvSlice)
		if err != nil {
			log.Printf("Error: failed to parse CSV record %s; %s", strings.Join(csvSlice, ","), err)
			continue
		}

		actions = append(actions, turnAction)
	}

	return actions, nil
}
