# Android SMS Gateway Server

This server acts as the backend component of the Android SMS Gateway, facilitating the sending of SMS messages through connected Android devices. It includes a RESTful API for message management, integration with Firebase Cloud Messaging (FCM), and a database for persistent storage.

## Table of Contents

- [Android SMS Gateway Server](#android-sms-gateway-server)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Running the Server](#running-the-server)
    - [Running with Docker](#running-with-docker)
  - [Contributing](#contributing)
  - [License](#license)

## Features

- Send SMS messages via a RESTful API.
- Schedule and perform periodic tasks.
- Integrate with Firebase Cloud Messaging for notifications.

## Prerequisites

- Go (for development and testing purposes)
- Docker and Docker Compose (for Docker-based setup)
- A configured MySQL/MariaDB database

## Installation

To set up the server on your local machine for development and testing purposes, follow these steps:

1. Clone the repository to your local machine.
2. Install Go (version 1.21 or newer) if not already installed.
3. Navigate to the cloned directory and install dependencies:

```bash
make init
```

4. Build the server binary:

```bash
make build
```

## Configuration

The server uses `yaml` for configuration with ability to override some values from environment variables. By default configuration is loaded from the `config.yml` file in the root directory. But path can be overridden with the `CONFIG_PATH` environment variable.

Below is a template for the `config.yml` file with environment variables in comments:

```yaml
http:
  listen: ":3000"   # HTTP__LISTEN
database:
  dialect: "mysql"  # DATABASE__DIALECT
  host: "localhost" # DATABASE__HOST
  port: 3306        # DATABASE__PORT
  user: "sms"       # DATABASE__USER
  password: "sms"   # DATABASE__PASSWORD
  database: "sms"   # DATABASE__DATABASE
  timezone: "UTC"   # DATABASE__TIMEZONE
fcm:
  credentials_json:  >
    {
    ...
    }
tasks:
  hashing:
    interval_seconds: 900
```

Replace the placeholder values with your actual configuration.

## Running the Server

### Running with Docker

For convenience, a Docker-based setup is provided. Please refer to the Docker prerequisites above before proceeding.

1. Prepare configuration file `config.yml`
2. Pull the Docker image: `docker pull capcom6/sms-gateway`
3. Apply database migrations: `docker run --rm -it -v ./config.yml:/app/config.yml capcom6/sms-gateway db:migrate`
4. Start the server: `docker run -p 3000:3000 -v ./config.yml:/app/config.yml capcom6/sms-gateway`

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the Apache-2.0 license. See [LICENSE](LICENSE) for more information.