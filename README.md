# Image Receiver API

The Image Receiver API, implemented using Golang with the Fiber web framework, serves as a tool to receive JPEG or PNG files, store them in an S3 bucket, and emit event notifications via RabbitMQ to a consumer application proccess them.

## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
   - [Uploading Images](#uploading-images)
   - [Event Notification](#event-notification)
3. [Configuration](#configuration)
4. [Dependencies](#dependencies)
5. [Contributing](#contributing)
6. [License](#license)

## Installation

To set up the Image Receiver API, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/image-receiver-api.git
   cd image-receiver-api
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```
3. Run the API:
   ```
   go run main.go
   ```

## Usage

### uploading-images

To upload an image, send a POST request to the /upload endpoint with the image file attached as form data. The API will save the image to the configured S3 bucket.
Example using cURL:

```
curl -X POST -H "Content-Type: multipart/form-data" -F "file=@/path/to/your/image.jpg" http://localhost:5000/api/v1/upload

```

![Upload Sequence](/receiver-api.png)

### event-notification

Upon successfully saving the image to the S3 bucket, the API will send a message event via RabbitMQ. Ensure that the RabbitMQ configuration is correctly set up to receive and process these events.

![Event Architecture](/fannoutDiagram.png)
