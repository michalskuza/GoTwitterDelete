package twitterDelete

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
)

//ExportToCSVFile --
func ExportToCSVFile(fileName string, tweets []twitter.Tweet) {
	file, err := os.Create(fileName + ".csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, tweet := range tweets {
		err := writer.Write([]string{tweet.IDStr, tweet.CreatedAt, tweet.Text})
		checkError("Cannot write to file", err)
	}

	defer writer.Flush()
}

//ExporToJSONTFile --
func ExporToJSONTFile(fileName string, tweets []twitter.Tweet) {
	file, err := os.Create(fileName + ".json")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, tweet := range tweets {
		marshaled, _ := json.Marshal(tweet)
		err := writer.Write([]string{tweet.IDStr, string(marshaled)})
		checkError("Cannot write to file", err)
	}

	defer writer.Flush()
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
