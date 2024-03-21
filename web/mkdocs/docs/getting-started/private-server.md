# Getting Started

## Private Server

<div align="center">
    <img src="/assets/private-server-arch.png" alt="Architecture of the Private Server mode">
</div>

To enhance privacy, you can host a private server within your own infrastructure, ensuring that all messages remain solely on devices you control. The only required external network connection is for sending push notifications via the public server at `sms.capcom.me`. This setup eliminates the need to configure Firebase Cloud Messaging (FCM) or rebuild the Android app, but it does demand some technical know-how.

### Prerequisites

- A MySQL or MariaDB database server with an empty database, and a user granted full access to that database.
- A Virtual Private Server (VPS) running Linux with Docker installed.
- A reverse proxy with a valid SSL certificate and HTTPS enabled.

### Run the Server

1. Create a `config.yml` file based on the template provided in [config.example.yml](https://github.com/capcom6/sms-gateway/blob/master/configs/config.example.yml). Pay special attention to the `database`, `http`, and `gateway` sections. Environment variables can be used to override values in the config file.
   - Set `gateway.mode` to `private`.
   - Define `gateway.private_token` as the access token for device registration in private mode. Ensure this token matches the one on the devices set to private mode.
2. Start the server in Docker with the following command: 
   ```sh
   docker run -p 3000:3000 -v $(pwd)/config.yml:/app/config.yml capcom6/sms-gateway:latest
   ```
3. Configure your reverse proxy, enable SSL, and modify your firewall settings to permit Internet access to the server.

Refer to the server's [README.md](https://github.com/capcom6/sms-gateway/blob/master/README.md) for more information.

See also:

- [Installation Example with Ubuntu, Docker, and Nginx Proxy Manager](https://github.com/capcom6/android-sms-gateway/discussions/50)
- [Docker Compose Quickstart for Single File Deployment](https://github.com/capcom6/sms-gateway/tree/master/deployments/docker-compose-proxy)

### Configure the Android App

<div align="center">
    <img src="/assets/private-server.png" alt="Example settings for Private Server mode">
</div>

*Note*: Changing the server will invalidate current credentials, and the device will be re-registered with new ones.

1. Navigate to the Settings tab.
2. In the Cloud Server section, enter the API URL and Private token, ensuring they match those in the server configuration. Note that you should include the full URL with the path, such as `https://private.example.com/api/mobile/v1`.
3. Switch to the Home tab.
4. Activate the Cloud server option.
5. Apply the new configuration by stopping and starting the app using the button at the bottom of the screen.

If everything is configured correctly, the new credentials for the private server will be displayed in the Cloud Server section on the Home tab.
