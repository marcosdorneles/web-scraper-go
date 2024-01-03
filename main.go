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

	emailPassword := goDotEnvVariable("EMAIL_PASSWORD")
	dialerEmail := goDotEnvVariable("DIALER_EMAIL")
	fromEmail := goDotEnvVariable("FROM_EMAIL")
	toEmail := goDotEnvVariable("TO_EMAIL")

	scrape := func() {
		c := colly.NewCollector()
		var isLiverpoolPlaying bool
		c.OnHTML("div.FixtureTitle_name__Wirsw", func(e *colly.HTMLElement) {
			times := e.Text
			fmt.Println("Times que jogam hoje:", times)
			if strings.Contains(times, "Liverpool") {
				isLiverpoolPlaying = true
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
		c.Visit(goDotEnvVariable("URL"))
	}

	now := time.Now()
	fmt.Println(now)
	sendTime := time.Date(now.Year(), now.Month(), now.Day(), 15, 40, 0, 0, now.Location())
	fmt.Print(sendTime)

	waitDuration := sendTime.Sub(now)
	if waitDuration < 0 {
		sendTime = sendTime.Add(24 * time.Hour)
		waitDuration = sendTime.Sub(now)
	}

	timer := time.NewTimer(waitDuration)
	<-timer.C

	scrape()
}
