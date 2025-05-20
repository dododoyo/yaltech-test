package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	prematchData, err := readPreMatchData("cricket_prematch.json")
	if err != nil {
		log.Fatalf("Error reading prematch data: %v", err)
	}

	resultData, err := readResultData("cricket_result.json")
	if err != nil {
		log.Fatalf("Error reading result data: %v", err)
	}

	selections := createSampleSelections(prematchData)

	evaluateSelections(selections, resultData)

	printResults(selections)
}

func readPreMatchData(filename string) (*PreMatchData, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var prematch PreMatchData
	err = json.Unmarshal(data, &prematch)
	if err != nil {
		return nil, err
	}

	return &prematch, nil
}

func readResultData(filename string) (*ResultData, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var result ResultData
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func createSampleSelections(prematch *PreMatchData) []*Selection {
	var selections []*Selection

	if len(prematch.Results) == 0 {
		return selections
	}

	match := prematch.Results[0]

	for _, odd := range match.Main.SP.ToWinTheMatch.Odds {
		odds, err := ParseOdds(odd.Odds)
		if err != nil {
			continue
		}

		handicap, err := ParseHandicap(odd.Handicap)
		if err != nil {
			continue
		}

		selections = append(selections, &Selection{
			Market:   "Win/Draw/Win",
			Team:     odd.Name,
			Odds:     odds,
			Handicap: handicap,
		})
	}

	for _, odd := range match.Main.SP.DoubleChance.Odds {
		odds, err := ParseOdds(odd.Odds)
		if err != nil {
			continue
		}

		handicap, err := ParseHandicap(odd.Handicap)
		if err != nil {
			continue
		}

		selections = append(selections, &Selection{
			Market:   "Double Chance",
			Team:     odd.Name,
			Odds:     odds,
			Handicap: handicap,
		})
	}

	for _, odd := range match.Main.SP.MatchTotals.Odds {
		odds, err := ParseOdds(odd.Odds)
		if err != nil {
			continue
		}

		handicap, err := ParseHandicap(odd.Handicap)
		if err != nil {
			continue
		}

		team := "Over"
		if strings.ToLower(odd.Header) == "under" {
			team = "Under"
		}

		selections = append(selections, &Selection{
			Market:   "Over/Under",
			Team:     team,
			Odds:     odds,
			Handicap: handicap,
		})
	}

	return selections
}

func evaluateSelections(selections []*Selection, result *ResultData) {
	if len(result.Results) == 0 {
		return
	}

	match := result.Results[0]

	for _, selection := range selections {
		switch selection.Market {
		case "Win/Draw/Win":
			EvaluateWinDrawWin(selection, match.SS)
		case "Double Chance":
			EvaluateDoubleChance(selection, match.SS)
		case "Over/Under":
			EvaluateOverUnder(selection, match.SS)
		case "Correct Score":
			EvaluateCorrectScore(selection, match.SS)
		}
	}
}

func printResults(selections []*Selection) {
	fmt.Println("\nBetting Selections Results:")
	fmt.Println("===========================")
	
	for _, selection := range selections {
		result := "LOST"
		if selection.IsWinner {
			result = "WON"
		}

		handicapStr := ""
		if selection.Handicap != 0 {
			handicapStr = fmt.Sprintf(" (%.1f)", selection.Handicap)
		}

		fmt.Printf("Market: %s\n", selection.Market)
		fmt.Printf("Selection: %s%s\n", selection.Team, handicapStr)
		fmt.Printf("Odds: %.2f\n", selection.Odds)
		fmt.Printf("Result: %s\n", result)
		fmt.Println("---------------------------")
	}
} 