package stations

import (
	"fmt"
	"regexp"

	"tobloggan/code/contracts"
)

type BaseURLRewriter struct {
	baseURL string
	pattern *regexp.Regexp
}

func NewBaseURLRewriter(baseURL string) contracts.Station {
	return &BaseURLRewriter{
		baseURL: fmt.Sprintf(`href="%s/$1`, baseURL),
		pattern: regexp.MustCompile(`href="/([^/])`),
	}
}

func (this *BaseURLRewriter) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.Page:
		input.Content = this.pattern.ReplaceAllString(input.Content, this.baseURL)
		output(input)
	default:
		output(input)
	}
}
