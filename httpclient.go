package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	"flag"
)

const api_key = "WsJyksJKHPWzjL7dWwhRGx8PXp6VJkZV"
func main()  {
	var gifs_limit = flag.Int("gifs_limit", 5, "gifs limit number")
	flag.Parse()
	type Data []struct {
		Type string `json:"type""`
		Id string `json:"id"`
		Title string `json:"title"`
		Url string `json:"url"`
	}
	type Trenging struct{
		Data Data `json:"data"`
	}
	var api_key = "WsJyksJKHPWzjL7dWwhRGx8PXp6VJkZV"
	var url = fmt.Sprintf("http://api.giphy.com/v1/gifs/trending?api_key=%s&limit=%d", api_key, *gifs_limit)
	resp, error := http.Get(url)
	if error != nil {
		fmt.Printf("Error during get request: %s", error)
		return
	}
	defer resp.Body.Close()
	contents, error := ioutil.ReadAll(resp.Body)
	if error!= nil {
		fmt.Printf("Error during reading response: %s", error)
		return
	}
	file, error := os.Create("data")
	if error!= nil {
		fmt.Printf("Error during file creation: %s", error)
		return
	}
	defer file.Close()
	bytes_written, error := file.Write(contents)
	if error!= nil {
		fmt.Printf("Error during file update: %s", error)
		return
	}
	fmt.Printf("wrote %d bytes to file\n",bytes_written)

	var gifs Trenging
	err := json.Unmarshal(contents, &gifs)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, gif := range(gifs.Data){
		fmt.Printf("Gif Name: %s, Gif URL: %s\n", gif.Title, gif.Url)
	}
}