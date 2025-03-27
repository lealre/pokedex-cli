package commands

import "fmt"

func commandPokedex(_ string, storage *Storage) error {
	if len(storage.Storage) == 0 {
		fmt.Println("You have not caught any Pokemon yet.")
		return nil
	}

	for name := range storage.Storage {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
