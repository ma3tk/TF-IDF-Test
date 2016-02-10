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
		// TFIDF
		tfidfResult := getTFIDFResult(key, data)
		output(key, "TFIDF", tfidfResult)

		// TF
		tfResult := getTFResult(key, data)
		output(key, "TF", tfResult)

		// IDF
		idfResult := getIDFResult(key, data)
		output(key, "IDF", idfResult)
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

func output(key string, typeName string, result utils.ListFloat) {
	// 出力
	fmt.Println(key, ",", typeName, "========================")

	for _, val := range result {
		fmt.Println(val.GetName(), ",", val.GetValue())
	}

	fmt.Println("")
}



func getTFIDFResult(key string, data map[string]string) utils.ListFloat {
	tfidf := models.CreateTFIDF(
		data[key],
		data,
		key,
	)

	return utils.MorphNounSortFloat(
		tfidf.GetTFIDFList(),
	)
}

func getTFResult(key string, data map[string]string) utils.ListFloat {
	tfidf := models.CreateTFIDF(
		data[key],
		data,
		key,
	)

	return utils.MorphNounSortFloat(
		tfidf.GetTF().GetTFList(),
	)
}

func getIDFResult(key string, data map[string]string) utils.ListFloat {
	tfidf := models.CreateTFIDF(
		data[key],
		data,
		key,
	)

	return utils.MorphNounSortFloat(
		tfidf.GetIDF().GetIDFListByKey(key),
	)
}


