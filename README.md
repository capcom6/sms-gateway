[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Apache 2.0 License][license-shield]][license-url]

# SMS Gateway for Android™ Server

This server acts as the backend component of the [SMS Gateway for Android](https://github.com/capcom6/android-sms-gateway), facilitating the sending of SMS messages through connected Android devices. It includes a RESTful API for message management, integration with Firebase Cloud Messaging (FCM), and a database for persistent storage.

## Table of Contents

- [SMS Gateway for Android™ Server](#sms-gateway-for-android-server)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [Quickstart](#quickstart)
  - [Work modes](#work-modes)
  - [Contributing](#contributing)
  - [License](#license)
  - [Legal Notice](#legal-notice)

## Features

- **SMS Messaging**: Dispatch SMS messages through a RESTful API.
- **Message Status**: Retrieve status for sent messages.
- **Device Management**: View information about connected Android devices.
- **Webhooks**: Configure webhooks for event-driven notifications.
- **Health Monitoring**: Access health check endpoints to ensure system integrity.
- **Access Control**: Operate in either public mode for open access or private mode for restricted access.

## Prerequisites

- Go (for development and testing purposes)
- Docker and Docker Compose (for Docker-based setup)
- A configured MySQL/MariaDB database

## Quickstart

The easiest way to get started with the server is to use the Docker-based setup in Private Mode. In this mode device registration endpoint is protected, so no one can register a new device without knowing the token.

1. Set up MySQL or MariaDB database.
2. Create config.yml, based on [config.example.yml](configs/config.example.yml). The most important sections are `database`, `http` and `gateway`. Environment variables can be used to override values in the config file.
   1. In `gateway.mode` section set `private`.
   2. In `gateway.private_token` section set the access token for device registration in private mode. This token must be set on devices with private mode active.
3. Start the server in Docker: `docker run -p 3000:3000 -v ./config.yml:/app/config.yml capcom6/sms-gateway:latest`.
4. Set up private mode on devices.
5. Use started private server with the same API as the public server at [sms.capcom.me](https://sms.capcom.me).

See also [docker-composee.yml](deployments/docker-compose/docker-compose.yml) for Docker-based setup.

## Work modes

The server has two work modes: public and private. The public mode allows anonymous device registration and used at [sms.capcom.me](https://sms.capcom.me). Private mode can be used to send sensitive messages and running server in local infrastructure.

In most operations public and private modes are the same. But there are some differences:

- `POST /api/mobile/v1/device` endpoint is protected by API key in private mode. So it is not possible to register a new device on private server without knowing the token.
- FCM notifications from private server are sent through `sms.capcom.me`. Notifications don't contain any sensitive data like phone numbers or message text.

See also [private mode discussion](https://github.com/capcom6/android-sms-gateway/issues/20).

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

## Legal Notice

Android is a trademark of Google LLC.

[contributors-shield]: https://img.shields.io/github/contributors/capcom6/sms-gateway.svg?style=for-the-badge
[contributors-url]: https://github.com/capcom6/sms-gateway/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/capcom6/sms-gateway.svg?style=for-the-badge
[forks-url]: https://github.com/capcom6/sms-gateway/network/members
[stars-shield]: https://img.shields.io/github/stars/capcom6/sms-gateway.svg?style=for-the-badge
[stars-url]: https://github.com/capcom6/sms-gateway/stargazers
[issues-shield]: https://img.shields.io/github/issues/capcom6/sms-gateway.svg?style=for-the-badge
[issues-url]: https://github.com/capcom6/sms-gateway/issues
[license-shield]: https://img.shields.io/github/license/capcom6/sms-gateway.svg?style=for-the-badge
[license-url]: https://github.com/capcom6/sms-gateway/blob/master/LICENSE