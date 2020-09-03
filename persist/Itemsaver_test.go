package persist

import (
	"Crawler/model"
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Name:       "安静的雪",
		Gender:     "女",
		Age:        18,
		Height:     159,
		Weight:     88,
		Income:     "2000-3000元",
		Marriage:   "离异",
		Education:  "本科",
		Occupation: "人事/行政",
		Hokou:      "山东菏泽",
		Xinzou:     "牡羊座",
		House:      "已购房",
		Car:        "未购车",
	}

	id, err := Save(expected)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index(ES_INDEX).
		Type(ES_TYPE).
		Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("got %v; expected  %v", actual, expected)
	}
}
