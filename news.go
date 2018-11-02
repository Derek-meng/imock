package main

import (
	. "news/engine"
	"news/zhenai/parser"
)

func main() {
	Run(Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

