package main

import (
	"./models"
	"./utils"
	"fmt"
)

func main() {

	data := createData()

	keys := []string{"doc1", "doc2", "doc3", "doc4"}

	for _, key := range keys {

		result := getResult(key, data)

		// 出力
		fmt.Println(key , "========================")
		for _, val := range result {
			fmt.Println(val.GetName(), ",", val.GetValue())
		}
		fmt.Println("")

	}
}

func createData() map[string]string {
	data := map[string]string{}

	data["doc1"] = "今日は、いい天気です。天気。"
	data["doc2"] = "今日は、悪い天気です。悪い。悪い。悪い。悪い。悪い。悪い。悪い。"
	data["doc3"] = "今日は、そこそこの天気です。天気。天気。天気。天気。天気。"
	data["doc4"] = "今日は、そこそこの天気みたいです。だね。"

	return data
}



func getResult(key string, data map[string]string) utils.ListFloat {
	tfidf := models.CreateTFIDF(
		data[key],
		data,
		key,
	)

	return utils.MorphNounSortFloat(tfidf.GetTFIDFList())
}
