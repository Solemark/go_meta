package employeeRouter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func writeEmployeeList(data []Employee) {
	file, err := os.Create("data/employees.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	output := [][]string{}

	for _, e := range data {
		row := []string{e.FirstName, e.LastName, e.EmailAddress, strconv.FormatBool(e.Visible)}
		output = append(output, row)
	}
	writer.WriteAll(output)
}
