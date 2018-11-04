package parser

import (
	"imock/engine"
	"imock/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08="">([\d]+岁)</div>`)
var allRE = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08="">([^<]+)</div>`)

func ParseProfile(contents [] byte,name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name=name
	age,err := strconv.Atoi(extractString(contents,ageRe))
	if err!=nil {
		profile.Age=age
	}
	result:= engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return  result
}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return  string(match[1])
	}else {
		return ""
	}
}
