package main

import (
	crawler "myCrawler/crawler/icorating"
	"myCrawler/misc"
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
