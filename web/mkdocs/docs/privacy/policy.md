# Privacy

We believe in transparency and the importance of privacy. Here's how we handle information in our app:

## Local Mode

- **No Data Sent to Cloud**: The app operates independently without sending any data to a cloud server.
- **Public IP Retrieval**: The only external network communication made is to [ipify](https://www.ipify.org) to obtain your public IP address, and no personal data is shared during this process.

## Cloud Mode

- **Encrypted Communication**: Communication between the app and the cloud server is encrypted using secure protocols to protect your data in transit.
- **End-to-End Encryption**: We have implemented optional AES-based end-to-end encryption to ensure that all messages and phone numbers can be encrypted before being sent to the API. This means that data is encrypted before transmission and decrypted on the user's device before sending the SMS, ensuring that no one – including us as the service provider, the hosting provider, or any other party – can access the content and recipients of the messages.
- **Message Handling**: If end-to-end encryption is not used, after your device confirms receipt, message content and recipients are converted into a SHA256 hash within 15 minutes, ensuring it is not stored in clear form.
- **Limited Data Sharing**: Only essential data such as the device manufacturer, model, app version, and Firebase Cloud Messaging (FCM) token is sent to the server to enable cloud functionality.
  
## No Collection of Usage Statistics

- **No Tracking**: We do not collect any usage statistics, including crash reports. Your usage of the app remains private and untracked.