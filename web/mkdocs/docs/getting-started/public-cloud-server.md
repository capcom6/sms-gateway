# Getting Started

## Cloud Server

<div align="center">
    <img src="/assets/cloud-server-arch.png" alt="Architecture of the Cloud Server mode">
</div>

Use the cloud server mode when dealing with dynamic or shared device IP addresses. The best part? No registration, email, or phone number is required to start using it.

1. Launch the app on your device.
2. Toggle the `Cloud Server` switch to the "on" position.
3. Tap the `Online` button located at the bottom of the screen to connect to the cloud server.
4. In the `Cloud Server` section, the credentials for basic authentication will be displayed.
   <div align="center">
      <img src="/assets/cloud-server.png" alt="Example settings for Cloud Server mode">
   </div>
5. To send a message via the cloud server, perform a `curl` request with a command similar to the following, substituting `<username>` and `<password>` with the actual values obtained in step 4:

    ```sh
    curl -X POST -u <username>:<password> -H "Content-Type: application/json" -d '{ "message": "Hello, world!", "phoneNumbers": ["79990001234", "79995556677"] }' https://sms.capcom.me/api/3rdparty/v1/message
    ```
