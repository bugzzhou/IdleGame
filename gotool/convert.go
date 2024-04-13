package gotool

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// 示例  新建一个结构体，然后写一个转化函数
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ConvertJson2interface(filePath string) interface{} {
	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var person Person

	err = json.Unmarshal(jsonBytes, &person)
	if err != nil {
		fmt.Printf("failed to unmarshal, and err is: %s\n", err.Error())
		return nil
	}
	return person
}

// 示例
