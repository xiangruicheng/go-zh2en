package check

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"unicode"
)

// containsChinese 检查字符串是否包含中文字符
func containsChinese(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// checkFileForChinese 检查文件内容中的中文字符并输出所在行号
func checkFileForChinese(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if containsChinese(line) {
			TotalRows += 1
			fmt.Printf("File %s contains Chinese characters on line %d: %s\n", filePath, lineNumber, line)
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
	}
}

// walkDir 递归遍历目录，检查文件
func walkDir(root string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			// 过滤非文本文件（如果需要）

			checkFileForChinese(path)

		}
		return nil
	})
	return err
}

var TotalRows int

func Run(directory string) {
	// 替换为你想要检查的文件夹路径
	err := walkDir(directory)
	fmt.Printf("total rows:%d\n", TotalRows)
	if err != nil {
		fmt.Printf("Error walking the directory %s: %v\n", directory, err)
	}
}
