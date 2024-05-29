# Custom Gateway Setup

By setting up a custom gateway, you gain full control over communications and data. However, this process requires technical skills and infrastructure. Before starting, please consider the [Private Server](./private-server.md) mode, as it offers additional privacy features and is not significantly different from a self-built solution, with much fewer setup operations required.

## Before You Begin

As of the time of writing this document, the backend in its default configuration exposes public endpoints. Therefore, anyone who knows the server address can connect a mobile app to it and use the API.

Please be aware that these instructions are simplified and do not cover detailed security configurations, such as firewall settings and HTTPS setup. Implementing these security measures is the user's responsibility.

**Use at your own risk.**

## Prerequisites

1. Android Studio with the Android SDK installed.
2. A Google Firebase account and a project with Cloud Messaging activated.
3. A MySQL or MariaDB database.
4. A VPS (Virtual Private Server) with a public IP address.
5. A reverse proxy such as Nginx or Traefik with HTTPS support.
6. Go programming language with the latest version installed (only necessary for using private server mode).

## Backend

If you do not plan to use the private server mode, there is no need to rebuild the backend. You can follow the instructions for the [Private Server](./private-server.md) mode with the following changes:

1. Set the `gateway.mode` value to `public`.
2. Update the `fcm.credentials_json` with the content of your Firebase project's credentials JSON file.

### Private Server Mode

To use the private server mode, you must rebuild the backend after modifying the file at `internal/sms-gateway/modules/push/upstream/client.go` to set your main server address. You can then build the binary by executing `make build` or the Docker image by running `make docker-build`.

## Android App

To build a custom version of the Android application that will communicate with your server, follow these steps:

1. Clone the repository at [https://github.com/capcom6/android-sms-gateway](https://github.com/capcom6/android-sms-gateway) and open it in Android Studio.
2. Navigate to `app/src/main/java/me/capcom/smsgateway/modules/gateway/GatewayApi.kt` and update the `BASE_URL` constant to your server's URL.
3. Modify the `applicationId` in `app/build.gradle` to a unique value.
4. Refer to the instructions at [https://firebase.google.com/docs/android/setup](https://firebase.google.com/docs/android/setup) to register your application and generate the `google-services.json` file.
5. Build and run the application.

## See Also

* [GitHub Discussion](https://github.com/capcom6/android-sms-gateway/discussions/71)