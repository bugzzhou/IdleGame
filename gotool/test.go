package gotool

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tidwall/gjson"
)

func TestGjson() {
	jsonBytes, err := ioutil.ReadFile("./json/1.json")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	// 将字节切片转换为字符串
	jsonString := string(jsonBytes)

	name := gjson.Get(jsonString, "name")
	fmt.Println("last name:", name.String())

	age := gjson.Get(jsonString, "age")
	fmt.Println("age:", age.Int())

	add2 := gjson.Get(jsonString, "add.2.city")
	fmt.Println("age:", add2.String())

}
