package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	for {
		toEmail := goDotEnvVariable("TO_EMAIL")
		fmt.Println(toEmail)

		scrape := func() {
			ctx, cancel := chromedp.NewContext(context.Background())
			defer cancel()

			var buf []byte
			url := goDotEnvVariable("URL")
			if err := chromedp.Run(ctx, fullScreenshot(url, 90, &buf)); err != nil {
				log.Fatal(err)
			}

			// Save the screenshot to a file
			screenshotPath := "screenshot.png"
			if err := os.WriteFile(screenshotPath, buf, 0644); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Screenshot saved as screenshot.png")

			if err := sendEmail(toEmail, screenshotPath); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Screenshot sent to email")
		}

		scrape()
		time.Sleep(24 * time.Hour)
	}
}

func fullScreenshot(url string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(2 * time.Second),
		chromedp.FullScreenshot(res, quality),
	}
}

func sendEmail(toEmail, attachmentPath string) error {
	fromEmail := goDotEnvVariable("FROM_EMAIL")
	emailPassword := goDotEnvVariable("EMAIL_PASSWORD")
	smtpHost := goDotEnvVariable("SMTP_HOST")
	smtpPortStr := goDotEnvVariable("SMTP_PORT")
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatalf("Invalid SMTP_PORT: %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Daily Screenshot")
	m.SetBody("text/plain", "Please find the attached screenshot.")
	m.Attach(attachmentPath)

	d := gomail.NewDialer(smtpHost, smtpPort, fromEmail, emailPassword)

	return d.DialAndSend(m)
}
