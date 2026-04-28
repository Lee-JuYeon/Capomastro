---
ide: terminal
appName: Terminal
description: Node.js REST API (Fastify + TypeScript)
---

fastify-cli generate creates a Fastify project with TypeScript, ESLint, Prettier, and Jest.
npm run dev starts the server on port 3000 with hot-reload.

```json
[
  { "action": "shell", "cmd": "npm init fastify@latest -- {name} --lang=ts", "cwd": "{outputDir}", "timeout": 60000 },
  { "action": "shell", "cmd": "npm install",                                   "cwd": "{outputDir}/{name}", "timeout": 90000 }
]
```
