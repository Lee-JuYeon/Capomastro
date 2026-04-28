---
ide: terminal
appName: Terminal
description: Chrome Browser Extension (Manifest V3 + TypeScript)
---

create-chrome-ext scaffolds a Manifest V3 extension with a popup, background service worker,
and content script. Uses Vite for bundling.

```json
[
  { "action": "shell", "cmd": "npx create-chrome-ext@latest {name}", "cwd": "{outputDir}", "timeout": 60000 },
  { "action": "shell", "cmd": "npm install",                          "cwd": "{outputDir}/{name}", "timeout": 90000 }
]
```
