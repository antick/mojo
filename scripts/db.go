package scripts

import (
	"github.com/dgraph-io/badger/v3"
	"log"
	"os"
	"path/filepath"
)

var db *badger.DB

func InitDB(dbPath string) {
	var err error
	opts := badger.DefaultOptions(dbPath).WithLogger(nil)
	db, err = badger.Open(opts)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
}

func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			return
		}
	}
}

func SetKeyValue(key, value string) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
}

func GetKeyValue(key string) (string, error) {
	var valCopy []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		valCopy, err = item.ValueCopy(nil)
		return err
	})
	return string(valCopy), err
}

func GetDBPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, "Library", "Application Support", "antick", "mojo.db")
}
