package main

import (
	"encoding/json"
	_ "errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"someurl.com/factorymethods"
)

func GetLastDatabase() string {
	entries, err := os.ReadDir(factorymethods.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}

	current := ""
	for _, e := range entries {
		if e.IsDir() {
			if e.Name() > current {
				current = e.Name()
			}
		}
	}
	return current
}

func main() {
	lastDB := GetLastDatabase()
	if lastDB == "" {
		fmt.Println("No database to serve")
		return
	}
	repo, err := factorymethods.InitShardedDb(factorymethods.GetDatabaseFileTemplateFromTimestamp(lastDB), false)
	if err != nil {
		fmt.Printf("Could not start server. %s", err)
		return
	}

	http.HandleFunc("/promotions/", func(w http.ResponseWriter, r *http.Request) {
		id := path.Base(r.URL.Path)
		obj := repo.Get(id)
		if obj.Id == id {
			u, err := json.Marshal(obj)
			if err == nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(200)
				w.Write(u)
			}
			w.WriteHeader(500)
		} else {
			w.WriteHeader(404)
		}
	})

	http.HandleFunc("/updateDatabase/", func(w http.ResponseWriter, r *http.Request) {
		timestamp := path.Base(r.URL.Path)
		lr, le := factorymethods.InitShardedDb(factorymethods.GetDatabaseFileTemplateFromTimestamp(timestamp), false)
		if le == nil {
			repo = lr
			w.Write([]byte("Database updated"))
		} else {
			w.Write([]byte("Database was not updated"))
		}
	})

	err = http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Printf("Could not start server. %s", err)
		return
	}
}
