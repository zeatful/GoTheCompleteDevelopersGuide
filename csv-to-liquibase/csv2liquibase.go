package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Object represents the row of the CSV as an object to use for writing and parsing data
type Object struct {
	Discriminator string
	Key           string
	Text          string
}

// used to replace special characters with their html encoding
var htmlEscaper = strings.NewReplacer(
	`&`, "&amp;",
	`'`, "&#39;", // "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	`<`, "&lt;",
	`>`, "&gt;",
	`"`, "&#34;", // "&#34;" is shorter than "&quot;".
	`§`, "&sect;",
)

// parse CSV and read each line into a object type
func parseCSV(path string) ([]Object, error) {
	csvFile, _ := os.Open(path)
	reader := csv.NewReader(csvFile)
	var objects []Object

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		objects = append(objects, Object{
			Discriminator: line[0],
			Key:           line[1],
			Text:          line[3],
		})
	}

	return objects, nil
}

// split up workload
func splitWork(objects []Object, author string) {
	jobcount := runtime.NumCPU() // total jobs to run
	length := len(objects)       // number of elements
	divided := length / jobcount // how many items should exist in each job

	// setup channels for jobs
	jobs := make(chan string)

	// iterate for each job
	for i := 0; i < jobcount; i++ {
		go createLiquiBase(i, objects[i*divided:(i+1)*divided], jobs, author)
	}

	// for each job block until all complete
	for i := 0; i < jobcount; i++ {
		<-jobs
	}
	close(jobs)
}

// handle a workload
func createLiquiBase(count int, objects []Object, jobs chan string, author string) {
	// generate the custom id based on time
	timestamp := time.Now().Format("200601021504")
	id := timestamp + "_update_" + strconv.Itoa(count)

	// use a string builder to be efficient
	var liquiBase strings.Builder

	// append liquiBase databaseChangelog and changeSet information
	liquiBase.WriteString("<databaseChangeLog xmlns='http://www.liquiBase.org/xml/ns/dbchangelog'\n")
	liquiBase.WriteString("\t\txmlns:xsi='http://www.w3.org/2001/XMLSchema-instance'\n")
	liquiBase.WriteString("\t\txsi:schemaLocation='http://www.liquiBase.org/xml/ns/dbchangelog http://www.liquiBase.org/xml/ns/dbchangelog/dbchangelog-2.0.xsd'>\n\n")
	liquiBase.WriteString("\t<changeSet author='")
	liquiBase.WriteString(author)
	liquiBase.WriteString("' id='")
	liquiBase.WriteString(id)
	liquiBase.WriteString("' runOnChange='false'>\n")

	// loop through each object and append a liquiBase update statement
	for _, object := range objects {
		liquiBase.WriteString("\t\t<update tableName='")
		liquiBase.WriteString(object.Discriminator)
		liquiBase.WriteString("'>\n")
		liquiBase.WriteString("\t\t\t<column name='TEXT' value='")

		// trim trailing whitespace and html encode special characters
		liquiBase.WriteString(htmlEscaper.Replace(strings.TrimSpace(object.Text)))
		liquiBase.WriteString("'/>\n")
		liquiBase.WriteString("\t\t\t<where>KEY = '")
		liquiBase.WriteString(object.Key)
		liquiBase.WriteString("'</where>\n")
		liquiBase.WriteString("\t\t</update>\n")
	}

	// append liquiBase change set and databaseChangelog closing tags
	liquiBase.WriteString("\t</changeset>\n")
	liquiBase.WriteString("</databaseChangeLog>")

	// write the liquiBase to a file
	writeLiquiBase(id, liquiBase.String())

	jobs <- id
}

// takes a change set id and the liquiBase body
func writeLiquiBase(id string, liquibase string) {
	var filename = id + ".xml"

	// attempt to create file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// write file
	_, er := file.WriteString(liquibase)

	if er != nil {
		log.Fatal(er)
	}
}

func main() {
	inputFile := os.Args[1]
	author := os.Args[2]

	// confirm number of CPU threads available
	fmt.Println("Threads Used: " + strconv.Itoa(runtime.NumCPU()))

	// parse the CSV into an object slice
	objects, _ := parseCSV(inputFile)

	// attempt to split the work to write out liquiBase for the objects
	splitWork(objects, author)
}
