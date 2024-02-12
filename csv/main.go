package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Record struct {
	Date time.Time
	Open float64
}

func main() {

	http.HandleFunc("/", web)
	http.ListenAndServe(":8080", nil)

}

func web(resp http.ResponseWriter, req *http.Request) {

	records := csvFile("table.csv")

	tpl, err := template.ParseFiles("hw.gohtml")
	if err != nil {
		fmt.Println(err)
	}

	err = tpl.Execute(resp, records)
	if err != nil {
		fmt.Println(err)
	}

}

func csvFile(filePath string) []Record {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	c := csv.NewReader(file)

	data, err := c.ReadAll()

	records := make([]Record, 0, len(data))

	for _, row := range data {
		date, _ := time.Parse("2006-02-01", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}
	return records
}
