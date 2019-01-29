package parser

import (
	"github.com/Kirk-Wang/Hello-Gopher/15.7/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/15.7/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)kg</div>`)

var incomeRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入:([^<]+)</div>`)

// var genderRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入:[^<]+</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>([^<]+车)</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)</div></div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>籍贯:([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>([^<]+房)</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)</div><div class="m-btn purple"[^>]*>[\d]+岁</div>`)
var occupationRe = regexp.MustCompile(`月收入:[^<]+</div><div class="m-btn purple"[^>]*>([^<]+)</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+座)[^<]+</div>`)

// var guessRe = regexp.MustCompile(``)

func ParseProfile(contents []byte, name string, gender string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = gender

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(contents, incomeRe)
	// profile.Gender = extractString(contents, genderRe)
	profile.Car = extractString(contents, carRe)
	profile.Education = extractString(contents, educationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.House = extractString(contents, houseRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents) // 只查找第一个

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
