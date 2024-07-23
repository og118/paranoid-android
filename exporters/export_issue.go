package exporters

import (
	"bytes"
	"fmt"
	"paranoid_android/types"
	"paranoid_android/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-shiori/go-epub"
	"github.com/gocolly/colly/v2"
)

const (
	CLARKESWORLD_BASE_URL = "https://clarkesworldmagazine.com/"
)

func ExportIssue(c *gin.Context, C *colly.Collector) {
	mailTo := c.Query("mail_to")
	issueID := c.Param("issue_id")
	// Handle the /issue/{issue_id} route here
	C.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	C.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
	})

	var article_ids []string = make([]string, 0)
	var month, year string
	C.OnHTML(".issue", func(e *colly.HTMLElement) {
		// Find and visit all links
		issue := e.Text
		month_year := strings.Split(strings.TrimSpace(strings.Split(issue, "–")[1]), " ")

		month = month_year[0]
		switch month {
		case "January":
			month = "01"
		case "February":
			month = "02"
		case "March":
			month = "03"
		case "April":
			month = "04"
		case "May":
			month = "05"
		case "June":
			month = "06"
		case "July":
			month = "07"
		case "August":
			month = "08"
		case "September":
			month = "09"
		case "October":
			month = "10"
		case "November":
			month = "11"
		case "December":
			month = "12"
		}
		year = month_year[1][2:]
		fmt.Println("Month:", month, "Year:", year)
	})
	C.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasSuffix(link, "_"+month+"_"+year) {
			articleID := strings.Split(link, "/")[3]
			article_ids = append(article_ids, articleID)
		}
	})

	C.Visit(CLARKESWORLD_BASE_URL + "/prior/issue_" + issueID)
	C.Wait()
	var articles []types.Article = make([]types.Article, 0)
	for _, article_id := range article_ids {
		article, err := utils.ScrapeArticle(C, article_id)
		if err != nil {
			fmt.Println("Failed to scrape article", article_id)
		} else {
			articles = append(articles, article)
		}
	}

	e, err := epub.NewEpub("Clarkesworld Magazine: Issue " + issueID)
	if err != nil {
		fmt.Println(err)
	}

	cover, err := e.AddImage(fmt.Sprintf(CLARKESWORLD_BASE_URL+"covers/cw_%s_800.jpg", issueID), "cover.jpg")
	if err != nil {
		fmt.Println(err)
	}
	e.SetCover(cover, "")

	css, err := e.AddCSS("./styles.css", "")
	if err != nil {
		fmt.Println(err)
	}

	for _, article := range articles {
		article.Description = fmt.Sprintf("<h1>%s</h1><i>%s (Clarkesworld Magazine %s)</i>",
			article.Title, article.Author, article.Issue)
		e.AddSection(article.Description+"<br/><br/>"+article.Text, article.Title, "", css)
	}

	var buf bytes.Buffer
	_, err = e.WriteTo(&buf)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "Failed to create epub"})
		c.Abort()
	}
	utils.SendMail(mailTo, buf.Bytes(), "Clarkesworld Magazine: Issue "+issueID+".epub")
}
