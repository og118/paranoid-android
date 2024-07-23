package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
)

var C *colly.Collector

func Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	C = colly.NewCollector(
		colly.AllowedDomains("clarkesworldmagazine.com"),
	)
}
