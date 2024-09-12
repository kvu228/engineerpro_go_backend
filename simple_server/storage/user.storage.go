package storage

import (
	"encoding/json"
	"io"
	"os"
	"simple_server/models"
	"sync"
)

var (
	Users    map[string]models.User
	dataFile = "users.json"
	Mutex    sync.Mutex
)

// LoadUsers loads user data from the file
func LoadUsers() error {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			Users = make(map[string]models.User)
			return nil
		}
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Users)
	return err
}

// SaveUsers saves user data to the file
func SaveUsers() error {
	Mutex.Lock()
	defer Mutex.Unlock()

	data, err := json.Marshal(Users)
	if err != nil {
		return err
	}

	err = os.WriteFile(dataFile, data, 0644)
	return err
}
