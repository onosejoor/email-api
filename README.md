**Email API Service: Multi-Language Gateway** 📧

A robust and scalable Email API service designed to streamline transactional email sending. ✨ This project offers dual implementations in Go and Node.js, demonstrating language-agnostic backend development for efficient and reliable email dispatch. Perfect for integrating email capabilities into any application! 🚀

---

## 📝 Table of Contents

- [✨ Features](#-features)
- [🛠️ Technologies Used](#%F0%9F%9B%A0%EF%B8%8F-technologies-used)
- [🚀 Getting Started](#-getting-started)
  - [Go Server](#go-server)
    - [Installation](#installation)
    - [Environment Variables](#environment-variables)
    - [Running the Server](#running-the-server)
  - [Node.js Server](#nodejs-server)
    - [Installation](#installation-1)
    - [Environment Variables](#environment-variables-1)
    - [Running the server](#running-the-server-1)
- [📄 API Documentation (Go Server)](#-api-documentation-go-server)
- [📄 API Documentation (Node.js Server)](#-api-documentation-nodejs-server)
- [💡 Usage Example](#-usage-example)
- [🤝 Contributing](#-contributing)
- [👤 Author](#-author)
- [📜 License](#-license)
- [🏅 Badges](#-badges)

---

## ✨ Features

- **Dual Language Support**: Backend implementations available in both Go and Node.js for flexibility and diverse project needs.
- **Transactional Email Dispatch**: Reliably sends emails via a configured SMTP service, ideal for notifications, confirmations, and alerts.
- **RESTful API**: Clean and intuitive endpoints for easy integration into web and mobile applications.
- **Environment Configuration**: Secure handling of sensitive credentials through environment variables.
- **Vercel Optimized**: Project structure is suitable for deployment as serverless functions on platforms like Vercel.

---

## 🛠️ Technologies Used

| Technology             | Description                                                               |
| :--------------------- | :------------------------------------------------------------------------ |
| **Go**                 | High-performance, concurrent programming language                         |
| **Node.js**            | JavaScript runtime for server-side applications                           |
| **Express.js**         | Web framework for Node.js (implicit for Vercel functions)                 |
| **`gopkg.in/mail.v2`** | Email sending package for Go                                              |
| **Nodemailer**         | Module for Node.js applications to send emails                            |
| **`dotenv`**           | Loads environment variables from a `.env` file                            |
| **Vercel**             | Platform for frontend frameworks and static sites (serverless deployment) |

---

## 🚀 Getting Started

Follow these steps to set up and run the email API services locally.

### Go Server

#### Installation

1.  **Clone the Repository**:
    ```bash
    git clone https://github.com/onosejoor/email-api.git
    cd email-api/go-server
    ```
2.  **Initialize Go Modules**:
    ```bash
    go mod tidy
    ```

#### Environment Variables

Create a `.env` file in the `go-server` directory with the following variables:

```env
GMAIL_USER=your_gmail_address@gmail.com
GMAIL_APP_PASSWORD=your_gmail_app_password
```

_Note_: For `GMAIL_APP_PASSWORD`, you'll need to generate an App Password from your Google Account settings if you have 2-Step Verification enabled.

#### Running the Server

The Go server files (`api/index.go`, `api/send-email.go`) are designed as HTTP handlers, typically for a serverless environment or to be mounted on a router. For local testing, you would typically have a `main.go` file to run these handlers. Since a `main.go` is not provided in the snippet, we assume these are deployed as Vercel functions. If running locally, you'd integrate them into a custom HTTP server:

```go
// Example main.go for local testing (not part of the provided code)
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	handler "github.com/onosejoor/email-api/api" // Adjust path based on your module setup
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/api", handler.HomeHandler)
	http.HandleFunc("/api/send-email", handler.Handler) // assuming handler.Handler is the send-email handler

	log.Println("Go server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

To run this example `main.go` (if you create it):

```bash
go run main.go
```

The server will then be accessible at `http://localhost:8080`.

### Node.js Server

#### Installation

1.  **Clone the Repository**:
    ```bash
    git clone https://github.com/onosejoor/email-api.git
    cd email-api/node-server
    ```
2.  **Install Dependencies**:
    ```bash
    npm install
    ```

#### Environment Variables

Create a `.env` file in the `node-server` directory with the following variables:

```env
GMAIL_USER=your_gmail_address@gmail.com
GMAIL_APP_PASSWORD=your_gmail_app_password
```

_Note_: As with the Go server, `GMAIL_APP_PASSWORD` requires an App Password from Google Account settings.

#### Running the server

The Node.js server files (`api/index.js`, `api/send-email.js`) are designed as Vercel serverless functions.
To run the Node.js server locally (which primarily uses `index.js` for `npm run dev`):

```bash
npm run dev
```

The server will typically run on `http://localhost:3000` or a similar port.

---

## 📄 API Documentation (Go Server)

### Base URL

- **Local Development**: `http://localhost:8080`
- **Vercel Deployment**: `https://your-go-deployment.vercel.app` (if deployed as Vercel functions)

### Endpoints

#### GET /api

**Overview**: Provides a basic health check and welcome message.
**Request**:
No payload required.

**Response**:

```json
{
  "success": true,
  "message": "Hello World!"
}
```

**Errors**:

- `405 Method Not Allowed`: If the request method is not GET.

#### POST /api/send-email

**Overview**: Dispatches an email using the configured SMTP server.
**Request**:

```json
{
  "to": "recipient@example.com",
  "subject": "Important Update",
  "html": "<p>Hello there, this is an <b>HTML formatted</b> email from your Go API!</p>",
  "from": "Your App Name"
}
```

**Response**:

```json
{
  "success": true,
  "message": "Email sent successfully"
}
```

**Errors**:

- `405 Method Not Allowed`: If the request method is not POST.
- `400 Bad Request`: If the JSON body is invalid, or missing required fields (`to`, `subject`, `html`, `from`).
- `500 Internal Server Error`: If there is an issue sending the email (e.g., SMTP failure, invalid credentials, network issues).

---

## 📄 API Documentation (Node.js Server)

### Base URL

- **Local Development**: `http://localhost:3000`
- **Vercel Deployment**: `https://your-node-deployment.vercel.app` (if deployed as Vercel functions)

### Endpoints

#### GET /api

**Overview**: Provides a basic health check and welcome message.
**Request**:
No payload required.

**Response**:

```json
{
  "message": "Hello from Node.js on Vercel!"
}
```

**Errors**:
_(The Node.js `index.js` handler does not explicitly check for HTTP methods, so it will respond to any method with a 200 OK.)_

#### POST /api/send-email

**Overview**: Dispatches an email using the configured SMTP server.
**Request**:

```json
{
  "to": "recipient@example.com",
  "subject": "Important Update",
  "text": "Hello there, this is a plain text email from your Node.js API!",
  "html": "<p>Hello there, this is an <b>HTML formatted</b> email from your Node.js API!</p>"
}
```

**Note**: Either `text` or `html` (or both) can be provided. The `from` field in the request body is ignored; the sender is fixed as `"Zendo" <GMAIL_USER>`.

**Response**:

```json
{
  "success": true,
  "message": "Email sent successfully"
}
```

**Errors**:

- `405 Method Not Allowed`: If the request method is not POST.
- `400 Bad Request`: If missing required fields (`to`, `subject`, and either `text` or `html`).
- `500 Internal Server Error`: If there is an issue sending the email (e.g., SMTP failure, invalid credentials, network issues).

---

## 💡 Usage Example

Below are cURL examples demonstrating how to interact with the `POST /api/send-email` endpoint for both Go and Node.js servers.

### Go Server Example (assuming local server at `http://localhost:8080`)

```bash
curl -X POST \
     -H "Content-Type: application/json" \
     -d '{
           "to": "test@example.com",
           "subject": "Hello from Go!",
           "html": "<p>This is a test email sent from the <b>Go backend</b>.</p>",
           "from": "My Go App"
         }' \
     http://localhost:8080/api/send-email
```

### Node.js Server Example (assuming local server at `http://localhost:3000`)

```bash
curl -X POST \
     -H "Content-Type: application/json" \
     -d '{
           "to": "test@example.com",
           "subject": "Hello from Node.js!",
           "html": "<p>This is a test email sent from the <b>Node.js backend</b>.</p>"
         }' \
     http://localhost:3000/api/send-email
```

---

## 🤝 Contributing

We welcome contributions to enhance this project! To contribute:

- 🍴 Fork the repository.
- 🌿 Create a new branch (`git checkout -b feature/your-feature`).
- 💡 Implement your changes.
- ✅ Ensure your code passes any existing tests and adheres to coding standards.
- 💬 Commit your changes with a clear message (`git commit -m 'feat: Add new feature'`).
- ⬆️ Push to your branch (`git push origin feature/your-feature`).
- 📝 Open a pull request.

---

## 👤 Author

**Onos Ejoot**

- Twitter: [@DevText16]
- Portfolio: [https://onos-ejoor.vercel.app]

---

## 📜 License

This project does not currently have an explicit license file. Please refer to the repository owner for licensing information.

---

## 🏅 Badges

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Node.js](https://img.shields.io/badge/Node.js-339933?style=for-the-badge&logo=nodedotjs&logoColor=white)
![Express.js](https://img.shields.io/badge/Express.js-000000?style=for-the-badge&logo=express&logoColor=white)
![Nodemailer](https://img.shields.io/badge/Nodemailer-007bff?style=for-the-badge&logo=nodemailer&logoColor=white)
![Vercel](https://img.shields.io/badge/Vercel-000000?style=for-the-badge&logo=vercel&logoColor=white)

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)
