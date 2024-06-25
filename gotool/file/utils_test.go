package file

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestConvertJson2Struct(t *testing.T) {
	file := "./json/simple.json"
	a := GetBytesByPath(file)

	var person Person
	err := json.Unmarshal(a, &person)
	if err != nil {
		fmt.Printf("failed to unmarshal, and err is: %s\n", err.Error())
		return
	}
	fmt.Printf("person is: %#v\n", person)
}
