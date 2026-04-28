---
ide: terminal
appName: Terminal
description: Tauri Desktop App (Rust backend + TypeScript frontend)
---

create-tauri-app scaffolds a lightweight desktop app. Frontend: Vanilla TypeScript (swap to React/Vue post-creation).
Requires Rust toolchain + macOS Xcode Command Line Tools.

```json
[
  { "action": "shell", "cmd": "npm create tauri-app@latest {name} -- --template vanilla-ts --manager npm --yes", "cwd": "{outputDir}", "timeout": 120000 }
]
```
