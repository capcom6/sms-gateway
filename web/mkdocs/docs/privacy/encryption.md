# Encryption

The application supports end-to-end encryption by encrypting message text and recipients' phone numbers before sending them to the API and decrypting them on the device. This ensures that no one – including us as the service provider, the hosting provider, or any third parties – can access the content and recipients of the messages.

## Requirements

1. Fields `message` and every value in the `phoneNumbers` field must be encrypted.
2. The `isEncrypted` field of the message object must be set to `true`.
3. On the device, the same passphrase must be specified as in step 1.

## Algorithm

1. Select a passphrase that will be used for encryption and specify it on the device.
2. Generate a random salt, with 16 bytes being the recommended size.
3. Create a secret key using the PBKDF2 algorithm with SHA1 hash function, key size of 256 bits, and an iteration count of 300,000.
4. Encrypt the message text and recipients' phone numbers using the AES-256-CBC algorithm and encode the result as Base64.
5. Add the Base64-encoded salt to the front of the encrypted data, separated by a dot.

Or use one of the following realization:

### [PHP](https://github.com/capcom6/android-sms-gateway-php/blob/master/src/Encryptor.php)

```php
class Encryptor {
    protected string $passphrase;

    public function __construct(
        string $passphrase
    ) {
        $this->passphrase = $passphrase;
    }

    public function Encrypt(string $data): string {
        $salt = $this->generateSalt();
        $secretKey = $this->generateSecretKeyFromPassphrase($this->passphrase, $salt);

        return base64_encode($salt) . '.' . openssl_encrypt($data, 'aes-256-cbc', $secretKey, 0, $salt);
    }

    public function Decrypt(string $data): string {
        list($saltBase64, $encryptedBase64) = explode('.', $data, 2);

        $salt = base64_decode($saltBase64);
        $secretKey = $this->generateSecretKeyFromPassphrase($this->passphrase, $salt);

        return openssl_decrypt($encryptedBase64, 'aes-256-cbc', $secretKey, 0, $salt);
    }

    protected function generateSalt(int $size = 16): string {
        return random_bytes($size);
    }

    protected function generateSecretKeyFromPassphrase(
        string $passphrase,
        string $salt,
        int $keyLength = 32,
        int $iterationCount = 300000
    ): string {
        return hash_pbkdf2('sha1', $passphrase, $salt, $iterationCount, $keyLength, true);
    }
}
```