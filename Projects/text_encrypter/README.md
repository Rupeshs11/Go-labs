# 🔠 Text Encrypter (Caesar Cipher)

A lightweight encryption tool written in **Go** that demonstrates the fundamental concepts of cryptography by executing a classic **Substitution Cipher** (similar to the famous Caesar Cipher).

---

## 🛠 Features

- **Alphabet Substitution**: Performs character mapping by intelligently shifting the standard `A-Z` alphabet based on a specific numerical `key` offset.
- **Bi-directional Encryption**: Native functions for both encrypting cleartext and reversing ciphertext back to its original value.
- **Rune Manipulation**: Demonstrates how to cleanly handle string internals, leveraging Go's robust `rune` slice casting to safely index string data without breaking UTF-8 boundaries.

---

## 💻 Code Explanations

### The Hashing Mechanism (`hashLetterFn`)

```go
const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashLetterFn(key int, letter string) (result string) {
    runes := []rune(letter)
    lastLetterKey := string(runes[len(letter)-key : len(letter)])
    leftOversLetter := string(runes[0 : len(letter)-key])

    return fmt.Sprintf(`%s%s`, lastLetterKey, leftOversLetter)
}
```

- Rather than shifting letters mathematically by their ASCII values, this function slices the original alphabet string cleanly into two parts using a specified key, then swaps their execution placements to generate the new "hashed" alphabet cipher map.

### The Index Mapper (`strings.Map`)

```go
findOne := func(r rune) rune {
    pos := strings.Index(originalLetter, string([]rune{r}))
    if pos != -1 {
        // Find position and map substitution
        // ...
        return r
    }
    return r
}
strings.Map(findOne, plainText)
```

- The `encrypt` and `decrypt` pipelines use `strings.Map` to dynamically iterate across the target string, sending each character (`rune`) into the `findOne` execution hook to locate its index within the alphabet and map it cleanly to the parallel position on the shifted array.

---

## 🚀 How to Run

1. Clone the repository and navigate to the project directory:

   ```bash
   cd Projects/text_encrypter
   ```

2. Run the application:

   ```bash
   go run main.go
   ```

3. **Input Format**: Note that because it utilizes standard `fmt.Scan()`, the program expects a single string sequence (no spaces), and natively maps against capital strings `A-Z`.

### Example Execution

```bash
Enter the text to encrypt :
KNOX
Plain Text KNOX
Encrypted FJIH
Decrypted KNOX
```

---

## Dependencies

- **Zero External Dependencies!** Built entirely using Go's excellent built-in `strings` and `fmt` manipulation packages.
