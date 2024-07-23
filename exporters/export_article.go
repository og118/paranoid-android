package exporters

import (
	"bytes"
	"fmt"
	"paranoid_android/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-shiori/go-epub"
	"github.com/gocolly/colly/v2"
)

func ExportIssueArticle(g *gin.Context, C *colly.Collector) {
	articleID := g.Param("article_id")
	mailTo := g.Query("mail_to")

	article, err := utils.ScrapeArticle(C, articleID)
	if err != nil {
		fmt.Println(err)
		g.JSON(500, gin.H{"error": "Failed to scrape article"})
		g.Abort()
	}

	e, err := epub.NewEpub(article.Title)
	e.SetAuthor(article.Author)
	if err != nil {
		fmt.Println(err)
	}

	issueNo := strings.FieldsFunc(strings.TrimSpace(article.Issue), func(r rune) bool {
		return r == '–' || r == ' '
	})[1]

	cover, err := e.AddImage(fmt.Sprintf(CLARKESWORLD_BASE_URL+"covers/cw_%s_800.jpg", issueNo), "cover.jpg")
	if err != nil {
		fmt.Println(err)
	}
	e.SetCover(cover, "")

	css, err := e.AddCSS("./styles.css", "")
	if err != nil {
		fmt.Println(err)
	}

	article.Description = fmt.Sprintf("<h1>%s</h1><i>%s (Clarkesworld Magazine %s)</i>",
		article.Title, article.Author, article.Issue)
	e.AddSection(article.Description+"<br/><br/>"+article.Text, "", "", css)

	var buf bytes.Buffer
	_, err = e.WriteTo(&buf)
	if err != nil {
		fmt.Println(err)
		g.JSON(500, gin.H{"error": "Failed to create epub"})
		g.Abort()
	}

	err = utils.SendMail(mailTo,
		buf.Bytes(),
		strings.Join(strings.Split(article.Title+" "+article.Author, " "), "_")+".epub")
	if err != nil {
		fmt.Println(err)
		g.JSON(400, gin.H{"error": "Failed to send mail"})
		g.Abort()
	}
	g.JSON(200, gin.H{"message": "Article sent to your mail"})

}
