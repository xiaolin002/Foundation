package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type XKCDComic struct {
	Num        int    `json:"num"`
	Title      string `json:"title"`
	Img        string `json:"img"`
	Transcript string `json:"transcript"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a search term.")
		os.Exit(1)
	}

	searchTerm := strings.Join(os.Args[1:], " ")

	// 加载索引
	index, err := loadIndex("xkcd_index.json")
	if err != nil {
		fmt.Println("Error loading index:", err)
		os.Exit(1)
	}

	// 搜索索引
	results := searchIndex(index, searchTerm)

	// 打印结果
	if len(results) == 0 {
		fmt.Println("No results found.")
	} else {
		fmt.Println("Matching comics:")
		for _, comic := range results {
			fmt.Printf("URL: %s\n", comic.Img)
		}
	}
}

func loadIndex(filename string) ([]XKCDComic, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var index []XKCDComic
	err = json.Unmarshal(data, &index)
	if err != nil {
		return nil, err
	}

	return index, nil
}

func searchIndex(index []XKCDComic, searchTerm string) []XKCDComic {
	var results []XKCDComic
	for _, comic := range index {
		if strings.Contains(comic.Title, searchTerm) || strings.Contains(comic.Transcript, searchTerm) {
			results = append(results, comic)
		}
	}
	return results
}
