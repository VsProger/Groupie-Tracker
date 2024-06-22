# Groupie Trackers

Groupie Trackers is a project that uses a provided API to create a user-friendly website displaying information about various bands and artists. The project focuses on data visualization and client-server interactions, utilizing Go for the backend development.

## Table of Contents

- [Objectives](#objectives)
- [API Structure](#api-structure)
- [Project Requirements](#project-requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Testing](#testing)

## Objectives

The main objectives of this project are:
- To create a website that displays band and artist information using various data visualization techniques.
- To implement client-server interactions for dynamic data retrieval and display.
- To ensure the backend, written in Go, is stable and reliable.
- To follow best coding practices and include unit tests for code reliability.

## API Structure

The provided API consists of four main parts:

1. **Artists**: Contains information about bands and artists, including:
   - Names
   - Images
   - Year they began their activity
   - Date of their first album
   - Band members

2. **Locations**: Contains information about the locations of past and upcoming concerts.

3. **Dates**: Contains information about the dates of past and upcoming concerts.

4. **Relation**: Links the artists, dates, and locations components, integrating the data across the API.

## Project Requirements

- **Data Visualization**: Display band and artist information using blocks, cards, tables, lists, pages, graphics, etc.
- **Client-Server Interaction**: Implement a feature that triggers client-server communication, sending a request from the client to the server and receiving a response.
- **Backend in Go**: Develop a stable backend using Go, ensuring no crashes and correct functionality of all pages.
- **Best Practices**: Follow coding best practices and include unit tests.

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/VsProger/Groupie-Tracker.git
   cd groupie-trackers
   ```
2. Install dependencies:
    ```
    go mod tidy
    ```
3. Start the server:
   ```
   go run main.go
   ```
## Usage

To use Groupie Trackers, follow these steps:

1. **Access the Website**: Open your web browser and navigate to [http://localhost:8080](http://localhost:8080).

2. **Explore Band Information**: Browse through the displayed information about various bands and artists.

3. **Interact with the Website**: Trigger client-server interactions by interacting with the website's interface to view dynamic updates and data visualizations.

## Features

- **Dynamic Data Visualization**: Display artist and concert information using various visual formats such as blocks, cards, tables, lists, and graphs.
  
- **Client-Server Interaction**: Implement actions that communicate with the server to fetch and display data in response to user interactions.
  
- **Robust Backend**: Utilize Go for backend development to ensure stability, reliability, and efficient handling of data requests.
  
- **Error Handling**: Gracefully manage errors to provide a seamless user experience and prevent crashes.

## Testing

To run unit tests for Groupie Trackers:

1. Ensure you have Go installed on your system.

2. Navigate to the project directory in your terminal.

3. Run the following command to execute all unit tests:

   ```
   go test ./...
   ```
