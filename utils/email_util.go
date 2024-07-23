package utils

import (
	"encoding/base64"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendMail(to string, epub []byte, bookName string) error {
	godotenv.Load()
	from := os.Getenv("USER_EMAIL")
	password := os.Getenv("USER_PASSWORD")

	subject := bookName
	body := "Export Request for " + bookName + " has been completed. Please find the attached file."

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: multipart/mixed; boundary=\"***BOUNDARY***\"\n" +
		"\n--***BOUNDARY***\n" +
		"Content-Type: text/plain; charset=\"utf-8\"\n" +
		"Content-Transfer-Encoding: 7bit\n" +
		"\n" + body + "\n" +
		"--***BOUNDARY***\n" +
		"Content-Type: application/epub+zip; name=\"" + bookName + "\"\n" +
		"Content-Transfer-Encoding: base64\n" +
		"Content-Disposition: attachment; filename=\"" + bookName + "\"\n" +
		"\n" + base64.StdEncoding.EncodeToString(epub) + "\n" +
		"--***BOUNDARY***--"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		return err
	}

	return nil
}
