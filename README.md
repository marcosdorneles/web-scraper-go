# Web Scraper and Email Notifier

This project is a Go application that takes a screenshot of a specified webpage and sends it via email. The application runs periodically (every 24 hours) and uses Docker for containerization.

## Prerequisites

- Docker
- Go (if you want to run the application outside Docker)

## Setup

### Environment Variables

Create a `.env` file in the root directory of the project with the following content:

```properties
EMAIL_PASSWORD=your-email-password
DIALER_EMAIL=your-email@example.com
FROM_EMAIL=your-email@example.com
TO_EMAIL=recipient-email@example.com
URL=https://example.com
SMTP_HOST=smtp.example.com
SMTP_PORT=587
```

Replace the placeholder values with your actual email credentials and the URL you want to capture.

Docker
Build the Docker Image:

Run the Docker Container:

Running Locally
If you prefer to run the application outside Docker, follow these steps:

Install Dependencies:

Run the Application:

Project Structure
main.go: The main application file that contains the logic for taking a screenshot and sending an email.
.env: Environment variables file containing email credentials and the target URL.
Dockerfile: Docker configuration file for containerizing the application.
How It Works
The application loads environment variables from the .env file.
It sets up a headless browser using chromedp to navigate to the specified URL and take a screenshot.
The screenshot is saved as screenshot.png.
The application sends the screenshot as an email attachment using the gomail package.
The process repeats every 24 hours.
Dependencies
chromedp: For headless browser automation.
gomail: For sending emails.
godotenv: For loading environment variables from the .env file.
Troubleshooting
Ensure that the .env file is correctly formatted and contains valid values.
If using Gmail, make sure "Less secure app access" is enabled for the account.
Check network connectivity and firewall settings if you encounter connection issues with the SMTP server.
License
This project is licensed under the MIT License.


### Summary:

- **Environment Variables**: Ensure the [.env](http://_vscodecontentref_/#%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fmarcosguilhermedornelesstaub%2FDocuments%2Fweb-scraper-go%2F.env%22%2C%22path%22%3A%22%2FUsers%2Fmarcosguilhermedornelesstaub%2FDocuments%2Fweb-scraper-go%2F.env%22%2C%22scheme%22%3A%22file%22%7D%7D) file contains the correct values.
- **Docker**: Build and run the Docker container using the provided commands.
- **Running Locally**: Install dependencies and run the application using Go commands.
- **Project Structure**: Describes the main components of the project.
- **How It Works**: Explains the workflow of the application.
- **Dependencies**: Lists the main dependencies used in the project.
- **Troubleshooting**: Provides tips for resolving common issues.

