# Tem Jogo Hoje - Web Scraper Worker

This project is a web scraper written in Go for the "Tem Jogo Hoje" application. It checks the database for users' emails and favorite teams. If a user's favorite team is playing, the application sends an email notification to the user.

## Features

- Scrapes web data to check if a user's favorite team is playing.
- Sends email notifications to users if their favorite team has a match.

## Installation

1. Clone the repository:
  ```sh
  git clone https://github.com/yourusername/tem-jogo-hoje-worker.git
  ```
2. Navigate to the project directory:
  ```sh
  cd tem-jogo-hoje-worker
  ```
3. Install dependencies:
  ```sh
  go mod tidy
  ```

## Usage

1. Set up your database and email configurations in the `.env` file.
2. Run the application:
  ```sh
  go run main.go
  ```

