package parser

import (
	"regexp"

	"github.com/Kirk-Wang/Hello-Gopher/re/14.6/crawler/model"

	"github.com/Kirk-Wang/Hello-Gopher/re/14.6/crawler/engine"
)

// "basicInfo":["未婚","28岁","射手座(11.22-12.21)","173cm","工作地:阿坝茂县","月收入:3-5千","高中及以下"]
var basicInfoRe = regexp.MustCompile(`"basicInfo":([^\]]+])`)
var fieldRe = regexp.MustCompile(`[^\[",\]]+`)
var genderRe = regexp.MustCompile(`"genderString":"([^"]+)"`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	basicInfo := basicInfoRe.FindSubmatch(contents)
	basicFields := fieldRe.FindAll(basicInfo[1], -1)
	gender := genderRe.FindSubmatch(contents)

	profile.Name = name
	profile.Gender = string(gender[1])
	profile.Marriage = string(basicFields[0])
	profile.Age = string(basicFields[1])
	profile.Xinzuo = string(basicFields[2])
	profile.Height = string(basicFields[3])

	return engine.ParseResult{
		Items: []interface{}{profile},
	}
}
