---
ide: terminal
appName: Terminal
description: Rust CLI Tool (Clap + Tokio)
---

cargo new creates the binary crate. cargo add clap fetches the CLI argument parser with derive macros.
Build with: cargo build --release. Binary output: target/release/{name}

```json
[
  { "action": "shell", "cmd": "cargo new {name} --bin", "cwd": "{outputDir}", "timeout": 15000 },
  { "action": "shell", "cmd": "cargo add clap --features derive && cargo add tokio --features full", "cwd": "{outputDir}/{name}", "timeout": 90000 }
]
```
