package main

import (
	"fmt"
	"io"
	"os"
)

// 利用io.Copy()实现文件拷贝

func CopyFile(src, dst string) (int64, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

func main() {
	_, err := CopyFile("chapter8_rwData/fileCopy/source.txt", "chapter8_rwData/fileCopy/target.txt")
	if err != nil {
		fmt.Println("copy error", err)
		return
	}
	fmt.Println("copy done")

}
