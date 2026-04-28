---
ide: terminal
appName: Terminal
description: Node.js Real-time Server (WebSocket + Socket.io)
---

Creates a Socket.io server with TypeScript. Client and server typings included.
Run with: npm run dev (listens on :3000 with /socket.io endpoint).

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "npm init -y && npm install socket.io && npm install -D typescript ts-node-dev @types/node", "cwd": "{outputDir}/{name}", "timeout": 60000 }
]
```
