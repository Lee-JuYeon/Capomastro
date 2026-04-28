---
ide: terminal
appName: Terminal
description: Rust REST API (Axum + Tokio)
---

cargo new creates the crate; cargo add fetches Axum and Tokio.
Run with: cargo run (listens on :3000). Requires Rust toolchain (rustup).

```json
[
  { "action": "shell", "cmd": "cargo new {name}", "cwd": "{outputDir}", "timeout": 15000 },
  { "action": "shell", "cmd": "cargo add axum tokio --features tokio/full", "cwd": "{outputDir}/{name}", "timeout": 90000 }
]
```
