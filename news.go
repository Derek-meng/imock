package main

import (
	"imock/engine"
	"imock/zhenai/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		simple: parser.ParseCityList,
	})
}

