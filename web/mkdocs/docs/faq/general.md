# FAQ

## How can I send an SMS using the second SIM card?

To send an SMS using a non-default SIM card, you can specify the SIM card slot number in the `simNumber` field of request. For instance, the following request will send an SMS using the second SIM card:

```sh
curl -X POST \
    -u <username>:<password> \
    -H 'content-type: application/json' \
    https://sms.capcom.me/api/3rdparty/v1/message \
    -d '{
        "message": "Hello from SIM2",
        "phoneNumbers": ["79990001234"],
        "simNumber": 2
    }'
```

## Does the app require power saving mode to be turned off to function without interruptions?

### Local Mode
- **Power saving settings:** No changes needed. The app uses a foreground service with a wake lock to operate, which allows it to run without power saving mode.
- **Battery impact:** Using wake lock may lead to quicker battery drain.

### Cloud Mode
- **Power saving settings:** No changes needed. The app uses Firebase Cloud Messaging (FCM) for push notifications, which works without disabling power saving mode.
- **Potential delays:** High message rates could lead to occasional delays when the device is in power-saving modes because of limits on high-priority FCM notifications.

### Recommendations
- **Testing:** It's advisable to test the app with and without power saving mode activated to see how it performs on your specific device and Android version.
- **Device manufacturers:** Behavior might vary by device manufacturer.
- **Local + Cloud:** For better responsiveness, consider running a local server alongside the cloud server connection.
  
See also issue [#17](https://github.com/capcom6/android-sms-gateway/issues/17).

## How do I enable or disable delivery reports for messages?

As of version [1.3.0](https://github.com/capcom6/android-sms-gateway/releases/tag/v1.3.0), you have the option to enable or disable delivery reports for each message. By default, the delivery report feature is turned on. If you prefer not to receive delivery reports, you can disable them by setting the `withDeliveryReport` field to `false` in the JSON body of your message request. Here is an example of how to send a message without requesting a delivery report:

```json
{
  "message": "Your message text here",
  "phoneNumbers": ["79990001234", "79995556677"],
  "withDeliveryReport": false
}
```

## Can I use long or non-standard phone numbers?

Yes, starting from [1.6.1](https://github.com/capcom6/android-sms-gateway/releases/tag/v1.6.1) of the app, our system allows the use of long or non-standard phone numbers, which may be common with M2M (machine-to-machine) SIM cards or other special cases. To bypass the standard phone number validation, simply add the query parameter `skipPhoneValidation=true` to your API request. Please note that with validation disabled, you are responsible for ensuring the correctness of the phone numbers. They should still follow the E.164 format, beginning with a '+' and containing only digits.

## How can I avoid mobile operator restrictions?

The application provides two features to help you avoid mobile operator restrictions:

### Random delay between messages

You can introduce a random delay between messages by specifying a maximum delay time. The application will then randomly select a delay within this range for each message sent. This helps to reduce the likelihood of messages being flagged as spam by simulating a more human-like sending pattern. This option is available in the "Messages" section of the device's settings and is named "Delay between messages".

### Limiting the number of messages sent per period

The app offers a feature to restrict the number of messages sent within a specified period—be it a minute, hour, or day. You can find this option under the "Limits" section in the device settings. When the limit is reached, the app will pause sending messages until the limit period resets. It's important to note that this feature should not be used for time-sensitive messages, such as sending authorization codes, where delays could cause issues.

## How to hide messages on device?

It's not possible to completely hide messages on the device without major changes in the app due to Android's limitations. Messages sent will appear in the default messaging app. The best way to keep messages private is to restrict physical access to the phone.

