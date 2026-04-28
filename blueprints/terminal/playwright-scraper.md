---
ide: terminal
appName: Terminal
description: Web Scraper / Crawler (TypeScript + Playwright)
---

Playwright supports Chromium, Firefox, and WebKit. Headless by default; set headless: false for debug.
Run: npx playwright test (or ts-node scraper.ts for scripted crawls).

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "npm init -y && npm install playwright @playwright/test && npm install -D typescript ts-node @types/node && npx playwright install chromium", "cwd": "{outputDir}/{name}", "timeout": 180000 }
]
```
