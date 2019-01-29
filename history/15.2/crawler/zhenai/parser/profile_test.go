package parser

import (
	"github.com/Kirk-Wang/Hello-Gopher/15.2/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "初心", "女士")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name:       "初心",
		Gender:     "女士",
		Age:        30,
		Height:     160,
		Weight:     46,
		Income:     "1.2-2万",
		Marriage:   "未婚",
		Education:  "高中及以下",
		Occupation: "医疗管理",
		Hokou:      "重庆",
		Xinzuo:     "射手座",
		House:      "租房",
		Car:        "未买车",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
