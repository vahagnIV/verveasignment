package main

import (
	"fmt"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
	"someurl.com/datarepository"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := openFile("/home/vahagn/GolandProjects/verveTest/data/promotions.csv")
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		// do something
	}

	db.AutoMigrate(&datarepository.MyObject{})

	for scanner.Scan() {
		split_values := strings.Split(scanner.Text(), ",")

		price, err := strconv.ParseFloat(split_values[1], 64)

		if err != nil {
			//do someting
		}
		tm, err := time.Parse("2006-01-02 15:04:05 +0200 CEST", split_values[2])

		if err != nil {
			//do someting
		}
		insertObject := &datarepository.MyObject{Id: split_values[0], Price: price, ExpirationDate: tm}

		db.Create(insertObject)
	}
	fmt.Println("Hello, World!")
}
