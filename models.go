package main

type PreMatchData struct {
	Success int `json:"success"`
	Results []struct {
		FI       string `json:"FI"`
		EventID  string `json:"event_id"`
		Main     struct {
			UpdatedAt string `json:"updated_at"`
			SP        struct {
				ToWinTheMatch struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"to_win_the_match"`
				DoubleChance struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"double_chance"`
				MatchTotals struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Header   string `json:"header"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"match_totals"`
			} `json:"sp"`
		} `json:"main"`
	} `json:"results"`
}

type ResultData struct {
	Success int `json:"success"`
	Results []struct {
		ID         string `json:"id"`
		SportID    string `json:"sport_id"`
		TimeStatus string `json:"time_status"`
		SS         string `json:"ss"`
		Home       struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			ImageID string `json:"image_id"`
			CC      string `json:"cc"`
		} `json:"home"`
		Away struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			ImageID string `json:"image_id"`
			CC      string `json:"cc"`
		} `json:"away"`
	} `json:"results"`
}

type Selection struct {
	Market    string
	Team      string
	Odds      float64
	Handicap  float64
	IsWinner  bool
	Evaluated bool
} 