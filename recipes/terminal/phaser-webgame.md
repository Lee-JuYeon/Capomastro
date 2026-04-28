---
ide: terminal
appName: Terminal
description: Web Game (Phaser 3 + TypeScript)
---

Phaser 3 Vite template: full TypeScript setup with hot-reload dev server.
Includes Scene management, asset loader, physics (Arcade), and input handling.

```json
[
  { "action": "shell", "cmd": "npm create vite@latest {name} -- --template vanilla-ts", "cwd": "{outputDir}", "timeout": 30000 },
  { "action": "shell", "cmd": "npm install phaser", "cwd": "{outputDir}/{name}", "timeout": 60000 }
]
```
