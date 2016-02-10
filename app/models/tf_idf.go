package models

import (
	"math"
)

type TFIDF struct {
	TF              *TF
	IDF             *IDF
	TFIDFList       map[string]float64
}
func (t *TFIDF) GetTF() *TF {
	return t.TF
}
func (t *TFIDF) SetTF(value *TF) {
	t.TF = value
}
func (t *TFIDF) GetIDF() *IDF {
	return t.IDF
}
func (t *TFIDF) SetIDF(value *IDF) {
	t.IDF = value
}
func (t *TFIDF) GetTFIDFList() map[string]float64 {
	return t.TFIDFList
}
func (t *TFIDF) SetTFIDFList(value map[string]float64) {
	t.TFIDFList = value
}

func CreateTFIDF(text string, allText map[string]string, target string) *TFIDF {
	tf := CreateTF(
		text,
	)

	idf := CreateIDF(
		allText,
	)

	return &TFIDF{
		TF: tf,
		IDF: idf,
		TFIDFList: createTFIDFList(tf, idf, target),
	}
}
func createTFIDFList(tfData *TF, idfData *IDF, target string) map[string]float64 {
	res := map[string]float64{}

	for term, tf := range tfData.GetTFList() {
		idfList := idfData.GetIDFList()
		idfTarget := idfList[target]
		idf := idfTarget[term]

		res[term] = tf * idf
	}

	return res
}


type TF struct {
	TFList          map[string]float64
}
func (t *TF) GetTFList() map[string]float64 {
	return t.TFList
}
func (t *TF) SetTFList(value map[string]float64) {
	t.TFList = value
}

func CreateTF(doc string) *TF {

	termData := createTermData(doc)

	totalTermNumber := len(CreateMorpheme(doc).GetMorphemeList())

	return &TF{
		TFList: createTermFrequency(termData, totalTermNumber),
	}
}

func createTermFrequency(termData map[string]int, totalTermNumber int) map[string]float64 {

	res := map[string]float64{}

	for term, num := range termData {
		res[term] = float64(num) / float64(totalTermNumber)
	}

	return res
}


// 単語と出現回数
func createTermData(doc string) map[string]int {

	morph := CreateMorpheme(doc)

	res := map[string]int{}

	for _, morphShard := range morph.GetMorphemeList() {
		surface := morphShard.GetSurface()
		_, keyExistInRes := res[surface]
		if keyExistInRes {
			res[surface] = res[surface] + 1
		} else {
			res[surface] = 1
		}
	}

	return res
}



type IDF struct {
	IDFList         map[string]map[string]float64
}
func (i *IDF) GetIDFList() map[string]map[string]float64 {
	return i.IDFList
}
func (i *IDF) SetIDFList(value map[string]map[string]float64) {
	i.IDFList = value
}

func CreateIDF( allDoc map[string]string) *IDF {

	df := createDocumentFrequency(allDoc)

	totalDocNumber := len(allDoc)

	return &IDF{
		IDFList: CreateInverseDocumentFrequency(df, totalDocNumber),
	}
}

// 単語と出現回数
func createDocumentFrequency(allDoc map[string]string) map[string]map[string]float64 {

	// 単語 = 1 の入ったリスト
	// morphList["Doc1"]["単語1"] = 1
	// morphList["Doc2"]["単語2"] = 1
	morphList:= map[string]map[string]int{}
	for name, doc := range allDoc {
		morphList[name] = CreateTermExistList(doc)
	}

	res := map[string]map[string]float64{}

	for name, existList := range morphList {
		// 最終結果はここに入れる
		data := map[string]float64{}

		for term, _ := range existList {
			data[term] = 0
			for _, eList := range morphList {
				_, existTerm := eList[term]
				if existTerm {
					data[term]++
				}
			}
		}
		res[name] = data
	}

	return res
}

func CreateInverseDocumentFrequency(df map[string]map[string]float64, totalDocNUmber int) map[string]map[string]float64 {

	res := map[string]map[string]float64{}

	for name, termDFList := range df {
		data := map[string]float64{}

		for term, frequency := range termDFList {
			// http://qiita.com/ynakayama/items/300460aa718363abc85c#3-3
			// log(N/df) + 1
			data[term] = math.Log(float64(totalDocNUmber) / frequency) + 1
		}

		res[name] = data
	}

	return res
}


// 単語が出てきたら 1が入ってるリスト
func CreateTermExistList(doc string) map[string]int {

	morph := CreateMorpheme(doc)

	res := map[string]int{}

	for _, morphShard := range morph.GetMorphemeList() {
		surface := morphShard.GetSurface()
		res[surface] = 1
	}

	return res
}
