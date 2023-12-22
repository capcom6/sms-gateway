# Privacy

We believe in transparency and the importance of privacy. Here's how we handle information in our app:

## Local Mode

- **No Data Sent to Cloud**: The app operates independently without sending any data to a cloud server.
- **Public IP Retrieval**: The only external network communication made is to [ipify](https://www.ipify.org) to obtain your public IP address, and no personal data is shared during this process.

## Cloud Mode

- **Encrypted Communication**: Communication between the app and the cloud server is encrypted.
- **Limited Data Sharing**: Only essential data such as the device manufacturer, model, app version, and Firebase Cloud Messaging (FCM) token is sent to the server to enable cloud functionality.
- **Message Handling**: Message content and recipient phone numbers are stored on the server only until your device confirms receipt. Afterwards, this information is converted into a SHA256 hash within 15 minutes, ensuring it is not stored in clear form.

## No Collection of Usage Statistics

- **No Tracking**: We do not collect any usage statistics, including crash reports. Your usage of the app remains private and untracked.
