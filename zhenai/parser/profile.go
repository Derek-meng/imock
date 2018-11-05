package parser

import (
	"fmt"
	"imock/engine"
	"imock/model"
	"regexp"
	"strconv"
)

const (
	MARRIAGE   = 0
	XINZUO     = 2
	HEIGHT     = 3
	OCCUPATION = 7
	EDUCATION  = 8
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>([\d]+)岁</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>([\d]+)kg</div>`)
var allRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>([^<]+)</div>`)
var IncomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>>月收入:([^<]+)</div>`)
var houseRe =regexp.MustCompile(`<div class="m-btn pink" data-v-ff544c08="">([^<]+)</div>`)

func ParseProfile(contents [] byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		profile.Weight = weight
	}
	income := extractString(contents, IncomeRe)
	profile.Income = income
	allSubMatch := allRe.FindAllSubmatch(contents, -1)
	var matchResult string
	for key, match := range allSubMatch {
		matchResult = string(match[1])
		switch key {
		case MARRIAGE:
			profile.Marriage = matchResult
		case XINZUO:
			profile.Xinzuo = matchResult
		case HEIGHT:
			height, err := strconv.Atoi(matchResult)
			if err != nil {
				profile.Height = height
			}
		case OCCUPATION:
			profile.Occupation = matchResult
		case EDUCATION:
			profile.Education = matchResult
		}
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	fmt.Println(profile)
	return result
}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
