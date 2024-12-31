package wal

import (
	"fmt"
	"os"
)

type WALWriter struct {
	logFile string
}

func NewWALWriter(logFile string) *WALWriter {
	return &WALWriter{logFile: logFile}
}

func (w *WALWriter) Write(entry string) (int64, error) {
	// Open file in append mode
	file, err := os.OpenFile(w.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, fmt.Errorf("could not open WAL file: %v", err)
	}
	defer file.Close()

	// Write the log entry as plain text
	_, err = fmt.Fprintf(file, "%s\n", entry)
	if err != nil {
		return 0, fmt.Errorf("could not write log entry: %v", err)
	}

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, fmt.Errorf("could not get file info: %v", err)
	}

	return fileInfo.Size(), nil

}
