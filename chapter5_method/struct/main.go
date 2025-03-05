package main

import (
	"fmt"
	"time"
)

// Address 结构体定义了一个地址的详细信息
type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

// Vcard 结构体定义了一个人的基本信息，包括一个指向 Address 结构体的指针
type Vcard struct {
	Name        string
	AddressID   int
	DateOfBirth time.Time
	Image       []byte
	Address     *Address // 使用指针类型
}

func main() {
	// 创建一个 Address 实例
	address := &Address{
		Street:  "123 Main St",
		City:    "Anytown",
		State:   "CA",
		ZipCode: "12345",
	}

	// 创建一个 Vcard 实例
	vcard := Vcard{
		Name:        "John Doe",
		AddressID:   1,
		DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		Image:       []byte{0x12, 0x34, 0x56, 0x78}, // 示例图像数据
		Address:     address,
	}

	// 打印 Vcard 的内容
	fmt.Printf("Name: %s\n", vcard.Name)
	fmt.Printf("Address ID: %d\n", vcard.AddressID)
	fmt.Printf("Date of Birth: %s\n", vcard.DateOfBirth.Format("2006-01-02"))
	fmt.Printf("Image: %v\n", vcard.Image)
	fmt.Printf("Address: %v\n", vcard.Address)
}
