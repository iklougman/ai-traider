package parser

import (
	"github.com/gocolly/colly"
)

type Parser struct {
	Url    string
	client *colly.Collector
}

func NewParser(url string) *Parser {
	return &Parser{
		Url:    url,
		client: colly.NewCollector(),
	}
}

func (p Parser) Parse() {

}
