# Groupie-Tracker

## Created by:
1. bsaliyev (Baurzhan)
2. nsergaziy (Nurkhat)

## Overview
Groupie-Tracker is a web application built with Go that allows users to track musical artists and their details, such as members, creation date, first album, concert locations, concert dates, and relations. The application fetches artist data from an external API and presents it in a user-friendly interface.

## Usage

1. Clone the repository: git clone git@git.01.alem.school:bsaliyev/groupie-tracker.git
2. Navigate to the project directory: cd groupie-tracker
3. Start the server: `go run cmd/main.go`
4. Open your browser and navigate to http://localhost:8080 to use the application.

## Objective

The Groupie-Tracker project aims to:
1. Develop skills in web development using Go.
2. Learn to interact with external APIs and handle JSON data.
3. Practice creating and using templates for dynamic HTML content.
4. Implement robust error handling and input validation.

## Components

### Project Structure

The Groupie-Tracker project is structured as follows:

```
/Groupie-Tracker
├── /cmd
│   └── main.go
├── /handlers
│   ├── artist.go
│   ├── handlers.go
│   ├── main_test.go
│   ├── server.go
│   └── utils.go
├── /static
│   ├── /css
│   │   └── styles.css
├── /templates
│   ├── error.html
│   ├── artist.html
│   └── home.html
├── go.mod
└── README.md

```

### File Descriptions

- `main.go`: The entry point of the application, starting the server.
- `artist.go`: Contains the Artist struct, functions to fetch artist data from the API, and the ArtistHandler.
- `handlers.go`: Defines the Home handler for the home page.
- `server.go`: Initializes and starts the HTTP server, setting up routes and template parsing.
- `utils.go`: Provides utility functions, such as displayError, for error handling.


## Running Tests

To run the unit tests for the Groupie-Tracker project:

1. Navigate to the project root directory.
2. Run `go test ./handlers` in the terminal. This will execute tests defined in main_test.go, ensuring the handlers work as expected.