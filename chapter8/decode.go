package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

type Post struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error openong JSON file: ", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("Error decoding JSON: ", err)
		return
	}
	return
}

func unmarsha1(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}
	json.Unmarshal(jsonData, &post)
	return
}

func main () {
	_, err := decode("post.json")
	if err != nil {
		fmt.Println("Error:", err)
	}
}