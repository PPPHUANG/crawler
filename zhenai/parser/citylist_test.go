package parser

import (
	"testing"
	"crawler/fetcher"
	"log"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err !=  nil {
		panic(err)
	}
	log.Printf("%s\n", contents)
}
