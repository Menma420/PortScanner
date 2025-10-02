# Mini Port Scanner 🔍

A fast, lightweight, and concurrent **TCP port scanner** written in **Go**. Designed for quick reconnaissance of hosts and ports with human-friendly output formats.

It supports:

* Scanning single or multiple ports, including ranges (`22,80,8000-8010`)
* Resolving hostnames to IPs
* Concurrency for fast scanning
* Output in **table** or **JSON**
* Optional confirmation before scanning

---

## ⚡ Features

* Multi-threaded scanning with configurable concurrency
* Timeout control for slow or unreachable ports
* Clear CLI table output
* JSON output for easy integration with other tools
* Works with single IPs, hostnames, and multiple ports

---

## 🛠 Installation

1. **Install Go** if you haven’t already:
   [https://go.dev/doc/install](https://go.dev/doc/install)

2. **Clone the repo**:

```bash
git clone https://github.com/Menma420/PortScanner.git
cd mini-portscanner
```

3. **Install dependencies** (for table output):

```bash
go get github.com/olekukonko/tablewriter
```

---

## 🚀 Usage

```bash
go run ./cmd/mini-portscanner --target <IP/hostname> --ports <ports> --output <table/json> --confirm
```

### Options

| Flag            | Description                                       | Default |
| --------------- | ------------------------------------------------- | ------- |
| `--target`      | Target IP address or hostname (required)          | ""      |
| `--ports`       | Ports to scan (single, comma-separated, or range) | 80      |
| `--concurrency` | Number of threads/workers                         | 100     |
| `--timeout`     | Timeout per port in seconds                       | 1.5     |
| `--output`      | `table` or `json`                                 | table   |
| `--confirm`     | Required to start scan                            | false   |

### Examples

* Scan a single port:

```bash
go run ./cmd/mini-portscanner --target 127.0.0.1 --ports 80 --confirm
```

* Scan multiple ports:

```bash
go run ./cmd/mini-portscanner --target google.com --ports 22,80,443 --output json --confirm
```

* Scan a range of ports:

```bash
go run ./cmd/mini-portscanner --target 192.168.1.1 --ports 8000-8010 --concurrency 50 --confirm
```

---

## 📝 Sample Output

**Table format:**

| IP        | Port | Open  | Latency | Error              |
| --------- | ---- | ----- | ------- | ------------------ |
| 127.0.0.1 | 22   | false | 1.03ms  | connection refused |
| 127.0.0.1 | 80   | true  | 0.52ms  | nil              |
| 127.0.0.1 | 8000 | true  | 0.49ms  | nil              |

**JSON format:**

```json
[
  {
    "ip": "127.0.0.1",
    "port": 80,
    "open": true,
    "latency": "0.49ms",
    "err": null
  },
  {
    "ip": "127.0.0.1",
    "port": 22,
    "open": false,
    "latency": "1.03ms",
    "err": "dial tcp 127.0.0.1:22: connect: connection refused"
  }
]
```

---

## ⚙️ Future Enhancements

* Automatic banner grabbing for open ports (HTTP, SSH, FTP…)
* CLI animation / spinner during scanning
* Export results to CSV or integrate with other tools

---

## ⚠️ Disclaimer

This tool is intended for **educational purposes only**. Only scan systems that you own or have explicit permission to test. Unauthorized scanning may be illegal and is **not condoned**.
