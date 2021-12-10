package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct { //#a
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"` //attrはモードフラグ
	Content string   `xml:"content"` //モードフラグなしはタグに紐付けられる
	Author  Author   `xml:"author"`
	Xml     string   `xml:"innerxml"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Post2 struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   string    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:comments>comment"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opeing XML file: ", err)
		return
	}
	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data: ", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)

	xmlFile, err = os.Open("post2.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	xmlData, err = ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data: ", err)
		return
	}

	var post2 Post2
	xml.Unmarshal(xmlData, &post2)
}
