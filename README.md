# 🌌 Stardate CLI 🚀

A simple command-line tool to calculate **Star Trek Stardates** based on The Next Generation (TNG) formula.

## ⭐ Features
- 📅 Supports date input in **DD-MM-YYYY** format
- 📌 Defaults to the **current date** if no input is given
- 🔢 Accurate calculation considering **leap years**
- 💻 Lightweight, written in **Go**

- **Date Conversion:**  
  - Convert a human date to a stardate.
  - Convert a stardate back to a human date.

- **Customizable Base Year:**  
  - **Default Base Year:** 2323.
  - Temporarily override the base year for a single conversion using the `-base` (or `-b`) flag.
  - Persistently update the base year for all future conversions using the `-set-base` flag.
  - Check the current persistent base year with the `-show-base` flag.

- **Locale-Aware:**  
  All conversions use your local timezone for accurate date handling.

- **Short Flags:**  
  - `-d` for `--date`
  - `-s` for `--stardate`
  - `-b` for `--base`
  - Standard help flag (`-h` or `--help`)

- **Minimal Default Output:**  
  When no flags are provided, the CLI displays the current persistent base year and the current local date, with a hint to run `stardate --help` for more details.
---

## 📥 Installation

### **1. Download the Installer File**
#### **For Linux/macOS**
```sh
wget https://github.com/arcapol/stardate-cli/releases/latest/download/installer.sh
chmod +x installer.sh
sudo sh installer.sh
```
  The installer automatically:
   - Detects your operating system and architecture.
   - Downloads the appropriate tar.gz asset from the latest release.
   - Extracts and installs the `stardate` binary to `/usr/local/bin`.

---

### **2. Install via `go install` (For Go Developers)**
```sh
go install github.com/arcapol/stardate-cli@latest
```

---

## 🗑️ Uninstallation

### **Download the Uninstaller File**
#### **For Linux/macOS**
```sh
wget https://github.com/arcapol/stardate-cli/releases/latest/download/uninstaller.sh
chmod +x uninstaller.sh
sudo sh uninstaller.sh
```
---

## 📌 Usage

### **Get Current Stardate**
```sh
stardate
```
📌 **Example Output:**
```
Stardate: 142.47
```

- **Convert Human Date to Stardate:**  
  Specify a date with `-date` (or `-d`). If omitted, the current date is used.
  ```sh
  stardate -date 21-02-2025
  ```
  
- **Convert Stardate to Human Date:**  
  Use the `-stardate` (or `-s`) flag:
  ```sh
  stardate -stardate 45000
  ```

- **Temporary Base Year Override:**  
  Use `-base` (or `-b`) to set a base year for just this conversion:
  ```sh
  stardate -date 21-02-2025 -base 2300
  ```

- **Persistently Update Base Year:**  
  Update the stored base year (affects future conversions) with:
  ```sh
  stardate -set-base 2300
  ```

- **Show Current Persistent Base Year:**  
  ```sh
  stardate -show-base
  ```

- **Help:**  
  For detailed usage information, run:
  ```sh
  stardate --help
  ```

---

### **Calculate Stardate for a Specific Date**
```sh
stardate -date 21-12-2025
```
📌 **Example Output:**
```
Stardate: 972.60
```

---

## 📜 Stardate Formula
This CLI follows the **Star Trek: The Next Generation (TNG)** Stardate system:

$$
\text{Stardate} = 1000 \times (\text{Year} - 2025) + \frac{\text{Day of the Year}}{\text{Total Days in the Year}} \times 1000
$$

- **Reference Year:** 2025
- **Leap Year Handling:** Yes ✅

---

## 🛠 Development & Contribution

### **Clone the Repo**
```sh
git clone https://github.com/arcapol/stardate-cli.git
cd stardate-cli
```

### **Build & Run**
```sh
go build -o stardate
./stardate -date 01-01-2500
```

### **Submit a PR**
1. Fork the repository.
2. Create a new branch (`git checkout -b feature-xyz`).
3. Commit your changes (`git commit -m "Added feature XYZ"`).
4. Push and create a **Pull Request**.

---

## 📜 License
This project is open-source under the **Apache License 2.0**.

---

## 🚀 Live Long and Prosper! 🖖

