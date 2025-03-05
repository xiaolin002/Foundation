package main

import (
	"flag"
	"fmt"
	"gopl.io/ch4/github"
	"log"
	"time"
)

func main() {
	var query []string

	// 使用 flag.Var 函数将命令行参数解析到 query 变量中
	flag.Var((*funcVar)(&query), "q", "Query terms")
	flag.Parse()
	result, err := github.SearchIssues(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	//定义时间分类的阈值
	// 当前的时间减去一个月的时间
	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	// 当前的时间减去一年的时间
	oneYearAgo := time.Now().AddDate(-1, 0, 0)

	//分类打印问题
	fmt.Println("Issues within the last month:")
	printIssuesWithinTimeRange(result.Items, oneMonthAgo, time.Now())

	fmt.Println("\nIssues within the last year:")
	printIssuesWithinTimeRange(result.Items, oneYearAgo, oneMonthAgo)

	fmt.Println("\nIssues over a year old:")
	printIssuesWithinTimeRange(result.Items, time.Time{}, oneYearAgo)

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

// funcVar 是一个自定义类型，实现了 flag.Value 接口
type funcVar []string

func (i *funcVar) String() string {
	return fmt.Sprint(*i)
}

func (i *funcVar) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// printIssuesWithinTimeRange 打印在指定时间范围内的问题
func printIssuesWithinTimeRange(items []*github.Issue, start time.Time, end time.Time) {
	for _, item := range items {
		//如果创建的时间点在（当前时间减去一个月）之后   在当前时间之前
		if item.CreatedAt.After(start) && item.CreatedAt.Before(end) {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
