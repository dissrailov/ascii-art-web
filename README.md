# ASCII-Art-Web

## Description
ASCII Art Web is a web application that allows users to create ASCII art using different banners such as shadow, standard, and thinkertoy. The application provides a graphical user interface (GUI) accessible through a web browser.

## Authors
- disrailov
- tsadvaka

## Usage: How to Run
1. Make sure you have Go installed on your system.
2. Clone this repository: `git clone git@git.01.alem.school:disrailov/ascii-art-web.git`
3. Navigate to the project directory: `cd ascii-art-web`
4. Run the server: `go run ./cmd`
5. Open your web browser and go to `http://localhost:8080` to access the application.

## Implementation Details: Algorithm
The application is implemented in Go and uses HTML templates for rendering the web pages. The server handles two HTTP endpoints:

### 1. GET /
- Renders the main page with a text input, radio buttons to switch between banners, and a button to send a POST request.
- Uses Go templates to receive and display data from the server.

### 2. POST /
- Accepts data (text and banner type) from the client using form tags.
- Processes the data and generates ASCII art based on the selected banner.
- The result can be displayed either on the `/generate` page or appended to the home page, depending on the implementation.

## Instructions
- The main page should have a text input for the ASCII art text, radio buttons (or select objects) to switch between banners, and a button to send a POST request.
- The application should handle HTTP status codes appropriately:
  - OK (200) for successful requests.
  - Not Found (404) for missing templates or banners.
  - Bad Request (400) for incorrect requests.
  - Internal Server Error (500) for unhandled errors.
