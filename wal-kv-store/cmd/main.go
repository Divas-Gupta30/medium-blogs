package main

import (
	"flag"
	"fmt"
	"github.com/Divas-Gupta30/medium-blogs/wal-kv-store/state"
	"github.com/Divas-Gupta30/medium-blogs/wal-kv-store/storage"
	"github.com/Divas-Gupta30/medium-blogs/wal-kv-store/wal"
	"log"
)

func main() {
	action := flag.String("action", "put", "Action to perform: put, get, del")
	key := flag.String("key", "k1", "Key for the operation")
	value := flag.String("value", "v6", "Value for the put operation")

	flag.Parse()

	if *action == "" || *key == "" {
		log.Fatal("Please provide an action (put, get, del) and a key.")
		return
	}

	// Initialize components
	storageService := storage.NewFileStorage("state.json")
	stateManager := state.NewStateManager(storageService)
	walManager := wal.NewManager("wal.log")

	// Perform action based on the CLI input
	switch *action {
	case "put":
		if *value == "" {
			log.Fatal("Please provide a value for the put operation with -value flag.")
		}
		err := walManager.Put(*key, *value, stateManager)
		if err != nil {
			log.Fatalf("Error during put operation: %v", err)
		}
		fmt.Printf("Key '%s' set to value '%s'.\n", *key, *value)

	case "get":
		value, err := walManager.Get(*key, stateManager)
		if err != nil {
			log.Fatalf("Error during get operation: %v", err)
		}
		if value == "" {
			fmt.Printf("Key '%s' not found.\n", *key)
		} else {
			fmt.Printf("Key '%s' has value '%s'.\n", *key, value)
		}

	case "del":
		err := walManager.Delete(*key, stateManager)
		if err != nil {
			log.Fatalf("Error during delete operation: %v", err)
		}
		fmt.Printf("Key '%s' deleted.\n", *key)

	default:
		log.Fatal("Invalid action. Supported actions are: put, get, del.")
	}
}
