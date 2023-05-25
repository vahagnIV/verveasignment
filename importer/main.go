package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"someurl.com/datarepository"
	"someurl.com/factorymethods"
	"strconv"
	"strings"
	"time"
)

const batchSize int = 1000

func importCsv(filename string, repo datarepository.Repo) error {
	scanner, err := openFile(filename)
	if err != nil {
		return err
	}
	counter := 0
	i := 0
	var arr []datarepository.DataRow

	for scanner.Scan() {
		counter += 1
		i++
		if counter%batchSize == 0 {
			fmt.Println("Line: ", i)
			counter = 0
			repo.BatchInsert(arr)
			arr = []datarepository.DataRow{}
		}
		splitValues := strings.Split(scanner.Text(), ",")

		price, err := strconv.ParseFloat(splitValues[1], 64)

		if err != nil {
			fmt.Printf("Could not parse value %s\n", splitValues[1])
			continue
		}
		tm, err := time.Parse("2006-01-02 15:04:05 +0200 MST", splitValues[2])

		if err != nil {
			fmt.Printf("Could not parse value %s, %s\n", splitValues[2], err)
			continue
		}
		arr = append(arr, datarepository.DataRow{Id: splitValues[0], Price: price, ExpirationDate: tm})
	}
	if len(arr) != 0 {
		repo.BatchInsert(arr)
	}

	return nil
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: app path/to/import.csv")
		return
	}

	filepath := os.Args[1]

	if !factorymethods.FileExists(filepath) {
		fmt.Println("Could not read import file")
		return
	}

	startTime := time.Now()

	timestamp := strconv.FormatInt(startTime.Unix(), 10)

	dbFileTemplate := factorymethods.GetDatabaseFileTemplateFromTimestamp(timestamp)
	err := os.MkdirAll(path.Dir(dbFileTemplate), os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory %s", err)
	}

	repo, err := factorymethods.InitShardedDb(dbFileTemplate, true)
	//repo, err := factorymethods.InitSqlShard("/home/vahagn/GolandProjects/verveTest/data/db/gorm")

	if err != nil {
		print("Could not initialize databases")
		return
	}

	err = importCsv(filepath, repo)
	if err == nil {
		serverUpdated := false
		for i := 0; i < 3; i++ {
			_, err := http.Get(factorymethods.ServerUpdateUrl + timestamp)
			if err == nil {
				serverUpdated = true
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
		if !serverUpdated {
			fmt.Println("Could not update server")
		}
	}

	endTime := time.Now()

	timeDiff := endTime.Sub(startTime).Milliseconds()

	fmt.Printf("Elapsed time: %d\n", timeDiff)
}
