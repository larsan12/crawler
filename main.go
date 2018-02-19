package main

import (
	crawler "github.com/visheratin/ico-crawler/crawler/icorating"
	"github.com/visheratin/ico-crawler/misc"
)

func main() {
	misc.InitLog()
	config := misc.ReadConfig("config.json")
	manager := crawler.ICORatingCrawler{}
	err := manager.Init(config)
	if err != nil {
		misc.LogError(err)
	}
}
