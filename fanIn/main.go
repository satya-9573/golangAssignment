package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

// MarksData structure to hold marks data
type MarksData struct {
	Name  string
	Marks int
}

// AgeData structure to hold age data
type AgeData struct {
	Name string
	Age  int
}

type Data struct {
	Name  string
	Age   int
	Marks int
}

// getCSVData reads data from the provided CSV file and sends it to the appropriate channel
func getCSVData(file string, wg *sync.WaitGroup, marksChan chan<- MarksData, ageChan chan<- AgeData, mu *sync.Mutex) {
	defer wg.Done() // Decrements the WaitGroup counter when the goroutine completes

	f, err := os.Open(file) // Opens the file for reading
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if len(records) == 0 {
		log.Println("No data found in file:", file)
		return
	}

	headers := records[0]
	if isMarksDataFile(headers) {
		for _, row := range records[1:] {
			if len(row) < 2 {
				continue
			}
			marks, err := strconv.Atoi(row[1])
			if err != nil {
				log.Println("Error converting marks:", err)
				continue
			}
			marksChan <- MarksData{Name: row[0], Marks: marks}
		}
	} else {
		for _, row := range records[1:] {
			if len(row) < 2 {
				continue
			}
			age, err := strconv.Atoi(row[1])
			if err != nil {
				log.Println("Error converting age:", err)
				continue
			}
			ageChan <- AgeData{Name: row[0], Age: age}
		}
	}
}

// isMarksDataFile determines if the file contains marks data based on headers
func isMarksDataFile(headers []string) bool {
	for _, header := range headers {
		if header == "Marks" {
			return true
		}
	}
	return false
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	marksChan := make(chan MarksData)
	ageChan := make(chan AgeData)
	AllData := make([]*Data, 0)

	files := []string{"Age.csv", "Marks.csv"}

	// Start a goroutine for each file to read data
	for _, file := range files {
		wg.Add(1)
		go getCSVData(file, &wg, marksChan, ageChan, &mu)
	}

	go func() {
		for marks := range marksChan {
			var found bool
			for _, data := range AllData {
				if data.Name == marks.Name {
					data.Marks = marks.Marks
					found = true
					break
				}
			}
			if !found {
				AllData = append(AllData, &Data{Name: marks.Name, Marks: marks.Marks})
			}
		}
	}()

	go func() {
		for age := range ageChan {
			mu.Lock()
			var found bool
			for _, data := range AllData {
				if data.Name == age.Name {
					data.Age = age.Age
					found = true
					break
				}
			}
			if !found {
				AllData = append(AllData, &Data{Name: age.Name, Age: age.Age})
			}
			mu.Unlock()
		}
	}()

	// Wait for all data to be read
	wg.Wait()

	// Close channels after all data has been sent
	close(marksChan)
	close(ageChan)

	// Print processed data
	for _, data := range AllData {
		fmt.Printf("Person:%s,Marks:%v,Age:%v\n", data.Name, data.Marks, data.Age)
	}
}
