package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRoutes(r *gin.Engine) {
	r.GET("/city/:name", Dummy)
}

//Dummy function
func Dummy(c *gin.Context) {
	name, ok := c.Params.Get("name")
	records := readCsvFile("./read.csv")

	cityName := getCity(records, name)
	if ok == false {
		res := gin.H{
			"error": "name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	/*
		city := ""
	*/
	res := gin.H{
		"name": name,
		"city": cityName,
	}
	c.JSON(http.StatusOK, res)
}

func getCity(records [][]string, input string) string {
	var cityName string
	for i := 0; i < len(records); i++ {

		//fmt.Println(records[0][0], i)
		if records[i][0] == input {
			cityName = records[i][1]
			break
		}

	}
	return cityName
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
