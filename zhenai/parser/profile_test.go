package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)


func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents,"安靜的雪")
	fmt.Println(result)
}
