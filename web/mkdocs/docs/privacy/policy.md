# Privacy

We are committed to transparency and the protection of user privacy. This document outlines how we handle information within our app.

## Local Server

- **No Data Sent to Cloud**: The app functions autonomously, without transmitting any data to cloud servers.
- **Public IP Retrieval**: The sole external network request is made to [ipify](https://www.ipify.org) to retrieve your public IP address. No personal data is disclosed in this process.

## Cloud Server

- **Encrypted Communication**: All communication between the app and the cloud server is protected with secure encryption protocols, safeguarding your data during transit.
- **End-to-End Encryption**: Optional AES-based end-to-end encryption is available. With this enabled, messages and phone numbers are encrypted prior to being sent to the API. Consequently, data is encrypted on transmission and only decrypted on the user's device when sending the SMS. This ensures that no entity, including us as the service provider, the hosting provider, or any intermediaries, can access message contents or recipient details.
- **Message Handling**: If end-to-end encryption is not utilized, the message content and recipient information are transformed into a SHA256 hash within 15 minutes of device acknowledgment, preventing the storage of this information in plain text.
- **Limited Data Sharing**: Only necessary data such as device manufacturer, model, app version, and Firebase Cloud Messaging (FCM) token are shared with the server to facilitate cloud services.

## Private Server

- **Push Notifications**: Push notifications are routed through the cloud server. These notifications carry solely the device's FCM token and omit any information regarding SMS content or recipients.

## No Collection of Usage Statistics

- **No Tracking**: We abstain from collecting any usage statistics, including crash reports. Your interaction with the app remains confidential and unmonitored.
