package wal

import (
	"fmt"
	"github.com/Divas-Gupta30/medium-blogs/wal-kv-store/state"
	"os"
	"strings"
)

const CompactionThreshold = 64 // Size of wal.log for compaction , if size exceeds this start compaction

type WALManager struct {
	logFile string
	writer  *WALWriter
	reader  *WALReader
}

func NewManager(logFile string) *WALManager {
	return &WALManager{
		logFile: logFile,
		writer:  NewWALWriter(logFile),
		reader:  NewWALReader(logFile),
	}
}

type LogEntry struct {
	Action string
	Key    string
	Value  string
	Index  int
}

func (w *WALManager) Put(key, value string, stateManager *state.StateManager) error {
	// Get the current state
	stateData, err := stateManager.LoadState()
	if err != nil {
		return err
	}

	// Increment index for new entry
	newIndex := stateData.LastAppliedIndex + 1

	// Write to WAL in plain text format
	logEntry := fmt.Sprintf("%s %s %s %d", "put", key, value, newIndex)
	fileSize, err := w.writer.Write(logEntry)
	if err != nil {
		return err
	}

	if fileSize > CompactionThreshold {
		_ = w.Compact(stateData.LastAppliedIndex)
	}

	// Update the state
	stateData.Store[key] = value
	stateData.LastAppliedIndex = newIndex
	return stateManager.SaveState(stateData)
}

func (w *WALManager) Get(key string, stateManager *state.StateManager) (string, error) {
	// Load current state
	stateData, err := stateManager.LoadState()
	if err != nil {
		return "", err
	}

	value, exists := stateData.Store[key]
	if !exists {
		return "", nil
	}
	return value, nil
}

func (w *WALManager) Delete(key string, stateManager *state.StateManager) error {
	// Get the current state
	stateData, err := stateManager.LoadState()
	if err != nil {
		return err
	}

	// Increment index for new entry
	newIndex := stateData.LastAppliedIndex + 1

	logEntry := fmt.Sprintf("%s %s %d", "del", key, newIndex)
	_, err = w.writer.Write(logEntry)
	if err != nil {
		return err
	}

	// Update the state
	delete(stateData.Store, key)
	stateData.LastAppliedIndex = newIndex
	return stateManager.SaveState(stateData)
}

func (w *WALManager) Compact(lastAppliedIndex int) error {
	logEntries, err := w.reader.Read()
	if err != nil {
		return fmt.Errorf("unable to read log file for compaction: %v", err)
	}

	// Filter out entries with an index less than or equal to lastAppliedIndex
	var newEntries []string
	for _, entry := range logEntries {

		var newEntry string
		if entry.Action == "put" {
			newEntry = fmt.Sprintf("%s %s %s %d", entry.Action, entry.Key, entry.Value, entry.Index)
		} else {
			newEntry = fmt.Sprintf("%s %s %d", entry.Action, entry.Key, entry.Index)

		}
		if entry.Index > lastAppliedIndex {
			newEntries = append(newEntries, newEntry)
		}
	}
	// this is added to ensure a \n after last value
	newEntries = append(newEntries, "")

	// Write the remaining entries back to the wal.log file
	err = os.WriteFile(w.logFile, []byte(strings.Join(newEntries, "\n")), 0644)
	if err != nil {
		return fmt.Errorf("could not write to WAL file: %v", err)
	}

	return nil
}
