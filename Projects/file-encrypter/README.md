# 🔐 CryptoGo (File Encrypter)

A resilient, military-grade File Encryption CLI (CryptoGo) built entirely in **Go**. It encrypts and decrypts files seamlessly using `AES-256-GCM` authenticated encryption.

---

## 🛠 Features

- **Military-Grade Encryption**: Protects any target file using `AES-256` symmetric encryption in `Galois/Counter Mode (GCM)`.
- **Password Obfuscation**: CLI passwords are never echoed in the terminal, thanks to the robust `golang.org/x/term` library.
- **Key Stretching (PBKDF2)**: Ensures raw passwords are computationally immune against brute-force/dictionary attacks.
- **In-place Operations**: Safely rewrites file contents securely, appending the cryptographic Nonce required to unlock it at the tail of the payload!

---

## 💻 Under the Hood (Cryptography logic)

### Safely deriving a 32-bit key (`PBKDF2`)

```go
dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)
```

- Passes your raw password through `4096` iterations of `SHA-1` via the `PBKDF2` algorithm.
- Extrapolates a secure, unpredictable 32-byte (256-bit) encryption key from your simple password to plug directly into `AES`.

### Encrypting & Payload Masking (GCM)

```go
block, _ := aes.NewCipher(dk)
aesgcm, _ := cipher.NewGCM(block)

ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

// Magic step: Append 12-byte nonce to the very end of file
ciphertext = append(ciphertext, nonce...)
```

- Uses GCM which provides **Authenticated Encryption with Associated Data (AEAD)**—meaning it encrypts and simultaneously protects against malicious tampering!
- Instead of forcing the user to manually memorize/pass around the 12-byte initialization vector (`nonce`), CryptoGo simply appends it securely to the very last 12 bytes of the file itself!

### Decrypting

```go
// Extacting the Salt/Nonce hidden at the tail of the payload
salt := ciphertext[len(ciphertext)-12:]
str := hex.EncodeToString(salt)
nonce, _ := hex.DecodeString(str)

// Run the AEAD Open function
plaintext, err := aesgcm.Open(nil, nonce, ciphertext[:len(ciphertext)-12], nil)
```

- Snips the last 12 bytes to recover the `nonce`, discarding it from the ciphertext payload before throwing the remaining data into `aesgcm.Open(...)`.

---

## 🚀 How to Run

1. Clone the repository and navigate to the project directory:

   ```bash
   cd Projects/file-encrypter
   ```

2. Download dependencies (`golang.org/x/term` and `golang.org/x/crypto/pbkdf2`):

   ```bash
   go mod tidy
   ```

3. Build the application binary:
   ```bash
   go build -o file-encrypter.exe main.go
   ```

### Command Examples

**1. Encrypt a file (e.g., sample image):**

```bash
./file-encrypter.exe encrypt deadpool.jpg
# You will be prompted to enter & confirm a hidden password
```

**2. Decrypt a file:**

```bash
./file-encrypter.exe decrypt deadpool.jpg
# Enter your password to retrieve the original logic!
```

**3. Display the help menu:**

```bash
./file-encrypter.exe help
```

---

## Dependencies

| Package                                                 | Purpose                                                         |
| ------------------------------------------------------- | --------------------------------------------------------------- |
| [`x/term`](https://golang.org/x/term)                   | Safely read passwords invisibly within the CLI prompt           |
| [`x/crypto/pbkdf2`](https://golang.org/x/crypto/pbkdf2) | Password-Based Key Derivation Function 2 for AES key stretching |
