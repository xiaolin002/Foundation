package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopl.io/ch4/github"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var (
	createFlag = flag.Bool("create", false, "Create a new issue")
	readFlag   = flag.Bool("read", false, "Read an issue")
	updateFlag = flag.Bool("update", false, "Update an issue")
	closeFlag  = flag.Bool("close", false, "Close an issue")
	issueID    = flag.Int("id", 0, "Issue ID")
)

func main() {
	flag.Parse()

	if *createFlag {
		createIssue()
	} else if *readFlag {
		readIssue(*issueID)
	} else if *updateFlag {
		updateIssue(*issueID)
	} else if *closeFlag {
		closeIssue(*issueID)
	} else {
		fmt.Println("Please specify an action: -create, -read, -update, or -close")
		flag.PrintDefaults()
	}
}

func createIssue() {
	// 打开默认编辑器让用户输入issue信息
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}
	cmd := exec.Command(editor)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// 读取编辑器中的内容
	issueContent, err := ioutil.ReadAll(cmd.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// 解析issue信息
	var issue github.Issue
	err = json.Unmarshal(issueContent, &issue)
	if err != nil {
		log.Fatal(err)
	}

	// 创建issue
	createdIssue, _, err := github.CreateIssue(issue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created issue #%d: %s\n", createdIssue.Number, createdIssue.Title)
}

func readIssue(id int) {
	issue, _, err := github.GetIssue(id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("#%-5d %9.9s %.55s\n",
		issue.Number, issue.User.Login, issue.Title)
}

func updateIssue(id int) {
	// 打开默认编辑器让用户输入更新后的issue信息
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}
	cmd := exec.Command(editor)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// 读取编辑器中的内容
	issueContent, err := ioutil.ReadAll(cmd.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// 解析更新后的issue信息
	var updatedIssue github.Issue
	err = json.Unmarshal(issueContent, &updatedIssue)
	if err != nil {
		log.Fatal(err)
	}

	// 更新issue
	_, _, err = github.UpdateIssue(id, updatedIssue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated issue #%d\n", id)
}

func closeIssue(id int) {
	// 获取当前issue的状态
	issue, _, err := github.GetIssue(id)
	if err != nil {
		log.Fatal(err)
	}

	// 如果issue已经关闭，则直接返回
	if issue.State == "closed" {
		fmt.Printf("Issue #%d is already closed\n", id)
		return
	}

	// 更新issue状态为关闭
	issue.State = "closed"
	_, _, err = github.UpdateIssue(id, issue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Closed issue #%d\n", id)
}
