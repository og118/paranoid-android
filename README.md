# paranoid-android

A [Clarkesworld Magazine](https://clarkesworldmagazine.com/) scraper and exporter exposed as a Go-Gin based RESTful API server. Currently only capable of exporting EPUBs via email.

## How to run the project locally?
1. Install Golang
2. Install Dependencies
```bash
go mod tidy
```
3. Run the server
```bash
go run .
```

## Routes
- `/issue/{issue_id}`: exports all articles of a magazine issue with the given `id`, which is the issue number.
- `issue/{issue_id}/article/{article_id}`: exports a specific article with given `id`

## Emails
- Check spam
- Create [Application Password](https://myaccount.google.com/apppasswords) for your Google Account for sending using gmail-smtp

## Contributing
Feel free to contribute, suggest and discuss ideas, improvements and report bugs ofc.
