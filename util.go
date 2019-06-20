package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var pods []Response

func ReadJSONFromFile(filename string) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	_ = json.Unmarshal(byteValue, &pods)
}
