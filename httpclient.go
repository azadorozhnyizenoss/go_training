package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// заховай ключ через аргументи прапорця
// https://golang.org/doc/effective_go.html#mixed-caps
const api_key = "WsJyksJKHPWzjL7dWwhRGx8PXp6VJkZV"

// виносим змінну в heap
var gifs_limit = flag.Int("gifs_limit", 5, "gifs limit number")

type Data []struct {
	Type  string `json:"type""`
	Id    string `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}
type Trenging struct {
	Data Data `json:"data"`
}

func main() {
	flag.Parse()

	var url = fmt.Sprintf("http://api.giphy.com/v1/gifs/trending?api_key=%s&limit=%d", api_key, *gifs_limit)
	resp, error := http.Get(url)
	if error != nil {
		fmt.Printf("Error during get request: %s", error)
		return
	}
	defer resp.Body.Close()

	contents, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		fmt.Printf("Error during reading response: %s", error)
		return
	}
	file, error := os.Create("data")
	if error != nil {
		fmt.Printf("Error during file creation: %s", error)
		return
	}
	defer file.Close()

	// https://golang.org/doc/effective_go.html#mixed-caps
	bytes_written, error := file.Write(contents)
	if error != nil {
		fmt.Printf("Error during file update: %s", error)
		return
	}
	fmt.Printf("wrote %d bytes to file\n", bytes_written)

	var gifs Trenging
	err := json.Unmarshal(contents, &gifs)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, gif := range gifs.Data {
		fmt.Printf("Gif Name: %s, Gif URL: %s\n", gif.Title, gif.Url)
	}
}
