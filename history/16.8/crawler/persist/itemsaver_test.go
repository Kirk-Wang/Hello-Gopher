package persist

import (
	"context"
	"encoding/json"
	"github.com/Kirk-Wang/Hello-Gopher/history/16.8/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/history/16.8/crawler/model"
	"github.com/olivere/elastic"
	"testing"
)

func TestSave(t *testing.T) {
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

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// Save expected item
	err = save(client, index, expected)

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	// t.Logf("%+v", resp)
	// t.Logf("%s", resp.Source)
	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v;", actual, expected)
	}
}
