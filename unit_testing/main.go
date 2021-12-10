package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opengin json: ", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		err = decoder.Decode(&post)
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			fmt.Println("Error decoding json: ", err)
			return
		}
	}
	return
}

func unmardhal(filename string) (post Post, err error) {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON: ", err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data: ", err)
		return
	}

	json.Unmarshal(jsonData, &post)
	return
}

func main() {
	_, err := decode("post.json")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
