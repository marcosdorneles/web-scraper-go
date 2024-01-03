package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	// "net/smtp"
	"strings"

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
	c := colly.NewCollector()

	var isLiverpoolPlaying bool

	c.OnHTML("div.FixtureTitle_name__Wirsw ", func(e *colly.HTMLElement) {
		times := e.Text
		fmt.Println("Times que jogam hoje:", times)

		// Verifica se a lista de times contém "Liverpool"
		if strings.Contains(times, "Liverpool") {
			isLiverpoolPlaying = true
		}
	})

	c.OnScraped(func(r *colly.Response) {
		if isLiverpoolPlaying {
			fmt.Println("O Liverpool está jogando hoje!")
			m := gomail.NewMessage()
			m.SetHeader("From", "FROM_EMAIL")
			m.SetHeader("To", "TO_EMAIL")
			m.SetHeader("Subject", "Hello!")
			m.SetBody("text/plain", "Hello, Liverpool is playing today!!")

			d := gomail.NewDialer("smtp.gmail.com", 587, "DIALER_EMAIL", "EMAIL_PASSWORD")

			d.TLSConfig = &tls.Config{InsecureSkipVerify: true} 

			if err := d.DialAndSend(m); err != nil {
				panic(err)
			}
		} else {
			fmt.Println("O Liverpool não está jogando hoje.")
		}
	})

	c.Visit("https://jogoshoje.com/")

}
