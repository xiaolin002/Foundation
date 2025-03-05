package main

import (
	"fmt"
	"os"
)

/*
编写一个save方法，将title作为文件名，将body作为文件内容，写入到文本文件中
再编写load方法，接收的参数是title 该函数读取title对应的文件内容，使用*Page作为参数 并且使用ioutil包

*/

// Page 结构体
type Page struct {
	Title string
	Body  []byte
}

// save 方法将 Page 的内容保存到文件
func (p *Page) save() error {
	// 使用 Title 作为文件名
	filename := p.Title + ".txt"
	// 将 Body 写入文件
	// 注意将文件的路径加上 负责不会创建chapter8_rwData/eq/2/
	return os.WriteFile("chapter8_rwData/eq/2/"+filename, p.Body, 0666)
}

// load 方法从文件中读取内容并填充到 Page 结构体
func (p *Page) load(title string) error {
	// 使用 Title 作为文件名
	filename := title + ".txt"
	// 读取文件内容
	body, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	// 填充 Page 结构体
	p.Title = title
	p.Body = body
	return nil
}

func main() {
	// 创建一个 Page 实例
	page := &Page{Title: "example", Body: []byte("This is an example body.")}

	// 保存 Page 到文件
	if err := page.save(); err != nil {
		fmt.Println("Error saving page:", err)
		return
	}
	fmt.Println("Page saved successfully.")

	// 创建一个新的 Page 实例用于加载
	loadedPage := &Page{}
	if err := loadedPage.load("example"); err != nil {
		fmt.Println("Error loading page:", err)
		return
	}

	// 输出加载的内容
	fmt.Printf("Loaded Page Title: %s\n", loadedPage.Title)
	fmt.Printf("Loaded Page Body: %s\n", loadedPage.Body)
}
