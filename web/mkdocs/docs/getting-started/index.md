# Getting Started

The Android SMS Gateway can operate in three distinct modes, all utilizing the same API:

1. [**Local Server**](./local-server.md): Operate entirely within your local network by running the server on your Android device. This mode is perfect for quick setups with no need for Internet access. For remote access, you may need to configure your network settings and implement additional security measures.
2. [**Public Cloud Server**](./public-cloud-server.md): Easily connect via the Internet using the public server at `sms.capcom.me`. Messages are routed through this server to your devices, simplifying the setup without requiring network adjustments. This mode is suitable for non-sensitive data only. For more secure communication, please refer to the [end-to-end encryption section](../privacy/encryption.md).
3. [**Private Server**](./private-server.md): Deploy your own server instance and connect your Android devices to ensure maximum privacy. We will not have access to your message content, making this the preferred option for sensitive communication. However, this option requires setting up and maintaining your own infrastructure, including a database and server application.

To begin with any of these modes, you must first install the Android SMS Gateway app on your device, as described in the [Installation](../installation.md) section.

For more details on how to use the API, please consult the [API](../api.md) section.

## Building Your Own Gateway

Building your own gateway is an option that allows you to create an independent infrastructure without any connection to the public server at `sms.capcom.me`. In most cases, this is not necessary, and you can use the [Private Server](./private-server.md) mode with all of its privacy features. However, if you require full control, please see [Custom Gateway Setup](./custom-gateway.md) section.