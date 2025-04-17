# Go-Meter
Go-Meter is high performing Load Testing tool written in Go, built specifically to test TCP based services. It enables developers to simulate massive concurrent traffic, validate custom binary protocols, and measure throughput, latency, and system resilience — all without leaving the Go ecosystem.

---

## 🔧 Features
- 🔌 TCP client load generation using custom payloads
- 🧵 Configurable number of concurrent connections
- ⏱ Adjustable test duration and send intervals
- 📂 YAML-based config file
- 📦 Hex-encoded payload support for binary-safe transmission
- 🚦 Graceful shutdown via context
- 🧱 Modular, idiomatic Go project structure

---

## 📦 Installation
```bash
git clone https://github.com/yourusername/gometer.git
cd gometer
go mod tidy
```

---

## 🚀 Usage
### 1. Create a `config.yaml` file
```yaml
target_host: "127.0.0.1"
target_port: 3386
duration: 10s
connections: 10
payload_path: "payload.hex"
interval: 500ms
```

### 2. Create a `payload.hex` file
A sample payload for `PING\n`:
```
50494e470a
```

### 3. Run a test TCP server (optional)
```bash
nc -lk 3386
```

> Or use the included dev TCP server in `dev/echo_server.go`

### 4. Run gometer
```bash
go run ./cmd/gometer
```

Set a custom config file:
```bash
GOMETER_CONFIG=./my-test.yaml go run ./cmd/gometer
```

---

## 📁 Project Structure
```
gometer/
├── cmd/gometer         # Entry point
├── internal/
│   ├── config          # Config loading and parsing
│   ├── core            # Load test runner
│   ├── protocol        # TCP client logic
│   └── zlog            # Logging wrapper around zerolog
```

---

## 🧪 Example
```bash
$ go run ./cmd/gometer
15:04PM INF Starting load test target=127.0.0.1 port=3386 duration=10s connections=10
15:04PM DBG Worker started worker=3
...
15:04PM INF All workers completed
15:04PM INF [gometer] Exited cleanly.
```

---

## 📄 License
MIT

---

## 💡 TODO (Coming Soon)
- Metrics reporting (bytes sent, MBps, etc)
- JSON report output
- Custom payload per connection
- Prometheus integration

---

## 👋 Contributions Welcome!
Open an issue or PR if you'd like to extend `gometer`. This is just the beginning.