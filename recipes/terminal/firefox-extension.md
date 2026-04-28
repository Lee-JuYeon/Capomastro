---
ide: terminal
appName: Terminal
description: Firefox Browser Extension (WebExtensions API + TypeScript)
---

web-ext scaffold creates a Firefox-compatible extension. Manifest V2/V3 both supported.
Uses the same Vite-based template as Chrome — cross-browser compatible out of the box.

```json
[
  { "action": "shell", "cmd": "npx create-chrome-ext@latest {name}", "cwd": "{outputDir}", "timeout": 60000 },
  { "action": "shell", "cmd": "npm install",                          "cwd": "{outputDir}/{name}", "timeout": 90000 }
]
```
