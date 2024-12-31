package wal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type WALReader struct {
	logFile string
}

func NewWALReader(logFile string) *WALReader {
	return &WALReader{logFile: logFile}
}

func (r *WALReader) Read() ([]LogEntry, error) {
	var entries []LogEntry
	file, err := os.Open(r.logFile)
	if err != nil {
		return nil, fmt.Errorf("could not open WAL file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 3 {
			continue
		}

		var entry LogEntry
		entry.Action = parts[0]
		entry.Key = parts[1]
		if entry.Action == "put" {
			if len(parts) != 4 {
				continue
			}
			entry.Value = parts[2]
			fmt.Sscanf(parts[3], "%d", &entry.Index)
		} else {
			fmt.Sscanf(parts[2], "%d", &entry.Index)
		}
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading WAL file: %v", err)
	}

	return entries, nil
}
