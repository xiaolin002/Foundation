package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Movie struct {
	Title  string `json:"Title"`
	Poster string `json:"Poster"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a movie title.")
		os.Exit(1)
	}

	movieTitle := strings.Join(os.Args[1:], " ")

	// 调用OMDB API获取电影信息
	movie, err := getMovieInfo(movieTitle)
	if err != nil {
		fmt.Println("Error getting movie info:", err)
		os.Exit(1)
	}

	// 下载海报
	err = downloadPoster(movie.Poster, movie.Title+".jpg")
	if err != nil {
		fmt.Println("Error downloading poster:", err)
		os.Exit(1)
	}

	fmt.Printf("Poster eleven_1 '%s' downloaded successfully.\n", movie.Title)
}

func getMovieInfo(title string) (*Movie, error) {
	apiKey := "YOUR_OMDB_API_KEY" // 替换为你的OMDB API密钥
	url := fmt.Sprintf("http://www.omdbapi.com/?t=%s&apikey=%s", title, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var movie Movie
	err = json.Unmarshal(body, &movie)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func downloadPoster(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
