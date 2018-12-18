package persist

import (
	"context"
	"encoding/json"
	"github.com/Kirk-Wang/Hello-Gopher/16.5/crawler/model"
	"github.com/olivere/elastic"
	"testing"
)

func TestSave(t *testing.T) {
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

	id, err := save(expected)

	if err != nil {
		panic(err)
	}

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	// t.Logf("%+v", resp)
	// t.Logf("%s", resp.Source)
	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)

	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("got %v; expected %v;", actual, expected)
	}
}
