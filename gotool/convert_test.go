package gotool

import "testing"

func TestConvertJson2Struct(t *testing.T) {
	file := "./json/simple.json"
	ConvertJson2interface(file)
}
