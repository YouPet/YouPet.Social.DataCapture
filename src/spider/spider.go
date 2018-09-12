package spider

import (
	"errors"
)

type Blog struct {
	Title    string
	Content  string
	Url      string
	ImageUrl string
	Tags     string
}

type Spider interface {
	SpiderUrl(url string) ([]Blog, error)
}

func NewSpider(from string) (Spider, error) {
	switch from {
	case "chongbaba":
		return new(ChongbabaSpider), nil
	default:
		return nil, errors.New("No support")
	}
}
