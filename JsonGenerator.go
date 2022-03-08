package main

import (
	"JsonGenerator/Tools"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

var s = flag.String("Json", "", "Json字符串")

var f = flag.String("FilePath", "", "Json格式文件路径")

func main() {
	flag.Parse()
	jsonStr := ""

	if len(*s) > 0 {
		jsonStr = *s
	} else if len(*f) > 0 {
		data, err := os.ReadFile(*f)
		if err != nil {
			log.Fatalf("Input Json File Path Error:%s", err)
			return
		}
		jsonStr = string(data)
	} else {
		fmt.Println("使用方式：")
		//fmt.Println("-Json   : Json字符串")
		fmt.Println("-FilePath   : Json格式文件路径")
		return
	}

	var info map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &info)

	if err != nil {
		log.Fatalf("decode err:%s", err)
		return
	}

	Tools.JsonInfo.Init()
	Tools.Check("Info", info)

	log.Println("----------------------------------")

	Tools.JsonInfo.Gen()
}
