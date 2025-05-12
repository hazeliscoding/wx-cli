# wx-cli 🌦️  
Simple, cross‑platform CLI for a fast 7‑day weather forecast

![Go version](https://img.shields.io/badge/Go-1.21%2B-blue)
![License](https://img.shields.io/badge/license-MIT-green)

`wx-cli` pulls live data from **WeatherAPI.com** and presents it in a tidy, Unicode‑powered table.  
It auto‑detects your location via **ipinfo.io**, so a single command is all you need:

```bash
wx-cli forecast
````

---

## ✨ Features

* **One‑command forecast** – no flags required
* **Auto‑detect location** using your public IP
* 7‑day forecast with sunrise/sunset, temps & rain chance
* Works anywhere Go runs (Windows, Linux, macOS)
* Zero runtime dependencies beyond the compiled binary

---

## 📸 Quick Demo

```bash
$ wx-cli forecast
┌────────────────────────────────────────────────────────────┐
│ 🌍 Location: Houston, Texas, United States
│ 📅 Date: 05/12 | Monday
├────────────────────────────────────────────────────────────┤
│ 🌡️ Current: partly cloudy, 78°F (Feels like 80°F)
│ 🔽 Min: 72°F | 🔼 Max: 86°F
│ 💧 Humidity: 64%
└────────────────────────────────────────────────────────────┘

📅 6‑Day Forecast:
┌───────────┬───────────┬─────────────┬──────────────────────┬──────────┬──────────┐
│   Date    │ Temp (°F) │ Rain Chance │      Condition       │ Sunrise  │  Sunset  │
├───────────┼───────────┼─────────────┼──────────────────────┼──────────┼──────────┤
│ 05/12 Mon │       78  │  0 %        │ partly cloudy        │ 06:20 AM │ 08:04 PM │
│ …         │           │             │                      │          │          │
└───────────┴───────────┴─────────────┴──────────────────────┴──────────┴──────────┘
```

---

## 🚀 Installation

### Prerequisites

| Requirement        | Notes                                                                    |
| ------------------ | ------------------------------------------------------------------------ |
| **Go 1.21+**       | In your `PATH`                                                           |
| **WeatherAPI key** | Sign up free at [https://www.weatherapi.com](https://www.weatherapi.com) |

### 1. Install via `go install` (recommended)

```bash
go install github.com/hazeliscoding/wx-cli@latest
```

This drops a binary in **`$GOPATH/bin`** (commonly `~/go/bin`).
Add that directory to your *PATH* so you can call `wx-cli` anywhere.

### 2. (or) Build from source

```bash
git clone https://github.com/hazeliscoding/wx-cli.git
cd wx-cli
go build -o wx-cli .
```

On Windows you’ll get `wx-cli.exe`.

---

## 🔧 Configuration

`wx-cli` looks for an environment variable **`WEATHERAPI_API_KEY`**.
Create a small `.env` file so you don’t have to export it every session.

| OS                       | One‑liner                                                                                                                                                         |
| ------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Linux / macOS**        | `bash mkdir -p ~/.config/wx-cli && echo "WEATHERAPI_API_KEY=<YOUR_KEY>" > ~/.config/wx-cli/.env `                                                                 |
| **Windows (PowerShell)** | `pwsh New-Item -ItemType Directory -Force "$env:USERPROFILE\.config\wx-cli"; Set-Content "$env:USERPROFILE\.config\wx-cli\.env" 'WEATHERAPI_API_KEY=<YOUR_KEY>' ` |

> Prefer a global variable?
> • Linux/macOS `export WEATHERAPI_API_KEY=…`
> • Windows `setx WEATHERAPI_API_KEY …`

---

## 🏃 Usage

```bash
wx-cli forecast         # normal usage – auto‑detects city & region
go run . forecast       # if you’re hacking inside the repo
```

Exit codes: `0` = success · `1` = config/network error.

---

## 🤝 Contributing

1. Fork & create a feature branch
2. `go fmt ./...` before committing
3. Open a PR – GitHub Actions will run checks

---

## 🗺 Roadmap

* `--city` / `--zip` flags to override auto‑detect
* Metric/°C output
* Caching for offline viewing
* Test suite & CI badges

---

## 📝 License

This project is released under the **MIT License**.
See the [LICENSE](LICENSE) file for details.

<br>

Made with ❤️ by [Hazel Granados](https://github.com/hazeliscoding)