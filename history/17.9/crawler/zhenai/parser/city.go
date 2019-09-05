package parser

import (
	"encoding/json"
	"regexp"

	"github.com/Kirk-Wang/Hello-Gopher/history/17.9/crawler/config"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.9/crawler/engine"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)

var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(string(m[2]), string(m[3])),
		})
	}

	// 下一页
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}

	return result
}

type ProfileParser struct {
	userName   string
	userGender string
}

type ProfileParams struct {
	Name   string
	Gender string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName, p.userGender)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	params := ProfileParams{
		Name:   p.userName,
		Gender: p.userGender,
	}
	s, _ := json.Marshal(params)
	// return "ProfileParser", string(s)
	return config.ParseProfile, string(s)
}

func NewProfileParser(name string, gender string) *ProfileParser {
	return &ProfileParser{
		userName:   name,
		userGender: gender,
	}
}
