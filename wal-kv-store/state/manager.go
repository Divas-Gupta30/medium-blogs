package state

import (
	"encoding/json"
	"fmt"
	"github.com/Divas-Gupta30/medium-blogs/wal-kv-store/storage"
)

type State struct {
	LastAppliedIndex int               `json:"last_applied_index"`
	Store            map[string]string `json:"store"`
}

type StateManager struct {
	storage storage.Storage
}

func NewStateManager(storage storage.Storage) *StateManager {
	return &StateManager{storage: storage}
}

func (s *StateManager) LoadState() (*State, error) {
	data, err := s.storage.Load()
	if err != nil {
		return nil, err
	}

	var state *State
	if len(data) > 0 {
		err = json.Unmarshal(data, &state)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal state: %v", err)
		}
	}
	if state == nil {
		state = &State{
			LastAppliedIndex: 0,
			Store:            make(map[string]string),
		}
	}
	return state, nil
}

func (s *StateManager) SaveState(state *State) error {
	data, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("could not marshal state: %v", err)
	}
	return s.storage.Save(data)
}
