package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
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
		emailPassword := goDotEnvVariable("EMAIL_PASSWORD")
		dialerEmail := goDotEnvVariable("DIALER_EMAIL")
		fromEmail := goDotEnvVariable("FROM_EMAIL")
		toEmail := goDotEnvVariable("TO_EMAIL")
		fmt.Println(toEmail)

		scrape := func() {
			c := colly.NewCollector()
			c.OnResponse(func(r *colly.Response) {
				fmt.Println("Visited:", r.Request.URL)

			})

			c.OnError(func(r *colly.Response, err error) {
				fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "Error:", err)

			})
			var isLiverpoolPlaying bool
			c.OnHTML("div.FixtureTitle_name__Wirsw", func(e *colly.HTMLElement) {
				times := e.Text
				fmt.Println("Times que jogam hoje:", times)
				if strings.Contains(times, "Liverpool") {
					isLiverpoolPlaying = true
					fmt.Println("true")
				} else {
					fmt.Println("false")
				}
			})
			c.OnScraped(func(r *colly.Response) {
				if isLiverpoolPlaying {
					fmt.Println("O Liverpool está jogando hoje!")
					m := gomail.NewMessage()
					m.SetHeader("From", fromEmail)
					m.SetHeader("To", toEmail)
					m.SetHeader("Subject", "Liverpool!")
					m.SetBody("text/plain", "Hello, Liverpool is playing today!!")

					d := gomail.NewDialer("smtp.gmail.com", 587, dialerEmail, emailPassword)
					d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
					if err := d.DialAndSend(m); err != nil {
						panic(err)
					}
				} else {
					fmt.Println("O Liverpool não está jogando hoje!")

				}
			})
			if err := c.Visit(goDotEnvVariable("URL")); err != nil {
				log.Printf("Error visiting the URL: %s", err)
			}
		}


		scrape()
		time.Sleep(24 * time.Hour)

	}
}
