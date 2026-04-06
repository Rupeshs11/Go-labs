# 📧 Email Checker

A highly efficient, concurrent-friendly Command Line Tool written in **Go** to bulk verify the DNS configurations and security standings of email domains.

---

## 🛠 Features

- **Bulk Domain Verification**: Reads domain names directly from standard input (`stdin`), parsing them interactively or via redirected files.
- **Deep DNS Lookups**:
  - **MX (Mail Exchange)**: Confirms the domain is actively configured to receive emails.
  - **SPF (Sender Policy Framework)**: Validates if the domain protects against email spoofing.
  - **DMARC (Domain-based Message Authentication)**: Detects rigorous sender authenticity and reporting configurations.
- **CSV Formatted Output**: Spits out results sequentially via comma-separated values, perfectly built to be ingested by datasets, bash scripts, or Excel sheets.

---

## 💻 Code Explanations

### Reading via Standard Input (`bufio.Scanner`)

```go
scanner := bufio.NewScanner(os.Stdin)
fmt.Printf("domain,hasMX,hasSPF,sprRecord,hasDMARC,dmarcRecord\n")

for scanner.Scan() {
    checkDomain(scanner.Text())
}
```

- Listens actively on the console until standard input is closed.
- Prints a structured CSV-header immediately payload generation starts.

### Checking Mail Exchange (MX)

```go
mxRecords, err := net.LookupMX(domain)
if len(mxRecords) > 0 {
    hasMX = true
}
```

- Reaches out to external DNS servers natively using `net.LookupMX` to locate Mail Exchange records. If the slice contains records, we assume the server can legitimately receive emails.

### Validating SPF and DMARC (TXT Records)

```go
txtRecords, err := net.LookupTXT(domain)
for _, record := range txtRecords {
    if strings.HasPrefix(record, "v=spf1") {
        hasSPF = true
        spfRecord = record
        break
    }
}
```

- Performs a highly specific `net.LookupTXT` query against the provided domain.
- Slices through the arbitrary strings to strictly match standard `"v=spf1"` patterns, capturing exactly the framework parameters.
- Utilizes the exact same strategy, but explicitly queries `"_dmarc." + domain`, hunting for `"v=DMARC1"`.

---

## 🚀 How to Run

1. Clone the repository and navigate to the project directory:

   ```bash
   cd Projects/Email-checker
   ```

2. Generate the binary executable:

   ```bash
   go build -o check_email main.go
   ```

3. Run the application interactively:
   ```bash
   ./check_email
   github.com
   google.com
   ```

### 🗃️ Running in Bulk (Piping Input)

Because it runs completely on standard input pipelines, you can feed text files packed with domains directly into it!

```bash
# Given a text file (domains.txt) containing lines of domains
cat domains.txt | ./check_email > validated_results.csv
```

---

## Dependencies

- **Zero External Dependencies!** Built entirely using the pristine `net` and `bufio` standard Go packages.
