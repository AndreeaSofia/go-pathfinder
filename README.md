<h1 align="center">🔎 Go Pathfinder</h1>

<p align="center">
  A small TCP port scanner built in Go.<br>
  Built to explore Goroutines, timeouts (and structuring something that isn't a script, for once).
</p>

---

## ✨ What is Go Pathfinder?

**Go Pathfinder** is a basic TCP port scanner that takes a host and a port range, and scans everything in between using Goroutines and a timeout dial.

> ⚠️ **Disclaimer:** **Go Pathfinder** is intended for learning & to be used **ethically and responsibly**, only on networks one owns or has permission to scan.

I wanted to try out Go’s concurrency properly and apply it to something small, but complete.

The main goal was to keep the code organized from the start, split across files properly, readable, and useful to me while I was writing it.

---

## 🧠 How it works

**Go Pathfinder** takes a host and a port range, then spins up a Goroutine for each port to scan it concurrently.

Each Goroutine attempts a TCP connection using a short timeout. If the connection succeeds, the port is marked as open. Everything else is silent,  closed ports just time out and move on.

All of the scanning runs in parallel using a `sync.WaitGroup` to wait for all Goroutines to finish. No external libraries, just Go’s `net`, `time`, and `sync` packages.

---

## 💻 How to run

1. Clone the repository:

```
git clone https://github.com/AndreeaSofia/go-pathfinder.git
```

2. Navigate to the project folder:

```
cd go-pathfinder
```

3. Build the binary:

```
go build -o pathfinder
```

4. Run the tool:

```
./pathfinder -host 127.0.0.1 -ports 20-80
```

☝️ That example scans ports 20 to 80 on your local machine. Try changing the host or port range to suit your needs!

## ✨ Features

- 🔍 Scans open ports on a target host (using concurrency for speed).
- 🎯 Supports flexible port ranges (e.g. `1-100`, `22,80,443`).
- ⚙️ Simple command-line interface with clear flags.
- 📦 Built with Go (no dependencies, just one small binary).
- 🧩 Modular code structure (easy to read and extend).

## 📜 License

This project is licensed under the [MIT License](LICENSE).

That means you’re free to use, modify, and share it, just keep the original license and give credit where it’s due.
