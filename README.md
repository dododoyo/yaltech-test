# Sports Betting Odds Evaluation System

This Go-based system evaluates sports betting selections by comparing pre-match odds data against actual match results. It supports multiple sports (currently cricket and volleyball) and provides accurate evaluation of various betting markets.

## Features

### Supported Markets

1. **Win/Draw/Win (1X2)**
   - Home team win (1)
   - Away team win (2)
   - Draw (X) - where applicable

2. **Over/Under**
   - Total score over a specified handicap
   - Total score under a specified handicap
   - Supports decimal handicaps (e.g., 2.5, 3.5)

3. **Correct Score**
   - Exact match score prediction
   - Format: "home-away" (e.g., "2-1")

4. **Double Chance**
   - Home Win or Draw (1X)
   - Away Win or Draw (X2)
   - Either Team to Win (12)

## Requirements

- Go 1.16 or higher
- JSON input files (pre-match data and results)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/dododoyo/yaltech-test
cd yaltech-test
```

2. Initialize the Go module:
```bash
go mod init yaltech-test
```

3. Build the program:
```bash
go build
```

## Usage

### Input Files

The system requires two JSON files for each sport:

1. **Pre-match Data** (e.g., `cricket_prematch.json`, `volleyball_prematch.json`)
   - Contains available markets
   - Odds for each selection
   - Event information

2. **Result Data** (e.g., `cricket_result.json`, `volleyball_result.json`)
   - Final scores
   - Match outcome
   - Additional statistics

### Running the Program

1. Place your JSON files in the project directory
2. Run the program:
```bash
./yaltech-test
```

### Output Format

The program outputs evaluation results in the following format:
```
Betting Selections Results:
===========================
Market: Win/Draw/Win
Selection: 1
Odds: 2.50
Result: LOST
---------------------------
Market: Win/Draw/Win
Selection: 2
Odds: 1.53
Result: WON
---------------------------
```

## Market Evaluation Rules

### Win/Draw/Win (1X2)
- Home win (1): Wins if home team score > away team score
- Away win (2): Wins if away team score > home team score

### Over/Under
- Over: Wins if total score > handicap
- Under: Wins if total score < handicap

### Correct Score
- Wins only on exact score match
- Format must be "home-away" (e.g., "2-1")

### Double Chance
- 1X: Wins if home team wins or draws
- X2: Wins if away team wins or draws
- 12: Wins if either team wins (no draw)

## Error Handling

The system handles various error cases:
- Invalid score formats
- Missing data
- Invalid market types
- Parse errors in odds/handicaps

## Testing

Run the test suite:
```bash
go test -v
```

The test suite covers:
- All market types
- Edge cases
- Invalid inputs
- Score parsing
- Odds calculation

## Project Structure

- `main.go`: Program entry point and core logic
- `models.go`: Data structures for JSON parsing
- `odds.go`: Market evaluation functions
- `odds_test.go`: Comprehensive test suite
