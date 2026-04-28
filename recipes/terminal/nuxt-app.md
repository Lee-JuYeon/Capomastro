---
ide: terminal
appName: Terminal
description: Nuxt 3 Vue App (TypeScript)
---

nuxi init creates a minimal Nuxt 3 project. --no-install defers npm install so it runs offline-safe.
Run `npm install` inside the project after creation.

```json
[
  { "action": "shell", "cmd": "npx nuxi@latest init {name} --no-install", "cwd": "{outputDir}", "timeout": 60000 },
  { "action": "shell", "cmd": "npm install",                               "cwd": "{outputDir}/{name}", "timeout": 120000 }
]
```
