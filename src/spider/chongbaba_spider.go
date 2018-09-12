package spider

import (
	"bwkernel/youpet/youpet.social.datacapture/src/common"
	"github.com/PuerkitoBio/goquery"
)

type ChongbabaSpider struct {
}

func (self *ChongbabaSpider) SpiderUrl(url string) ([]Blog, error) {
	var blogs []Blog
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	doc.Find(".deanartice li").Each(func(i int, contentSelection *goquery.Selection) {
		blog := Blog{}
		blog.Title = common.GbkToUtf8(contentSelection.Find(".deanarttitles").Text())
		blog.Content = contentSelection.Find(".deanarticersummary").Text()
		blog.ImageUrl, _ = contentSelection.Find(".deanartice img").Attr("src")
		blog.Url, _ = contentSelection.Find(".deanarttitles").Attr("href")

		blogs = append(blogs, blog)
	})

	return blogs, nil
}
