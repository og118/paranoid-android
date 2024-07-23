package utils

import (
	"fmt"
	"paranoid_android/types"
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	CLARKESWORLD_BASE_URL = "https://clarkesworldmagazine.com/"
)

func ScrapeArticle(c *colly.Collector, articleID string) (types.Article, error) {
	var issue, title, description, text, author string
	var errBool bool

	c.OnHTML(".content-section .issue-heading", func(e *colly.HTMLElement) {
		issue = e.Text
	})
	c.OnHTML(".content-section .story-title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.OnHTML(".content-section .story-description", func(e *colly.HTMLElement) {
		description = e.Text
	})

	c.OnHTML(".content-section .story-text", func(e *colly.HTMLElement) {
		html, err := e.DOM.Html()
		if err != nil {
			fmt.Println(err)
		}
		text = strings.Split(html, `<h3 class="about"> </h3>`)[0]
	})

	c.OnHTML(".content-section .story-author", func(e *colly.HTMLElement) {
		author = e.Text
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
		errBool = true
	})

	c.Visit(CLARKESWORLD_BASE_URL + articleID)
	c.Wait()
	fmt.Println("Scraped article")
	if errBool {
		return types.Article{}, fmt.Errorf("failed to scrape article")
	}
	return types.Article{
		Title:       title,
		Description: description,
		Author:      author,
		Text:        text,
		Issue:       issue,
	}, nil
}
