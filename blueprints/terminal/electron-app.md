---
ide: terminal
appName: Terminal
description: Electron Desktop App (TypeScript + Vite)
---

electron-vite scaffolds a cross-platform desktop app (macOS/Windows/Linux) with TypeScript + Vite.
Main process, preload, and renderer are all TypeScript out of the box.

```json
[
  { "action": "shell", "cmd": "npm create @quick-start/electron@latest {name} -- --template vanilla-ts", "cwd": "{outputDir}", "timeout": 60000 },
  { "action": "shell", "cmd": "npm install",                                                               "cwd": "{outputDir}/{name}", "timeout": 90000 }
]
```
