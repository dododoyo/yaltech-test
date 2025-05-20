package main

import "testing"

func TestEvaluateDoubleChance(t *testing.T) {
	tests := []struct {
		name     string
		bet      *Selection
		score    string
		expected bool
	}{
		{
			name: "1X - Home Win",
			bet: &Selection{
				Market: "Double Chance",
				Team:   "1X",
			},
			score:    "2-1",
			expected: true,
		},
		{
			name: "1X - Draw",
			bet: &Selection{
				Market: "Double Chance",
				Team:   "1X",
			},
			score:    "1-1",
			expected: true,
		},
		{
			name: "1X - Away Win (Loss)",
			bet: &Selection{
				Market: "Double Chance",
				Team:   "1X",
			},
			score:    "1-2",
			expected: false,
		},
		{
			name: "X2 - Away Win",
			bet: &Selection{
				Market: "Double Chance",
				Team:   "X2",
			},
			score:    "1-2",
			expected: true,
		},
		{
			name: "X2 - Draw",
			bet: &Selection{
				Market: "Double Chance",
				Team:   "X2",
			},
			score:    "2-2",
			expected: true,
		},
		{
			name: "X2 - Home Win (Loss)",
			bet: &Selection{
				Market: "Double Chance",
				Team:   "X2",
			},
			score:    "2-1",
			expected: false,
		},
		{
			name: "12 - Home Win",
			bet: &Selection{
				Market: "Double Chance",
				Team:   "12",
			},
			score:    "2-1",
			expected: true,
		},
		{
			name: "12 - Away Win",
			bet: &Selection{
				Market: "Double Chance",
				Team:   "12",
			},
			score:    "1-2",
			expected: true,
		},
		{
			name: "12 - Draw (Loss)",
			bet: &Selection{
				Market: "Double Chance",
				Team:   "12",
			},
			score:    "1-1",
			expected: false,
		},
		{
			name: "Invalid Score Format",
			bet: &Selection{
				Market: "Double Chance",
				Team:   "1X",
			},
			score:    "invalid",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := EvaluateDoubleChance(tt.bet, tt.score)
			if err != nil {
				t.Errorf("EvaluateDoubleChance() error = %v", err)
				return
			}

			if tt.bet.IsWinner != tt.expected {
				t.Errorf("EvaluateDoubleChance() got = %v, want %v", tt.bet.IsWinner, tt.expected)
			}

			if !tt.bet.Evaluated {
				t.Error("EvaluateDoubleChance() selection was not marked as evaluated")
			}
		})
	}
}

func TestEvaluateWinDrawWin(t *testing.T) {
	tests := []struct {
		name     string
		bet      *Selection
		score    string
		expected bool
	}{
		{
			name: "Home Win - Correct",
			bet: &Selection{
				Market: "Win/Draw/Win",
				Team:   "1",
			},
			score:    "2-1",
			expected: true,
		},
		{
			name: "Home Win - Incorrect",
			bet: &Selection{
				Market: "Win/Draw/Win",
				Team:   "1",
			},
			score:    "1-2",
			expected: false,
		},
		{
			name: "Away Win - Correct",
			bet: &Selection{
				Market: "Win/Draw/Win",
				Team:   "2",
			},
			score:    "1-2",
			expected: true,
		},
		{
			name: "Away Win - Incorrect",
			bet: &Selection{
				Market: "Win/Draw/Win",
				Team:   "2",
			},
			score:    "2-1",
			expected: false,
		},
		{
			name: "Invalid Score Format",
			bet: &Selection{
				Market: "Win/Draw/Win",
				Team:   "1",
			},
			score:    "invalid",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := EvaluateWinDrawWin(tt.bet, tt.score)
			if err != nil {
				t.Errorf("EvaluateWinDrawWin() error = %v", err)
				return
			}

			if tt.bet.IsWinner != tt.expected {
				t.Errorf("EvaluateWinDrawWin() got = %v, want %v", tt.bet.IsWinner, tt.expected)
			}

			if !tt.bet.Evaluated {
				t.Error("EvaluateWinDrawWin() selection was not marked as evaluated")
			}
		})
	}
}

func TestEvaluateOverUnder(t *testing.T) {
	tests := []struct {
		name     string
		bet      *Selection
		score    string
		handicap float64
		expected bool
	}{
		{
			name: "Over - Correct",
			bet: &Selection{
				Market:   "Over/Under",
				Team:     "Over",
				Handicap: 2.5,
			},
			score:    "2-1",
			expected: true,
		},
		{
			name: "Over - Incorrect",
			bet: &Selection{
				Market:   "Over/Under",
				Team:     "Over",
				Handicap: 3.5,
			},
			score:    "2-1",
			expected: false,
		},
		{
			name: "Under - Correct",
			bet: &Selection{
				Market:   "Over/Under",
				Team:     "Under",
				Handicap: 3.5,
			},
			score:    "1-1",
			expected: true,
		},
		{
			name: "Under - Incorrect",
			bet: &Selection{
				Market:   "Over/Under",
				Team:     "Under",
				Handicap: 2.5,
			},
			score:    "2-1",
			expected: false,
		},
		{
			name: "Invalid Score Format",
			bet: &Selection{
				Market:   "Over/Under",
				Team:     "Over",
				Handicap: 2.5,
			},
			score:    "invalid",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := EvaluateOverUnder(tt.bet, tt.score)
			if err != nil {
				t.Errorf("EvaluateOverUnder() error = %v", err)
				return
			}

			if tt.bet.IsWinner != tt.expected {
				t.Errorf("EvaluateOverUnder() got = %v, want %v", tt.bet.IsWinner, tt.expected)
			}

			if !tt.bet.Evaluated {
				t.Error("EvaluateOverUnder() selection was not marked as evaluated")
			}
		})
	}
}

func TestEvaluateCorrectScore(t *testing.T) {
	tests := []struct {
		name     string
		bet      *Selection
		score    string
		expected bool
	}{
		{
			name: "Correct Score - Match",
			bet: &Selection{
				Market: "Correct Score",
				Team:   "2-1",
			},
			score:    "2-1",
			expected: true,
		},
		{
			name: "Correct Score - No Match",
			bet: &Selection{
				Market: "Correct Score",
				Team:   "2-1",
			},
			score:    "1-1",
			expected: false,
		},
		{
			name: "Invalid Score Format",
			bet: &Selection{
				Market: "Correct Score",
				Team:   "2-1",
			},
			score:    "invalid",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := EvaluateCorrectScore(tt.bet, tt.score)
			if err != nil {
				t.Errorf("EvaluateCorrectScore() error = %v", err)
				return
			}

			if tt.bet.IsWinner != tt.expected {
				t.Errorf("EvaluateCorrectScore() got = %v, want %v", tt.bet.IsWinner, tt.expected)
			}

			if !tt.bet.Evaluated {
				t.Error("EvaluateCorrectScore() selection was not marked as evaluated")
			}
		})
	}
}

func TestParseOdds(t *testing.T) {
	tests := []struct {
		name    string
		odds    string
		want    float64
		wantErr bool
	}{
		{
			name:    "Valid odds",
			odds:    "2.50",
			want:    2.50,
			wantErr: false,
		},
		{
			name:    "Invalid odds",
			odds:    "invalid",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseOdds(tt.odds)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseOdds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseOdds() = %v, want %v", got, tt.want)
			}
		})
	}
} 