# Getting Started

## Private Server

<div align="center">
    <img src="/assets/private-server.png" alt="Example settings for Private Server mode">
</div>

To enhance privacy, you can host a private server within your own infrastructure, ensuring that your messages reside only on devices you control. The only required external network connection is for sending push notifications through the public server at `sms.capcom.me`. This architecture eliminates the need to set up Firebase Cloud Messaging (FCM) and rebuild the Android app, though it does require some technical expertise.

### Prerequisites

- A MySQL or MariaDB database server with an empty database and a user granted full access to that database.
- A VPS (Virtual Private Server) running Linux with Docker installed.
- A reverse proxy with an SSL certificate and HTTPS enabled.

### Run the Server

1. Create a `config.yml` file based on the [config.example.yml](https://github.com/capcom6/sms-gateway/blob/master/configs/config.example.yml). The critical sections to configure are `database`, `http`, and `gateway`. Environment variables can be used to override values in the config file.
   - In the `gateway.mode` section, set the value to `private`.
   - In the `gateway.private_token` section, specify the access token for device registration in private mode. This token must also be set on devices operating in private mode.
2. Launch the server in Docker with the following command: 
   ```sh
   docker run -p 3000:3000 -v $(pwd)/config.yml:/app/config.yml capcom6/sms-gateway:latest
   ```
3. Set up your reverse proxy, configure SSL, and adjust your firewall to allow access to the server from the Internet.

For additional details, refer to the server's [README.md](https://github.com/capcom6/sms-gateway/blob/master/README.md).

### Configure Android App

To be continued...
