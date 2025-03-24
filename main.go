package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

type Match struct {
	Email string `json:"email"`
	Team  string `json:"team"`
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func getMatches() ([]Match, error) {
	resp, err := http.Get("http://localhost:3000/matches")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var matches []Match
	if err := json.Unmarshal(body, &matches); err != nil {
		return nil, err
	}

	return matches, nil
}

func main() {
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	for {
		matches, err := getMatches()
		if err != nil {
			log.Fatalf("Error getting matches: %v", err)
		}

		scrape := func() {
			ctx, cancel := chromedp.NewContext(context.Background())
			defer cancel()

			var pageContent string
			url := "https://jogoshoje.com/"
			if err := chromedp.Run(ctx, chromedp.Tasks{
				chromedp.Navigate(url),
				chromedp.Sleep(2 * time.Second),
				chromedp.OuterHTML("html", &pageContent),
			}); err != nil {
				log.Fatal(err)
			}

			for _, match := range matches {
				log.Printf("Checking team: %s", match.Team)
				if strings.Contains(pageContent, match.Team) {
					log.Printf("Page contains team: %s", match.Team)
					log.Printf("Sending email to: %s", match.Email)
					if err := sendEmail(match.Email, match.Team); err != nil {
						log.Printf("Failed to send email to %s about team %s: %v", match.Email, match.Team, err)
					} else {
						log.Printf("Email sent to %s about team %s", match.Email, match.Team)
					}
				} else {
					log.Printf("Page does not contain team: %s", match.Team)
				}
			}
		}

		scrape()

		<-ticker.C
	}
}

func sendEmail(toEmail, team string) error {
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
	m.SetHeader("Subject", "Seu time irá jogar hoje")
	m.SetBody("text/plain", fmt.Sprintf("Seu time %s irá jogar hoje.", team))

	d := gomail.NewDialer(smtpHost, smtpPort, fromEmail, emailPassword)

	return d.DialAndSend(m)
}
