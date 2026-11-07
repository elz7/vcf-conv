# 📇 vcf-conv

**Turn phone number lists into contacts.**

A lightweight CLI tool that converts `.txt` and `.csv` files into `.vcf` (vCard) files, ready to import into your phone's address book.

Got a list of phone numbers and want to check which ones are associated with publicly visible WhatsApp profiles? Import them into your contacts in seconds.

### 🔎 Use cases

- **OSINT:** Prepare collected phone number lists for contact-based research.
- **WhatsApp:** Import phone numbers into your address book and check publicly visible profile information, such as profile photos.
- **Contact discovery:** Quickly turn raw phone number lists into actual contacts.
- **Data processing:** Convert TXT and CSV datasets into a standard contact format.

### ✨ Features

- 📄 TXT and CSV input support.
- 📇 VCF (vCard) output.
- ⚡ Batch contact conversion.
- 🛠 Custom output file path.
- 🖥 Simple command-line interface.

---

## 🚀 Getting started

### Requirements

- [Go](https://go.dev/dl/) 1.24.2 or later.

### Build from source

```bash
git clone https://github.com/elz7/vcf-conv.git
cd vcf-conv

go build -o vcf-conv .
```

## 💻 Usage

```bash
./vcf-conv [flags] <input-file>
```

### Examples

Convert a TXT file:

```bash
./vcf-conv phones.txt
```

Convert a CSV file:

```bash
./vcf-conv contacts.csv
```

Specify a custom output file:

```bash
./vcf-conv --out imported-contacts.vcf phones.csv
```

By default, the output filename is derived from the input filename.

For example:

```text
phones.txt → phones.vcf
```

### ⚙️ Available flags

| Flag | Description |
|---|---|
| `--out` | Specify the output VCF file. |
| `--print-csv-fields` | Display the expected CSV field order. |

To see the expected CSV structure:

```bash
./vcf-conv --print-csv-fields
```

## 📋 Input formats

### TXT

One phone number per line:

```text
+12025550123
+12025550124
+12025550125
```

### CSV

CSV fields must follow this order:

```text
phone,email,firstname,lastname,note
```

**Important:** This is the field order, not a header row. Do not include the header in your CSV file.

Example:

```csv
+12025550123,alice@example.com,Alice,Smith,Example contact
+12025550124,,John,Doe,
+12025550125,,,,
```

## 📱 Importing contacts on Android

Once you've generated your `.vcf` file:

1. Transfer the VCF file to your Android device.
2. Open the **Contacts** app.
3. Navigate to **Organize → Import from file**.
4. Select the account where you want to save your contacts.
5. Find and select your `.vcf` file.

Your contacts are now imported and ready to use.

[📖 Official Google guide: Import contacts from a VCF file](https://support.google.com/contacts/answer/15147365?co=GENIE.Platform%3DAndroid)

<!-- TODO: Add an Android import demonstration video here -->

---

## ⚠️ Disclaimer

This tool is intended for contact management, data processing, and legitimate OSINT research.

Use responsibly and respect other people's privacy.

---

### 🏷️ Tags

`#OSINT` `#vCard` `#VCF` `#WhatsApp` `#Contacts` `#Golang` `#CLI`