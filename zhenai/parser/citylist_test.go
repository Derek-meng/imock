package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)
	const  resultSize  =470
	if len(result.Requests) != resultSize{
		t.Errorf("result should have %d requests ; but had %d", resultSize ,len(result.Requests))
	}
	if len(result.Items)!= resultSize {
		t.Errorf("result should have %d requests ; but had %d",resultSize,len(result.Items))
	}

}
