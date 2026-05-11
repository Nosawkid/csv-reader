package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Config struct {
	file     string
	severity string
	from     string
	to       string
}

type LogEntry struct {
	Timestamp time.Time
	Severity  string
	Source    string
	Message   string
}

func parseFlags() Config {
	filePtr := flag.String("file", "logs.csv", "Path to CSV log file")
	severityPtr := flag.String("severity", "", "Filter by severity: INFO, WARN, ERROR")
	fromPtr := flag.String("from", "", "Start date YYYY-MM-DD")
	toPtr := flag.String("to", "", "End date YYYY-MM-DD")
	flag.Parse()
	return Config{*filePtr, *severityPtr, *fromPtr, *toPtr}
}

func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("reading csv: %w", err)
	}
	return records, nil
}

func parseRow(row []string) (LogEntry, error) {
	ts, err := time.Parse("2006-01-02 15:04:05", row[0])
	if err != nil {
		return LogEntry{}, fmt.Errorf("parsing timestamp %q: %w", row[0], err)
	}
	return LogEntry{Timestamp: ts, Severity: row[1], Source: row[2], Message: row[3]}, nil
}

func filterLogs(entries []LogEntry, severity string, from, to time.Time) []LogEntry {
	var result []LogEntry
	for _, e := range entries {
		if e.Severity == severity &&
			!e.Timestamp.Before(from) &&
			!e.Timestamp.After(to) {
			result = append(result, e)
		}
	}
	return result
}

func summarize(entries []LogEntry) map[string]int {
	result := make(map[string]int)
	for _, e := range entries {
		result[e.Severity]++
	}
	return result
}

func printSummary(summary map[string]int, total int) {
	fmt.Printf("\n=== Log Analysis Summary ===\n")
	fmt.Printf("Total matched: %d\n\n", total)
	for _, sev := range []string{"ERROR", "WARN", "INFO"} {
		fmt.Printf("  %-8s %d\n", sev, summary[sev])
	}
}

func main() {
	cfg := parseFlags()

	records, err := readCSV(cfg.file)
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	var entries []LogEntry
	for _, row := range records[1:] {
		entry, err := parseRow(row)
		if err != nil {
			fmt.Printf("skipping bad row: %v\n", err)
			continue
		}
		entries = append(entries, entry)
	}

	from, _ := time.Parse("2006-01-02", cfg.from)
	to, _ := time.Parse("2006-01-02", cfg.to)

	filtered := filterLogs(entries, cfg.severity, from, to)
	summary := summarize(filtered)
	printSummary(summary, len(filtered))
}
