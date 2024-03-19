# Getting Started

## Local Server

<div align="center">
    <img src="/assets/local-server-arch.png" alt="Architecture of the Local Server mode">
</div>

This mode is ideal for sending messages from a local network.

1. Launch the app on your device.
2. Toggle the `Local Server` switch to the "on" position.
3. Tap the `Offline` button located at the bottom of the screen to activate the server.
4. The `Local Server` section will display your device's local and public IP addresses, as well as the credentials for basic authentication. Please note that the public IP address is only accessible if you have a public (also known as "white") IP and your firewall is configured appropriately.
    <div align="center">
        <img src="/assets/local-server.png" alt="Example settings for Local Server mode">
    </div>
5. To send a message from within the local network, execute a `curl` command like the following. Be sure to replace `<username>`, `<password>`, and `<device_local_ip>` with the actual values provided in the previous step:

    ```sh
    curl -X POST -u <username>:<password> -H "Content-Type: application/json" -d '{ "message": "Hello, world!", "phoneNumbers": ["79990001234", "79995556677"] }' http://<device_local_ip>:8080/message
    ```
