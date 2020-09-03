package persist

import (
	"Crawler/engine"
	"Crawler/model"
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "https://album.zhenai.com/u/1241698784",
		Type: "zhenai",
		Id:   "1241698784",
		Payload: model.Profile{
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
		},
	}

	err := Save(expected)
	if err != nil {
		panic(err)
	}

	// TODO: 记得启动es docker
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expected {
		t.Errorf("got %v; expected  %v", actual, expected)
	}
}
