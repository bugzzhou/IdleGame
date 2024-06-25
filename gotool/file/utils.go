package file

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetBytesByPath(filePath string) []byte {
	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	return jsonBytes

}

func IsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func Create(filePath string, data interface{}) error {
	var file *os.File
	var err error

	if IsExist(filePath) {
		// 文件存在，打开文件进行写入
		file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	} else {
		// 文件不存在，创建文件
		file, err = os.Create(filePath)
	}

	if err != nil {
		return fmt.Errorf("failed to open or create file: %w", err)
	}
	defer file.Close()

	// 将结构体数据转换为 JSON 格式
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// 写入文件
	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	fmt.Printf("Data written to file %s successfully.\n", filePath)
	return nil
}
