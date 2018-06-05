package crawler

import (
	"time"

	"myCrawler/misc"

	"github.com/PuerkitoBio/goquery"
)

var mainLink = "https://icodrops.com/category/ended-ico/"

type ICORatingCrawler struct {
	workers []*ICORatingWorker
}

func (manager *ICORatingCrawler) Init(config misc.Configuration) error {
	links, err := manager.GetEntitiesLinks(mainLink)
	if err != nil {
		return err
	}
	workersLinks, err := misc.SplitLinks(links, config.WorkersNumber)
	if err != nil {
		return err
	}
	for i := 0; i < config.WorkersNumber; i++ {
		worker := &ICORatingWorker{
			id:    i,
			links: workersLinks[i],
		}
		manager.workers = append(manager.workers, worker)
		go func() {
			worker.Start()
		}()
	}
	timeout, err := time.ParseDuration(config.UpdateTimeout)
	if err != nil {
		timeout, _ = time.ParseDuration("5m")
	}
	for {
		time.Sleep(timeout)
		workersFinished := true
		for _, worker := range manager.workers {
			if !worker.finished {
				workersFinished = false
				break
			}
		}
		if workersFinished {
			break
		}
	}
	return nil
}

func (crawler *ICORatingCrawler) GetEntitiesLinks(mainPageLink string) ([]string, error) {
	doc, err := goquery.NewDocument(mainPageLink)
	if err != nil {
		return nil, err
	}
	result := []string{}
	doc.Find("div.col-md-12").Each(func(i int, s *goquery.Selection) {
		aBlock := s.Find("a")
		href, found := aBlock.Attr("href")
		if found {
			result = append(result, href)
		}
	})
	return result, nil
}
