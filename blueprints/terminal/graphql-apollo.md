---
ide: terminal
appName: Terminal
description: GraphQL API (Apollo Server 4 + TypeScript)
---

Creates an Apollo Server 4 project with TypeScript, type-graphql schema, and nodemon for dev.

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "npm init -y && npm install @apollo/server graphql type-graphql reflect-metadata && npm install -D typescript ts-node-dev @types/node", "cwd": "{outputDir}/{name}", "timeout": 90000 }
]
```
