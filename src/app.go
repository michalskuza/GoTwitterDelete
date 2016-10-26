package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"./twitterDelete"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	routePath := "/twitterDelete"

	router.HandleFunc(routePath + "/export", ExportTweetsToFile).Methods("POST")
	router.HandleFunc(routePath + "/deleteAll", DeleteAllTweets).Methods("POST")
	router.HandleFunc(routePath + "/deleteAllExceptFavorited", DeleteAllTweetsExceptFavorited).Methods("POST")

	port := 8080
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	fmt.Printf("Go Twitter Delete listening on port %v!\n", port)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), router))
}

//ExportTweetsToFile --
func ExportTweetsToFile(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	exportRequest := new(twitterDelete.ExportRequest)
	json.Unmarshal(body, &exportRequest)

	client := twitterDelete.CreateTwitterClient(exportRequest.Credentials)
	tweets := twitterDelete.GetAllTweetsList(*client)
	dirStatic := "static/"
	dirPath := "exported/"

	fileName := exportRequest.Username + "_exported"

	if strings.EqualFold(exportRequest.Extension, "csv") {
		twitterDelete.ExportToCSVFile(dirStatic + dirPath + fileName, tweets)
	} else if strings.EqualFold(exportRequest.Extension, "json") {
		twitterDelete.ExporToJSONTFile(dirStatic + dirPath + fileName, tweets)
	}

	fmt.Fprintln(w, fileName + "." + exportRequest.Extension)
}

//DeleteAllTweets --
func DeleteAllTweets(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	credentials := new(twitterDelete.Credentials)
	json.Unmarshal(body, &credentials)
	client := twitterDelete.CreateTwitterClient(*credentials)
	tweets := twitterDelete.GetAllTweetsList(*client)
	twitterDelete.DeleteAllTweets(*client, tweets)
	fmt.Fprintln(w, "All tweets deleted")
}

//DeleteAllTweetsExceptFavorited --
func DeleteAllTweetsExceptFavorited(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	credentials := new(twitterDelete.Credentials)
	json.Unmarshal(body, &credentials)
	client := twitterDelete.CreateTwitterClient(*credentials)
	tweets := twitterDelete.GetAllTweetsList(*client)
	twitterDelete.DeleteAllTweetsExceptFavorited(*client, tweets)
	fmt.Fprintln(w, "All tweets except favorited deleted")
}
