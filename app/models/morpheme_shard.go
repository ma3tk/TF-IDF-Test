package models

import (
	"github.com/yukihir0/mecab-go"
)

type MorphemeShard struct {
	mecab.ParseResult
}
func (m *MorphemeShard) GetSurface() string {
	return m.Surface
}
func (m *MorphemeShard) SetSurface(value string) {
	m.Surface = value
}
func (m *MorphemeShard) GetFeature() string {
	return m.Feature
}
func (m *MorphemeShard) SetFeature(value string) {
	m.Feature = value
}
func (m *MorphemeShard) GetPos() string {
	return m.Pos
}
func (m *MorphemeShard) SetPos(value string) {
	m.Pos = value
}
func (m *MorphemeShard) GetPos1() string {
	return m.Pos1
}
func (m *MorphemeShard) SetPos1(value string) {
	m.Pos1 = value
}
func (m *MorphemeShard) GetPos2() string {
	return m.Pos2
}
func (m *MorphemeShard) SetPos2(value string) {
	m.Pos2 = value
}
func (m *MorphemeShard) GetPos3() string {
	return m.Pos3
}
func (m *MorphemeShard) SetPos3(value string) {
	m.Pos3 = value
}
func (m *MorphemeShard) GetCform() string {
	return m.Cform
}
func (m *MorphemeShard) SetCform(value string) {
	m.Cform = value
}
func (m *MorphemeShard) GetCtype() string {
	return m.Ctype
}
func (m *MorphemeShard) SetCtype(value string) {
	m.Ctype = value
}
func (m *MorphemeShard) GetBase() string {
	return m.Base
}
func (m *MorphemeShard) SetBase(value string) {
	m.Base = value
}
func (m *MorphemeShard) GetRead() string {
	return m.Read
}
func (m *MorphemeShard) SetRead(value string) {
	m.Read = value
}
func (m *MorphemeShard) GetPron() string {
	return m.Pron
}
func (m *MorphemeShard) SetPron(value string) {
	m.Pron = value
}

func CreateMorphemeShard(result mecab.ParseResult) *MorphemeShard {
	return &MorphemeShard{
		result,
	}
}


