# URL Shortener

A simple URL shortener built with Go with allows users to input a long URL and receive a shortened form of it. They will be redirected to the original URL when visiting the shortened URL.

# Getting Started
## Prerequisites
Make sure that you have Go 1.16+ installed on your machine.

## Installation

1. Cloning the repositary:
   git clone https://github.com/Mijahla/URL-Shortener-go/tree/main
   cd URL-Shortener-go

2. Running the application:
   go build main.go
   go run main.go

3. Accessing the service:
   Open the web browser and navigate to:
   http://localhost:8080

4. Running the tests:
   In the project directory in terminal:
   cd URL-Shortener-go
   go test ./...

   This command will run all the test case in the main_test.go file and display the results in the terminal.
   
5. Using the URL Shortener
 - In the webpage, Enter a valid URL in the input field.
 - Click "Shorten URL" button.
 - You will receive a shortened URL which will redirect to the original URL.
