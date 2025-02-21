# 🌌 Stardate CLI 🚀

A simple command-line tool to calculate **Star Trek Stardates** based on The Next Generation (TNG) formula.

## ⭐ Features
- 📅 Supports date input in **DD-MM-YYYY** format
- 📌 Defaults to the **current date** if no input is given
- 🔢 Accurate calculation considering **leap years**
- 💻 Lightweight, written in **Go**

---

## 📥 Installation

### **1. Download the Installer File**
#### **For Linux/macOS**
```sh
wget https://github.com/arcapol/stardate-cli/releases/latest/download/installer.sh
chmod +x installer.sh
sudo sh installer.sh
```

---

### **2. Install via `go install` (For Go Developers)**
```sh
go install github.com/arcapol/stardate-cli@latest
```

---

## 📥 Uninstallation

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

\[
\text{Stardate} = 1000 \times (\text{Year} - 2025) + \frac{\text{Day of the Year}}{\text{Total Days in the Year}} \times 1000
\]

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

