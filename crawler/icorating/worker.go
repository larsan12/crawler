package crawler

import (
	"fmt"
	"strings"

	model "myCrawler/model/icorating"
	"myCrawler/writer"

	"github.com/PuerkitoBio/goquery"
)

type ICORatingWorker struct {
	id       int
	finished bool
	pageType string
	links    []string
}

func (worker *ICORatingWorker) Start() error {
	for _, link := range worker.links {
		entity, _ := worker.GetDetails(link)
		outputPath := "./data/icorating/"
		outFilename := entity.Title + ".json"
		writer.WriteToFS(outputPath, outFilename, entity)
	}
	return nil
}

func (worker *ICORatingWorker) GetDetails(detailsLink string) (model.ICORatingCompany, error) {
	doc, err := goquery.NewDocument(detailsLink)
	fmt.Println(detailsLink)
	if err != nil {
		return model.ICORatingCompany{}, err
	}
	result := model.ICORatingCompany{}

	titleNode := doc.Find(".ico-main-info h3")
	if len(titleNode.Nodes) > 0 {
		result.Title = titleNode.First().Text()
	}

	category := doc.Find(".ico-main-info span.ico-category-name")
	if len(category.Nodes) > 0 {
		category.Contents().Each(func(i int, s *goquery.Selection) {
			if goquery.NodeName(s) == "#text" {
				result.Category = s.Text()
			}
		})
	}

	description := doc.Find("span.ico-category-name div.ico-description")
	if len(description.Nodes) > 0 {
		result.Description = strings.Trim(description.First().Text(), "\t\n")
	}

	sites := doc.Find(".ico-right-col a")
	if len(sites.Nodes) > 0 {
		sites.Each(func(i int, s *goquery.Selection) {
			value, exist := s.Attr("href")
			if exist {
				note := s.Find("div.button")
				if len(note.Nodes) > 0 {
					text := note.First().Text()
					if text == "WEBSITE" {
						result.Website = value
					} else if text == "WHITEPAPER" {
						result.Whitepaper = value
					}
				} else {
					result.Refs = append(result.Refs, value)
				}
			}
		})
	}

	if result.ICO == nil {
		result.ICO = make(map[string]string)
	}

	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		span := s.Find("span")
		if len(span.Nodes) > 0 {
			value := ""
			key := clearText(span.Text())
			s.Contents().Each(func(i int, s1 *goquery.Selection) {
				if goquery.NodeName(s1) == "#text" {
					value = s1.Text()
				}
			})
			if value != "" && key != "" && key != "Whitelist" {
				result.ICO[key] = clearText(value)
			}
		}
	})

	if result.Raitings == nil {
		result.Raitings = make(map[string]string)
	}

	doc.Find(".rating-item .rating-box").Each(func(i int, s *goquery.Selection) {
		if len(s.Nodes) > 0 {
			key := clearText(s.Find("p").First().Text())
			value := clearText(s.Find("p.rate").First().Text())
			if key != "" && value != "" {
				result.Raitings[key] = value
			}
		}
	})

	return result, nil
}

func clearText(input string) string {
	output := strings.Replace(input, "\n", "", -1)
	output = strings.Replace(output, ":", "", -1)
	output = strings.TrimSpace(output)
	return output
}

// func (worker *ICORatingWorker) GetNews(link string) (interface{}, error) {

// }

// func (worker *ICORatingWorker) GetReview(link string) (interface{}, error) {

// }
