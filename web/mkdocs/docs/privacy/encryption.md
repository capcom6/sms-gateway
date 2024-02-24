# Encryption

The application supports end-to-end encryption by encrypting message text and recipients' phone numbers before sending them to the API and decrypting them on the device. This ensures that no one – including us as the service provider, the hosting provider, or any third parties – can access the content and recipients of the messages.

Please note that using encryption will increase device battery usage.

## Requirements

1. Fields `message` and every value in the `phoneNumbers` field must be encrypted.
2. The `isEncrypted` field of the message object must be set to `true`.
3. On the device, the same passphrase must be specified as in step 1.

## Algorithm

1. Select a passphrase that will be used for encryption and specify it on the device.
2. Generate a random salt, with 16 bytes being the recommended size.
3. Create a secret key using the PBKDF2 algorithm with SHA1 hash function, key size of 256 bits, and recommended iteration count of 75,000.
4. Encrypt the message text and recipients' phone numbers using the AES-256-CBC algorithm and encode the result as Base64.
5. Format result as `$aes-256-cbc/pbkdf2-sha1$i=<iteration count>$<base64 encoded salt>$<base 64 encoded encrypted data>`. The format is inspired by [PHC](https://github.com/P-H-C/phc-string-format/blob/master/phc-sf-spec.md).

Or use one of the following realization:

### [PHP](https://github.com/capcom6/android-sms-gateway-php/blob/master/src/Encryptor.php)

```php
class Encryptor {
    protected string $passphrase;
    protected int $iterationCount;

    /**
     * Encryptor constructor.
     * @param string $passphrase Passphrase to use for encryption
     * @param int $iterationCount Iteration count
     */
    public function __construct(
        string $passphrase,
        int $iterationCount = 75000
    ) {
        $this->passphrase = $passphrase;
        $this->iterationCount = $iterationCount;
    }

    public function Encrypt(string $data): string {
        $salt = $this->generateSalt();
        $secretKey = $this->generateSecretKeyFromPassphrase($this->passphrase, $salt, 32, $this->iterationCount);

        return sprintf(
            '$aes-256-cbc/pbkdf2-sha1$i=%d$%s$%s',
            $this->iterationCount,
            base64_encode($salt),
            openssl_encrypt($data, 'aes-256-cbc', $secretKey, 0, $salt)
        );
    }

    public function Decrypt(string $data): string {
        list($_, $algo, $paramsStr, $saltBase64, $encryptedBase64) = explode('$', $data);

        if ($algo !== 'aes-256-cbc/pbkdf2-sha1') {
            throw new \RuntimeException('Unsupported algorithm');
        }

        $params = $this->parseParams($paramsStr);
        if (empty($params['i'])) {
            throw new \RuntimeException('Missing iteration count');
        }

        $salt = base64_decode($saltBase64);
        $secretKey = $this->generateSecretKeyFromPassphrase($this->passphrase, $salt, 32, intval($params['i']));

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

    /**
     * @return array<string, string>
     */
    protected function parseParams(string $params): array {
        $keyValuePairs = explode(',', $params);
        $result = [];
        foreach ($keyValuePairs as $pair) {
            list($key, $value) = explode('=', $pair, 2);
            $result[$key] = $value;
        }
        return $result;
    }
}
```

### TypeScript

Please note, that Bun's implementation of `crypto` package is not optimized, so it is much slower than Node's implementation.

```typescript
import crypto from "crypto";

class Encryptor {
    constructor(protected readonly passphrase: string, protected readonly iterations: number = 75_000) {

    }

    public Decrypt(input: string): string {
        const parts = input.split("$");
        if (parts.length !== 5) {
            throw new Error("Invalid encrypted text");
        }

        if (parts[1] !== "aes-256-cbc/pbkdf2-sha1") {
            throw new Error("Unsupported algorithm");
        }

        const paramsStr = parts[2];
        const params = this.parseParams(paramsStr);
        if (!params.has("i")) {
            throw new Error("Missing iteration count");
        }

        const iterations = parseInt(params.get("i")!);
        const salt = Buffer.from(parts[3], "base64");
        const encryptedText = Buffer.from(parts[4], "base64");

        const secretKey = this.generateSecretKeyFromPassphrase(this.passphrase, salt, 32, iterations);
        const decryptedText = this.decryptString(encryptedText, secretKey, salt);

        return decryptedText.toString("utf8");
    }

    protected parseParams(params: string): Map<string, string> {
        const keyValuePairs = params.split(",");

        const result = new Map<string, string>();
        keyValuePairs.forEach(pair => {
            const [key, value] = pair.split("=");
            result.set(key, value);
        });
        return result;
    }

    protected decryptString(input: Buffer, secretKey: Buffer, iv: Buffer): Buffer {
        const decipher = crypto
            .createDecipheriv("aes-256-cbc", secretKey, iv);
        return Buffer.concat([decipher.update(input), decipher.final()]);
    }

    public Encrypt(input: string): string {
        const salt = this.generateSalt();
        const secretKey = this.generateSecretKeyFromPassphrase(this.passphrase, salt, 32, this.iterations);
        const encryptedText = this.encryptString(Buffer.from(input, "utf8"), secretKey, salt);
        return `$aes-256-cbc/pbkdf2-sha1$i=${this.iterations}$${salt.toString("base64")}$${encryptedText.toString("base64")}`;
    }

    protected encryptString(input: Buffer, secretKey: Buffer, iv: Buffer): Buffer {
        const cypher = crypto
            .createCipheriv("aes-256-cbc", secretKey, iv);

        return Buffer.concat([cypher.update(input), cypher.final()]);
    }

    protected generateSalt(size: number = 16): Buffer {
        return crypto.randomBytes(size);
    }

    protected generateSecretKeyFromPassphrase(passphrase: string, salt: Buffer, keyLength: number = 32, iterations: number = 75_000): Buffer {
        return crypto.pbkdf2Sync(passphrase, salt, iterations, keyLength, "sha1");
    }
}
```

### Python

Based on [pycryptodome](https://pypi.org/project/pycryptodome/)

```python
from Crypto.Cipher import AES
from Crypto.Hash import SHA1
from Crypto.Protocol.KDF import PBKDF2
from Crypto.Random import get_random_bytes
from Crypto.Util.Padding import pad, unpad

class AESEncryptor(BaseEncryptor):
    def encrypt(self, cleartext: str) -> str:
        saltBytes = self._generate_salt()
        key = self._generate_key(saltBytes, self.iterations)

        cipher = AES.new(key, AES.MODE_CBC, iv=saltBytes)

        encrypted_bytes = cipher.encrypt(pad(cleartext.encode(), AES.block_size))

        salt = base64.b64encode(saltBytes).decode("utf-8")
        encrypted = base64.b64encode(encrypted_bytes).decode("utf-8")

        return f"$aes-256-cbc/pbkdf2-sha1$i={self.iterations}${salt}${encrypted}"

    def decrypt(self, encrypted: str) -> str:
        chunks = encrypted.split("$")

        if len(chunks) < 5:
            raise ValueError("Invalid encryption format")

        if chunks[1] != "aes-256-cbc/pbkdf2-sha1":
            raise ValueError("Unsupported algorithm")

        params = self._parse_params(chunks[2])
        if "i" not in params:
            raise ValueError("Missing iteration count")

        iterations = int(params["i"])
        salt = base64.b64decode(chunks[-2])
        encrypted_bytes = base64.b64decode(chunks[-1])

        key = self._generate_key(salt, iterations)
        cipher = AES.new(key, AES.MODE_CBC, iv=salt)

        decrypted_bytes = unpad(cipher.decrypt(encrypted_bytes), AES.block_size)

        return decrypted_bytes.decode("utf-8")

    def _generate_salt(self) -> bytes:
        return get_random_bytes(16)

    def _generate_key(self, salt: bytes, iterations: int) -> bytes:
        return PBKDF2(
            self.passphrase,
            salt,
            count=iterations,
            dkLen=32,
            hmac_hash_module=SHA1,
        )

    def _parse_params(self, params: str) -> t.Dict[str, str]:
        return {k: v for k, v in [p.split("=") for p in params.split(",")]}
```