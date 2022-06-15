package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	input := "manish"
	records := readCsvFile("./read.csv")

	city:=list(records,input)
	fmt.Println(city)
	
	//for


	//fmt.Println(records)
}
func list(records [][]string,input string)string{
	// function body.....
	var i, j int
	var city string
for i = 0; i < len(records); i++ {
	for j = 0; j < len(records[i]); j++ {
		//fmt.Println(records[i][j], i, j)
		if records[i][j] == input {
			city= records[i][j+1]
			break
		}
	}
}
return city
}


func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
