✉️ **Robust Email API Service**

This project provides a versatile and efficient API for sending emails, implemented in both Go and Node.js. It offers a standardized interface to facilitate email delivery for various applications, featuring rate limiting and API key authentication for secure and controlled access. The Go implementation showcases two distinct approaches to credential management, demonstrating flexibility in deployment scenarios.

## Features

-   **Go-based Email API**: A high-performance email sending service built with Go, leveraging `gopkg.in/mail.v2` for reliable SMTP communication.
-   **Node.js Email API**: An alternative implementation using Node.js and `Nodemailer` for easy integration into JavaScript ecosystems.
-   **Rate Limiting**: Protects against abuse and ensures stable service by limiting the number of requests per client.
-   **API Key Authentication**: Secures endpoints by requiring a valid `X-API-KEY` header for all email sending requests.
-   **Flexible Credential Management**: The Go version offers a `go-server` that fetches SMTP credentials from environment variables and a `go-server-without-creds` that accepts them directly in the request payload.
-   **Modular Mailer Service**: A dedicated `pkg/mailer` module encapsulates email sending logic, promoting reusability and clean architecture.

## Getting Started

To get a copy of this project up and running on your local machine, follow these steps.

### Installation

First, clone the repository:

```bash
git clone https://github.com/onosejoor/email-api.git
cd email-api
```

There are three distinct server implementations within this repository: `go-server`, `go-server-without-creds`, and `node-server`.

#### Go Server (Credential Management via Environment Variables)

1.  Navigate to the `go-server` directory:
    ```bash
    cd go-server
    ```
2.  Install Go dependencies:
    ```bash
    go mod tidy
    ```

#### Go Server (Credential Management via Request Body)

1.  Navigate to the `go-server-without-creds` directory:
    ```bash
    cd go-server-without-creds
    ```
2.  Install Go dependencies:
    ```bash
    go mod tidy
    ```

#### Node.js Server

1.  Navigate to the `node-server` directory:
    ```bash
    cd node-server
    ```
2.  Install Node.js dependencies:
    ```bash
    npm install
    ```

### Environment Variables

Each server requires specific environment variables to function correctly. Create a `.env` file in the root directory of each respective server (`go-server`, `go-server-without-creds`, `node-server`).

For `go-server` and `node-server`:

```env
EMAIL_API_TOKEN=your_strong_api_key_here
GMAIL_USER=your_gmail_address@gmail.com
GMAIL_APP_PASSWORD=your_gmail_app_password
```

For `go-server-without-creds`:

```env
EMAIL_API_TOKEN=your_strong_api_key_here
# GMAIL_USER and GMAIL_APP_PASSWORD are provided in the request body for this server
```

**Note**: A Gmail App Password is required if you have 2-Factor Authentication enabled on your Google account. You can generate one from your Google Account security settings.

## Usage

Each server offers endpoints for sending emails.

### Starting the Servers

#### Go Server (with environment credentials)

Navigate to the `go-server` directory and run:

```bash
go run api/send-email.go # This will typically run on http://localhost:8080 or a similar port
```

#### Go Server (with request body credentials)

Navigate to the `go-server-without-creds` directory and run:

```bash
go run api/send-email.go # This will typically run on http://localhost:8081 or a similar port
```

#### Node.js Server

Navigate to the `node-server` directory and run:

```bash
npm run dev # For development, using nodemon
# or
npm start # For production
```

By default, the Node.js server will run on `http://localhost:3000` (or Vercel for serverless deployment).

### API Documentation

#### Go Server Endpoints (with environment credentials)

This server primarily focuses on the `/api/send-email` endpoint. The base URL will depend on where your server is hosted (e.g., `http://localhost:8080`).

#### POST /api/send-email
Sends an email using credentials configured in environment variables.

**Request**:
```json
{
  "to": "recipient@example.com",
  "subject": "Hello from Go Server",
  "html": "<h1>Welcome!</h1><p>This is a test email from the Go server.</p>",
  "from": "Your Name"
}
```
**Headers**:
`X-API-KEY: your_strong_api_key_here`

**Response**:
```json
{
  "success": true,
  "message": "Email sent successfully"
}
```

**Errors**:
-   `400 Bad Request`: Missing required fields (`to`, `subject`, `html`, `from`) or invalid JSON body.
-   `401 Unauthorized`: Invalid or missing `X-API-KEY` header.
-   `405 Method Not Allowed`: Attempted a method other than POST.
-   `429 Too Many Requests`: Rate limit exceeded.
-   `500 Internal Server Error`: Failed to send email due to SMTP issues or internal errors.

#### Go Server without Credentials Endpoints (with request body credentials)

This server also primarily focuses on the `/api/send-email` endpoint. The base URL will depend on where your server is hosted (e.g., `http://localhost:8081`).

#### POST /api/send-email
Sends an email using credentials provided directly in the request body.

**Request**:
```json
{
  "to": ["recipient1@example.com", "recipient2@example.com"],
  "subject": "Dynamic Email from Go Server",
  "html": "<p>This email was sent with credentials in the payload.</p>",
  "from": "Dynamic Sender",
  "gmail_user": "dynamic_sender@gmail.com",
  "gmail_app_password": "dynamic_gmail_app_password"
}
```
**Headers**:
`X-API-KEY: your_strong_api_key_here`

**Response**:
```json
{
  "success": true,
  "message": "Email sent successfully"
}
```

**Errors**:
-   `400 Bad Request`: Missing required fields (`to`, `subject`, `html`, `from`, `gmail_user`, `gmail_app_password`) or invalid JSON body.
-   `401 Unauthorized`: Invalid or missing `X-API-KEY` header.
-   `405 Method Not Allowed`: Attempted a method other than POST.
-   `429 Too Many Requests`: Rate limit exceeded.
-   `500 Internal Server Error`: Failed to send email due to SMTP issues or internal errors.

#### Node.js Server Endpoints

This server primarily focuses on the `/api/send-email` endpoint. The base URL will depend on where your server is hosted (e.g., `http://localhost:3000`).

#### POST /api/send-email
Sends an email using Nodemailer and credentials configured in environment variables.

**Request**:
```json
{
  "to": "recipient@example.com",
  "subject": "Greetings from Node.js",
  "html": "<p>Hello, this is an HTML email!</p>",
  "text": "Hello, this is a plain text email!"
}
```
**Headers**:
_No `X-API-KEY` required for this Node.js implementation based on the code provided._

**Response**:
```json
{
  "success": true,
  "message": "Email sent successfully"
}
```

**Errors**:
-   `400 Bad Request`: Missing required fields (`to`, `subject`, and either `text` or `html`).
-   `405 Method Not Allowed`: Attempted a method other than POST.
-   `500 Internal Server Error`: Failed to send email due to Nodemailer configuration or SMTP issues.

---

## Technologies Used

| Technology    | Description                                       |
| :------------ | :------------------------------------------------ |
| **Go**        | High-performance backend language.                |
| **Node.js**   | JavaScript runtime for server-side applications.  |
| **Express**   | Fast, unopinionated, minimalist web framework for Node.js. |
| **Gomail**    | Go package for sending emails.                    |
| **Nodemailer**| Node.js module for sending emails.                |
| **Dotenv**    | Loads environment variables from a `.env` file.   |
| **Go Mod**    | Go's dependency management system.                |

---

## Author

**Onos**

-   LinkedIn: [https://linkedin.com/in/yourusername](https://linkedin.com/in/yourusername)
-   Twitter: [https://twitter.com/yourusername](https://twitter.com/yourusername)
-   Portfolio: [https://yourportfolio.com](https://yourportfolio.com)

---

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)