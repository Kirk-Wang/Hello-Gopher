package parser

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/15.2/crawler/engine"
	"regexp"
)

// const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]+)</td>`

// <a href="http://album.zhenai.com/u/1320662004" target="_blank">微微一笑</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>男士</td>
// <a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]+)</td>
func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		gender := string(m[3])
		result.Items = append(result.Items, "User "+name+" Gender "+gender)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			// ParserFunc: engine.NilParser,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name, gender)
			},
		})
	}
	return result
}
