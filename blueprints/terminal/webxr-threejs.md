---
ide: terminal
appName: Terminal
description: WebXR / AR / VR Web App (Three.js + WebXR Device API)
---

Vite + Three.js scaffold with WebXR session support. Works in Chrome on Quest (WebVR) or as AR on mobile.
npx vite (dev server) → open in WebXR-enabled browser.

```json
[
  { "action": "shell", "cmd": "npm create vite@latest {name} -- --template vanilla-ts", "cwd": "{outputDir}", "timeout": 30000 },
  { "action": "shell", "cmd": "npm install three @types/three", "cwd": "{outputDir}/{name}", "timeout": 60000 }
]
```
