# Getting Started

First of all, you need to install the Android SMS Gateway app on your device as described in the [Installation](installation.md).

The Android SMS Gateway can work in two modes: with a local server started on the device or with a cloud server at [sms.capcom.me](https://sms.capcom.me). The basic API is the same for both modes and is documented on the [Android SMS Gateway API Documentation](https://capcom6.github.io/android-sms-gateway/).

## Local server

<div align="center">
    <img src="/assets/local-server.png" alt="Local server example settings">
</div>

This mode is recommended for sending messages from local network.

1. Start the app on the device.
2. Activate the `Local server` switch.
3. Tap the `Offline` button at the bottom of the screen.
4. In the `Local server` section, the local and public addresses of the device, along with the credentials for basic authentication, will be displayed. Please note that the public address is only usable if you have a "white" IP address and have correctly configured your firewall.
5. Make a `curl` call from the local network using a command like the following, replacing `<username>`, `<password>`, and `<device_local_ip>` with the values obtained in step 4: 

    ```
    curl -X POST -u <username>:<password> -H "Content-Type: application/json" -d '{ "message": "Hello, world!", "phoneNumbers": ["79990001234", "79995556677"] }' http://<device_local_ip>:8080/message
    ```

## Cloud server

<div align="center">
    <img src="/assets/cloud-server.png" alt="Cloud server example settings">
</div>

If you need to send messages with dynamic or shared device IP, you can use the cloud server. The best part? No registration, email, or phone number is required to start using it.

1. Start the app on the device.
2. Activate the `Cloud server` switch.
3. Tap the `Offline` button at the bottom of the screen.
4. In the `Cloud server` section, the credentials for basic authentication will be displayed.
5. Make a curl call using a command like the following, replacing `<username>` and `<password>` with the values obtained in step 4:

    ```
    curl -X POST -u <username>:<password> -H "Content-Type: application/json" -d '{ "message": "Hello, world!", "phoneNumbers": ["79990001234", "79995556677"] }' https://sms.capcom.me/api/3rdparty/v1/message
    ```
