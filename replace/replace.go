package replace

import (
	"bufio"
	"fmt"
	"go-zh2en/youdao"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	// 假设的Translate函数，用于翻译汉字段落
	// 你需要实现这个函数，或者调用一个翻译API
)

// Translate 假设的翻译函数
func Translate(chinese string) (string, error) {
	// 这里调用翻译API的代码
	// 返回一个英文翻译和可能的错误

	str := youdao.Fanyi(chinese)
	fmt.Println(chinese)
	fmt.Println(str)
	return str, nil
	return "translated text", nil // 示例中省略了实际实现
}

// replaceChineseParagraphs 翻译并替换文件中的汉字段落
func replaceChineseParagraphs(filePath string) error {
	// 读取文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// 正则表达式，匹配汉字段落（这里只是一个示例，实际可能需要更复杂的表达式）
	re := regexp.MustCompile(`[\p{Han}]+`) // 匹配连续的中文字符

	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1) // 查找所有匹配的汉字段落

		for _, match := range matches {
			translation, err := Translate(match)
			if err != nil {
				return err
			}
			// 替换原文中的汉字段落为英文翻译
			line = strings.ReplaceAll(line, match, translation)
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// 将修改后的行合并回文件内容
	newContent := strings.Join(lines, "\n")

	// 写回文件
	return ioutil.WriteFile(filePath, []byte(newContent), 0644)
}

// traverseDirectory 遍历目录并处理文件
func TraverseDirectory(root string) error {
	// 遍历目录
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			// 调用函数处理文件
			if err := replaceChineseParagraphs(path); err != nil {
				fmt.Printf("Error processing file %s: %v\n", path, err)
			}
		}
		return nil
	})
	return err
}
