package utils

import (
	"sort"
)

type EntryFloat struct {
	name  string
	value float64
}
func (e *EntryFloat) GetName() string {
	return e.name
}
func (e *EntryFloat) GetValue() float64 {
	return e.value
}

type ListFloat []EntryFloat

func (l ListFloat) Len() int {
	return len(l)
}

func (l ListFloat) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ListFloat) Less(i, j int) bool {
	if l[i].value == l[j].value {
		return (l[i].name > l[j].name)
	} else {
		return (l[i].value > l[j].value)
	}
}

func MorphNounSortFloat(morphNounFrequency map[string]float64) ListFloat {
	a := ListFloat{}

	for k, v := range morphNounFrequency {
		e := EntryFloat{k, v}
		a = append(a, e)
	}

	sort.Sort(a)

	return a
}
