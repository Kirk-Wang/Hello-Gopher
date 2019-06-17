package parser

import (
	"regexp"
	"strings"

	"github.com/Kirk-Wang/Hello-Gopher/re/14.5/crawler/engine"
)

const cityListRe = `{linkContent:"([^"征婚]+)",linkURL:"([^"]+)"}`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[1]))

		result.Requests = append(result.Requests, engine.Request{
			Url:        strings.Replace(string(m[2]), "m.", "www.", 1),
			ParserFunc: ParseCity})
	}
	return result
}
