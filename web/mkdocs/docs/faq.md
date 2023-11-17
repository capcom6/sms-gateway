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
