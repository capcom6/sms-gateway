# Getting Started

The Android SMS Gateway can operate in three distinct modes, all using the same API:

1. [**Local Server**](./local-server.md): Operate entirely within your local network by running the server directly on your Android device. This is perfect for a quick setup and use where Internet access isn't necessary. For remote access, you may need to adjust your network settings and consider additional security measures.
2. [**Public Cloud Server**](./public-cloud-server.md): Connect easily via the Internet using our public server at `sms.capcom.me`. Messages are routed through this server to your devices, which simplifies the setup without the need for network adjustments. This is suitable for non-sensitive data only — for more secure communication, please check out the [end-to-end encryption section](../privacy/encryption.md).
3. [**Private Server**](./private-server.md): Deploy your own server instance and connect your Android devices to ensure maximum privacy. We won't have access to your message content, making this the preferred option for sensitive communication. However, it requires setting up and maintaining your own infrastructure, which includes a database and server application.

For any of these modes, you'll first need to install the Android SMS Gateway app on your device as described in the [Installation](../installation.md) section.

For more information on how to use the API, please refer to the [API](../api.md) section.
