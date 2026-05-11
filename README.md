# CSV Log Analyzer

A command-line tool written in Go that reads structured CSV log files, filters entries by severity and date range, and prints a summary report.

Built as part of a Go learning roadmap to practice slice operations, file I/O, error handling, and idiomatic Go patterns.

---

## What it does

- Reads a `.csv` log file into memory
- Parses each row into a typed `LogEntry` struct with a real `time.Time` timestamp
- Filters entries by severity level (`INFO`, `WARN`, `ERROR`) and an optional date range
- Outputs a summary showing total matched entries and a breakdown by severity

---

## Project structure

```
csv-log-analyzer/
├── main.go       # All logic: parseFlags, readCSV, parseRow, filterLogs, summarize, printSummary
├── main_test.go  # Table-driven tests for core functions
├── logs.csv      # Sample log file to try the tool immediately
├── Makefile      # Shortcuts for run, build, test
└── go.mod        # Module definition
```

---

## Getting started

**Prerequisites:** Go 1.21 or later. Check with:

```bash
go version
```

**Clone and run:**

```bash
git clone https://github.com/yourusername/csv-log-analyzer.git
cd csv-log-analyzer
go run main.go --file logs.csv --severity ERROR --from 2024-01-01 --to 2024-01-03
```

Or using Make:

```bash
make run
```

---

## Usage

```
go run main.go [flags]

Flags:
  --file       Path to the CSV log file        (default: logs.csv)
  --severity   Filter by severity level        (INFO | WARN | ERROR)
  --from       Start date in YYYY-MM-DD format (inclusive)
  --to         End date in YYYY-MM-DD format   (inclusive)
```

### Examples

```bash
# All errors across all days
go run main.go --file logs.csv --severity ERROR --from 2024-01-01 --to 2024-01-03

# Warnings on a single day
go run main.go --file logs.csv --severity WARN --from 2024-01-02 --to 2024-01-02

# All info logs in January
go run main.go --file logs.csv --severity INFO --from 2024-01-01 --to 2024-01-31
```

### Sample output

```
=== Log Analysis Summary ===
Total matched: 10

  ERROR    10
```

---

## CSV format

The tool expects a CSV file with the following columns and a header row:

```
timestamp,severity,source,message
2024-01-01 08:12:34,INFO,auth-service,User login successful for user_id=1042
2024-01-01 08:21:45,ERROR,payment-service,Stripe API timeout after 30s
```

| Column    | Format                  | Example                  |
|-----------|-------------------------|--------------------------|
| timestamp | `YYYY-MM-DD HH:MM:SS`   | `2024-01-01 08:12:34`    |
| severity  | `INFO`, `WARN`, `ERROR` | `ERROR`                  |
| source    | service name string     | `payment-service`        |
| message   | free text               | `Connection pool exhausted` |

A sample `logs.csv` with 51 entries across 3 days and 4 services is included in the repo.

---

## Running tests

```bash
go test ./...

# With race detector (recommended)
go test -race ./...

# With coverage
go test -cover ./...
```

---

## Building a binary

```bash
# Build for your current OS
make build

# Output binary
./bin/log-analyzer --file logs.csv --severity ERROR --from 2024-01-01 --to 2024-01-03
```

---

## What I learned building this

- How Go slices work under the hood — length vs capacity, how `append` allocates a new backing array when capacity is exceeded, and why functions that filter should return a new slice rather than mutate the original
- Idiomatic Go error handling — returning `error` as a second value, wrapping errors with `fmt.Errorf("context: %w", err)`, and using `log.Fatalf` at the top level instead of `panic`
- File I/O patterns — `os.Open`, `defer file.Close()`, and `encoding/csv` for structured reading
- The `time` package — parsing timestamps with Go's reference time layout (`2006-01-02 15:04:05`) and doing inclusive date range comparisons with `!t.Before(from) && !t.After(to)`

---

## License

MIT