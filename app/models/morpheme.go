package models

import (
	"github.com/yukihir0/mecab-go"
)

type Morpheme struct {
	MorphemeList           []*MorphemeShard
}
func (m *Morpheme) GetMorphemeList() []*MorphemeShard {
	return m.MorphemeList
}
func (m *Morpheme) SetMorphemeList(value []*MorphemeShard) {
	m.MorphemeList = value
}

func CreateMorpheme(str string) *Morpheme {
	resultList, err := mecab.Parse(str)

	if err != nil {
		panic(err)
	}

	morphemeList := []*MorphemeShard{}
	for _, parseResult := range resultList {
		morphemeShard := CreateMorphemeShard(parseResult)
		morphemeList = append(morphemeList, morphemeShard)
	}

	return &Morpheme{
		MorphemeList: morphemeList,
	}
}
