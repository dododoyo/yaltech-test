package main

import (
	"strconv"
	"strings"
)

func EvaluateWinDrawWin(selection *Selection, score string) error {
	scores := strings.Split(score, "-")
	if len(scores) != 2 {
		selection.IsWinner = false
		selection.Evaluated = true
		return nil
	}

	homeScore, err := strconv.Atoi(scores[0])
	if err != nil {
		selection.IsWinner = false
		selection.Evaluated = true
		return nil
	}

	awayScore, err := strconv.Atoi(scores[1])
	if err != nil {
		selection.IsWinner = false
		selection.Evaluated = true
		return nil
	}

	switch selection.Team {
	case "1": // home
		selection.IsWinner = homeScore > awayScore
	case "2": // away
		selection.IsWinner = awayScore > homeScore
	default:
		selection.IsWinner = false
	}

	selection.Evaluated = true
	return nil
}

func EvaluateDoubleChance(selection *Selection, score string) error {
	scores := strings.Split(score, "-")
	if len(scores) != 2 {
		selection.IsWinner = false
		selection.Evaluated = true
		return nil
	}

	homeScore, err := strconv.Atoi(scores[0])
	if err != nil {
		selection.IsWinner = false
		selection.Evaluated = true
		return nil
	}

	awayScore, err := strconv.Atoi(scores[1])
	if err != nil {
		selection.IsWinner = false
		selection.Evaluated = true
		return nil
	}

	switch selection.Team {
	case "1X": // home win or draw
		selection.IsWinner = homeScore >= awayScore
	case "X2": // away win or draw
		selection.IsWinner = awayScore >= homeScore
	case "12": // either team wins (no draw)
		selection.IsWinner = homeScore != awayScore
	}

	selection.Evaluated = true
	return nil
}

func EvaluateOverUnder(selection *Selection, score string) error {
	scores := strings.Split(score, "-")
	if len(scores) != 2 {
		selection.IsWinner = false
		selection.Evaluated = true
		return nil
	}

	homeScore, err := strconv.Atoi(scores[0])
	if err != nil {
		selection.IsWinner = false
		selection.Evaluated = true
		return nil
	}

	awayScore, err := strconv.Atoi(scores[1])
	if err != nil {
		selection.IsWinner = false
		selection.Evaluated = true
		return nil
	}

	totalScore := float64(homeScore + awayScore)

	switch selection.Team {
	case "Over":
		selection.IsWinner = totalScore > selection.Handicap
	case "Under":
		selection.IsWinner = totalScore < selection.Handicap
	default:
		selection.IsWinner = false
	}

	selection.Evaluated = true
	return nil
}

func EvaluateCorrectScore(selection *Selection, score string) error {
	selection.IsWinner = selection.Team == score
	selection.Evaluated = true
	return nil
}

func ParseOdds(odds string) (float64, error) {
	return strconv.ParseFloat(odds, 64)
}

func ParseHandicap(handicap string) (float64, error) {
	if handicap == "" {
		return 0, nil
	}
	return strconv.ParseFloat(handicap, 64)
} 