package parser

import (
	"github.com/Kirk-Wang/Hello-Gopher/17.9/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/17.9/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "http://album.zhenai.com/u/1113093123", "心想事成", "女士")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1113093123",
		Type: "zhenai",
		Id:   "1113093123",
		Payload: model.Profile{
			Name:       "心想事成",
			Gender:     "女士",
			Age:        26,
			Height:     159,
			Weight:     47,
			Income:     "8千-1.2万",
			Marriage:   "未婚",
			Education:  "硕士",
			Occupation: "金融/银行/保险",
			Hokou:      "重庆",
			Xinzuo:     "射手座",
			House:      "已购房",
			Car:        "未买车",
		},
	}

	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}
}
