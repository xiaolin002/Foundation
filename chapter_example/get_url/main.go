package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	url := "baidu.com"

	prefix := strings.HasPrefix(url, "http://")
	if !prefix {

		//url = strings.Join([]string{"http://", url}, "")
		//fmt.Println(url)
		var builder strings.Builder
		builder.WriteString("http://")
		builder.WriteString(url)
		url = builder.String()
		fmt.Println(url)
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	//b, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(resp.Status)
	//defer resp.Body.Close()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	//	os.Exit(1)
	//}

	//fmt.Printf("%s", b)

}
